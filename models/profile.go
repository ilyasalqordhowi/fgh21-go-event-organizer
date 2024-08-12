package models

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)
type Profile struct {
	Id            int    `json:"id"`
	Picture       *string `json:"picture" form:"picture" db:"picture"`
	FullName      string `json:"fullName" form:"fullName" db:"full_name"`
	Birthdate     *string `json:"birtdate" form:"birthdate" db:"birth_date"`
	Gender        *int    `json:"gender" form:"gender" db:"gender"`
	PhoneNumber   *string `json:"phoneNumber" form:"phoneNumber" db:"phone_number"`
	Profession    *string `json:"profession" form:"profession" db:"profession"`
	NationalityId *int `json:"nationalityId" form:"nationalityId" db:"nationality_id"`
	UserId        int `json:"userId" form:"userId" db:"user_id"`
}
type JoinRegist struct {
    Id       int    `json:"id"`
    Email    string `json:"email" form:"email" db:"email"`
    Password string `json:"-" form:"password" db:"password"`
    Results  Profile
}

func CreateProfile(joinRegist JoinRegist) ( *Profile , error) {
   db := lib.DB()
    defer db.Close(context.Background())

		joinRegist.Password = lib.Encrypt(joinRegist.Password)

    var userId int
    err := db.QueryRow(
        context.Background(),
        `INSERT INTO "users" ("email", "password") VALUES ($1, $2) RETURNING "id"`,
        joinRegist.Email, joinRegist.Password,
    ).Scan(&userId)
    if err != nil {
        return nil, fmt.Errorf("failed to insert into users table: %v", err)
    }
    fmt.Println("-----")
    fmt.Println(err)

    var profile Profile
    err = db.QueryRow(
        context.Background(),
        `INSERT INTO "profile" ("picture", "full_name", "birth_date", "gender", "phone_number", "profession", "nationality_id", "user_id") 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, picture, full_name, birth_date, gender, phone_number, profession, nationality_id, user_id`,
        joinRegist.Results.Picture, joinRegist.Results.FullName, joinRegist.Results.Birthdate, joinRegist.Results.Gender,
        joinRegist.Results.PhoneNumber, joinRegist.Results.Profession, joinRegist.Results.NationalityId, userId,
    ).Scan(
        &profile.Id, &profile.Picture, &profile.FullName, &profile.Birthdate,
        &profile.Gender, &profile.PhoneNumber, &profile.Profession, &profile.NationalityId, &profile.UserId,
    )

    if err != nil {
        return nil, fmt.Errorf("failed to insert into profile table: %v", err)
    }

    return &profile, nil
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
