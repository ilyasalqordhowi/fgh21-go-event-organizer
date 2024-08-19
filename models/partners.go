package models

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)
type Partners struct{
	Id            int`json:"id"`
	Name          string `json:"name" form:"name" db:"name"`
	Image       *string `json:"img" form:"img" db:"img"`
}
func FindAllPartner() []Partners {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "partners" order by "id" asc`,
	)
	partner, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Partners])
	if err != nil {
		fmt.Println(err)
	}

	return partner
}