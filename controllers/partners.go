package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/repository"
)

func ListAllPartner(c *gin.Context) {
	results := repository.FindAllPartner()
	lib.HandlerOk(c, "List All Partners",nil, results)
}