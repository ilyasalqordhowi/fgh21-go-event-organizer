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
func FindSectionsByEvent(id int) EventSections {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "event_sections"`,
	)
	eventSection, err := pgx.CollectRows(rows, pgx.RowToStructByPos[EventSections])
	if err != nil {
		fmt.Println(err)
	}
	dataSections := EventSections{}
	for _, i := range eventSection {
		if i.Id == id {
			dataSections = i
		}
	}
	return dataSections
}