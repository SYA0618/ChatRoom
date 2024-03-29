package models

import (
	"webserver/service/mysql"
)

type User struct {
	Id       int    `form:"id,omitempty"`
	User     string `form:"user_name"`
	Password string `form:"password"`
}

func (u *User) CheckUser() bool {
	db := mysql.GetDB()
	var queryUser User
	row := db.QueryRow("SELECT id, user_name, password FROM `account` WHERE user_name = ? and password = ?", u.User, u.Password)
	if err := row.Scan(&queryUser.Id, &queryUser.User, &queryUser.Password); err != nil {
		return false
	}
	return true
}
