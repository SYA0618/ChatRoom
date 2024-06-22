package hub

import (
	"chatsocket/conn/conn_data"
	"encoding/json"
)

type hub struct {
	//当前在线connection信息
	Conn map[*conn_data.Connection]bool
	//删除connection
	Un_conn chan *conn_data.Connection
	//传递数据
	Msg_chan chan []byte
	//加入connection
	En_conn chan *conn_data.Connection
}

var Center = hub{
	Conn:     make(map[*conn_data.Connection]bool),
	Un_conn:  make(chan *conn_data.Connection),
	Msg_chan: make(chan []byte),
	En_conn:  make(chan *conn_data.Connection),
}

func (h *hub) Run() {
	for {
		select {
		case conn_in := <-h.En_conn:
			h.Conn[conn_in] = true
			conn_in.Data.Ip = conn_in.Ws.RemoteAddr().String()
			conn_in.Data.Type = "handshake"
			conn_in.Data.UserList = conn_data.User_list
			data_b, _ := json.Marshal(conn_in.Data)
			conn_in.Sc <- data_b

		case conn_out := <-h.Un_conn:
			if _, ok := h.Conn[conn_out]; ok {
				delete(h.Conn, conn_out)
				close(conn_out.Sc)
			}

		case data := <-h.Msg_chan:
			for c := range h.Conn {
				select {
				case c.Sc <- data:

				default:
					delete(h.Conn, c)
					close(c.Sc)
				}
			}
		}
	}
}
