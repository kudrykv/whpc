package types

import "net/http"

type Req struct {
	Header http.Header `json:"headers"`
	Body   []byte      `json:"body"`
}
