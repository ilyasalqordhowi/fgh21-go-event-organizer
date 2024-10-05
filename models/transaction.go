package models

type ResultDetail struct {
	Id             int       `json:"id"`
	FullName       string    `json:"fullName" form:"fullName" db:"full_name"`
	Title     	   string    `json:"eventTitle" form:"eventTitle" db:"title"`
	LocationId     *int      `json:"location_id" form:"location_id" db:"location"`
	Date           string 	 `json:"date" form:"date" db:"date"`
	PaymentId      string    `json:"PaymentId" form:"PaymentId" db:"payment_method_id"`
	SectionName    []string  `json:"sectionName" form:"sectionName" db:"name"`
	TicketQuantity []int     `json:"TicketQuantity" form:"TicketQuantity" db:"tick	et_qty"`
}