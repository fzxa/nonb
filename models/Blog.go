package models

import (
	"github.com/astaxie/beego/orm"
)

type MyBlog struct {
	Id int
	Title string
	Url_title string
	Content string
	Files string
	Date string
	View string
}


func init() {
	orm.RegisterModel(new(MyBlog))
}