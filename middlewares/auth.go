package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
)
func tokenfiled(ctx *gin.Context){
	if e:=recover();e != nil {
		ctx.JSON(http.StatusBadRequest, lib.Message{
		Success:  false,
		Message: "Unauthorized",
	})
	ctx.Abort()
	}
}
func AuthMiddleware() gin.HandlerFunc{
	return func (ctx *gin.Context)  {
		defer tokenfiled(ctx)
		token := ctx.GetHeader("Authorization")[7:]
		isValidated, userId := lib.ValidateToken(token)
		if isValidated {
			ctx.Set("userId",userId)
			ctx.Next()
		}else{
			panic("Error: token invalid")
		}
	}
}