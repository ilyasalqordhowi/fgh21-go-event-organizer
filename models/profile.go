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
	FullName      string `json:"full_name"  binding:"required" form:"full_name" db:"full_name"`
	Birthdate     *string `json:"birtdate" form:"birthdate" db:"birth_date"`
	Gender        *int    `json:"gender" form:"gender" db:"gender"`
	PhoneNumber   *string `json:"phoneNumber" form:"phone_number" db:"phone_number"`
	Profession    *string `json:"profession" form:"profession" db:"profession"`
	NationalityId *int `json:"nationalityId" form:"nationalityId" db:"nationality_id"`
	UserId        int `json:"userId" form:"userId" db:"user_id"`
}
type JoinRegist struct {
    Id       int    `json:"id"`
    Email    *string `json:"email"  binding:"required,email" form:"email" db:"email"`
    Password string `json:"-" form:"password"  binding:"required" db:"password"`
    Results  Profile
}
type Nationality struct{
    Id int `json:"id"`
    Name string `json:"nationalities" db:"name"`
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

   profile := Profile{
        UserId:  userId,
        FullName: joinRegist.Results.FullName,
    }
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
    fmt.Println(&profile,"masuk")
    return &profile, nil
}
func FindOneProfile(id int) Profile {
   db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "profile" where "user_id" = $1`, id,
	)
	profile, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[Profile])
	fmt.Println(err)
	if err != nil {

		fmt.Println(err)

	}
	fmt.Println(profile)

	return profile
}

func FindAllProfile() []Profile {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "profile" order by "id" asc`,
	)
	profile, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Profile])
	if err != nil {
		fmt.Println(err)
	}
	return profile
}
func EditProfile(data Profile, Id int) error{
	db := lib.DB()
    defer db.Close(context.Background())

    dataSql := `update "profile" set ("picture", "full_name", "birth_date", "gender", "phone_number", "profession", "nationality_id") = ($1, $2, $3, $4, $5, $6, $7) where "user_id" = $8`
    db.Exec(context.Background(), dataSql,data.Picture, data.FullName, data.Birthdate, data.Gender, data.PhoneNumber, data.Profession, data.NationalityId, Id)
	return nil
}
func FindAllNational() []Nationality {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "nationalities" order by "id" asc`,
	)
	national, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Nationality])
	if err != nil {
		fmt.Println(err)
	}
	return national
}
	func FindOneNational(id int) []Nationality {
		db := lib.DB()
		defer db.Close(context.Background())
	
		rows, _ := db.Query(context.Background(),
			`select * from "nationalities" where "id" = $1`,id,
		)
		nationality, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Nationality])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(nationality)
		return nationality
	}


func UpdateProfileImage(data Profile,id int) (Profile,error) {
		db := lib.DB()
		defer db.Close(context.Background())
	
		sql := `UPDATE profile SET "picture" = $1 WHERE user_id=$2 returning *`
	
		row, err := db.Query(context.Background(), sql, data.Picture, id)
		if err != nil {
			return Profile{}, nil
		}
	
		profile, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Profile])
		if err != nil {
			return Profile{}, nil
		}
	
		return profile, nil
	}