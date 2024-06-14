package main

type Offer struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
}

type Review struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Content     string `json:"content"`
    Rating      int    `json:"rating"`
}

type Blog struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
}
