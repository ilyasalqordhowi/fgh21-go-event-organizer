package models
type PaymentMethod struct{
	Id int `json:"id"`
	Name string `json:"name" db:"name"`
}