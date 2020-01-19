package cache

import (
	"github.com/gin-gonic/gin"
	"httpcache/lusilu/Class03/service/cacheuser"
	"httpcache/lusilu/Class03/service/cacheuser/impl"
)

type Controller struct {
	service cacheuser.Interface
}

func New() *Controller {
	return &Controller{service: impl.New()}
}

func (c *Controller) Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	token, err := c.service.Login(username, password)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code": 1,
			"msg":  "登录失败",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":  0,
		"msg":   "登录成功",
		"token": token,
	})

}

func (c *Controller) Register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	code, msg := c.service.Register(username, password)
	if code == 1 {
		ctx.JSON(200, gin.H{
			"code": code,
			"msg":  msg,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func (c *Controller) Logout(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	code, msg := c.service.Logout(token)
	if code == 1 {
		ctx.JSON(200, gin.H{
			"code": code,
			"msg":  msg,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func (c *Controller) Delete(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	code, msg := c.service.Delete(token)
	if code == 1 {
		ctx.JSON(200, gin.H{
			"code": code,
			"msg":  msg,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
	})
	return
}
