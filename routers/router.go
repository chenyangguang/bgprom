package routers

import (
	"github.com/chenyangguang/bgprom/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
