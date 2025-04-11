package core

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQConn struct {
	Broker *amqp091.Connection
	Ch     *amqp091.Channel
	Err    string
}

func GetRabbitMQConn() (*RabbitMQConn, error) {
	error := ""
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}
	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")
	host := os.Getenv("RABBITMQ_IP")
	port := os.Getenv("RABBITMQ_PORT")

	rabbitURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

	//Conexión a RabbitMQ
	conn, err := amqp091.Dial(rabbitURL)

	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error al abrir un canal")
	}

	return &RabbitMQConn{Broker: conn, Ch: ch, Err: error}, nil
}
func (conn *RabbitMQConn) FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}