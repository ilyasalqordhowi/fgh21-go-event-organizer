package models
type Nationality struct{
    Id int `json:"id"`
    Name string `json:"nationalities" db:"name"`
}