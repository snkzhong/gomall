package libs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 0
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse() *Response {
	return &Response{}
}

func SendResult(response *Response, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response)
}

func (r *Response) Ok(ctx *gin.Context) {
	r.OkFull("", map[string]interface{}{}, ctx)
}

func (r *Response) OkWithMsg(msg string, ctx *gin.Context) {
	r.OkFull(msg, map[string]interface{}{}, ctx)
}

func (r *Response) OkWithData(data interface{}, ctx *gin.Context) {
	r.OkFull("", data, ctx)
}

func (r *Response) OkFull(msg string, data interface{}, ctx *gin.Context) {
	r.Code = SUCCESS
	r.Msg = msg
	r.Data = data
	SendResult(r, ctx)
}
