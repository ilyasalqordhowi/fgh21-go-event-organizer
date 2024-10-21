package models

type EventLocation struct {
	Id          int    `json:"id"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Location    *string `json:"location"`
	CreatedBy   *int   `json:"created_by" db:"created_by"`
}