package main

import (
	"log"
	"net/http"

	"certainwager-be/db"
	"certainwager-be/router"
)

func main() {
	db.InitDB()

	r := router.NewRouter()

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
