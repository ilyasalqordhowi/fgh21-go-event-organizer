package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/controllers"
)

func AuthRouter(r *gin.RouterGroup){
	r.POST("/login",controllers.AuthLogin)
	r.POST("/register",controllers.CreateProfile)
}