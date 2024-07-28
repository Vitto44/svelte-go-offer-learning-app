package handlers

import (
	"certainwager-be/db"
	"certainwager-be/models"
	"encoding/json"
	"net/http"
)

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	var blogs []models.Blog
	rows, err := db.DB.Query("SELECT id, title, content FROM blogs")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var blog models.Blog
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		blogs = append(blogs, blog)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}
