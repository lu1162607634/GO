package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"httpcache/lusilu/Class03/model"
)

func Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("token")
		redis := model.ConnRedis()
		value, err := redis.Get(token).Result()

		fmt.Println("valuev", value)
		fmt.Println(err)
		if len(value) == 0 {
			context.JSON(200, gin.H{
				"code": 1,
				"msg":  "缓存不存在",
			})
			context.Abort()
			return
		}
		if err != nil {
			context.JSON(200, gin.H{
				"code": 1,
				"msg":  "中间价失败",
			})
			context.Abort()
			return
		}
		context.Next()
	}
}
