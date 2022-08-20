package app

import (
	projectcontroller "soniacheung/mint-raffle/cmd/mint-raffle/app/controllers/project_controller"
	submissioncontroller "soniacheung/mint-raffle/cmd/mint-raffle/app/controllers/submission_controller"

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
	router.POST("/projects", projectController.PostProject)

	// Submission related routes
	submissionController := submissioncontroller.NewSubmissionController(engine)
	router.GET("/submissions", submissionController.GetSubmissions)
	router.POST("/submissions", submissionController.PostSubmission)

	return router, nil
}
