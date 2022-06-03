package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RespCode 响应码
type RespCode int

type Exception struct {
	HttpStatus int      `json:"-"`
	Code       RespCode `json:"code,omitempty"`
	Message    string   `json:"message,omitempty"`
}

func NewException(code RespCode, message ...string) *Exception {

	if len(message) > 0 {
		return &Exception{
			Code:    code,
			Message: message[0],
		}
	}

	return &Exception{
		Code: code,
	}
}

const (
	Error       RespCode = iota
	Success     RespCode = http.StatusOK
	BadRequest  RespCode = http.StatusBadRequest
	ServerError RespCode = http.StatusInternalServerError
)

const (
	jCode       string = "code"
	jMsg        string = "msg"
	jData       string = "data"
	jMsgSuccess string = "success"
)

// genGinH 生成响应数据结构
func genGinH(code RespCode, message string, data ...any) gin.H {
	if len(data) > 0 {
		return gin.H{
			jCode: code,
			jMsg:  message,
			jData: data[0],
		}
	}
	return gin.H{
		jCode: code,
		jMsg:  message,
	}
}

type response struct {
	request *gin.Context
}

// New 生成句柄
func New(c *gin.Context) *response {
	return &response{request: c}
}

// ResponseOK 正常响应
func (r *response) ResponseOK(data ...any) {
	if len(data) > 0 {
		r.request.JSON(http.StatusOK, genGinH(Success, jMsgSuccess, data[0]))
	} else {
		r.request.JSON(http.StatusOK, genGinH(Success, jMsgSuccess))
	}
	r.request.Abort()
}

// ResponseException 异常响应
func (r *response) ResponseException(data *Exception) {
	r.request.JSON(data.HttpStatus, genGinH(data.Code, data.Message))
	r.request.Abort()
}
