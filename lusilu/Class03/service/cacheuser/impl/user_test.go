package impl

import (
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {
	service := New()
	username := "lusilu01"
	password := "1234"

	code, _ := service.Register(username, password)
	if code == 0 {
		t.Log("注册测试成功")
	} else {
		t.Error("注册测试失败")
	}
	username2 := "lusilu02"
	password2 := "12345"
	code2, _ := service.Register(username2, password2)
	if code2 == 0 {
		t.Log("注册测试成功")
	} else {
		t.Error("注册测试失败")
	}
}

func TestLogin(t *testing.T) {
	service := New()
	username := "lusilu01"
	password := "1234"

	token, err := service.Login(username, password)
	//fmt.Println(token)
	if token == "81dc9bdb52d04dc20036dbd8313ed055" {
		t.Log("登录测试成功")
	}
	if err != nil {
		t.Error("登录测试失败")
	}
}

func TestDelete(t *testing.T) {
	service := New()
	token := "81dc9bdb52d04dc20036dbd8313ed055"
	code, msg := service.Delete(token)
	fmt.Println(msg)
	if code == 0 {
		t.Log("注销测试成功")
	} else {
		t.Error("注销测试失败")
	}
}

func TestLogout(t *testing.T) {
	service := New()
	token := "81dc9bdb52d04dc20036dbd8313ed055"
	code, _ := service.Logout(token)
	if code == 0 {
		t.Log("登出测试成功")
	} else {
		t.Error("登出测试失败")
	}
}
