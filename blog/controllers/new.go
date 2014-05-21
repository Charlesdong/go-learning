package controllers

import (
	"github.com/astaxie/beego"
    "blog/models"
	"time"
	"fmt"
)

type NewController struct {
		beego.Controller
}
func (this *NewController) Get() {
		this.Layout = "layout.tpl"
		this.TplNames = "new.tpl"
}
func (this *NewController) Post() {
		inputs := this.Input()
		var blog models.Blog
		blog.Title = inputs.Get("title")
		fmt.Println("title=",inputs.Get("title"))
		blog.Content = inputs.Get("content")
		fmt.Println("content=",inputs.Get("content"))
		blog.Created = time.Now()
		models.SaveBlog(blog)
		this.Ctx.Redirect(302, "/")
}
