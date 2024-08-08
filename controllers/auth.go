package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
)

type Token struct{
	JWToken string `json:"token"`
}
func AuthLogin(ctx *gin.Context) {
    var user models.User
    ctx.Bind(&user)

    found := models.FindOneUserByEmail(user.Email)
	fmt.Println(found,"tes")

    // if found == (models.User{}) {
    //     ctx.JSON(http.StatusUnauthorized,
    //         lib.Message{
    //             Success: false,
    //             Message: "Wrong Email or Password",
    //         })
    //     return
    // }

    isVerified := lib.Verify(user.Password, found.Password)
	// fmt.Println(isVerified)

    if isVerified {
        JWToken := lib.GenerateUserIdToken(found.Id)
        ctx.JSON(http.StatusOK,
            lib.Message{
                Success: true,
                Message: "Login success",
                Results: Token{JWToken},
            })
    } else {
		JWToken := lib.GenerateUserIdToken(found.Id)
        ctx.JSON(http.StatusUnauthorized,
            lib.Message{
                Success: false,
                Message: "Wrong Email or Password",
				Results: Token{JWToken},
            })
    }
}