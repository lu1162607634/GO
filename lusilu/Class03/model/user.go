package model

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type User struct {
	Id       int64  `xorm:"pk autoincr"`
	Username string `xorm:"unique username"`
	Password string `xorm:"password"`
}

func (u *User) TableName() string {
	//返回数据库表名
	return "user"
}

func ConnMysql() *xorm.Engine {
	x, err := xorm.NewEngine("mysql", "root:root@(127.0.0.1:3306)/lesson9?charset=utf8")
	if err != nil {
		fmt.Println("mysql连接失败")
		return &xorm.Engine{}
	}
	return x
}

func ConnRedis() *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr: "localhost:16379",
	})
	return c
}
