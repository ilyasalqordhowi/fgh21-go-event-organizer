package models
type Locations struct {
    Id    int    `json:"id" db:"id"`
    Name  string `json:"name" form:"name" db:"name"`
    Lat   string `json:"lat" form:"lat" db:"lat"`
    Long  string `json:"long" form:"long" db:"long"`
    Img string `json:"img" form:"img" db:"img"`
}
