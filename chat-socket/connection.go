package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// 用戶連接
type connection struct {
	ws   *websocket.Conn
	sc   chan []byte
	data *Data
}

// 用戶在線名單
var user_list = []string{}

type Data struct {
	Ip       string   `json:"ip"`
	User     string   `json:"user"`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"user_list"`
}

var wu = &websocket.Upgrader{ReadBufferSize: 512,
	WriteBufferSize: 512, CheckOrigin: func(r *http.Request) bool { return true }}

// websocket服務
func myws(w http.ResponseWriter, r *http.Request) {
	//協議upgrade
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	//創建連接
	c := &connection{sc: make(chan []byte, 256), ws: ws, data: &Data{}}
	//connection加入hub管理
	h.r <- c
	go c.writer()
	c.reader()
	//推出登入
	defer logout(c)
}

// 數據寫入
func (c *connection) writer() {
	//取出發送訊息並寫入
	for message := range c.sc {
		fmt.Println(message, "\n")
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

// 數據讀取
func (c *connection) reader() {
	for {
		//接收websocket訊息
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			h.r <- c
			break
		}
		json.Unmarshal(message, &c.data)
		//解析訊息類型
		switch c.data.Type {
		//User登入
		case "login":
			c.data.User = c.data.Content
			c.data.From = c.data.User
			//在線人數更新
			user_list = append(user_list, c.data.User)
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			//訊息發送
			h.b <- data_b
		case "user":
			c.data.Type = "user"
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		//用戶登出
		case "logout":
			c.data.Type = "logout"
			//在線人數減少
			user_list = del(user_list, c.data.User)
			data_b, _ := json.Marshal(c.data)
			//刪除連接
			h.b <- data_b
			//發送用戶登出訊息
			h.r <- c
		default:
			fmt.Print("========default================")
		}
	}
}

// 刪除登出用戶，維護用戶列表
func del(slice []string, user string) []string {
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}
	var n_slice = []string{}
	for i := range slice {
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			n_slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return n_slice
}

// 退出
func logout(c *connection) {
	c.data.Type = "logout"
	user_list = del(user_list, c.data.User)
	c.data.UserList = user_list
	c.data.Content = c.data.User
	data_b, _ := json.Marshal(c.data)
	h.b <- data_b
	h.r <- c
}
