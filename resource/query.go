package resource

import (
	"io"

	"github.com/pubgo/lava/pkg/utils"
)

func Component(kind string, names ...string) *baseQuery {
	name := utils.GetDefault(names...)
	check(kind, name)
	return &baseQuery{kind: kind, name: join(kind, name)}
}

type baseQuery struct {
	kind, name string
}

func (t *baseQuery) Get() io.Closer {
	if val, ok := sources.Load(t.name); ok {
		return val.(Resource).getObj()
	}
	return nil
}

func (t *baseQuery) Load() (interface{}, Release) {
	if val, ok := sources.Load(t.name); ok {
		return val.(Resource).LoadObj()
	}
	return nil, nil
}

// Get 根据类型和名字获取一个资源
func Get(kind string, name string) Resource {
	check(kind, name)
	if val, ok := sources.Load(join(kind, name)); ok {
		return val.(Resource)
	}
	return nil
}

// GetByKind 通过资源类型获取资源列表
func GetByKind(kind string) map[string]Resource {
	check(kind, "check")
	var ss = make(map[string]Resource)
	sources.Range(func(name string, val interface{}) bool {
		if val.(Resource).Kind() == kind {
			ss[name] = val.(Resource)
		}
		return true
	})
	return ss
}

// GetOne 根据类型获取一个资源
func GetOne(kind string) Resource {
	check(kind, "check")
	var ss Resource
	sources.Range(func(_ string, val interface{}) bool {
		if val.(Resource).Kind() == kind {
			ss = val.(Resource)
			return false
		}
		return true
	})
	return ss
}

// GetAllKind 获取所有的资源类型
func GetAllKind() []string {
	var ss []string
	var set = make(map[string]struct{})
	sources.Range(func(_ string, val interface{}) bool {
		set[val.(Resource).Kind()] = struct{}{}
		return true
	})

	for k := range set {
		ss = append(ss, k)
	}
	return ss
}
