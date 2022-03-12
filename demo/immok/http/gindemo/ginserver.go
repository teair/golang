package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

const RequestId = "requestId"

func main() {
	r := gin.Default()

	logger, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}

	// middleware

	r.Use(func(gin *gin.Context) {
		s := time.Now()

		// 随机设置一个id
		rid := rand.Int()
		gin.Set(RequestId, rid)

		gin.Next()

		/*info :=  zap.Logger{

		}*/
		/*info := Info{
			msg:"incoming",
			path: gin.Request.URL.Path,
			status: gin.Writer.Status(),
			elapsed: time.Now().Sub(s),
		}*/

		// log latency、response code、 path
		logger.Info(
			"incoming",
			zap.String("Path", gin.Request.URL.Path),
			zap.Int("status", gin.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)),
			zap.Int(RequestId, rid),
		)
	})

	r.GET("/ping", func(c *gin.Context) {

		h := gin.H{
			"message": "pong",
		}
		if requestId, ok := c.Get(RequestId); ok {
			h[RequestId] = requestId
		}

		c.JSON(200, h)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
