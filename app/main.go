package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"encoding/json"
	"fmt"
	"flag"
	"bytes"
	"time"
)

type Req struct {
	Header http.Header `json:"headers"`
	Body   []byte      `json:"body"`
}

type Config struct {
	Host         string
	Channel      string
	Route        string
	Method       string
	PingInterval int
}

func main() {
	cfg := initFlags()

	if len(cfg.Channel) == 0 || len(cfg.Route) == 0 {
		flag.PrintDefaults()
		return
	}

	c, _, err := websocket.DefaultDialer.Dial(cfg.Host+cfg.Channel, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	go ping(c, cfg.PingInterval)

	for {
		_, bytesMessage, err := c.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		req := Req{}
		if err := json.Unmarshal(bytesMessage, &req); err != nil {
			fmt.Println(err)
			continue
		}

		r, err := http.NewRequest(cfg.Method, cfg.Route, bytes.NewReader(req.Body))
		if err != nil {
			fmt.Println(err)
			continue
		}

		r.Header = req.Header
		if _, err := http.DefaultClient.Do(r); err != nil {
			fmt.Println(err)
		}
	}
}

func initFlags() Config {
	host := flag.String("host", "wss://whps.herokuapp.com/websocket/", "Custom websocket server.")
	channel := flag.String("channel", "", "Unique random channel to listen on.")
	route := flag.String("route", "", "Where to reroute requests.")
	method := flag.String("method", "POST", "HTTP method.")
	pingInterval := flag.Int("pinginterval", 10, "Ping interval in seconds to send to the server to keep a connection open.")
	flag.Parse()

	return Config{
		Host:         *host,
		Channel:      *channel,
		Route:        *route,
		Method:       *method,
		PingInterval: *pingInterval,
	}
}

func ping(c *websocket.Conn, pingInterval int) {
	timer := time.NewTicker(time.Duration(pingInterval) * time.Second)

	for {
		<-timer.C

		err := c.WriteMessage(websocket.PingMessage, nil)
		if err != nil {
			break
		}
	}
}
