package api

import (
	"gf-demo/app"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

var Chat = chat{
	clientList: make(map[string]*ghttp.WebSocket),
}

type chat struct {
	clientList map[string]*ghttp.WebSocket
}

func (c *chat) Websocket(r *ghttp.Request) {
	ws, err := r.WebSocket()

	clientIP := r.GetClientIp()
	if err != nil {
		glog.Error(err)
		r.Exit()
	} else {
		app.DD(r)
		c.clientList[clientIP] = ws
	}

	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			return
		}
		temp := gconv.String(msg)
		temp = "[" + clientIP + "]: " + temp

		for _, v := range c.clientList {
			if err = v.WriteMessage(msgType, gconv.Bytes(temp)); err != nil {
				return
			}
		}

		//if err = ws.WriteMessage(msgType, msg); err != nil {
		//	return
		//}
	}
}
