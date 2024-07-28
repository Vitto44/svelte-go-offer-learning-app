package handlers

import (
	"certainwager-be/db"
	"certainwager-be/models"
	"encoding/json"
	"net/http"
)

func GetOffers(w http.ResponseWriter, r *http.Request) {
	var offers []models.Offer
	rows, err := db.DB.Query("SELECT id, title, description FROM offers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var offer models.Offer
		if err := rows.Scan(&offer.ID, &offer.Title); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		offers = append(offers, offer)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(offers)
}
