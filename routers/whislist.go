package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/controllers"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/middlewares"
)

func WhislistRouter(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.GET("/",controllers.ListAllWishlist)
	routerGroup.GET("/:id",controllers.ListOneWishlist)
	routerGroup.POST("/:id",controllers.Createwishlist)
	routerGroup.DELETE("/:id",controllers.DeleteUsers)

}	