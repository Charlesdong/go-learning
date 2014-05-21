package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type ViewController struct {
	beego.Controller
}

func (this *ViewController) Get() {
	//inputs := this.Input()
	id, _ := strconv.Atoi(this.Ctx.Input.Params[":id"])
	bl := models.GetBlog(id)
	st := bl.Created.String()
	sa := strings.Split(st, ":")
	this.Data["Post"] = bl
	this.Data["CreateTime"] = sa[0] + ":" + sa[1]
	this.Layout = "layout.tpl"
	this.TplNames = "view.tpl"
}
func dealbr(s string) (rs string) {
	rs = strings.Replace(s, "\r\n", "<br>", -1)
	rs = strings.Replace(rs, "\n", "<br>", -1)
	return
}
func init() {
	beego.AddFuncMap("dealbr", dealbr)
}
