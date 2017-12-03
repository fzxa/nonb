package main

import (
	"time"
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "hello/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"hello/models"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/fire_blog?charset=utf8")
}

func main() {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默认default

	// blog := models.MyBlog{Id:14}
	// err := o.Read(&blog)

	// if err == orm.ErrNoRows {
	// 	fmt.Println("查询不到")
	// } else if err == orm.ErrMissPK {
	// 	fmt.Println("找不到主键")
	// } else {
	// 	fmt.Println(blog.Date)
	// 	fmt.Println("测试一下。。。。")
	// }
	
	var blog models.MyBlog

	blog.Title = "Golang"
	blog.Content = "切换成Golange了"
	blog.Url_title = "golang-blog"
	blog.Date = time.Now().String()

	id, err := o.Insert(&blog)
	if err == nil {
		fmt.Println(id)
	}

	beego.Run()
}

