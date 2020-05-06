package main

import "github.com/gorilla/websocket"

/*
  入退室の機能を追加する
*/

type client struct {
	socket *websocket.Conn
	send   chan []byte
	room   *room
}

//読み出し
func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
}

//書き出し
func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
