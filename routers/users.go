package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/controllers"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/middlewares"
)

func UserRouter(routerGroup *gin.RouterGroup) {
	routerGroup.POST("",controllers.CreateUsers)
	routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.GET("", controllers.ListAllUsers)
	routerGroup.GET("/:id",controllers.DetailUsers)
	// routerGroup.PATCH("/:id",controllers.UpdateUsers)
	routerGroup.DELETE("/:id",controllers.DeleteUsers)
}	