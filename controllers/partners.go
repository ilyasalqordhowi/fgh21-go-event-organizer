package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
)

func ListAllPartner(r *gin.Context) {
	results := models.FindAllPartner()
	r.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "List All Partner",
		Results: results,
	})
}