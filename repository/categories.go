package repository

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
	"github.com/jackc/pgx/v5"
)

func FindAllCategories(search string, page int,limit int) ([]models.Categories) {
	db := lib.DB()
	defer db.Close(context.Background())
	offset := (page - 1) * limit

	sql :=	`SELECT * FROM "categories" where "name" ilike '%' || $1 || '%' offset $2 limit $3`
	rows, _ := db.Query(context.Background(),sql,search,offset,limit)
	categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Categories])

	fmt.Println(categories)


	if err != nil {
		fmt.Println(err)
	}
	
	return categories
}
func TotalCategory(search string)int{
	db := lib.DB()
	defer db.Close(context.Background())

	sql :=	`SELECT count(id) as "total" FROM "categories" where "name" ilike '%' || $1 || '%'`
	rows := db.QueryRow(context.Background(),sql,search)
	var results int
	rows.Scan(
		&results,
	)
	return results
}
func FindOneCategories(id int) models.Categories {
	db := lib.DB()
	defer db.Close(context.Background())

	rows,_ := db.Query(context.Background(),
		`select * from "categories"`,
	)
	categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Categories])
	if err != nil {
		fmt.Println(err)
	}
	dataCategories := models.Categories{}
	for _, i := range categories {
		if i.Id == id {
			dataCategories = i
		}
	}
	return dataCategories
}
func CreateCategories(Categories models.Categories, id int) error {
    db := lib.DB()
    defer db.Close(context.Background())

    _, err := db.Exec(
        context.Background(),
        `insert into "categories" (name) values ($1)`,
        Categories.Categories,
    )
	fmt.Println(err)
    if err != nil {
        return fmt.Errorf("failed to execute insert")
    }
    return nil
}
func RemoveCategories(id int) error {
    db := lib.DB()
    defer db.Close(context.Background())

    commandTag, err := db.Exec(
        context.Background(),
        `DELETE FROM "categories" WHERE "id" = $1`,
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
func EditCategories(Categories string) {
    db := lib.DB()
    defer db.Close(context.Background())

    dataSql := `update "categories" set (name)=($1) where "id" = $2`

    db.Exec(context.Background(), dataSql, Categories)
}
