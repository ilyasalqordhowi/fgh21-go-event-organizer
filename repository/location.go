package repository

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)


func FindAllLocations(search string, limit int, page int) []models.Locations {
    db := lib.DB()

    defer db.Close(context.Background())

    offset := (page - 1) * limit

    rows, _ := db.Query(
        context.Background(),
        `select * from "location" where "name" ilike '%' || $1 || '%' limit $2 offset $3`, search, limit, offset,
    )

    dataLocations, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Locations])

    if err != nil {
        fmt.Println(err)
    }
    return dataLocations
}

