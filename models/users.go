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
	Password string `json:"-" form:"password" db:"password"`
	Username *string `json:"username" form:"username" db:"username"`
}
type ChangePassword struct{
	OldPassword string `form:"oldPassword" json:"oldPassword"`
	NewPassword string `form:"newPassword" json:"newPassword"`
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
	 	`select * from "users" where "email" = $1`,
		email,
	)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	if err != nil {
		fmt.Println(err)
	}

	user := User{}
	for _, val := range users {
		if val.Email == email{
			user = val
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

func UpdatePassword(password string, id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	dataPassword := lib.Encrypt(password)
	
	dataSql := `UPDATE "users" SET password = $1 WHERE id = $2`
	_, err := db.Exec(context.Background(), dataSql, dataPassword, id)
	fmt.Println(dataSql,"halo")
	if err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}
	return nil
}
func FindOneUserByPassword(id int) ChangePassword {
	db := lib.DB()
	defer db.Close(context.Background())
   
	var cp ChangePassword
	err := db.QueryRow(context.Background(), `select password from "users" where id = $1`, id).Scan(&cp.OldPassword)
	if err != nil {
		fmt.Println(err)
	}
   
	return cp
   }
   func UpdateUsername(dataUser User, Id int) error{
	db := lib.DB()
	defer db.Close(context.Background())

	dataSql := `update "users" set ("email","username") = ($1,$2) where id= $3`
	_,err := db.Exec(context.Background(), dataSql, dataUser.Email, dataUser.Username, Id)
if  err != nil {
	return fmt.Errorf("failed %v",err)
}
	return nil
}
   