package router

import (
	"gin_demo/handlers"
	"gin_demo/middleware"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	r.Use(middleware.Cors())
	v1 := r.Group("/v1")
	{
		v1.GET("/onchain", handlers.PutOnChain)
	}
	r.Run(":8000")
}