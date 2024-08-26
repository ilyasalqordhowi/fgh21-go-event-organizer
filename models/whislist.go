package models

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Wishlist struct {
	Id       int `json:"id"`
	User_id  int `json:"user_id" form:"user_id"`
	Event_id int `json:"event_id" form:"event_id"`
}

func FindAllWishlist() []Wishlist {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "whislist" order by "id" asc`,
	)

	dataWhislist, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Wishlist])
	if err != nil {
		fmt.Println(err)
	}
	return dataWhislist
}
func FindOneWishlist(id int) ([]Wishlist, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	
	rows, err := db.Query(context.Background(),
		`SELECT * FROM "whislist" WHERE "user_id" = $1 ORDER BY "id" ASC`, id,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query wishlist: %w", err)
	}
	defer rows.Close() 
	wishlists, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Wishlist])
	if err != nil {
		return nil, fmt.Errorf("failed to collect wishlist rows: %w", err)
	}


	return wishlists, nil
}
func CreateWishlist(event_id int, id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	var exists bool
	err := db.QueryRow(
		context.Background(),
		`SELECT EXISTS (SELECT 1 FROM "whislist" WHERE user_id = $1 AND event_id = $2)`,
		id, event_id,
	).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check existing wishlist entry: %w", err)
	}

	if exists {
		return fmt.Errorf("wishlist entry already exists")
	}

	_, err = db.Exec(
		context.Background(),
		`INSERT INTO "whislist" (user_id, event_id) VALUES ($1, $2)`,
		id, event_id,
	)

	if err != nil {
		return fmt.Errorf("failed to insert whislist entry: %w", err)
	}

	return nil
}
func FindOneEventById(event_id int) (Event, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	var event Event
	err := db.QueryRow(context.Background(),
		`SELECT id, image, title, date, descriptions, location_id, created_by 
         FROM "events" WHERE id = $1`, event_id,
	).Scan(&event.Id, &event.Image, &event.Title, &event.Date, &event.Descriptions, &event.LocationId, &event.CreateBy)

	if err != nil {
		return Event{}, fmt.Errorf("failed to find event: %w", err)
	}

	return event, nil
}
func Deletewishlist(user_id int, event_id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	result, err := db.Exec(
		context.Background(),
		`DELETE FROM "whislist" WHERE "user_id" = $1 AND "event_id" = $2`,
		user_id, event_id,
	)

	if err != nil {
		return fmt.Errorf("failed to delete wishlist item: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("wishlist item not found")
	}

	return nil
}