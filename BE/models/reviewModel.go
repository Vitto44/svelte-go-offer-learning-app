package models

type Review struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
    Rating  int    `json:"rating"`
    Image  string `json:"image"`
    Link string `json:"link"`
    Created string `json:"created"`
    Offer Offer `json:"offer"`
    Description string `json:"description"`
    Disclaimer string `json:"disclaimer"`
    AvailibleCountries []string `json:"availibleCountries"`
    LicensedBy string `json:"licensedBy"`
    Games []string `json:"games"`
    Pros []string `json:"pros"`
    Cons []string `json:"cons"`
    Categories []string `json:"categories"`
}   
