package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/j23063519/clean_architecture/pkg/log"
	"github.com/j23063519/clean_architecture/pkg/util"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get response
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		// request data
		var requestBody []byte
		if c.Request.Body != nil {
			// c.Request.Body is a buffer only can read one time
			requestBody, _ = io.ReadAll(c.Request.Body)
			// if read it then must give it back
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// set start time
		start := time.Now()
		c.Next()

		// record log
		cost := time.Since(start)
		responStatus := c.Writer.Status()

		// define zap.Field
		logFields := []zap.Field{
			zap.Int("status", responStatus),
			zap.String("request", c.Request.Method+" "+c.Request.URL.String()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", util.MicrosecondsStr(cost)),
		}

		// [post、put、delete] request、response put into logFields
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			// request body
			logFields = append(logFields, zap.String("Request Body", string(requestBody)))

			// response body
			logFields = append(logFields, zap.String("Response Body", w.body.String()))
		}

		// log
		if responStatus >= 400 && responStatus <= 499 {
			log.Logger.Warn("HTTP Wanring "+cast.ToString(responStatus), logFields...)
		} else if responStatus >= 500 && responStatus <= 599 {
			log.Logger.Error("HTTP Error "+cast.ToString(responStatus), logFields...)
		} else {
			log.Logger.Debug("HTTP Access Log", logFields...)
		}
	}
}
