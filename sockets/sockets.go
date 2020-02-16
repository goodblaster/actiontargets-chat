package sockets

import (
	"net/http"
	"strconv"

	"golang.org/x/net/websocket"
)

// Run - Panic instead of returning error because this
// should be run in a go routine.
func Run(port int) {
	mux := http.NewServeMux()
	mux.Handle("/", websocket.Handler(socketHandler))
	s := http.Server{Addr: ":" + strconv.Itoa(port), Handler: mux}
	err := s.ListenAndServe()
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

//
func socketHandler(ws *websocket.Conn) {
	id := register(ws)
	defer deregister(id)

	for {
		var buffer []byte
		err := websocket.Message.Receive(ws, &buffer)
		if err != nil {
			_ = ws.Close()
			return
		}

		broadcast(ws.Request().RemoteAddr, buffer)
	}
}
