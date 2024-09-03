package models

type Court struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Location     string  `json:"location"`
	Availability string  `json:"availability"`
	Price        float64 `json:"price"`
}
