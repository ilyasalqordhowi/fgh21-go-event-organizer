package controllers

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/dtos"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/repository"
)

func CreateProfile(ctx *gin.Context) {
	account := dtos.JoinRegist{}
	if err := ctx.ShouldBind(&account); err != nil {
		lib.HandlerBadRequest(ctx, err.Error())
		return
	}
	profile, err := repository.CreateProfile(account)
	if *account.Email == "" && account.Password == "" && profile.FullName == "" {
		lib.HandlerBadRequest(ctx, "Data bad request")
		return
	}
	if err != nil {
		lib.HandlerBadRequest(ctx, err.Error())
		return
	}
	lib.HandlerOk(ctx, "Register User success", nil, gin.H{
		"id":       profile.Id,
		"fullName": profile.FullName,
		"email":    account.Email,
	})
}
func ListAllProfile(r *gin.Context) {
	results := repository.FindAllProfile()
	lib.HandlerOk(r, "List All Profile", nil, results)
}
func DetailUsersProfile(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	data,_ := repository.FindOneProfile(id)
	dataUser := repository.FindOneUser(id)
	fmt.Println(data, "helo")

	lib.HandlerOk(ctx, "Profile Found", nil, gin.H{
		"profile": data,
		"user":    dataUser,
	})
	
}
func ListOneNational(r *gin.Context) {
	id,_ := strconv.Atoi(r.Param("id"))
	results := repository.FindOneNational(id)
	lib.HandlerOk(r, "Id National", nil, results)
}
func ListAllNational(r *gin.Context) {
	results := repository.FindAllNational()
	lib.HandlerOk(r, "List All National", nil, results)
}
func UpdateProfile(c *gin.Context) {
	id := c.GetInt("userId")
	var form dtos.Profile
	var user dtos.User

	err := c.Bind(&form)
	errUser := c.Bind(&user)
	data,_ := repository.FindOneProfile(id)
	dataProfile := repository.FindOneUser(id)

	if err != nil {
		lib.HandlerBadRequest(c, "Invalid input data")
		return
	}

	if errUser != nil {
		lib.HandlerBadRequest(c, "Failed user")
		return
	}

	repository.EditProfile(form, id)
	repository.UpdateUsername(user, id)

	lib.HandlerOk(c, "Profile Found", nil, gin.H{
		"profile": data,
		"user":    dataProfile,
	})
			}
			func UploadProfileImage(c *gin.Context) {
				id := c.GetInt("userId")
				fmt.Println(id)
			
				file, err := c.FormFile("image")
				if err != nil {
					lib.HandlerBadRequest(c, "no files uploaded")
					return
				}
			
				allowExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
				fileExt := strings.ToLower(filepath.Ext(file.Filename))
				if !allowExt[fileExt] {
					lib.HandlerBadRequest(c, "invalid file extension")
					return
				}
			
				image := uuid.New().String() + fileExt
			
				root := "./img/profile/"
				if err := c.SaveUploadedFile(file, root+image); err != nil {
					lib.HandlerBadRequest(c, "Upload image failed")
					return
				}
			
				img := "http://103.93.58.89:21213/image/profile/" + image
				result, err := repository.UpdateProfileImage(dtos.Profile{Picture: &img}, id)
			
				if err != nil {
					lib.HandlerBadRequest(c, "Update image failed")
					return
				}
			
				lib.HandlerOk(c, "Upload image success", nil, result)
			}
			