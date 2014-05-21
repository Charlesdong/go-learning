package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	all := models.GetAll()
	this.Data["blogs"] = all
	this.Layout = "layout.tpl"
	this.TplNames = "index.tpl"
}

func Cutstring(in string, ln int) (out string) {
	index := ln
	u := []rune(in)
	if len(u) < ln {
		index = len(u)
	}
	ou := u[:index]
	return string(ou) 
}

func init() {
		beego.AddFuncMap("cutstring",Cutstring)
}
