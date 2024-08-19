package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	if account.Password == "" {
		  ctx.JSON(http.StatusBadRequest,
        lib.Message{
            Success: false,
            Message: "Password harus diisi",
            
        })
		return
	}
	if *account.Email == "" && account.Password == "" && profile.FullName == ""{
		  ctx.JSON(http.StatusBadRequest,
        lib.Message{
            Success: false,
            Message: "data harus diisi",
            
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
func DetailusersProfile(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	data := models.FindOneProfile(id)
	dataProfile := models.FindOneUser(id)
	fmt.Println(data,"helo")

	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Profile Found",
		Results: gin.H{
			"profile": data,
			"user":    dataProfile,
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