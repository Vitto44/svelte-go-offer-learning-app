package main

import (
	"certainwager-be/db"
	"certainwager-be/router"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	db.InitDB()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := router.NewRouter()

	port := os.Getenv("PORT")


	log.Println("Server started at port " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}