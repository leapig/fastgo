package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/leapig/fastgo/app/library/helper"
	"net/http"
)

const (
	ServiceError = iota + 4000
	ParameterError
	AuthError
)

var ErrMsg = map[int]string{
	ParameterError: "参数错误",
	AuthError:      "无权操作",
}

type BaseResponse struct {
	Data         interface{} `json:"data" example:"${响应数据}"`
	Success      bool        `json:"success" example:"false"`
	TraceId      string      `json:"traceId" example:"${请求ID}"`
	ErrorCode    int32       `json:"errorCode" example:"${错误编码}"`
	ErrorMessage string      `json:"errorMessage" example:"${错误信息}"`
}

// R 快捷返回
func R(c *gin.Context, data interface{}, err error) {
	if err == nil {
		traceId, _ := c.Get("traceId")
		c.JSON(http.StatusOK,
			BaseResponse{Success: true,
				TraceId:      traceId.(string),
				ErrorCode:    0,
				ErrorMessage: "ok",
				Data:         data})
		c.Abort()
	} else {
		FRD(c, err, data)
	}
}

// FR 失败返回
func FR(c *gin.Context, err error) {
	FRC(c, err, nil, ServiceError)
}

// FRD 失败返回(Data)
func FRD(c *gin.Context, err error, data interface{}) {
	FRC(c, err, data, ServiceError)
}

// FRP 参数错误
func FRP(c *gin.Context) {
	FRC(c, nil, nil, ParameterError)
}

// FRA 无权操作
func FRA(c *gin.Context) {
	FRC(c, nil, nil, AuthError)
}

// FRC 失败返回(Data&Code)
func FRC(c *gin.Context, err error, data interface{}, code int) {
	traceId, _ := c.Get("traceId")
	var msg string
	if code != ServiceError {
		msg = ErrMsg[code]
	} else {
		msg = err.Error()
	}
	c.JSON(http.StatusBadRequest,
		BaseResponse{Success: false,
			TraceId:      traceId.(string),
			ErrorCode:    int32(code),
			ErrorMessage: msg,
			Data:         data})
	c.Abort()
}

// TraceId trace中间件
func TraceId() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Set("traceId", helper.Rid(helper.UserT))
		c.Next()
	}
}
