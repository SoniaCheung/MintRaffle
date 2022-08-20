package app

import (
	projectcontroller "soniacheung/mint-raffle/cmd/mint-raffle/app/controllers/project_controller"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

func NewRouters(engine *xorm.Engine) (*gin.Engine, error) {
	router := gin.Default()
	router.Use(gin.Recovery())

	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, nil)
	})

	// Project related routes
	projectController := projectcontroller.NewProjectController(engine)

	router.GET("/projects", projectController.GetProjects)
	router.GET("/projects/:id", projectController.GetProjectById)
	return router, nil
}
