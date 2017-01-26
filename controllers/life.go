package controllers

import (
	"github.com/astaxie/beego"
)

type LifeController struct {
	beego.Controller
}

func (this *LifeController) Get() {
	this.LayoutSections = make(map[string]string)
	this.TplName = "public/life.html"
	this.Layout = "public/layout.html"
	this.LayoutSections["PageScripts"] = "public/scripts_life.html"

}
