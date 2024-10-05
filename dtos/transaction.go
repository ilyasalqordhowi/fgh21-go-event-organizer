package dtos

type Transaction struct{
	Id 				int `json:"id"`
	EventId 		int `json:"eventId" form:"event_id" db:"event_id"`
	PaymentMethodId int `json:"payment_method_id" form:"payment_method_id" db:"payment_method_id"`
	UserId 		  	int `json:"userId" db:"user_id"`
	SectionId     	[]int `json:"sectionId" form:"section_id" db:"section_id"`
	TicketQty     	[]int `json:"ticketQty" form:"ticket_qty" db:"ticket_qty"`
}
type TransactionDetail struct {
	Id             int `json:"id"`
	TransactionId  int `json:"transactionId" form:"transactionId" db:"transaction_id"`
	SectionId      int `json:"sectionId" form:"sectionId" db:"section_id"`
	TicketQuantity int `json:"ticketQuantity" form:"ticketQuantity" db:"ticket_qty"`
}
