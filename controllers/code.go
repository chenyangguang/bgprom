package controllers

import (
	"github.com/astaxie/beego"
	"github.com/chenyangguang/bgprom/monitor"
)

type CodeController struct {
	beego.Controller
}

func (c *CodeController) Get() {
	monitor.CodeCounter.Inc()

	c.Ctx.WriteString("hello")
	// c.TplName = "code.tpl"
}
