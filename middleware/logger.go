package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s::%s %s %s %s %s \n",
			param.ClientIP,
			param.Latency,
			param.StatusCode,
			param.Request,
			param.TimeStamp,
			param.Path)
	})
}
