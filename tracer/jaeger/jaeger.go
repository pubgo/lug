package jaeger

import (
	"github.com/opentracing/opentracing-go"
	"github.com/pubgo/lug"
	"github.com/pubgo/lug/tracer"
	"github.com/pubgo/x/merge"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"
)

func init() {
	xerror.Exit(tracer.Register(Name, NewWithMap))
}

func NewWithMap(cfgMap map[string]interface{}) (opentracing.Tracer, error) {
	var cfg = GetDefaultCfg()
	xerror.Panic(merge.MapStruct(&cfg, cfgMap))
	return New(cfg)
}

func New(cfg *Cfg) (opentracing.Tracer, error) {
	logOpt := config.Logger(&logger{logs: xlog.Named(cfg.ServiceName)})
	metrics := config.Metrics(prometheus.New())

	//r := jaeger.NewRemoteReporter(transport.NewIOTransport(os.Stdout, 1))
	//reporter := config.Reporter(r)

	trace, closer, err := cfg.NewTracer(logOpt, metrics)
	xerror.Panic(err)

	lug.AfterStop(func() {
		if closer != nil {
			_ = closer.Close()
		}
	})

	return trace, nil
}
