package models

type User struct {
	Id       int64  `xorm:"pk autoincr"`
	Username string `xorm:"unique username"`
	Password string `xorm:"password"`
}
func (u *User) TableName() string {
	//返回数据库表名
	return "user"
}