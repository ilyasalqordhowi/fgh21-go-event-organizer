package models

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Profile struct {
	Id            int    `json:"id"`
	Picture       string `json:"picture" form:"picture"`
	FullName      string `json:"fullName" form:"fullName"`
	Birthdate     string `json:"birtdate" form:"birthdate"`
	Gender        *int    `json:"gender" form:"gender"`
	PhoneNumber   string `json:"phoneNumber" form:"phoneNumber"`
	Profession    string `json:"profession" form:"profession"`
	NationalityId *int `json:"nationalityId" form:"nationalityId"`
	UserId        *int `json:"userId" form:"userId"`
}

func CreateProfile(newProfile Profile) Profile {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `insert into "profile" ("picture","fullName","birthdate","gender","phoneNumber","profession","nationalityId","userId") values ($1,$2,$3,$4,$5,$6,$7,$8) returning "id","picture","fullName","birthdate","gender","phoneNumber","profession","nationalityId","userId"`
	row := db.QueryRow(context.Background(), sql, newProfile.Picture, newProfile.FullName, newProfile.Birthdate, newProfile.Gender, newProfile.PhoneNumber, newProfile.Profession, newProfile.NationalityId, newProfile.UserId)
	var results Profile
	fmt.Println(results)
	row.Scan(
		&results.Id,
		&results.Picture,
		&results.FullName,
		&results.Birthdate,
		&results.Gender,
		&results.PhoneNumber,
		&results.Profession,
		&results.NationalityId,
		&results.UserId,
	)
	return results
}
func FindOneProfile(id int) Profile {
    db := lib.DB()
    defer db.Close(context.Background())

    rows,_ := db.Query(context.Background(),
        `select * from "profile"`,
    )
    profile, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Profile])
    if err != nil {
        fmt.Println(err)
    }
    dataProfile := Profile{}
    for _, i := range profile {
        if i.Id == id {
            dataProfile = i
        }
    }
    return dataProfile
}


