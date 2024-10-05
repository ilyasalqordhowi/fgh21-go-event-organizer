package repository

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func FindAllPartner() []models.Partners {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "partners" order by "id" asc`,
	)
	partner, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Partners])
	if err != nil {
		fmt.Println(err)
	}

	return partner
}