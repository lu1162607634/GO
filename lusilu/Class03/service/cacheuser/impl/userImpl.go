package impl

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/go-xorm/xorm"
	"httpcache/lusilu/Class02/models"
	"httpcache/lusilu/Class03/common"
	"httpcache/lusilu/Class03/model"
	"strconv"
)

type Service struct {
	x *xorm.Engine //初始化
	c *redis.Client
}

func New() *Service {
	return &Service{
		x: model.ConnMysql(),
		c: model.ConnRedis(),
	}
}

func (s *Service) Login(username string, password string) (token string, err error) {
	s.x.ShowSQL(true)
	u := new(models.User)
	md5password := common.GetMd5(password)
	userInfo, err := s.x.Where("username=? and password=?", username, md5password).Get(u)
	if err != nil || userInfo == false {
		return "无", err
	}
	fmt.Println(userInfo)
	//time.Now().Unix()
	m := common.GetMd5(password)
	s.c.Set(m, strconv.Itoa(int(u.Id)), 0)
	return m, err
}

func (s *Service) Register(username string, password string) (code int, msg string) {
	s.x.ShowSQL(true)
	u := new(models.User)
	md5password := common.GetMd5(password)
	has, err := s.x.Table("user").Where("username=?", username).Exist()
	if has || err != nil {
		return 1, "用户已注册或注册失败"
	}
	u.Username = username
	u.Password = md5password
	_, err = s.x.Insert(u)
	if err != nil {
	}
	return 0, "注册成功"

}

func (s *Service) Logout(token string) (code int, msg string) {
	s.c.Del(token)
	user_id := s.c.Get(token)
	fmt.Println("id", user_id)
	return 0, "登出成功"
}

func (s *Service) Delete(token string) (code int, msg string) {
	userId, err := s.c.Get(token).Result()
	fmt.Println("err", err)
	if err != nil {
		return 1, "获取redis异常"
	}
	user_id, err := strconv.Atoi(userId)

	if err != nil {
		return 1, "注销失败1"
	}
	u := new(models.User)
	info, err := s.x.Id(user_id).Delete(u)

	if err != nil {
		return 1, "注销失败2"
	}
	if info == 0 {
		return 1, "没有该用户3"
	}
	fmt.Println("sa")
	return 0, "注销成功"
}
