package models
type Partners struct{
	Id            int`json:"id"`
	Name          string `json:"name" db:"name"`
	Image       *string `json:"img"  db:"img"`
}