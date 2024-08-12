package controllers

import (
	"net/http"

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
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK,
        lib.Message{
            Success: true,
            Message: "Register User success",
            Results: profile,
        })
}