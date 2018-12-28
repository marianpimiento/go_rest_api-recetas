package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// getConnection --> Obteniene la conexion con la BD
func getConnection() *sql.DB {
	dsn := "postgresql://maxroach@localhost:26257/recetas?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error conectando con la Base de Datos: ", err)
	}
	return db
}