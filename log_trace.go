package log

import (
	"github.com/gin-gonic/gin"
)

// 日志跟踪关键词
const TraceID = "TRACE_ID"

// 日志埋点

const (
	// 收到请求
	ReqStart = "RS"
	// 请求处理中
	ReqDo = "RD"
	// 请求结束
	ReqEnd = "RE"
)

func GetLogTraceID(ctx *gin.Context) string {
	if nil == ctx {
		return "-"
	}

	tid := ctx.Query(TraceID)

	return tid
}
