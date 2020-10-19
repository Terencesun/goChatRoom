package usermgr

import (
	errorCode "server/errorcode"
)

// 用户管理模块
type user struct {
	username string
	password string
	gender string
	online bool
}

type userList struct {
	users map[string]*user
}

var list *userList = &userList{
	users: make(map[string]*user),
}

func (l *userList) List() map[string]*user {
	return l.users
}

func Create(username string, password string, gender string) (interface{}, error) {
	if _, ok := list.users[username]; ok {
		// 存在
		return nil, errorCode.USEREXISTERROR
	}
	newUser := new(user)
	newUser.username = username
	newUser.gender = gender
	newUser.password = password
	newUser.online = false
	list.users[username] = newUser
	return newUser, nil
}

func Login(username string, password string) (interface{}, error) {
	if v, ok := list.users[username]; ok {
		// 存在
		if password == v.password {
			// 登录成功
			return v, nil
		} else {
			return nil, errorCode.USERLOGINERROR
		}
	} else {
		return nil, errorCode.USERNOEXISTERROR
	}
}

func List() map[string]*user {
	return list.users
}
