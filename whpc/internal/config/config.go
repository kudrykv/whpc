package config

import "flag"

func Init() Config {
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
