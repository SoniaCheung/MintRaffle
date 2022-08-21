package responses

import "soniacheung/mint-raffle/cmd/mint-raffle/app/models"

type GetSubmissionsRsp struct {
	Submission models.Submission `json:"submission"`
}

type PostSubmissionRsp struct {
	Submission models.Submission `json:"submission"`
}
