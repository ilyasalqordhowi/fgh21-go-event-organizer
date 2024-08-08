package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/routers"
)

func main() {
	r := gin.Default()

	routers.RouterCombine(r)

	r.Run("localhost:8888")
}
