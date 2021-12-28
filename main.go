package main

import (
	_ "gf-demo/boot"
	_ "gf-demo/router"
	"github.com/gogf/gf/frame/g"
	"log"
	"os/exec"
	"runtime"
)

//cmd 初始化
func cmd_init() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("sh", "-c", "cd ./command&&gf build ./start.go -p . -a amd64 -s windows,linux &&cd ..")
		err := cmd.Run()
		if err != nil {
			log.Fatalf("______cmd.Run() failed with %s\n", err)
		}
	}
}

func main() {
	cmd_init()
	g.Server().Run()
}
