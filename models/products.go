package models

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Brand       string    `json:"brand"`
	Price       float64   `json:"price"`
	Rating      float64   `json:"rating"`
	Type        string    `json:"type"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Features    []string  `json:"features"`
	Reviews     []Reviews `json:"reviews"`
}
type Reviews struct {
	Author  string  `json:"author"`
	Rating  float64 `json:"rating"`
	Comment string  `json:"comment"`
}
