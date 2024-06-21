package conn

import (
	"chatsocket/conn/conn_data"
	"net/http"

	"github.com/gorilla/websocket"
)

var user_list = []string{}
var wu = &websocket.Upgrader{ReadBufferSize: 512,
	WriteBufferSize: 512, CheckOrigin: func(r *http.Request) bool { return true }}

func myws(w http.ResponseWriter, r *http.Request) {
	//协议升级
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	//创建连接
	c := &conn_data.Connection{Ws: ws, Sc: make(chan []byte, 256), Data: &conn_data.Data{}}
	//connection加入hub管理
	h.r <- c
	go c.writer()
	c.reader()
	//退出登录
	defer logout(c)
}
