package hello

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pubgo/lava/core/registry"

	"github.com/pubgo/lava"
	"github.com/pubgo/lava/core/debug/debug_plugin"
	"github.com/pubgo/lava/core/healthy"
	"github.com/pubgo/lava/example/protopb/proto/hello"
	"github.com/pubgo/lava/example/srv/hello/handler"
	"github.com/pubgo/lava/service"
	"github.com/pubgo/lava/service/gateway_plugin"
)

var name = "test-grpc"

func NewSrv() service.Service {
	srv := lava.NewService(name, "entry grpc test")

	registry.Enable(srv)
	debug_plugin.Enable(srv)
	gateway_plugin.Enable(srv)

	hello.RegisterTestApi(srv, handler.NewTestAPIHandler())

	// 健康检查
	healthy.Register(name, func(req *fiber.Ctx) error {
		return nil
	})

	return srv
}
