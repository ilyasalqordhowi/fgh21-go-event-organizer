package repository

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func FindSectionsByEvent(id int) ([]models.EventSections,error) {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "event_sections" where "events_id" = $1`,id,
	)
	eventSection, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.EventSections])
	if err != nil {
		return nil,fmt.Errorf("Error")
	}
	
	return eventSection,nil
}