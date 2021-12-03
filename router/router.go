package router

import (
	"fmt"
	"gf-demo/app/api"
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
)

func init() {
	s := g.Server()

	s.SetIndexFolder(true)
	s.SetServerRoot("/public")
	//s.SetIndexFiles([]string{"index.html","index.htm"})
	s.SetRewriteMap(g.MapStrStr{
		"/a": "/a.html",
	})

	//设置域名
	s.Domain("localhost3")

	fmt.Println("1231231", gfile.MainPkgPath())

	//api 接口
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)

		group.ALL("/", api.FirstCrl)
	})

	//websocket
	s.BindHandler("/ws", func(r *ghttp.Request) {
		ws, err := r.WebSocket()
		if err != nil {
			glog.Error(err)
			r.Exit()
		}
		for {
			msgType, msg, err := ws.ReadMessage()
			if err != nil {
				return
			}

			temp := gvar.New(msg, true)

			newstring := gvar.New("服务器返回: " + temp.String())

			msg = newstring.Bytes()

			if err = ws.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

}

func MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"goframe.org", "baidu.com"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
