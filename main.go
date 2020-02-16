package main

import (
	"os"
	"os/signal"

	"github.com/goodblaster/actiontargets-chat/sockets"
	"github.com/goodblaster/actiontargets-chat/web"
)

// Ideally these are configurable. For sake of time and simplicity, I'm ignoring that.
// One problem is that the JS file that opens the websocket has a hardcoded port number.
const (
	webPort    = 8081
	socketPort = 8082
)

func main() {
	// Server web Files.
	go web.Run(webPort)

	// Handle web Sockets.
	go sockets.Run(socketPort)

	// Run forever until app is killed.
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	<-done
}
