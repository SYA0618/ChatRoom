package conn_data

import (
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
