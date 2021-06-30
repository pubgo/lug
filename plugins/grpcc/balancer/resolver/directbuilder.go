package resolver

import (
	"strings"

	"github.com/pubgo/xerror"
	"google.golang.org/grpc/resolver"
)

// directBuilder creates a directBuilder which is used to factory direct resolvers.
// example:
//   direct://<authority>/127.0.0.1:9000,127.0.0.2:9000
type directBuilder struct{}

func (d *directBuilder) Scheme() string { return DirectScheme }

// Build [direct:///127.0.0.1,etcd:2379]
func (d *directBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	// 根据规则解析出地址
	endpoints := strings.FieldsFunc(target.Endpoint, func(r rune) bool { return r == EndpointSepChar })
	if len(endpoints) == 0 {
		return nil, xerror.Fmt("%v has not endpoint", target)
	}

	// 构造resolver address, 并处理副本集
	var addrs []resolver.Address
	for i := range endpoints {
		addr := endpoints[i]
		for j := 0; j < Replica; j++ {
			addrs = append(addrs, newAddr(addr, addr))
		}
	}

	return &baseResolver{cc: cc},
		xerror.WrapF(cc.UpdateState(newState(addrs)), "update resolver address: %v", addrs)
}
