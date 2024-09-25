package models

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

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
func CreateNewTransactions(tx pgx.Tx,data Transaction) (Transaction, error) {
    db := lib.DB()
    defer db.Close(context.Background())

 
    tx, err := db.BeginTx(context.Background(), pgx.TxOptions{})
    if err != nil {
        return Transaction{}, err 
    }

    sql := `insert into "transactions" ("event_id", "payment_method_id", "user_id") values ($1, $2, $3) returning "id", "event_id", "payment_method_id", "user_id"`
    row := tx.QueryRow(context.Background(), sql, data.EventId, data.PaymentMethodId, data.UserId)

    var results Transaction
    err = row.Scan(
        &results.Id,
        &results.EventId,
        &results.PaymentMethodId,
        &results.UserId,
    )
    if err != nil {
        return Transaction{}, err 
    }

    if commitErr := tx.Commit(context.Background()); commitErr != nil {
        return Transaction{}, commitErr 
    }

    fmt.Println(results, "hasil")
    return results, nil 
}


func CreateTransactionDetail(tx pgx.Tx,data TransactionDetail) (TransactionDetail,error) {
	db := lib.DB()
	defer db.Close(context.Background())
	tx, err := db.BeginTx(context.Background(), pgx.TxOptions{})
    if err != nil {
        return TransactionDetail{}, err 
    }

    sql := `insert into "transactions_details" (transaction_id, section_id, ticket_qty) values ($1, $2, $3) returning "id", "transaction_id", "section_id", "ticket_qty"`
    row := tx.QueryRow(context.Background(), sql, data.TransactionId, data.SectionId, data.TicketQuantity)

    var detail TransactionDetail
    err = row.Scan(
        &detail.Id,
        &detail.TransactionId,
        &detail.SectionId,
        &detail.TicketQuantity,
    )
    if err != nil {
        tx.Rollback(context.Background()) 
        return TransactionDetail{}, err 
    }

    if commitErr := tx.Commit(context.Background()); commitErr != nil {
        return TransactionDetail{}, commitErr 
    }

    fmt.Println(detail, "detail")
    return detail, nil
}


func DetailsTransaction(id int) ([]ResultDetail, error) {
    db := lib.DB()
    defer db.Close(context.Background())

    sql :=
        `select t.id, p.full_name, e.title as "event_title", e.location_id, e.date, pm.name as "payment_method",
        array_agg(es.name) as "section_name", array_agg(td.ticket_qty) as "ticket_qty"
        from "transactions" "t"
        join "users" "u" on u.id = t.user_id
        join "profile" "p" on p.user_id = u.id
        join "events" "e" on e.id = t.event_id
        join "payment_method" "pm" on pm.id = t.payment_method_id
        join "transactions_details" "td" on td.transaction_id = t.id
        join "event_sections" "es" on es.id = td.section_id
        where t.user_id = $1
        group by t.id, p.full_name, e.title, e.location_id, e.date, pm.name`

    send, _ := db.Query(
        context.Background(),
        sql,
        id,
    )

    row, err := pgx.CollectRows(send, pgx.RowToStructByPos[ResultDetail])
    if err != nil {
        fmt.Println(err)
    }
    return row, err
}

