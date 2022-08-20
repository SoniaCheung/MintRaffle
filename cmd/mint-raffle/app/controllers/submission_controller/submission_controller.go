package submissioncontroller

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

type SubmissionController struct {
	Engine *xorm.Engine
}

func NewSubmissionController(engine *xorm.Engine) *SubmissionController {
	return &SubmissionController{
		Engine: engine,
	}
}

func (s *SubmissionController) PostSubmission(ctx *gin.Context) {
	req := requests.PostSubmissionReq{}
	rsp := responses.PostSubmissionRsp{}
	var err error

	if err = ctx.Bind(&req); err != nil {
		api.ResponseJSON(ctx, err, nil)
		return
	}

	session := s.Engine.NewSession()
	defer session.Close()

	var project models.Project
	ok, err := session.Table(project.TableName()).ID(req.ProjectId).Get(&project)
	if !ok {
		api.ResponseJSON(ctx, errors.New("project does not exists"), nil)
		return
	}
	if err != nil {
		api.ResponseJSON(ctx, err, nil)
		return
	}

	submission := models.Submission{
		ProjectId:     req.ProjectId,
		WalletAddress: req.WalletAddress,
		Winner:        false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	var existingSubmission models.Submission
	exist, err := session.Table(submission.TableName()).Where("project_id = ? AND wallet_address = ?", req.ProjectId, req.WalletAddress).Exist(&existingSubmission)
	if exist {
		api.ResponseJSON(ctx, errors.New("wallet address already submitted"), nil)
		return
	}
	if err != nil {
		api.ResponseJSON(ctx, err, nil)
		return
	}

	if _, err = session.Table(submission.TableName()).Insert(&submission); err != nil {
		api.ResponseJSON(ctx, err, nil)
		return
	}

	rsp.Submission = submission
	api.ResponseJSON(ctx, nil, rsp)
}
