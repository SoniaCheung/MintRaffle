package responses

import "soniacheung/mint-raffle/cmd/mint-raffle/app/models"

type GetProjectsRsp struct {
	Projects []models.Project `json:"projects"`
}

type GetProjectByIdRsp struct {
	Project models.Project `json:"project"`
}

type PostProjectRsp struct {
	Project models.Project `json:"project"`
}
