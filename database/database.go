package database

import (
	"database/sql"
)

func InitDb() *sql.DB {
	connectionString := "root:admin@tcp(localhost:3306)/northwind"
	databaseConnection, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error()) //Error handling = Manejo de errores
	}
	return databaseConnection
}
