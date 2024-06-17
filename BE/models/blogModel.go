package models

type Blog struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
    Image   string `json:"image"`
    CardText string `json:"cardText"`
    Created string `json:"created"`
    Offer Offer `json:"offer"`

}
