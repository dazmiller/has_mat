package controllers

import (
	"github.com/astaxie/beego"
)

type ProfileController struct {
	beego.Controller
}

func (this *ProfileController) Get() {
	this.LayoutSections = make(map[string]string)
	this.TplName = "public/profile.html"
	this.Layout = "public/layout.html"
	this.LayoutSections["PageScripts"] = "public/scripts_profile.html"
}
