package router

import (
	"gf-demo/app/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	s.Use(MiddlewareErrorHandler)
	s.SetIndexFolder(true)
	s.SetServerRoot("/public")
	//s.SetIndexFiles([]string{"index.html","index.htm"})
	//s.SetRewriteMap(g.MapStrStr{
	//	"/a": "/a.html",
	//})

	//设置域名
	s.Domain("localhost3")

	//fmt.Println("", gfile.MainPkgPath())

	//api 接口
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)

		group.ALL("/firstcrl", api.FirstCrl)
		group.ALL("/chat", api.Chat)
	})

}

func MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"goframe.org", "baidu.com"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		// 记录到自定义错误日志文件
		g.Log("exception").Error(err)
		//返回固定的友好信息
		r.Response.ClearBuffer()
		//r.Response.Writef("%+v", err)
		r.Response.Writeln("服务器居然开小差了，请稍后再试吧！")
	}
}
