package main

import (
	"github.com/astaxie/beego"
	"github.com/chenyangguang/bgprom/monitor"
	_ "github.com/chenyangguang/bgprom/routers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	beego.Handler("/metrics", promhttp.Handler())
	monitor.RegisterMonitor()

	beego.Run()
}
