package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/controllers"
)

func UserRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", controllers.ListAllUsers)
	routerGroup.GET("/:id",controllers.DetailUsers)
	routerGroup.POST("/",controllers.CreateUsers)
	routerGroup.PATCH("/:id",controllers.UpdateUsers)
	routerGroup.DELETE("/:id",controllers.DeleteUsers)
}