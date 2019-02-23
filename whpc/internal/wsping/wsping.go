package wsping

import (
	"time"
	"github.com/gorilla/websocket"
)

func Ping(c *websocket.Conn, pingInterval int) {
	timer := time.NewTicker(time.Duration(pingInterval) * time.Second)

	for {
		<-timer.C

		err := c.WriteMessage(websocket.PingMessage, nil)
		if err != nil {
			break
		}
	}
}