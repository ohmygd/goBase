package mysql

import (
	//"github.com/ohmygd/mgo/model"
)

type User struct {
	//model.Mysql
	Name string `gorm:"column:name"`
	Mobile string `gorm:"mobile"`
	Openid string `gorm:"openid"`
}

func (u *User)TableName() string {
	return "user"
}