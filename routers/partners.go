package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/controllers"
)

func PartnersRouter(r *gin.RouterGroup){
    r.GET("", controllers.ListAllPartner)

}