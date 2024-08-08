package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
)

func DetailUsersProfile(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	Profile := models.FindOneProfile(id)

	if Profile.Id == id {
		ctx.JSON(http.StatusOK, lib.Message{
			Success: true,
			Message: "Profile OK",
			Results: Profile,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Message{
			Success: false,
			Message: "Profile Not Found",
		})
	}

}