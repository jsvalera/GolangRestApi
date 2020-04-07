package main

import (
	"GolangRestApi/database"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	databaseConnection := database.InitDb()
	defer databaseConnection.Close()
	fmt.Println(databaseConnection)
}
