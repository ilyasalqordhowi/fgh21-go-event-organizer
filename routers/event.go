package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/controllers"
)

func EventRouter(r *gin.RouterGroup){
	r.GET("", controllers.ListAllEvent)
	r.GET("/:id",controllers.DetailEvent)
	r.POST("",controllers.CreateEvent)
	r.PATCH("/:id",controllers.UpdateEvent)
	r.DELETE("/:id",controllers.DeleteEvent)
	r.GET("/section/:id", controllers.DetailEventSections)
	r.GET("/payment_method", controllers.ListAllPaymentMethod)
}