Beego project with prometheus monitor.

## Install  && import
+ Create a project  with bee tool.

```
bee new bgprom 
```

+ Add  promhttp

``` go
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
```

to the main.go.

##  Run

```
bee run 
```

Access  *http://localhost:8080/metrics* in browser.

## Add Counter 
New a Controller name CodeController or any other name same in router.go. Then import 
 ***"github.com/prometheus/client_golang/prometheus"*** . 

the test controller  should look like this: 

```
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
```

+ Visit browser by ***http://localhost:8080/metrics***  if there comes the result like below:
 
```
...
# HELP http_request_total_code total request code controller
# TYPE http_request_total_code counter
http_request_total_code 
...
```
it means works!

+ Access ***http://localhost:8080/code***, it will turn out that ***http_request_total_code*** increasing.
```
...
# HELP http_request_total_code total request code controller
# TYPE http_request_total_code counter
http_request_total_code 
...
```
Well done!
Other types are similar(like Summary, Gauge, Histogram, SummaryVec and so on).

## todo 
+ Add base prometheus metric
