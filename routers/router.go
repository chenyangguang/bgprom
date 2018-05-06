package routers

import (
	"github.com/astaxie/beego"
	"github.com/chenyangguang/bgprom/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/code", &controllers.CodeController{})
}
