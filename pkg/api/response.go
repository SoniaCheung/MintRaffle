package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

func ResponseJSON(ctx *gin.Context, err error, data interface{}) {
	if err == nil {
		ctx.JSON(http.StatusOK, &Response{
			Data: data,
		})
	} else {
		logrus.Error("%s response error: %s", ctx.Request.URL, err)
		ctx.JSON(http.StatusBadRequest, &Response{
			Error: fmt.Sprintf("Internal Server Error: %s", err),
		})
	}
}
