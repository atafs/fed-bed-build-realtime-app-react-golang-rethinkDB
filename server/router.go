package main

import (
	"github.com/gorilla/websocket"	
	"net/http"
	"time"	
)
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Router struct {}

func (e *Router) ServeHttp(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)	
}