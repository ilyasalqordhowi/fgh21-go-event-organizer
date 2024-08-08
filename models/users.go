package models

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"-" form:"password" binding:"required,min=8"`
	Username string `json:"username" form:"username" binding:"required"`
}

// var data = []User{
// 	{
// 		Id:       1,
// 		Email:    "ilyas@mail.com",
// 		Password: "1234",
// 		Username:  "ilyas",
// 	},
// }

func FindAllUsers() []User {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
	context.Background(),
	`SELECT * FROM "users"`)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	fmt.Println(users)

	if err != nil {
		fmt.Println(err)
	}

	return users
	// dataUsers := data
	// return dataUsers
}

func FindOneUser(id int) User {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
	context.Background(),
	`SELECT * FROM "users" where id=$1`,
	)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	fmt.Println(users)

	if err != nil {
		fmt.Println(err)
	} 
	
	user := User{}
	for _, v := range users{
		if v.Id == id {
			user = v
		}
	}
	return user
	// idUser := User{}
	// for _, getId := range data {
	// 	return getId
	// }
	// return idUser
}
func FindOneUserByEmail(email string) User {
    db := lib.DB()
    defer db.Close(context.Background())

    rows, _ := db.Query(
        context.Background(),
         `select * from "users" where "email"=$1`,
        email,
    )

    users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

    if err != nil {
        fmt.Println(err)
    }

    user := User{}
    for _, i := range users {
        if i.Email == email{
            i = user
        }
    }
    return user
	
}

func Create(newUser User) User {
	db := lib.DB()
	defer db.Close(context.Background())

	newUser.Password = lib.Encrypt(newUser.Password)
	
sql := `insert into "users" ("email","password","username") values ($1,$2,$3) returning "id","email","password","username"`
row := db.QueryRow(context.Background(),sql,newUser.Email,newUser.Password,newUser.Username)
var results User
row.Scan(
	&results.Id,
	&results.Email,
	&results.Password,
	&results.Username,
)
return results
	// newUser.Id = len(data) + 1
	// data = append(data, newUser)
	// return data
}
func DeleteUsers(id int) error {
    db := lib.DB()
    defer db.Close(context.Background())

    commandTag, err := db.Exec(
        context.Background(),
        `DELETE FROM "users" WHERE id = $1`,
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
func EditUser(email string, username string, password string, id string) {
    db := lib.DB()
    defer db.Close(context.Background())

    dataSql := `update "users" set (email , username, password) = ($1, $2, $3) where id=$4`

    db.Exec(context.Background(), dataSql, email, username, password, id)

}
// func DeleteUsers(id int) error {
// 	for i, user := range data {
// 		if user.Id == id {
// 			data = append(data[:i], data[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("Not Found")
// }
// func UpdateUsers(id int, newData User) error {

// 	for i, user := range data {
// 		if user.Id == id {
// 			newData.Id = user.Id
// 			data[i] = newData
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("user not found")
// }
