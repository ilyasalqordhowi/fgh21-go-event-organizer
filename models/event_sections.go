package models
type EventSections struct{
	Id int `json:"id"`
	Name string `json:"name" db:"name"`
	Price int `json:"price"  db:"price"`
	Quantity string `json:"quantity"  db:"quantity"`
	EventId int `json:"events_id" db:"events_id"`
}