package etcdv3

import (
	"context"

	client3 "go.etcd.io/etcd/client/v3"

	"github.com/pubgo/lava/pkg/lavax"
	"github.com/pubgo/lava/resource"
)

// Get 获取etcd client
func Get(names ...string) *Client {
	var c = resource.Get(Name, lavax.GetDefault(names...))
	if c != nil {
		return c.(*Client)
	}
	return nil
}

type Client struct {
	resource.Resource
}

func (c *Client) Kind() string         { return Name }
func (c *Client) Get() *client3.Client { return c.GetObj().(*client3.Client) }
func (c *Client) Load() (*client3.Client, context.CancelFunc) {
	var obj, cancel = c.LoadObj()
	return obj.(*client3.Client), cancel
}
