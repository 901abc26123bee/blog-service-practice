package middleware

import (
	"blog-service/global"
	"blog-service/pkg/logger"
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
)

/*
沒辦法非常直接的獲取到方法所返回的響應主體，
這時候我們需要巧妙利用Go interface 的特性，實際上在寫入流時，調用的是http.ResponseWriter
那麼我們只需要寫一個針對訪問日誌的Writer 結構體，實現我們特定的Write 方法就可以解決無法直接取到方法響應主體的問題了
我們在AccessLogWriter 的Write 方法中，實現了雙寫，因此我們可以直接通過AccessLogWriter 的body 取到值
*/
type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request": c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		global.Logger.WithFields(fields).Infof(
			"access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}