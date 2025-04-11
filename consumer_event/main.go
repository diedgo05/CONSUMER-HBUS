package main

import (
	"bytes"
	"consumer/models"
	"consumer/rabbit"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	
	rabbit := rabbit.NewRabbit()

	// obtiene los mensajes de rabbit
	msgs := rabbit.ReceiveContent()

	// Procesar los mensajes
	ProcessMessage(msgs)
}

// par procesar los mensajes de rabbit
func ProcessMessage(msgs <-chan amqp.Delivery) {
	forever := make(chan struct{})

	//  el goroutine para pdoer procesar los mensajes
	go func() {
		for d := range msgs {
			var bus models.Buses

			
			err := json.Unmarshal(d.Body, &bus)
			if err != nil {
				log.Printf("Error al decodificar el mensaje: %s", err)
				continue
			}

			log.Printf("[x] Cambio recibido: ID Bus %d, Disponible %d", bus.IdBus, bus.Disponible)

			// realiza la peticion para la validacion ala api2
			Fetch(bus)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

type updateDisponible struct {
	Disponible bool `json:"disponible"`
}

// realizar el post a la api2
func Fetch(bus models.Buses) {
	url := fmt.Sprintf("http://localhost:8081/buses/%d", bus.IdBus) 

	// vualve el objeto a JSON
	payload := updateDisponible{Disponible: true}
	
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error al serializar el cambio: %v", err)
	}
	
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalf("Error al crear la petición PUT: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error al hacer la petición: %v", err)
	}
	defer resp.Body.Close()
	

	//lee lo que viene de la api
	var result struct {
		Status string `json:"status"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalf("Error al decodificar la respuesta de la API: %v", err)
	}

	// respuesta hace::
	if result.Status == "aceptada" {
		log.Println("Inscripción aceptada")
	} else {
		log.Println("Inscripción rechazada")
	}
}