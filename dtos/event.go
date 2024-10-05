package dtos

type Event struct {
	Id            int`json:"id"`
	Image       string `json:"image" form:"image" db:"image"`
	Title      string `json:"title" form:"title" db:"title"`
	Date     string `json:"date" form:"date" db:"date"`
	Descriptions string`json:"descriptions" form:"descriptions" db:"descriptions"`
	LocationId *int `json:"locationId" form:"locationId" db:"location_id"`
	CreateBy  *int`json:"createBy" form:"createBy" db:"create_by"`
}