package game

import (
	"strings"

	"github.com/mcaci/ita-cards/card"
)

type Req struct {
	action       string
	origin       string
	data1, data2 string
}

func NewReq(request, origin string) *Req {
	var req Req
	fields := strings.Split(request, "#")
	if len(fields) > 0 {
		req.action = fields[0]
	}
	if len(fields) > 1 {
		req.data1 = fields[1]
	}
	if len(fields) > 2 {
		req.data2 = fields[2]
	}
	req.origin = origin
	return &req
}

func (r *Req) Action() string {
	return r.action
}

func (r *Req) From() string {
	return r.origin
}

func (r *Req) Card() (*card.Item, error) {
	return card.New(r.data1, r.data2)
}

func (r *Req) Value() string {
	return r.data1
}