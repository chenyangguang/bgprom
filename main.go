package main

import (
	_ "github.com/chenyangguang/bgprom/routers"
	"github.com/astaxie/beego"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    beego.Handler("/metrics", promhttp.Handler())
	beego.Run()
}

