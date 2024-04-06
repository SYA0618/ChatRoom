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

func (u *User) RegisterUser() bool {
	db := mysql.GetDB()
	_, err := db.Exec("INSERT INTO `account` (`user_name`, `password`)VALUES (?, ?)", u.User, u.Password)

	return err == nil
}
