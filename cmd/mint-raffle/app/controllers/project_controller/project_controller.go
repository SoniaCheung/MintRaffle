package projectcontroller

import (
	"errors"
	"soniacheung/mint-raffle/cmd/mint-raffle/app/models"
	"soniacheung/mint-raffle/cmd/mint-raffle/app/requests"
	"soniacheung/mint-raffle/cmd/mint-raffle/app/responses"
	"soniacheung/mint-raffle/pkg/api"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

type ProjectController struct {
	Engine *xorm.Engine
}

func NewProjectController(engine *xorm.Engine) *ProjectController {
	return &ProjectController{
		Engine: engine,
	}
}

func (p *ProjectController) GetProjects(ctx *gin.Context) {
	req := requests.GetProjectsReq{}
	rsp := responses.GetProjectsRsp{}
	var err error

	if err = ctx.BindQuery(&req); err != nil {
		api.ResponseJSON(ctx, err, nil)
		return
	}

	session := p.Engine.NewSession()
	defer session.Close()

	var project models.Project
	projects := make([]models.Project, 0)

	if req.Status == "" {
		if err = session.Table(project.TableName()).Desc("due_time").Find(&projects); err != nil {
			api.ResponseJSON(ctx, err, nil)
			return
		}
	} else {
		if req.Status == string(models.ProjectStatusOpening) ||
			req.Status == string(models.ProjectStatusPending) ||
			req.Status == string(models.ProjectStatusClosed) {
			if err = session.Table(project.TableName()).Where("status = ?", req.Status).Desc("due_time").Find(&projects); err != nil {
				api.ResponseJSON(ctx, err, nil)
				return
			}
		} else {
			api.ResponseJSON(ctx, errors.New("invalid status"), nil)
			return
		}
	}

	rsp.Projects = projects
	api.ResponseJSON(ctx, nil, rsp)
}

func (p *ProjectController) GetProjectById(ctx *gin.Context) {
	req := requests.GetProjectByIdReq{}
	rsp := responses.GetProjectByIdRsp{}
	var err error

	if err = ctx.BindUri(&req); err != nil {
		api.ResponseJSON(ctx, err, nil)
		return
	}

	session := p.Engine.NewSession()
	defer session.Close()

	var project models.Project
	ok, err := session.Table(project.TableName()).ID(req.Id).Get(&project)
	if !ok {
		api.ResponseJSON(ctx, errors.New("cannot find project"), nil)
		return
	}

	if err != nil {
		api.ResponseJSON(ctx, err, nil)
		return
	}

	rsp.Project = project
	api.ResponseJSON(ctx, nil, rsp)
}

func (p *ProjectController) PostProject(ctx *gin.Context) {
	req := requests.PostProjectReq{}
	rsp := responses.PostProjectRsp{}
	var err error

	if err = ctx.Bind(&req); err != nil {
		api.ResponseJSON(ctx, err, nil)
		return
	}

	session := p.Engine.NewSession()
	defer session.Close()

	project := models.Project{
		Name:        req.Name,
		Description: req.Description,
		OfficalLink: req.OfficalLink,
		MaxWinner:   req.MaxWinner,
		DueTime:     req.DueTime,
		Status:      models.ProjectStatusOpening,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if _, err = session.Table(project.TableName()).Insert(&project); err != nil {
		api.ResponseJSON(ctx, err, nil)
		return
	}

	rsp.Project = project
	api.ResponseJSON(ctx, nil, rsp)
}
