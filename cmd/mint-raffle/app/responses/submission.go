package responses

import "soniacheung/mint-raffle/cmd/mint-raffle/app/models"

type PostSubmissionRsp struct {
	Submission models.Submission `json:"submission"`
}
