package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/controllers"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/middlewares"
)

func UserRouter(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.POST("",controllers.CreateUsers)
	routerGroup.GET("", controllers.ListAllUsers)
	routerGroup.GET("/:id",controllers.DetailUsers)
	routerGroup.PATCH("/:id",controllers.Update)
	routerGroup.DELETE("/:id",controllers.DeleteUsers)
}	