package api

import (
	"gf-demo/app/service"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"time"
)

var FirstCrl = firstCrl{}

type firstCrl struct{}

// Index is a demonstration route handler for output "Hello World!".
func (*firstCrl) Index(r *ghttp.Request) {
	r.Response.Writeln("aaa")

	service.Jd.Test()

	r.Response.WriteJsonExit(gtime.New(time.Now()))

}

// Upload uploads files to ./tmp .
func (*firstCrl) Upload(r *ghttp.Request) {
	files := r.GetUploadFiles("file_name")
	names, err := files.Save("./tmp/")
	if err != nil {
		r.Response.WriteExit(err)
	}
	r.Response.WriteExit("upload successfully: ", names)
}
