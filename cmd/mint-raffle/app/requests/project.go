package requests

type GetProjectsReq struct {
	Status string `form:"status"`
}

type GetProjectByIdReq struct {
	Id int `uri:"id"`
}
