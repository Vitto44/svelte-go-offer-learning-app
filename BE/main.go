package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	initDB()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Casino Backend!")
	})

	r.HandleFunc("/offers", getOffers).Methods("GET")
	r.HandleFunc("/reviews", getReviews).Methods("GET")
	r.HandleFunc("/blogs", getBlogs).Methods("GET")

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getOffers(w http.ResponseWriter, r *http.Request) {
	var offers []Offer
	rows, err := db.Query("SELECT id, title, description FROM offers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var offer Offer
		if err := rows.Scan(&offer.ID, &offer.Title, &offer.Description); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		offers = append(offers, offer)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(offers)
}

func getReviews(w http.ResponseWriter, r *http.Request) {
	var reviews []Review
	rows, err := db.Query("SELECT id, title, content, rating FROM reviews")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var review Review
		if err := rows.Scan(&review.ID, &review.Title, &review.Content, &review.Rating); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reviews = append(reviews, review)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}

func getBlogs(w http.ResponseWriter, r *http.Request) {
	var blogs []Blog
	rows, err := db.Query("SELECT id, title, content FROM blogs")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var blog Blog
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		blogs = append(blogs, blog)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}
