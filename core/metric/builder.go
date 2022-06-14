package metric

import (
	"context"

	"github.com/pubgo/dix"
	"github.com/pubgo/xerror"
	"github.com/uber-go/tally"

	"github.com/pubgo/lava/config"
	"github.com/pubgo/lava/core/lifecycle"
	"github.com/pubgo/lava/core/runmode"
	"github.com/pubgo/lava/logging/logkey"
	"github.com/pubgo/lava/service"
)

func init() {
	dix.Register(func(c config.Config) *Cfg {
		var cfg = DefaultCfg()
		xerror.Panic(c.UnmarshalKey(Name, &cfg))
		xerror.Assert(cfg.Driver == "", "metric driver is null")
		return &cfg
	})

	dix.Register(func(m lifecycle.Lifecycle, cfg *Cfg, sopts map[string]*tally.ScopeOptions) Metric {
		var opts = sopts[cfg.Driver]
		if opts == nil {
			opts = &tally.ScopeOptions{Reporter: tally.NullStatsReporter}
		}
		opts.Tags = Tags{logkey.Project: runmode.Project}
		if cfg.Separator != "" {
			opts.Separator = cfg.Separator
		}

		scope, closer := tally.NewRootScope(*opts, cfg.Interval)
		m.BeforeStops(func() { xerror.Panic(closer.Close()) })

		registerVars(scope)
		return scope
	})

	dix.Register(func(m Metric) service.Middleware {
		return func(next service.HandlerFunc) service.HandlerFunc {
			return func(ctx context.Context, req service.Request, resp service.Response) error {
				return next(CreateCtx(ctx, m), req, resp)
			}
		}
	})
}
