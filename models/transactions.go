package models

import (
	"context"
	"fmt"
	"time"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
)

type Transaction struct{
	Id int `json:"id"`
	EventId int `json:"eventId" db:"event_id"`
	PaymentMethodId int `json:"payment_method" db:"payment_method_id"`
	UserId int `json:"userId" db:"user_id"`
}
type DetailTransaction struct {
	Id            int       `json:"id"`
	FullName      *string   `json:"fullName" form:"full name" db:"full_name"`
	Title         *string   `json:"title" form:"title" db:"event_title"`
	Date          time.Time `json:"date" form:"date" db:"date"`
	LocationId    *int      `json:"locationId" form:"locationId" db:"location_id"`
	PaymentMethod *string   `json:"paymentMethod" form:"payment_method" db:"payment_method"`
	TicketQyt 	[]int `json:"ticketQyt" db:"ticket_qyt"`
	Section		[]string `json:"section_name" db:"section_name"`
}
func CreateTransaction(transactions Transaction ,id int) (*Transaction, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	
sql := `insert into "transactions" ("event_id","payment_method","user_id") values ($1,$2,$3) returning "id","event_id","payment_method","user_id"`
row := db.QueryRow(context.Background(),sql,transactions.EventId,transactions.PaymentMethodId,transactions.UserId)
var results Transaction
row.Scan(
	&results.Id,
	&results.EventId,
	&results.PaymentMethodId,
	&results.UserId,
)
fmt.Println(results)
return &results,nil
}
func CreateDetailTransactions(id int)DetailTransaction{
	db := lib.DB()
	defer db.Close(context.Background())
	
	sql := `select t.id, p.full_name, e.title as "event_title",e.location_id,e."date", pm.name as "payment_method",array_agg(es.name) as "section_name",array_agg(td.ticket_qyt) as "ticket_qyt"
	form "transactions" t
	join users u on u.id = t.users_id
	join profile p on p.users_id = u.id
	join events e on e.id = t.event_id
	join paymnent_methods pm on pm.id = t.payment_method_id
	join transcation_details td on td.transaction_id = t.id
	join event_sections es on es.id = td.section_id
	group by t.id,p.full_name,e.title,e.location_id,e."date",pm.name;`
	row := db.QueryRow(context.Background(),sql,id)
	var results DetailTransaction
row.Scan(
	&results.Id,
	&results.FullName,
	&results.Title,
	&results.Date,
	&results.LocationId,
	
	&results.PaymentMethod,
	&results.TicketQyt,
	&results.Section,
)
fmt.Println(results)
return results
}


