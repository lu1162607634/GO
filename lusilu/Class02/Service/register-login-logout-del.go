package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/go-xorm/xorm"
	"httpcache/lusilu/Class02/models"
	"strconv"
)

type service struct {
	x *xorm.Engine //初始化
	c *redis.Client
}

//func init()  {
//	var err error
//赋值就好
//	x, err = xorm.NewEngine("mysql", "root:root@(127.0.0.1:3306)/lesson9?charset=utf8")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	x.ShowSQL(true)
//
//}

func getMd5(str string) string {

	hs := md5.New()
	hs.Write([]byte(str))
	return hex.EncodeToString(hs.Sum(nil))
}

func main() {
	s := &service{}
	x, _ := xorm.NewEngine("mysql", "root:root@(127.0.0.1:3306)/lesson9?charset=utf8")
	x.ShowSQL(true)
	s.x = x

	c := redis.NewClient(&redis.Options{
		Addr: "localhost:16379",
	})
	s.c = c
	router := gin.Default()
	//注册
	router.POST("/api/user/register", s.register)
	router.POST("/api/user/login", s.login)
	router.POST("/api/user/logout", s.middleware, s.logout)
	router.POST("/api/user/delete", s.middleware, s.delete)

	router.Run(":8081")
}

func (s *service) register(context *gin.Context) {

	username := context.PostForm("username")
	password := context.PostForm("password")

	u := new(models.User)
	has, err := s.x.Table("user").Where("username=?", username).Exist()
	fmt.Println(has)

	if has {
		outputErr(context, 200, 1, "用户已注册")
		return
	}
	u.Username = username
	u.Password = password
	_, err = s.x.Insert(u)
	if err != nil {
		outputErr(context, 200, 1, "注册失败")
		return
	}
	outputErr(context, 200, 0, "注册成功")
	return
}

func (s *service) login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	u := new(models.User)
	userInfo, err := s.x.Where("username=?", username).Get(u)
	if err != nil || userInfo == false {
		outputErr(c, 200, 1, "数据库查询失败")
		return
	}
	fmt.Println(userInfo)
	m := getMd5(username + password)
	s.c.Set(m, strconv.Itoa(int(u.Id)), 0)
	c.JSON(200, gin.H{
		"code":  0,
		"msg":   "登录成功",
		"token": m,
	})
	return
}

func (s *service) logout(c *gin.Context) {
	fmt.Println("exec logout")
	token := c.GetHeader("token")
	s.c.Del(token)
	outputErr(c, 200, 0, "登出成功")
	return
}

//注销
func (s *service) delete(c *gin.Context) {
	fmt.Println("delete login")

	token := c.GetHeader("token")
	userId, err := s.c.Get(token).Result()
	value, _ := s.c.Get("54f43988a2dac1d4aebfdbf0878a5a84").Result()
	fmt.Println(strconv.Atoi(value))

	if err != nil {
		outputErr(c, 200, 1, "注销失败2")
		return
	}
	id, _ := strconv.Atoi(userId)
	if id == 0 {
		outputErr(c, 200, 1, "没有存储该token对应用户信息")
		return
	}
	fmt.Println(strconv.Atoi(userId))
	u := new(models.User)
	info, err := s.x.Id(userId).Delete(u)
	fmt.Println("del", info)
	if err != nil {
		outputErr(c, 200, 1, "注销异常")
		return
	}
	if info == 0 {
		outputErr(c, 200, 1, "没有该用户")
		return
	}
	outputErr(c, 200, 0, "注销成功")
	return
}

func (s *service) middleware(c *gin.Context) {
	fmt.Println("exec middleware")
	token := c.GetHeader("token")
	val, err := s.c.Get(token).Result()

	if err != nil {
		outputErr(c, 200, 1, "获取token失败2")
		c.Abort()
		return
	}
	if val == "" {
		outputErr(c, 200, 1, "该token不存在，请重新登录")
		c.Abort()
		return
	}
	c.Next()
}

func outputErr(c *gin.Context, status int, code int, msg string) {
	c.JSON(status, gin.H{
		"code": code,
		"msg":  msg,
	})
}
