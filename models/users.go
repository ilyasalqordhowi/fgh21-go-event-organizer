package models

import (
	"context"
	"fmt"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" binding:"required,email" db:"email"`
	Password string `json:"-" form:"password" binding:"required,min=8" db:"password"`
	Username *string `json:"username" form:"username" binding:"required" db:"username"`
}


func FindAllUsers(search string ,page int,limit int) ([]User,int) {
	db := lib.DB()
	defer db.Close(context.Background())
	offset := (page - 1) * limit
	
	sql :=	`SELECT * FROM "users" where "email" ilike '%' || $1 || '%' offset $2 limit $3`
	rows, _ := db.Query(context.Background(),sql,search,offset,limit)
	
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	fmt.Println(users)
	
	if err != nil {
		fmt.Println(err)
	}
	result := TotalUsers(search)
	return users,result
}
func TotalUsers(search string)int{
	db := lib.DB()
	defer db.Close(context.Background())

	sql :=	`SELECT count(id) as "total" FROM "users" where "email" ilike '%' || $1 || '%'`
	rows := db.QueryRow(context.Background(),sql,search)
	var results int
	rows.Scan(
		&results,
	)
	return results
}
func FindOneUser(id int) User {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
		context.Background(),
	`SELECT * FROM "users"`,
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
}
func FindOneUserByEmail(email string) User {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
	context.Background(),
	`SELECT * FROM "users"`,
	)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	fmt.Println(users)

	if err != nil {
		fmt.Println(err)
	} 
	
	user := User{}
	for _, v := range users{
		if v.Email == email {
			user = v
		}
	}
	fmt.Println(user)
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
