package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/dtos"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/repository"
)

type Token struct{
	JWToken string `json:"token"`
}
func AuthLogin(ctx *gin.Context) {
    var user dtos.User
    ctx.Bind(&user)
    found := repository.FindOneUserByEmail(user.Email)
    fmt.Println(found)
    if found == (dtos.User{}) {
        ctx.JSON(http.StatusUnauthorized,
            lib.Message{
                Success: false,
                Message: "Wrong Email",
            })
        return
    }

    isVerified := lib.Verify(user.Password, found.Password)

    if isVerified {
        JWToken := lib.GenerateUserIdToken(found.Id)
        ctx.JSON(http.StatusOK,
            lib.Message{
                Success: true,
                Message: "Login success",
                Results: Token{JWToken},
            })
    } else {
        ctx.JSON(http.StatusUnauthorized,
            lib.Message{
                Success: false,
                Message: "Wrong Password",
            })
    }
}
