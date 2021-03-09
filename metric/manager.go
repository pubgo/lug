package metric

import (
	"github.com/pubgo/golug/consts"
	"github.com/pubgo/golug/types"
	"github.com/pubgo/xerror"
)

var reporters types.SMap

func Get(names ...string) Factory {
	val, ok := reporters.Load(consts.GetDefault(names...))
	if !ok {
		return nil
	}

	return val.(Factory)
}

func List() (dt map[string]Factory) {
	xerror.Panic(reporters.Map(&dt))
	return
}

func Register(name string, r Factory) (err error) {
	defer xerror.RespErr(&err)

	xerror.Assert(name == "" || r == nil, "[name,reporter] is null")
	xerror.Assert(reporters.Has(name), "reporter %s already exists", name)
	reporters.Set(name, r)
	return
}
