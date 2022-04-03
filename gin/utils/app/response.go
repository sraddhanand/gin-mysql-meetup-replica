package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type RtResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(statusCode int, data interface{}) {
	g.C.JSON(statusCode, RtResp{
		Code: statusCode,
		Msg:  http.StatusText(statusCode),
		Data: data,
	})
}

func (g *Gin) RespMsg(statusCode int, statusMsg string, data interface{}) {
	g.C.JSON(statusCode, RtResp{
		Code: statusCode,
		Msg:  statusMsg,
		Data: data,
	})
}
