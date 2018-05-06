package controllers

import (
	"github.com/astaxie/beego"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	codeCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_request_total_code",
			Help: "total request code controller",
		},
	)
)

type CodeController struct {
	beego.Controller
}

func init() {
	prometheus.MustRegister(codeCounter)
}

func (c *CodeController) Get() {
	codeCounter.Inc()
	c.TplName = "code.tpl"

}
