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

	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
)

func CreateProfile(ctx *gin.Context) {
    account := models.JoinRegist{}
    if err := ctx.ShouldBind(&account); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    profile, err := models.CreateProfile(account)
	if *account.Email == "" && account.Password == "" && profile.FullName == ""{
		  ctx.JSON(http.StatusBadRequest,
        lib.Message{
            Success: false,
            Message: "Data bad request",
            
        })
		return
	}
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK,
        lib.Message{
            Success: true,
            Message: "Register User success",
            Results: gin.H{
                "id":       profile.Id,
                "fullName": profile.FullName,
                "email":    account.Email,
            },
        })
}
func ListAllProfile(r *gin.Context) {
	results := models.FindAllProfile()
	r.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "List All Profile",
		Results: results,
	})
}
func DetailUsersProfile(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	data := models.FindOneProfile(id)
	dataUser := models.FindOneUser(id)
	fmt.Println(data,"helo")

	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Profile Found",
		Results: gin.H{
			"profile": data,
			"user":    dataUser,
		},
	})
	
}
func ListOneNational(r *gin.Context) {
	id,_ := strconv.Atoi(r.Param("id"))
	results := models.FindOneNational(id)
	r.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Id National",
		Results: results,
	})
}
func ListAllNational(r *gin.Context) {
	results := models.FindAllNational()
	r.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "List All National",
		Results: results,
	})
}
func UpdateProfile(c *gin.Context) {
	id := c.GetInt("userId")
	var form models.Profile
	var  user models.User
	err := c.Bind(&form)
	errUser := c.Bind(&user)
	data := models.FindOneProfile(id)
	dataProfile := models.FindOneUser(id)
	if err := c.ShouldBind(&form); err != nil {
        c.JSON(http.StatusBadRequest, lib.Message{
            Success: false,
            Message: "Invalid input data",
        })

        return
    }
    fmt.Println(errUser)
	if err != nil {
        c.JSON(http.StatusBadRequest, lib.Message{
            Success: false,
            Message: "Failed to update profile",
        })
        return
    }
		if errUser != nil {
			c.JSON(http.StatusBadRequest,
				lib.Message{
					Success: false,
					Message: "failed user",
				})
				return
			}

			models.EditProfile(form,id)
			models.UpdateUsername(user,id)
			c.JSON(http.StatusOK, lib.Message{
				Success: true,
				Message: "Profile Found",
				Results: gin.H{
					"profile": data,
					"user":    dataProfile,
				},
			})
				
			}
func UploadProfileImage(c *gin.Context) {
				id := c.GetInt("userId")
			
				maxFile := 100 * 1024 * 1024
				c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(maxFile))
	
				file, err := c.FormFile("profileImg")
				

		
				if err != nil {
					if err.Error() == "http: request body too large" {
						c.JSON(http.StatusRequestEntityTooLarge, lib.Message{
							Success: false,
							Message: "file size too large, max capacity 100 mb ",
							
						})
						return
					}
					c.JSON(http.StatusBadRequest,lib.Message{	
						Success: false,
						Message: "not file to upload",
					})
					return
				}
				if id == 0 {
					c.JSON(http.StatusBadRequest ,lib.Message{
						Success: false,
						Message: "User not found",
					})
					return
				}
			
				allowExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
				fileExt := strings.ToLower(filepath.Ext(file.Filename))
				if !allowExt[fileExt] {
					c.JSON(http.StatusBadRequest ,lib.Message{
						Success: false,
						Message: "extension file not validate",
					})
					return
				}
			
				newFile := uuid.New().String() + fileExt
			
				uploadDir := "./img/profile/"
				if err := c.SaveUploadedFile(file, uploadDir+newFile); err != nil {
					c.JSON(http.StatusBadRequest , lib.Message{
						Success: false,
						Message: "upload failed",
					})
					return
				}
			
				dataImg := "http://localhost:8888/img/profile/" + newFile
			
				delImgBefore := models.FindOneProfile(id)
				
				if delImgBefore.Picture != nil {
				   fileDel := strings.Split(*delImgBefore.Picture, "8000")[1]
				   fmt.Println("file :", fileDel)
				   os.Remove("." + fileDel)
				  
				}
			
				profile, err := models.UpdateProfileImage(models.Profile{Picture: &dataImg}, id)
				if err != nil {
				c.JSON(http.StatusBadRequest , lib.Message{
					Success: false,
					Message: "upload failed",
				})
					return
				}
			
				c.JSON(http.StatusOK ,lib.Message{
					Success: true,
					Message: "Upload success",
					Results: profile,
				})
			}
		