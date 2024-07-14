package main

import (
	"goTodo/database"
	"goTodo/router"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	database.ConnectDb()
}

func main() {
	router.ClientRoutes()
}
