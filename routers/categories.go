package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/controllers"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/middlewares"
)

func CategoriesRouter(r *gin.RouterGroup){
	r.Use(middlewares.AuthMiddleware())
	r.POST("",controllers.CreateCategory)
    r.GET("", controllers.ListAllCategory)
	r.GET("/:id",controllers.DetailCategory)
	r.PATCH("/:id",controllers.UpdateCategory)
	r.DELETE("/:id",controllers.DeleteCategory)
}