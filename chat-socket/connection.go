package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// 用户连接结构体
type connection struct {
	ws   *websocket.Conn
	sc   chan []byte
	data *Data
}

// 用户在线名单列表
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

// websocket服务
func myws(w http.ResponseWriter, r *http.Request) {
	//协议升级
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	//创建连接
	c := &connection{sc: make(chan []byte, 256), ws: ws, data: &Data{}}
	//connection加入hub管理
	h.r <- c
	go c.writer()
	c.reader()
	//退出登录
	defer logout(c)
}

// 数据写入器
func (c *connection) writer() {
	//取出发送信息并写入
	for message := range c.sc {
		fmt.Println(message, "\n")
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

// 数据读取器
func (c *connection) reader() {
	for {
		//接收ws信息
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			h.r <- c
			break
		}
		json.Unmarshal(message, &c.data)
		//解析信息类型
		switch c.data.Type {
		//用户登录
		case "login":
			c.data.User = c.data.Content
			c.data.From = c.data.User
			//在线人数增加
			user_list = append(user_list, c.data.User)
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			//发送信息
			h.b <- data_b
		case "user":
			c.data.Type = "user"
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		//用户登出
		case "logout":
			c.data.Type = "logout"
			//在线人数减少
			user_list = del(user_list, c.data.User)
			data_b, _ := json.Marshal(c.data)
			//删除连接
			h.b <- data_b
			//发送用户离线信息
			h.r <- c
		default:
			fmt.Print("========default================")
		}
	}
}

// 删除登出的用户，维护在线用户名单
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
