package router

import (
	"certainwager-be/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Casino Backend!")
	})

	r.HandleFunc("/offers", handlers.GetOffers).Methods("GET")
	r.HandleFunc("/reviews", handlers.GetReviews).Methods("GET")
	r.HandleFunc("/blogs", handlers.GetBlogs).Methods("GET")

	return r
}
