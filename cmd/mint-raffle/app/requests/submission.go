package requests

type GetSubmissionsReq struct {
	ProjectId     int    `form:"project_id"`
	WalletAddress string `form:"wallet_address"`
}

type PostSubmissionReq struct {
	ProjectId     int    `form:"project_id"`
	WalletAddress string `form:"wallet_address"`
}
