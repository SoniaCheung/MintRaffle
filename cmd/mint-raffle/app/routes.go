package app

import "github.com/gin-gonic/gin"

func NewRouters() (*gin.Engine, error) {
	router := gin.Default()
	router.Use(gin.Recovery())

	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, nil)
	})

	return router, nil
}
