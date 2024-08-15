package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/controllers"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/middlewares"
)

func TransactionRouter(r *gin.RouterGroup) {
	r.Use(middlewares.AuthMiddleware())
	r.POST("", controllers.CreateTransaction)
}