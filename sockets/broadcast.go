package sockets

import "golang.org/x/net/websocket"

type broadcastMessage struct {
	Type    string `json:"msgType"`
	Message string `json:"message"`
	ID      string `json:"id"`
}

// Send the messages in a JSON format that can be
// expanded in the future.
func broadcast(sender string, message []byte) {
	mutex.Lock()
	defer mutex.Unlock()

	for id, ws := range sockets {
		bm := broadcastMessage{
			Type:    "broadcast-message",
			Message: string(message),
			ID:      sender,
		}

		err := websocket.JSON.Send(ws, bm)
		if err != nil {
			// TODO: Output error
			_ = ws.Close()
			delete(sockets, id)
		}
	}
}
