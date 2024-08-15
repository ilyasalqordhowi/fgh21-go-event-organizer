package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/controllers"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/middlewares"
)

func EventRouter(r *gin.RouterGroup){
	r.Use(middlewares.AuthMiddleware())
	r.POST("",controllers.CreateEvent)
    r.GET("", controllers.ListAllEvent)
	r.GET("/:id",controllers.DetailEvent)
	r.PATCH("/:id",controllers.UpdateEvent)
	r.DELETE("/:id",controllers.DeleteEvent)
	r.GET("/section/:id", controllers.DetailEventSections)
	r.GET("/payment_method", controllers.ListAllPaymentMethod)
}