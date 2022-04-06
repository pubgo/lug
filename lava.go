package lava

import (
	"github.com/pubgo/lava/plugin"
	"github.com/pubgo/lava/service"
	"github.com/pubgo/lava/service/service_builder"
)

func Run(services ...service.Service) {
	service_builder.Run(services...)
}

func NewService(name string, desc string, plugins ...plugin.Plugin) service.Service {
	return service_builder.New(name, desc, plugins...)
}
