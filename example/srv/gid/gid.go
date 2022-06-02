package gid

import (
	"github.com/pubgo/lava"
	"github.com/pubgo/lava/example/protopb/gidpb"
	"github.com/pubgo/lava/example/srv/gid/handler"
	"github.com/pubgo/lava/service"
)

func NewSrv() service.Service {
	var srv = lava.NewSrv("gid", "gid generate")
	srv.Dix(gidpb.RegisterIdServer)
	srv.Dix(handler.NewId)
	return srv
}
