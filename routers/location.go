package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/controllers"
)

func LocationsRouter(r *gin.RouterGroup){
	r.GET("", controllers.ListAllLocations)
}