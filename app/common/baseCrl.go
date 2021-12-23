package common

import (
	"fmt"
	"gf-demo/app"
	"github.com/gogf/gf/net/ghttp"
)

type BaseCrl struct {
	Name string
}

func (*BaseCrl) GetName() {
	fmt.Println("basecrl")
}

/**
*参数效验
 */
func (*BaseCrl) CheckParam(r *ghttp.Request, param interface{}) {
	if err := r.Parse(param); err != nil {
		app.OutFalse(r, err.Error())
	}
}
