package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
)

func Event(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	Event := models.FindOneEvent(id)

	if Event.Id == id {
		ctx.JSON(http.StatusOK, lib.Message{
			Success: true,
			Message: "Event OK",
			Results: Event,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Message{
			Success: false,
			Message: "Event Not Found",
		})
	}

}