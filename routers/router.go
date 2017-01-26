package routers

import (
	"nohassls_material/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.Router("/profile", &controllers.ProfileController{}, "get:Get")
	beego.Router("/security", &controllers.SecurityController{}, "get:Get")
	beego.Router("/life", &controllers.LifeController{}, "get:Get")

	beego.Router("/blog", &controllers.BlogController{}, "get:Get")
	beego.Router("/blog/article/:id([0-9]+)", &controllers.BlogController{}, "get:GetOne")
	beego.Router("/blog/add", &controllers.BlogController{}, "get,post:Add")
	beego.Router("/blog/edit/:id([0-9]+)", &controllers.BlogController{}, "get,post:Edit")
	beego.Router("/blog/delete/:id([0-9]+)", &controllers.BlogController{}, "get:Delete")
}
