package sockets

import (
	"sync"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

var (
	sockets = make(map[uuid.UUID]*websocket.Conn)
	mutex   = sync.Mutex{}
)

func register(ws *websocket.Conn) uuid.UUID {
	mutex.Lock()
	defer mutex.Unlock()

	id := uuid.New()
	sockets[id] = ws

	return id
}

func deregister(id uuid.UUID) {
	mutex.Lock()
	defer mutex.Unlock()

	delete(sockets, id)
}
