package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	//執行Connection
	go h.run()
	//創建websocket服務
	router.HandleFunc("/ws", myws)
	//啟動http服務
	if err := http.ListenAndServe("0.0.0.0:8083", router); err != nil {
		fmt.Println("err:", err)
	}
}

// 用戶中心，維護多個用戶的connection
var h = hub{
	c: make(map[*connection]bool),
	u: make(chan *connection),
	b: make(chan []byte),
	r: make(chan *connection),
}

type hub struct {
	//當前在線connection訊息
	c map[*connection]bool
	//删除connection
	u chan *connection
	//傳遞數據
	b chan []byte
	//加入connection
	r chan *connection
}

func (h *hub) run() {
	for {
		select {
		//用戶連接，添加connection訊息
		case c := <-h.r:
			h.c[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			//發送給寫入器
			c.sc <- data_b
		//刪除指定用戶連接
		case c := <-h.u:
			if _, ok := h.c[c]; ok {
				delete(h.c, c)
				close(c.sc)
			}
		//向聊天室在線用戶發送訊息
		case data := <-h.b:
			for c := range h.c {
				select {
				//發送數據
				case c.sc <- data:
				//發送不成功則刪除connection訊息
				default:
					delete(h.c, c)
					close(c.sc)
				}
			}
		}
	}
}
