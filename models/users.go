package models

import (
	"fmt"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"-" form:"password" binding:"required,min=8"`
}

var data = []User{
	{
		Id:       1,
		Name:     "ilyas",
		Email:    "ilyas@mail.com",
		Password: "1234",
	},
}

func FindAllUsers() []User {
	dataUsers := data
	return dataUsers
}

func FindOneUser(id int) User {
	idUser := User{}
	for _, getId := range data {
		return getId
	}
	return idUser
}
func Create(newUser User) []User {
	newUser.Id = len(data) + 1
	data = append(data, newUser)
	return data
}
func DeleteUsers(id int) error {
	for i, user := range data {
		if user.Id == id {
			data = append(data[:i], data[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Not Found")
}
func UpdateUsers(id int, newData User) error {
	for i, user := range data {
		if user.Id == id {
			newData.Id = user.Id
			data[i] = newData
			return nil
		}
	}
	return fmt.Errorf("user not found")
}
