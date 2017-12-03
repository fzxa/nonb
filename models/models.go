package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int
	Name string
}


func init() {
	orm.RegisterModel(new(User))
}