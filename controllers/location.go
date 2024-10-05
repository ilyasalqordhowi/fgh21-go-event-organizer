package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/repository"
)

func ListAllLocations(ctx *gin.Context) {
	search := ctx.Query("search")
	limit,_ := strconv.Atoi(ctx.Query("limit"))
	page,_ := strconv.Atoi(ctx.Query("page"))

	if limit < 1 {
		limit = 7
	}

	if page < 1 {
		page = 1
	}

	results := repository.FindAllLocations(search, limit, page)

	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "List All Locations",
		Results: results,
	})
}