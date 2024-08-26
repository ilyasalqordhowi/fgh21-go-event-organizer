package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/controllers"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/middlewares"
)

func ProfileRouter(r *gin.RouterGroup){
	// r.GET("", controllers.ListAllProfile)
	r.GET("/national/:id", controllers.ListOneNational)
	r.GET("/national", controllers.ListAllNational)
	r.Use(middlewares.AuthMiddleware())
	r.GET("/",controllers.DetailusersProfile)
	r.PATCH("/update",controllers.UpdateProfile)
}