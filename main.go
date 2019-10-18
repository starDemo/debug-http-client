package main

import (
	_ "debug-http-client/boot"
	_ "debug-http-client/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
