package models

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)
type PaymentMethod struct{
	Id int `json:"id"`
	Name string `json:"name" db:"name"`
}

func FindAllPaymentMethod(search string, page int, limit int) ([]PaymentMethod, int) {
	db := lib.DB()
	defer db.Close(context.Background())
	offset := (page - 1) * limit

	sql := `SELECT * FROM "payment_method" where "name" ilike '%' || $1 || '%' offset $2 limit $3`
	rows, _ := db.Query(context.Background(), sql, search, offset, limit)
	payment, err := pgx.CollectRows(rows, pgx.RowToStructByPos[PaymentMethod])

	fmt.Println(payment)

	if err != nil {
		fmt.Println(err)
	}
	result := TotalPayment(search)
	return payment, result
}
func TotalPayment(search string) int {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT count(id) as "total" FROM "payment_method" where "name" ilike '%' || $1 || '%'`
	rows := db.QueryRow(context.Background(), sql, search)
	var results int
	rows.Scan(
		&results,
	)
	return results
}