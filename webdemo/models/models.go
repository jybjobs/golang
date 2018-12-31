package models

import (
	//"shopping/helper"

	"github.com/astaxie/beego/orm"
	//_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	Username string
	Password string
	Email    string
}

func init() {
	orm.RegisterModel(new(User)) //注册模型并使用表前缀
}

//func (u *User) Add(uid int, username string, password string) User {
func (u *User) Add(usr User) User {

	//o := orm.NewOrm()
	//	var user User
	// user.Id = uid
	// user.Username = username
	// user.Password = password
	// usr := User{Id: uid, Username: username, Password: password}
	//	id, err := o.Insert(&user)
	return usr
}
