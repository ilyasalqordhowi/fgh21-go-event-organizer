package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/controllers"
)

func EventRouter(r *gin.RouterGroup){
	r.POST("/event",controllers.Event)
}