package config

import (
	"fmt"
	"io"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/mitchellh/mapstructure"
	"github.com/pubgo/x/fx"
	"github.com/pubgo/x/typex"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"github.com/spf13/viper"
)

var (
	CfgType = "yaml"
	CfgName = "config"
	Home    = filepath.Join(xerror.PanicStr(filepath.Abs(filepath.Dir(""))), "home")
	CfgPath = ""
)

func GetCfg() *conf {
	xerror.AssertFn(cfg == nil, func() string { return "[config] please init config" })
	return cfg
}

var _ Config = (*conf)(nil)

var cfg = &conf{v: viper.New()}

type conf struct {
	rw sync.RWMutex
	v  *viper.Viper
}

func (t *conf) MergeConfig(in io.Reader) error {
	t.rw.Lock()
	defer t.rw.Unlock()

	return t.v.MergeConfig(in)
}

func (t *conf) AllKeys() []string {
	t.rw.RLock()
	defer t.rw.RUnlock()

	return t.v.AllKeys()
}

func (t *conf) ConfigFileUsed() string {
	t.rw.RLock()
	defer t.rw.RUnlock()

	return t.v.ConfigFileUsed()
}

func GetMap(keys ...string) map[string]interface{} {
	return GetCfg().GetStringMap(strings.Join(keys, "."))
}
func (t *conf) GetStringMap(key string) map[string]interface{} {
	t.rw.RLock()
	defer t.rw.RUnlock()

	return t.v.GetStringMap(key)
}

func (t *conf) Get(key string) interface{} {
	t.rw.RLock()
	defer t.rw.RUnlock()

	return t.v.Get(key)
}

func (t *conf) GetString(key string) string {
	t.rw.RLock()
	defer t.rw.RUnlock()

	return t.v.GetString(key)
}

func (t *conf) Set(key string, value interface{}) {
	t.rw.Lock()
	defer t.rw.Unlock()

	t.v.Set(key, value)
}

func (t *conf) UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	t.rw.RLock()
	defer t.rw.RUnlock()

	return t.v.UnmarshalKey(key, rawVal, opts...)
}

func Decode(name string, fn interface{}) (b bool) { return GetCfg().Decode(name, fn) }
func (t *conf) Decode(name string, fn interface{}) (b bool) {
	defer xerror.RespExit(name)

	xerror.Assert(name == "" || fn == nil, "[name,fn] should not be nil")
	if t.Get(name) == nil {
		xlog.Warnf("config key [%s] not found", name)
		return false
	}

	vfn := reflect.ValueOf(fn)
	switch vfn.Type().Kind() {
	case reflect.Func: // func(cfg *conf)
		xerror.Assert(vfn.Type().NumIn() != 1, "[fn] input num should be 1")

		mthIn := reflect.New(vfn.Type().In(0).Elem())
		ret := fx.WrapRaw(t.UnmarshalKey)(name, mthIn,
			func(cfg *mapstructure.DecoderConfig) { cfg.TagName = "json" })

		if !ret[0].IsNil() {
			xerror.PanicF(ret[0].Interface().(error),
				"config key %s decode error", name)
		}

		vfn.Call(typex.ValueOf(mthIn))
	case reflect.Ptr:
		xerror.PanicF(t.UnmarshalKey(name, fn,
			func(cfg *mapstructure.DecoderConfig) { cfg.TagName = "json" }),
			"config key %s decode error", name)
	default:
		xerror.AssertFn(true, func() string {
			return fmt.Sprintf("[fn] type error, refer: %#v", vfn)
		})
	}

	return true
}
