package models
type Categories struct {
	Id            int`json:"id"`
	Categories *string `json:"categories" form:"categories" db:"name"`
}