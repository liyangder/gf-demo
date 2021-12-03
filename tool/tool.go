package tool

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"reflect"
	"runtime"
)

func DD(param ...interface{}) {
	funcName, file, line, ok := runtime.Caller(1)
	if ok {
		if len(param) > 1 {
			for _, val := range param {
				fmt.Printf("file: %s	 line: %d  "+runtime.FuncForPC(funcName).Name()+"\n", file, line)
				fmt.Printf(" 	类型为 :  %T		变量为 :  %+v\n", val, val)
				g.Dump(val)
			}
		} else {
			fmt.Println(reflect.TypeOf(param))

			fmt.Printf("file: %s	 line: %d  "+runtime.FuncForPC(funcName).Name()+"\n", file, line)
			fmt.Printf(" 	类型为 :  %T		变量为 :  %+v\n", param, param)
			g.Dump(param)
		}

	}
}
