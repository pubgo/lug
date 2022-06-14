package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pubgo/lava/core/lifecycle"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

type Init interface {
	Init()
}

type Close interface {
	Close()
}

type Flags interface {
	Flags() []cli.Flag
}

type WebHandler interface {
	Router(r fiber.Router)
}

type Options struct {
	Id        string            `json:"id,omitempty"`
	Name      string            `json:"name,omitempty"`
	Version   string            `json:"version,omitempty"`
	Port      int               `json:"port,omitempty"`
	Addr      string            `json:"addr,omitempty"`
	Advertise string            `json:"advertise,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

type Command interface {
	Command() *cli.Command
	Start() error
	Stop() error
}

type AppInfo interface {
	Options() Options
}

type Service interface {
	lifecycle.Lifecycle
	grpc.ServiceRegistrar
	Command
	AppInfo
	Provider(provider interface{})
	SubCmd(cmd *cli.Command)
}
