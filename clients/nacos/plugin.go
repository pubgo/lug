package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/pubgo/x/merge"
	"github.com/pubgo/xerror"
	"go.uber.org/fx"

	"github.com/pubgo/lava/config"
	"github.com/pubgo/lava/consts"
	"github.com/pubgo/lava/logging"
	"github.com/pubgo/lava/module"
)

func init() {
	var cfgMap = make(map[string]*Cfg)
	xerror.Panic(config.Decode(Name, cfgMap))

	for name := range cfgMap {
		if name == consts.KeyDefault {
			name = ""
		}

		cfg := cfgMap[name]
		var ncp vo.NacosClientParam
		xerror.Panic(merge.CopyStruct(&ncp, cfg))
		module.Register(fx.Provide(fx.Annotated{
			Name: name,
			Target: func(log *logging.Logger) *Client {
				return &Client{
					cfg: xerror.PanicErr(clients.NewConfigClient(ncp)).(config_client.IConfigClient),
					srv: xerror.PanicErr(clients.NewNamingClient(ncp)).(naming_client.INamingClient),
				}
			},
		}))
	}
}
