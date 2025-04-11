package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func InitMySQL() (db *sql.DB, err error) {
	// Carga las variables del entorno desde el archivo .env
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error cargando el archivo .env: %v", err)
	}

	// Lee las credenciales de la base de datos
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// formatea el DSN para MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	// conectar a la base de datos
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// verificar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error al verificar la conexión a la base de datos: %v", err)
	}

	fmt.Println("Conexión exitosa a la base de datos")

	return db, nil
}
