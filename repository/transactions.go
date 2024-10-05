package repository

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/dtos"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)





func CreateNewTransactions(tx pgx.Tx,data dtos.Transaction) (dtos.Transaction, error) {
    db := lib.DB()
    defer db.Close(context.Background())

 
    tx, err := db.BeginTx(context.Background(), pgx.TxOptions{})
    if err != nil {
        return dtos.Transaction{}, err 
    }

    sql := `insert into "transactions" ("event_id", "payment_method_id", "user_id") values ($1, $2, $3) returning "id", "event_id", "payment_method_id", "user_id"`
    row := tx.QueryRow(context.Background(), sql, data.EventId, data.PaymentMethodId, data.UserId)

    var results dtos.Transaction
    err = row.Scan(
        &results.Id,
        &results.EventId,
        &results.PaymentMethodId,
        &results.UserId,
    )
    if err != nil {
        return dtos.Transaction{}, err 
    }

    if commitErr := tx.Commit(context.Background()); commitErr != nil {
        return dtos.Transaction{}, commitErr 
    }

    fmt.Println(results, "hasil")
    return results, nil 
}


func CreateTransactionDetail(tx pgx.Tx,data dtos.TransactionDetail) (dtos.TransactionDetail,error) {
	db := lib.DB()
	defer db.Close(context.Background())
	tx, err := db.BeginTx(context.Background(), pgx.TxOptions{})
    if err != nil {
        return dtos.TransactionDetail{}, err 
    }

    sql := `insert into "transactions_details" (transaction_id, section_id, ticket_qty) values ($1, $2, $3) returning "id", "transaction_id", "section_id", "ticket_qty"`
    row := tx.QueryRow(context.Background(), sql, data.TransactionId, data.SectionId, data.TicketQuantity)

    var detail dtos.TransactionDetail
    err = row.Scan(
        &detail.Id,
        &detail.TransactionId,
        &detail.SectionId,
        &detail.TicketQuantity,
    )
    if err != nil {
        tx.Rollback(context.Background()) 
        return dtos.TransactionDetail{}, err 
    }

    if commitErr := tx.Commit(context.Background()); commitErr != nil {
        return dtos.TransactionDetail{}, commitErr 
    }

    fmt.Println(detail, "detail")
    return detail, nil
}


func DetailsTransaction(id int) ([]models.ResultDetail, error) {
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

    row, err := pgx.CollectRows(send, pgx.RowToStructByPos[models.ResultDetail])
    if err != nil {
        fmt.Println(err)
    }
    return row, err
}

