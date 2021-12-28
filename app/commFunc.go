package app

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"runtime"
)

func DD(param ...interface{}) bool {
	funcName, file, line, ok := runtime.Caller(1)
	fmt.Println("\n\n\n\n \033[1;31;33m start  \033[0m ", gtime.Now())
	fmt.Printf("\033[1;31;32m 	file: \033[0m   %s		\033[1;31;34m line: \033[0m  %d  "+runtime.FuncForPC(funcName).Name()+"\n", file, line)

	if ok {
		if len(param) > 1 {
			for _, val := range param {
				fmt.Printf("\033[1;31;32m类型为 :\033[0m%T\n ", val)
				fmt.Printf("\033[1;31;32m 变量值为 : \033[0m  %+#v\n", val)
				g.Dump(val)
			}
		} else {
			//fmt.Println(reflect.TypeOf(param[0]))
			fmt.Printf("\033[1;31;32m类型为 :\033[0m%T\n ", param[0])
			fmt.Printf("\033[1;31;32m 变量值为 : \033[0m  %+#v\n", param[0])
		}

	}
	fmt.Println("_____________end_________________", gtime.Now())

	return true
}

func OutSuccess(r *ghttp.Request, data ...interface{}) {
	if len(data) == 0 {
		data = []interface{}{}
	}
	err := r.Response.WriteJsonExit(g.Map{
		"code":    0,
		"message": "success",
		"data":    data,
	})
	if err != nil {
		r.Response.Writeln(err)
	}
}

func OutFalse(r *ghttp.Request, data ...interface{}) {
	rmsg := "error"
	var rdata g.Map

	if len(data) == 1 {
		rmsg = data[0].(string)
	} else if len(data) > 1 {
		rdata = gconv.Map(data[1])
	}

	err := r.Response.WriteJsonExit(g.Map{
		"code":    500,
		"message": rmsg,
		"data":    rdata,
	})
	if err != nil {
		r.Response.Writeln(err)
	}
}

func Isset(r *ghttp.Request, param interface{}) bool {
	if err := r.GetError(); err != nil {
		return false
	}
	if param == nil {
		return false
	}
	return true
}
