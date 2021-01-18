package golug_db

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pubgo/dix"
	"github.com/pubgo/golug/golug_app"
	"github.com/pubgo/golug/golug_config"
	"github.com/pubgo/golug/golug_consts"
	"github.com/pubgo/golug/golug_types"
	"github.com/pubgo/golug/pkg/golug_utils"
	"github.com/pubgo/xerror"
	"xorm.io/xorm"
	xl "xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

var clientMap = golug_types.NewSyncMap()

func Get(names ...string) *xorm.Engine {
	var name = golug_consts.Default
	if len(names) > 0 && names[0] != "" {
		name = names[0]
	}

	val, ok := clientMap.Load(name)
	xerror.Assert(ok, "[db] %s not found", name)

	return val.(*xorm.Engine)
}

func List() map[string]*xorm.Engine {
	var data = make(map[string]*xorm.Engine)
	clientMap.Each(func(key string, value *xorm.Engine) { data[key] = value })
	return data
}

func initClient(name string, cfg Cfg) {
	source := golug_config.Template(cfg.Source)
	if strings.Contains(cfg.Driver, "sqlite") {
		if _dir := filepath.Dir(source); !golug_utils.PathExist(_dir) {
			_ = os.MkdirAll(_dir, 0644)
		}
	}

	engine := xerror.PanicErr(xorm.NewEngine(cfg.Driver, source)).(*xorm.Engine)
	engine.SetMaxOpenConns(cfg.MaxConnOpen)
	engine.SetMaxIdleConns(cfg.MaxConnIdle)
	engine.SetConnMaxLifetime(cfg.MaxConnTime)

	engine.Logger().SetLevel(xl.LOG_WARNING)
	if golug_app.IsDev() || golug_app.IsTest() {
		engine.Logger().SetLevel(xl.LOG_DEBUG)
	}

	xerror.Panic(engine.DB().Ping())
	engine.SetMapper(names.LintGonicMapper)

	clientMap.Set(name, engine)

	// 初始化完毕之后, 更新到对象管理系统
	updateEngine(name, engine)
}

func updateEngine(name string, engine *xorm.Engine) {
	xerror.Panic(dix.Dix(map[string]*xorm.Engine{name: engine}))
}
