package watcher

import (
	"context"

	"github.com/pubgo/x/jsonx"
	"github.com/pubgo/xerror"
)

type Factory func(cfg map[string]interface{}) (Watcher, error)

// Watcher ...
type Watcher interface {
	Watch(ctx context.Context, key string, opts ...OpOption) <-chan *Response
	Name() string
}

type OpOption func(*Op)
type Op struct{}

type CallBack func(event *Response) error
type Response struct {
	Event    string
	Key      string
	Value    []byte
	Revision int64
}

func (t *Response) OnPut(fn func()) {
	xerror.Panic(t.checkEventType())
	if t.Event == "PUT" {
		fn()
	}
}

func (t *Response) OnDelete(fn func()) {
	xerror.Panic(t.checkEventType())
	if t.Event == "DELETE" {
		fn()
	}
}

func (t *Response) Decode(val interface{}) (gErr error) {
	defer xerror.RespErr(&gErr)

	var err = jsonx.Unmarshal(t.Value, val)
	gErr = xerror.WrapF(err, "input: %#v, output: %t", t.Value, val)

	return
}

func (t *Response) checkEventType() error {
	switch t.Event {
	case "DELETE", "PUT":
		return nil
	default:
		return xerror.New("unknown type")
	}
}
