package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/repository"
)

func ListAllPartner(r *gin.Context) {
	results := repository.FindAllPartner()
	r.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "List All Partner",
		Results: results,
	})
}