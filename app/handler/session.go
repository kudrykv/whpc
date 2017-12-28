package handler

import (
	"encoding/json"
	"bytes"
	"io/ioutil"
	"github.com/gorilla/websocket"
	"github.com/kudrykv/whpc/app/internal/config"
	"github.com/kudrykv/whpc/app/types"
	"net/http"
	"github.com/kudrykv/whpc/app/internal/log"
	"github.com/sirupsen/logrus"
	"time"
)

type sessionHandler struct {
	config config.Config
}

func NewSessionHandler(config config.Config) *sessionHandler {
	return &sessionHandler{
		config: config,
	}
}

func (h *sessionHandler) StartWebsocketSessionLoop(c *websocket.Conn) {
	for {
		_, bytesMessage, err := c.ReadMessage()
		if err != nil {
			log.WithField("err", err).Error("failed to receive a message")
			return
		}

		req := types.Req{}
		if err := json.Unmarshal(bytesMessage, &req); err != nil {
			log.WithFields(logrus.Fields{
				"err":     err,
				"message": string(bytesMessage),
			}).Error("failed to unmarshal message")
			continue
		}

		r, err := http.NewRequest(h.config.Method, h.config.Route, bytes.NewReader(req.Body))
		if err != nil {
			log.WithField("err", err).Error("failed to create request")
			continue
		}

		r.Header = req.Header
		resp, err := http.DefaultClient.Do(r)
		if err != nil {
			log.WithField("err", err).Error("failed to perform a request")
			continue
		}

		respBodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.WithField("err", err).Error("failed to read response body")
			continue
		}

		log.WithFields(logrus.Fields{
			"status":   resp.Status,
			"body_len": len(respBodyBytes),
		}).Info("got response from the app")

		answer := types.Req{
			Id:     req.Id,
			Time:   types.JsonTime(time.Now()),
			Header: resp.Header,
			Body:   respBodyBytes,
		}

		bts, err := json.Marshal(answer)
		if err != nil {
			log.WithField("err", err).Error("failed to marshal the response")
			continue
		}

		if err := c.WriteMessage(websocket.TextMessage, bts); err != nil {
			log.WithField("err", err).Error("failed send back the response")
			continue
		}

		log.WithFields(logrus.Fields{
			"bytes_len": len(bts),
		}).Info("sent response back")
	}
}
