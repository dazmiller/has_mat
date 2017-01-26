package main

import (
	_ "nohassls_material/routers"

	"github.com/astaxie/beego"

	models "nohassls_material/models"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	mysqlReg := beego.AppConfig.String("mysqluser") + ":" +
		beego.AppConfig.String("mysqlpass") + "@tcp(127.0.0.1:3306)/" +
		beego.AppConfig.String("mysqldb") + "?charset=utf8&parseTime=true&loc=Australia%2FSydney"
	orm.RegisterDataBase("default", "mysql", mysqlReg)
	orm.RegisterModel(new(models.Articles))
}

func main() {
	name := "default"
	force := false
	verbose := true

	err := orm.RunSyncdb(name, force, verbose)

	if err != nil {
		beego.Debug(err)
	}

	beego.Run()
}
