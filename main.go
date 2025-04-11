package main

import (
	"consumer/src/buses/infraestructure/dependencies"
	"consumer/src/buses/infraestructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

 func main() {
 	r := gin.Default()

 	r.Use(cors.New(cors.Config{
 		AllowOrigins:     []string{"http://localhost:4200"}, // Origen de tu frontend Angular
 		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Métodos permitidos
 		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"}, // Encabezados permitidos
 		ExposeHeaders:    []string{"Content-Length"}, // Encabezados expuestos
		AllowCredentials: true, // Permitir credenciales (cookies, auth headers, etc.)
 	}))

 	routes.Routes(r)
	dependencies.InitBus()
	
 	r.Run(":8081")
 }