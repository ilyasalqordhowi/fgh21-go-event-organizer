package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

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
	data := repository.FindOneProfile(id)
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
	data := repository.FindOneProfile(id)
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
				if id == 0 {
					lib.HandlerBadRequest(c, "User not found")
					return
				}
			
			
				if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
					lib.HandlerBadRequest(c, "error parsing form")
					return
				}
			
				
				file, err := c.FormFile("profileImg")
				if err != nil {
					fmt.Println("Error while getting file:", err)
					lib.HandlerBadRequest(c, "no file to upload")
					return
				}
			
			
				if err := c.SaveUploadedFile(file, "./img/profile/"+file.Filename); err != nil {
					lib.HandlerBadRequest(c, "upload failed")
					return
				}
			
				lib.HandlerOk(c, "Upload success", nil, nil)
			}
			
		