package prometheus

import (
	"github.com/pubgo/x/merge"
	"github.com/pubgo/xerror"
	"github.com/uber-go/tally"
	"github.com/uber-go/tally/prometheus"
	"go.uber.org/zap"

	"github.com/pubgo/lug/logger"
	"github.com/pubgo/lug/metric"
)

const Name = "prometheus"

func init() {
	metric.Register(Name, func(cfg map[string]interface{}, opts *tally.ScopeOptions) (err error) {
		opts.Separator = prometheus.DefaultSeparator
		opts.SanitizeOptions = &prometheus.DefaultSanitizerOpts

		var proCfg = prometheus.Configuration{}
		xerror.Panic(merge.MapStruct(&cfg, &proCfg))

		opts.CachedReporter, err = proCfg.NewReporter(
			prometheus.ConfigurationOptions{
				OnError: func(e error) {
					zap.L().Error("metric error", logger.Err(e), logger.Pkg("metric.prometheus"))
				},
			})
		return xerror.Wrap(err)
	})
}
