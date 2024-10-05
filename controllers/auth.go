package controllers

import (
	"fmt"

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
        lib.HandlerUnauthorized(ctx, "Wrong Email")
		return
    }

    isVerified := lib.Verify(user.Password, found.Password)

    if isVerified {
        JWToken := lib.GenerateUserIdToken(found.Id)
		lib.HandlerOk(ctx, "Login Success", nil, Token{JWToken})
    } else {
        lib.HandlerUnauthorized(ctx, "Wrong Password")
		
    }
}
