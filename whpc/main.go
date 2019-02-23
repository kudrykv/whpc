package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"github.com/kudrykv/whpc/whpc/handler"
	"github.com/kudrykv/whpc/whpc/internal/config"
	"github.com/kudrykv/whpc/whpc/internal/log"
	"github.com/kudrykv/whpc/whpc/internal/wsping"
	"github.com/sirupsen/logrus"
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

	log.WithFields(logrus.Fields{
		"webhook_url": "https://whps.herokuapp.com/webhook/"+cfg.Channel,
	}).Info("host dialed")

	log.WithField("ping_interval", cfg.PingInterval).Info("setup ping routine")
	go wsping.Ping(c, cfg.PingInterval)

	log.Info("start process loop")
	sh := handler.NewSessionHandler(cfg)
	sh.StartWebsocketSessionLoop(c)
}
