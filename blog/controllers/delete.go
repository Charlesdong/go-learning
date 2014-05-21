package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"blog/models"
	"fmt"
)

type DeleteController struct {
		beego.Controller
}
func (this *DeleteController) Get() {
		//inputs := this.Input()
		id, _ := strconv.Atoi(this.Ctx.Input.Params[":id"])
		fmt.Println(id);
		bg := models.GetBlog(id)
		fmt.Println(bg);
//		this.Data["Post"] = models.DelBlog(bg)
		models.DelBlog(bg)
		this.Ctx.Redirect(302, "/")
}

