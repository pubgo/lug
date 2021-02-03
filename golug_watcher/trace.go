package golug_watcher

import (
	"github.com/pubgo/golug/golug_trace"
	"github.com/pubgo/xprocess/xutil"
)

func init() {
	golug_trace.Watch(Name+"_watcher_callback", func() interface{} {
		var dt []string
		callbackMap.Each(func(key string) { dt = append(dt, key) })
		return dt
	})

	golug_trace.Watch(Name+"_watcher", func() interface{} {
		var dt = make(map[string]string)
		for k, v := range List() {
			dt[k] = xutil.FuncStack(v)
		}
		return dt
	})
}
