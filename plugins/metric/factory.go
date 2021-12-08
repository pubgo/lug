package metric

import (
	"github.com/pubgo/xerror"
	"github.com/uber-go/tally"

	"github.com/pubgo/lava/pkg/lavax"
	"github.com/pubgo/lava/pkg/typex"
)

type Factory func(cfg map[string]interface{}, opts *tally.ScopeOptions) error

var reporters typex.SMap

func Get(names ...string) Factory {
	val, ok := reporters.Load(lavax.GetDefault(names...))
	if !ok {
		return nil
	}

	return val.(Factory)
}

func Register(name string, r Factory) {
	defer xerror.RespExit()
	xerror.Assert(name == "" || r == nil, "[name,r] is null")
	xerror.Assert(reporters.Has(name), "reporter [%s] already exists", name)
	reporters.Set(name, r)
	return
}
