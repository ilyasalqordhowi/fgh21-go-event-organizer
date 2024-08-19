package models

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
)

type TransactionDetail struct {
	Id            int `json:"id"`
	TransactionId int `json:"transactionId" db:"transaction_id"`
	SectionId     int `json:"sectionId" db:"section_id"`
	TicketQty     int `json:"ticketQty" db:"ticket_qty"`
}

func CreateTransactionDetail(data TransactionDetail) TransactionDetail{
	db := lib.DB()
	defer db.Close(context.Background())

	inputSQL := `insert into "transactions_details" (transaction_id, section_id, ticket_qty) values ($1, $2, $3) returning "id", "transaction_id", "section_id", "ticket_qty"`
	row := db.QueryRow(context.Background(), inputSQL, data.TransactionId, data.SectionId, data.TicketQty)

	var detail TransactionDetail

	row.Scan(
		&detail.Id,
		&detail.TransactionId,
		&detail.SectionId,
		&detail.TicketQty,
	)
	fmt.Println(detail,"kosonggg")
	return detail
}