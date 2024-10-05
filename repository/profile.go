package repository

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/dtos"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"

	"github.com/jackc/pgx/v5"
)



func CreateProfile(joinRegist dtos.JoinRegist) ( *dtos.Profile , error) {
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

   profile := dtos.Profile{
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
func FindOneProfile(id int) dtos.Profile {
   db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "profile" where "user_id" = $1`, id,
	)
	profile, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[dtos.Profile])
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(profile)

	return profile
}

func FindAllProfile() []dtos.Profile {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "profile" order by "id" asc`,
	)
	profile, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dtos.Profile])
	if err != nil {
		fmt.Println(err)
	}
	return profile
}
func EditProfile(data dtos.Profile, Id int) error{
	db := lib.DB()
    defer db.Close(context.Background())

    dataSql := `update "profile" set ("picture", "full_name", "birth_date", "gender", "phone_number", "profession", "nationality_id") = ($1, $2, $3, $4, $5, $6, $7) where "user_id" = $8`
    db.Exec(context.Background(), dataSql,data.Picture, data.FullName, data.Birthdate, data.Gender, data.PhoneNumber, data.Profession, data.NationalityId, Id)
	return nil
}
func FindAllNational() []models.Nationality {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "nationalities" order by "id" asc`,
	)
	national, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Nationality])
	if err != nil {
		fmt.Println(err)
	}
	return national
}
	func FindOneNational(id int) []models.Nationality {
		db := lib.DB()
		defer db.Close(context.Background())
	
		rows, _ := db.Query(context.Background(),
			`select * from "nationalities" where "id" = $1`,id,
		)
		nationality, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Nationality])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(nationality)
		return nationality
	}


func UpdateProfileImage(data dtos.Profile,id int) (dtos.Profile,error) {
		db := lib.DB()
		defer db.Close(context.Background())
	
		sql := `UPDATE profile SET "picture" = $1 WHERE user_id=$2 returning *`
		
		row, err := db.Query(context.Background(), sql, data.Picture, id)
		fmt.Println(row ,"ini modulnya")
		if err != nil {
			return dtos.Profile{}, nil
		}
	
		profile, err := pgx.CollectOneRow(row, pgx.RowToStructByName[dtos.Profile])
		if err != nil {
			return dtos.Profile{}, nil
		}
		return profile, nil
	}