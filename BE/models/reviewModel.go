package models

type Review struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
    Rating  int    `json:"rating"`
}