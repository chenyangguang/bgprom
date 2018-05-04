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


## todo 
+ Add base prometheus metric
