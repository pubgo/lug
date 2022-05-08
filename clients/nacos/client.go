package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/pubgo/xerror"
)

type Client struct {
	srv naming_client.INamingClient
	cfg config_client.IConfigClient
}

func (c Client) GetCfg() config_client.IConfigClient {
	xerror.Assert(c.cfg == nil, "please init config client")
	return c.cfg
}

func (c Client) GetRegistry() naming_client.INamingClient {
	xerror.Assert(c.srv == nil, "please init naming client")
	return c.srv
}
