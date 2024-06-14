package handlers

import (
	"certainwager-be/db"
	"certainwager-be/models"
	"encoding/json"
	"net/http"
)

func GetReviews(w http.ResponseWriter, r *http.Request) {
	var reviews []models.Review
	rows, err := db.DB.Query("SELECT id, title, content, rating FROM reviews")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.ID, &review.Title, &review.Content, &review.Rating); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reviews = append(reviews, review)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}
