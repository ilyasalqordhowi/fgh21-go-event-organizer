package controllers

import (
	"fmt"
	"net/http"
	"os"
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
				maxFile := 100 * 1024 * 1024
				c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(maxFile))
			
				file, err := c.FormFile("profileImg")
			
				if err != nil {
					if err.Error() == "http: request body too large" {
						lib.HandlerMaxFile(c, "file size too large, max capacity 100 mb")
						return
					}
					lib.HandlerBadRequest(c, "not file to upload")
					return
				}
				if id == 0 {
					lib.HandlerBadRequest(c, "User not found")
					return
				}
			
				allowExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
				fileExt := strings.ToLower(filepath.Ext(file.Filename))
				if !allowExt[fileExt] {
					lib.HandlerBadRequest(c, "extension file not validate")
					return
				}
			
				newFile := uuid.New().String() + fileExt
				uploadDir := "./img/profile/"
				if err := c.SaveUploadedFile(file, uploadDir+newFile); err != nil {
					lib.HandlerBadRequest(c, "upload failed")
					return
				}
			
				dataImg := "/img/profile/" + newFile
				delImgBefore := repository.FindOneProfile(id)
			
				if delImgBefore.Picture != nil {
					fileDel := strings.Split(*delImgBefore.Picture, "8000")[1]
					fmt.Println("file :", fileDel)
					os.Remove("." + fileDel)
				}
			
				profile, err := repository.UpdateProfileImage(dtos.Profile{Picture: &dataImg}, id)
				if err != nil {
					lib.HandlerBadRequest(c, "upload failed")
					return
				}
			
				lib.HandlerOk(c, "Upload success", nil, profile)
			}
		