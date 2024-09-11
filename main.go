package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/routers"
)

func main() {
	r := gin.Default()
	r.Static("/img/profile", "./img/profile")
	
    corsConfig := cors.DefaultConfig()
    corsConfig.AllowAllOrigins = true
    corsConfig.AllowHeaders = []string{
        "Origin", "Content-Type", "Authorization", "Content-Length",
    }
    r.Use(cors.New(corsConfig))
	routers.RouterCombine(r)
	r.Run("0.0.0.0:8888")
}