package models

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)
type EventSections struct{
	Id int `json:"id"`
	Name string `json:"name" db:"name"`
	Price int `json:"price"  db:"price"`
	Quantity string `json:"quantity"  db:"quantity"`
	EventId int `json:"events_id" db:"events_id"`
}
func FindSectionsByEvent(id int) ([]EventSections,error) {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "event_sections" where "events_id" = $1`,id,
	)
	eventSection, err := pgx.CollectRows(rows, pgx.RowToStructByPos[EventSections])
	if err != nil {
		return nil,fmt.Errorf("Error")
	}
	
	return eventSection,nil
}