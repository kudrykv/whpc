package main

import (
	"github.com/gorilla/websocket"
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/kudrykv/whpc/app/internal/config"
	"github.com/kudrykv/whpc/app/internal/wsping"
	"github.com/kudrykv/whpc/app/handler"
	"github.com/kudrykv/whpc/app/internal/log"
)

func main() {
	cfg := config.Init()

	if len(cfg.Channel) == 0 || len(cfg.Route) == 0 {
		flag.PrintDefaults()
		return
	}

	log.WithFields(logrus.Fields{
		"host":          cfg.Host,
		"channel":       cfg.Channel,
		"route":         cfg.Route,
		"method":        cfg.Method,
		"ping_interval": cfg.PingInterval,
	}).Info("config")

	c, _, err := websocket.DefaultDialer.Dial(cfg.Host+cfg.Channel, nil)
	if err != nil {
		log.WithField("err", err).Error("dial host failed")
		return
	}
	defer c.Close()
	log.Info("host dialed")

	log.WithField("ping_interval", cfg.PingInterval).Info("setup ping routine")
	go wsping.Ping(c, cfg.PingInterval)

	log.Info("start process loop")
	sh := handler.NewSessionHandler(cfg)
	sh.StartWebsocketSessionLoop(c)
}
