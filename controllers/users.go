package controllers

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/dtos"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/repository"
)
func ListAllUsers(c *gin.Context){
	search := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if page > 1 {
		page = (page - 1) * limit
	}

	listUsers, count := repository.FindAllUsers(search, page, limit)
	totalPage := math.Ceil(float64(count) / float64(limit))
	next := 0
	prev := 0

	if int(totalPage) > 1 {
		next = int(totalPage) - page
	}
	if int(totalPage) > 1 {
		prev = int(totalPage) - 1
	}

	totalInfo := lib.TotalInfo{
		TotalData: count,
		TotalPage: int(totalPage),
		Page:      page,
		Limit:     limit,
		Next:      next,
		Prev:      prev,
	}
	lib.HandlerOk(c, "success", totalInfo, listUsers)
	}
func DetailUsers(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := repository.FindOneUser(id)
	fmt.Println(id)

	if data.Id == id {
		lib.HandlerOk(c, "Users Found", nil, data)
		return
	}
	lib.HandlerNotFound(c, "Users Not Found")
}
func CreateUsers(c *gin.Context) {
	newUser := dtos.User{}

	if err := c.ShouldBind(&newUser); err != nil {
		lib.HandlerBadRequest(c, "Invalid input data")
		return
	}

	addUser := repository.Create(newUser)
	fmt.Println(addUser)
	lib.HandlerOk(c, "User created successfully", nil, addUser)
}

func DeleteUsers(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	dataUser := repository.FindOneUser(id)

	if err != nil {
		lib.HandlerBadRequest(c, "Invalid user ID")
		return
	}

	err = repository.DeleteUsers(id)
	if err != nil {
		lib.HandlerBadRequest(c, "Id Not Found")
		return
	}

	lib.HandlerOk(c, "User deleted successfully", nil, dataUser)
}

func Update(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	search := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if page > 1 {
		page = (page - 1) * limit
	}

	data, count := repository.FindAllUsers(search, page, limit)
	totalPage := math.Ceil(float64(count) / float64(limit))
	next := 0
	prev := 0

	if int(totalPage) > 1 {
		next = int(totalPage) - page
	}
	if int(totalPage) > 1 {
		prev = int(totalPage) - 1
	}

	totalInfo := lib.TotalInfo{
		TotalData: count,
		TotalPage: int(totalPage),
		Page:      page,
		Limit:     limit,
		Next:      next,
		Prev:      prev,
	}

	user := dtos.User{}
	err := c.Bind(&user)
	if err != nil {
		lib.HandlerBadRequest(c, "Invalid input data")
		return
	}

	result := dtos.User{}
	for _, v := range data {
		if v.Id == id {
			result = v
		}
	}

	if result.Id == 0 {
		lib.HandlerNotFound(c, "User with id "+param+" not found")
		return
	}

	idUser := 0
	for _, v := range data {
		idUser = v.Id
	}
	user.Id = idUser

	repository.EditUser(user.Email, user.Password, *user.Username, param)

	lib.HandlerOk(c, "User id "+param+" edit Success", totalInfo, user)
}
func UpdatePassword(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	var form dtos.ChangePassword
	err := ctx.Bind(&form)
	if err != nil {
		lib.HandlerBadRequest(ctx, "Invalid input data")
		return
	}
	found := repository.FindOneUserByPassword(id)

	isVerified := lib.Verify(form.OldPassword, found.OldPassword)
	fmt.Println(isVerified)
	if isVerified {
		err := repository.UpdatePassword(form.NewPassword, id)
		if err != nil {
			lib.HandlerBadRequest(ctx, "Failed to update password")
		} else {
			lib.HandlerOk(ctx, "Update password success", nil, nil)
		}
	} else {
		lib.HandlerBadRequest(ctx, "Wrong Password")
	}
	}