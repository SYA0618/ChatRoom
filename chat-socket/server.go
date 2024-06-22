package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	//开启协程启动connection服务管理中心
	go h.run()
	//创建ws服务
	router.HandleFunc("/ws", myws)
	//启动http服务
	if err := http.ListenAndServe("0.0.0.0:8083", router); err != nil {
		fmt.Println("err:", err)
	}
}

// 用户中心，维护多个用户的connection
var h = hub{
	c: make(map[*connection]bool),
	u: make(chan *connection),
	b: make(chan []byte),
	r: make(chan *connection),
}

type hub struct {
	//当前在线connection信息
	c map[*connection]bool
	//删除connection
	u chan *connection
	//传递数据
	b chan []byte
	//加入connection
	r chan *connection
}

func (h *hub) run() {
	for {
		select {
		//用户连接，添加connection信息
		case c := <-h.r:
			h.c[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			//发送给写入器
			c.sc <- data_b
		//删除指定用户连接
		case c := <-h.u:
			if _, ok := h.c[c]; ok {
				delete(h.c, c)
				close(c.sc)
			}
		//向聊天室在线人员发送信息
		case data := <-h.b:
			for c := range h.c {
				select {
				//发送数据
				case c.sc <- data:
				//发送不成功则删除connection信息
				default:
					delete(h.c, c)
					close(c.sc)
				}
			}
		}
	}
}
