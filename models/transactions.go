package models

import (
	"context"
	"fmt"
	"time"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Transaction struct{
	Id int `json:"id"`
	EventId int `json:"eventId" db:"event_id"`
	PaymentMethodId int `json:"payment_method" db:"payment_method_id"`
	UserId int `json:"userId" db:"user_id"`
}
type DetailTransaction struct {
	TransactionId 		int `json:"transactionId" db:"transaction_id"`
	FullName      string   `json:"fullName" db:"full_name"`
	Title         string   `json:"title"  db:"event_title"`
	LocationId    *int      `json:"locationId"  db:"location_id"`
	Date          time.Time `json:"date"  db:"date"`
	PaymentMethod string   `json:"paymentMethod"  db:"payment_method_id"`
	Section		[]string `json:"section_name" db:"name"`
	TicketQyt 	[]int `json:"ticketQyt" db:"ticket_qty"`
}
func CreateTransaction(data Transaction) Transaction {
db := lib.DB()
	defer db.Close(context.Background())

	sql := `insert into "transactions" ("event_id", "payment_method_id", "user_id") values ($1, $2, $3) returning "id", "event_id", "payment_method_id", "user_id"`
	row := db.QueryRow(context.Background(), sql, data.EventId, data.PaymentMethodId, data.UserId)

	var results Transaction
	row.Scan(
		&results.Id,
		&results.EventId,
		&results.PaymentMethodId,
		&results.UserId,
	)
	fmt.Println(results,"0iiiiiii")
	return results
}
func FindOneTransactionById(id int) Transaction {
	db := lib.DB()
	defer db.Close(context.Background())
	rows, _ := db.Query(context.Background(), `select * from "transactions" where "id" = $1`,
		id,
	)
	categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Transaction])

	if err != nil {
		fmt.Println(err)
	}

	category := Transaction{}
	for _, item := range categories {
		if item.Id == id {
			category = item
		}
	}
	return category
}
func CreateDetailTransactions()([]DetailTransaction,error){
	db := lib.DB()
	defer db.Close(context.Background())
	
	sql := `select "t"."id", "p"."full_name", "e"."title" as "event_title", "e"."location_id", "e"."date", "pm"."name" as "payment_method", array_agg("es"."name") as "ticket_section", array_agg("td"."ticket_qty") as "quantity"
    from "transactions" "t" 
    join "users" "u" on "u"."id" = "t"."user_id"
    join "profile" "p" on "p"."user_id" = "u"."id"
    join "events" "e" on "e"."id" = "t"."event_id"
    join "payment_method" "pm" on "pm"."id" = "t"."payment_method_id"
    join "transactions_details" "td" on "td"."transaction_id" = "t"."id"
    join "event_sections" "es" on "es"."id" = "td"."section_id"
    group by  "t"."id", "p"."full_name", "e"."title", "e"."location_id", "e"."date", "pm"."name";`
send, _ := db.Query(
		context.Background(),
		sql,
	)

	row, err := pgx.CollectRows(send, pgx.RowToStructByPos[DetailTransaction])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(row,err,"halo")
	return row, err
}
