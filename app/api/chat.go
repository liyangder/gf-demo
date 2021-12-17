package api

import (
	"fmt"
	"gf-demo/app"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var Chat = chat{
	clientList: make(map[string]*ghttp.WebSocket),
}

type chat struct {
	clientList map[string]*ghttp.WebSocket
}

func (c *chat) Index(r *ghttp.Request) {
	err := r.Response.WriteTpl("chat.html", g.Map{
		"id": 123,
	})
	if err != nil {

	}
}

func (c *chat) Websocket(r *ghttp.Request) {
	ws, err := r.WebSocket()

	app.DD(r.Session)
	fmt.Println(r.Session)
	clientIP := r.GetClientIp()
	if err != nil {
		glog.Error(err)
		r.Exit()
	} else {
		c.clientList[clientIP] = ws
	}
	fmt.Println(gtime.Now())
	fmt.Println(len(c.clientList))
	for {
		msgType, msg, err := ws.ReadMessage()

		if err != nil {
			return
		}
		temp := gconv.String(msg)
		temp = gtime.Now().String() + " [" + clientIP + "]: " + temp

		for k, _ := range c.clientList {
			if err = c.clientList[k].WriteMessage(msgType, gconv.Bytes(temp)); err != nil {
				return
			}
		}

		//if err = ws.WriteMessage(msgType, msg); err != nil {
		//	return
		//}
	}
}
