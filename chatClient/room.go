package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
  参加チャネル
  退出チャネル
  client管理
*/
type room struct {
	forward chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
}

/*
  入退室の機能を追加する
  参加用、退出用のチャネルを作成する
  在室しているクライアントを保存する
*/
func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			<-r.join
			r.clients[client] = true
		case client := <-r.leave:
			<-r.leave
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			for client := range r.clients {
				select {
				case client.send <- msg:
					//send message text
				default:
					//message of send  failed
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServerHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}