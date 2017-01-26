package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.LayoutSections = make(map[string]string)
	this.TplName = "public/homepage.html"
	this.Layout = "public/layout.html"
	this.LayoutSections["PageScripts"] = "public/scripts_dash.html"
}
