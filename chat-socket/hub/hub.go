package hub

import "chatsocket/conn/conn_data"

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

var h = hub{
	Conn:     make(map[*conn_data.Connection]bool),
	Un_conn:  make(chan *conn_data.Connection),
	Msg_chan: make(chan []byte),
	En_conn:  make(chan *conn_data.Connection),
}

func (h *hub) run() {
	for {
		select {
			case
		}
	}
}
