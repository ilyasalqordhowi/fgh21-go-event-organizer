package models

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Event struct {
	Id            int    `json:"id"`
	Image       string `json:"image" form:"image"`
	Title     string `json:"title" form:"title"`
	Date     string `json:"date" form:"date"`
	Descriptions string  `json:"descriptions" form:"descriptions"`
	LocationId  *int`json:"locationId" form:"locationId"`
	CreateBy    *int`json:"createBy" form:"createBy"`
}

func CreateEvent(newEvent Event) Event {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `insert into "event" ("image","title","date","descriptions","locationId","createBy") values ($1,$2,$3,$4,$5,$6) returning "id","image","title","date","descriptions","locationId","createBy"`
	row := db.QueryRow(context.Background(), sql, newEvent.Image, newEvent.Title, newEvent.Date, newEvent.Descriptions, newEvent.LocationId, newEvent.CreateBy)
	var results Event
	fmt.Println(results)
	row.Scan(
		&results.Id,
		&results.Image,
		&results.Title,
		&results.Date,
		&results.Descriptions,
		&results.LocationId,
		&results.CreateBy,
	)
	return results
}
func FindOneEvent(id int) Event {
    db := lib.DB()
    defer db.Close(context.Background())

    rows,_ := db.Query(context.Background(),
        `select * from "event"`,
    )
    event, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Event])
    if err != nil {
        fmt.Println(err)
    }
    dataEvent := Event{}
    for _, i := range event {
        if i.Id == id {
            dataEvent = i
        }
    }
    return dataEvent
}


