package app

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"reflect"
	"runtime"
)

func DD(param ...interface{}) bool {
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
	return true
}

func OutSuccess(r *ghttp.Request, data interface{}) {
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
