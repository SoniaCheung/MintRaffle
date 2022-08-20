package requests

import "time"

type GetProjectsReq struct {
	Status string `form:"status"`
}

type GetProjectByIdReq struct {
	Id int `uri:"id"`
}

type PostProjectReq struct {
	Name        string    `form:"name"`
	Description string    `form:"description"`
	OfficalLink string    `form:"offical_link"`
	MaxWinner   int       `form:"max_winner"`
	DueTime     time.Time `form:"due_time"`
}
