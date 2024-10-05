package models
type Wishlist struct {
	Id       int `json:"id"`
	User_id  int `json:"user_id" form:"user_id"`
	Event_id int `json:"event_id" form:"event_id"`
}