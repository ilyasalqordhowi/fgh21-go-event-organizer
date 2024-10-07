package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/controllers"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/middlewares"
)

func EventRouter(r *gin.RouterGroup){
	r.GET("", controllers.ListAllEvent)
	r.PATCH("/:id",controllers.UpdateEvent)
	r.DELETE("/:id",controllers.DeleteEvent)
	r.GET("/:id",controllers.DetailEvent)
	r.POST("/img",controllers.UploadImage)
	r.Use(middlewares.AuthMiddleware())
	r.POST("",controllers.CreateEvent)
	r.GET("/data",controllers.DetailCreateEvent)
	r.GET("/section/:id", controllers.DetailEventSections)
	r.GET("/payment_method", controllers.ListAllPaymentMethod)
	r.GET("/category/:id", controllers.FindEventsByCategory)
}