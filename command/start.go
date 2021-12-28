package main

import (
	"fmt"
	"github.com/gogf/gf/os/gcmd"
)

type abc struct {
}

func (*abc) help() {
	fmt.Println(123)
}
func (*abc) test() {
	fmt.Println("test")

}

var Abc = abc{}

func main() {
	gcmd.BindHandle("help", Abc.test)
	gcmd.AutoRun()
}
