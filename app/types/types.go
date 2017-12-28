package types

import "net/http"

type Req struct {
	Id     string      `json:"id"`
	Time   JsonTime    `json:"time"`
	Header http.Header `json:"headers"`
	Body   []byte      `json:"body"`
}
