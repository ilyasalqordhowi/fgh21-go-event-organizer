package models

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)
type Event struct {
	Id            int`json:"id"`
	Image       string `json:"image" form:"image" db:"image"`
	Title      string `json:"title" form:"title" db:"title"`
	Date     string `json:"date" form:"date" db:"date"`
	Descriptions string`json:"descriptions" form:"descriptions" db:"descriptions"`
	LocationId *int `json:"locationId" form:"locationId" db:"location_id"`
	CreateBy  *int`json:"createBy" form:"createBy" db:"create_by"`
}


func FindAllEvent(search string, page int,limit int) ([]Event,int) {
	db := lib.DB()
	defer db.Close(context.Background())
	offset := (page - 1) * limit

	sql :=	`SELECT * FROM "events" where "title" ilike '%' || $1 || '%' offset $2 limit $3`
	rows, _ := db.Query(context.Background(),sql,search,offset,limit)
	events, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Event])
	if err != nil {
		fmt.Println(err)
	}
	result := TotalEvent(search)
	fmt.Println(events)
	return events, result
}

func TotalEvent(search string)int{
	db := lib.DB()
	defer db.Close(context.Background())

	sql :=	`SELECT count(id) as "total" FROM "events" where "title" ilike '%' || $1 || '%'`
	rows := db.QueryRow(context.Background(),sql,search)
	var results int
	rows.Scan(
		&results,
	)
	return results
}
func FindOneEvent(id int) Event {
	db := lib.DB()
	defer db.Close(context.Background())

	rows,_ := db.Query(context.Background(),
		`select * from "events"`,
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
func FindOneByEvent(id int) []Event {
	db := lib.DB()
    defer db.Close(context.Background())

    rows, err := db.Query(
        context.Background(),
        `select * from "events" where "created_by" = $1`, id,
    )
    if err != nil {
        fmt.Println(err)
    }
    events, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Event])

    if err != nil {
        fmt.Errorf("Error")
    }

    return events
}
func CreateEvents(event Event, id int) error {
    db := lib.DB()
    defer db.Close(context.Background())

    _, err := db.Exec(
        context.Background(),
        `INSERT INTO "events" ("image", "title", "date", "descriptions", "location_id", "created_by") VALUES ($1, $2, $3, $4, $5, $6)`,
        event.Image, event.Title, event.Date, event.Descriptions, event.LocationId, id,
    )
    fmt.Println(err)
    if err != nil {
        return fmt.Errorf("failed to execute insert")
    }
    return nil
}
func CreateEvent(event Event, id int) error {
    db := lib.DB()
    defer db.Close(context.Background())

    _, err := db.Exec(
        context.Background(),
        `insert into "events" (image, title, date, descriptions, location_id, created_by) values ($1, $2, $3, $4, $5, $6)`,
        event.Image, event.Title, event.Date, event.Descriptions, event.LocationId, id,
    )
    if err != nil {
		return fmt.Errorf("failed to execute insert")
    }
	fmt.Println(err,"models")
    return nil
}
func RemoveEvent(id int) error {
    db := lib.DB()
    defer db.Close(context.Background())

    commandTag, err := db.Exec(
        context.Background(),
        `DELETE FROM "events" WHERE "id" = $1`,
        id,
    )

    if err != nil {
        return fmt.Errorf("failed to execute delete")
    }

    if commandTag.RowsAffected() == 0 {
        return fmt.Errorf("no user found")
    }

    return nil
}
func EditEvent(Image string, Title string, Date int,Descriptions string, LocationId int, CreateBy int, id string) {
    db := lib.DB()
    defer db.Close(context.Background())

    dataSql := `update "events" set (image,title,date,descriptions, locatin_id, created_by) = ($1, $2, $3, $4, $5, $6) where "id" = $7`

    db.Exec(context.Background(), dataSql, Image, Title, Date,Descriptions, LocationId, CreateBy, id)
}


