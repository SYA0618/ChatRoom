package conn_data

import (
	"chatsocket/hub"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type Connection struct {
	Ws   *websocket.Conn
	Sc   chan []byte
	Data *Data
}

type Data struct {
	Ip       string   `json:"ip"`
	User     string   `json:"user"`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"user_list"`
}

var User_list = []string{}

// 数据写入器
func (c *Connection) writer() {
	//取出发送信息并写入
	for message := range c.Sc {
		fmt.Println(message, "\n")
		c.Ws.WriteMessage(websocket.TextMessage, message)
	}
	c.Ws.Close()
}

// 数据读取器
func (c *Connection) reader() {
	for {
		//接收ws信息
		_, message, err := c.Ws.ReadMessage()
		if err != nil {
			hub.Center.En_conn <- c
			break
		}
		json.Unmarshal(message, &c.Data)
		//解析信息类型
		switch c.Data.Type {
		//用户登录
		case "login":
			c.Data.User = c.Data.Content
			c.Data.From = c.Data.User
			//在线人数增加
			User_list = append(User_list, c.Data.User)
			c.Data.UserList = User_list
			data_b, _ := json.Marshal(c.Data)
			//发送信息
			hub.Center.Msg_chan <- data_b
		case "user":
			c.Data.Type = "user"
			data_b, _ := json.Marshal(c.Data)
			hub.Center.Msg_chan <- data_b
		//用户登出
		case "logout":
			c.Data.Type = "logout"
			//在线人数减少
			User_list = del(User_list, c.Data.User)
			data_b, _ := json.Marshal(c.Data)
			//删除连接
			hub.Center.Msg_chan <- data_b
			//发送用户离线信息
			hub.Center.En_conn <- c
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
func logout(c *Connection) {
	c.Data.Type = "logout"
	User_list = del(User_list, c.Data.User)
	c.Data.UserList = User_list
	c.Data.Content = c.Data.User
	data_b, _ := json.Marshal(c.Data)
	hub.Center.Msg_chan <- data_b
	hub.Center.En_conn <- c
}
