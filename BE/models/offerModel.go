package models

type Offer struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Image       string `json:"image"`
    Link       string `json:"link"`
    Rating      int    `json:"rating"`
    Category    string `json:"category"`
    Investment  int `json:"investment"`
    Reward      int `json:"reward"`
    RewardType  string `json:"rewardType"`
    Code        string `json:"code"`
    IsMoney    bool `json:"isMoney"`
    HowToClaim string `json:"howToClaim"`
    Disclaimer  string `json:"disclaimer"`
    ProviderID  int `json:"providerID"`
}
