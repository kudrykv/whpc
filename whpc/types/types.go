package types

import (
	"net/http"
	"time"
)

type Req struct {
	Id     string      `json:"id"`
	Time   time.Time   `json:"time"`
	Status int         `json:"status"`
	Header http.Header `json:"headers"`
	Body   []byte      `json:"body"`
}
