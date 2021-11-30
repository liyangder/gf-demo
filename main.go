package main

import (
	_ "gf-demo/boot"
	_ "gf-demo/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
