package test

import (
	"testing"
)

import "server/usermgr"

func TestUserCreate(t *testing.T)  {
	if user, err := usermgr.Create("test", "123456", "M"); err == nil {
		t.Logf("创建成功, %v\n", user)
		if user, err := usermgr.Login("test", "123456"); err == nil {
			t.Logf("登录成功, %v\n", user)
			usermgr.List()
		} else {
			t.Errorf("登录失败，%v\n", err)
		}
	} else {
		t.Errorf("创建失败，%v\n", err)
	}
}
