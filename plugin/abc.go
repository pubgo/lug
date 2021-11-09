package plugin

import (
	"github.com/pubgo/lava/types"
)

const Name = "plugin"

type Plugin interface {
	String() string
	UniqueName() string
	Flags() types.Flags
	Commands() *types.Command
	Init() error
	Watch() types.Watcher
	Vars(types.Vars) error
	Health() types.Healthy
	Middleware() types.Middleware
	BeforeStarts() []func()
	AfterStarts() []func()
	BeforeStops() []func()
	AfterStops() []func()
}

type Process interface {
	BeforeStart(fn func())
	AfterStart(fn func())
	BeforeStop(fn func())
	AfterStop(fn func())
}
