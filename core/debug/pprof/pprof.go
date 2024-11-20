package pprof

import (
	"errors"
	"net/http/pprof"

	"github.com/felixge/fgprof"
	fiber "github.com/gofiber/fiber/v3"
	"github.com/pubgo/lava/core/debug"
)

func init() {
	debug.Get("/gprof/", debug.Wrap(fgprof.Handler()))
	debug.Get("/pprof/", debug.WrapFunc(pprof.Index))
	debug.Get("/pprof/:name", func(ctx fiber.Ctx) error {
		name := ctx.Params("name")
		switch name {
		case "cmdline":
			return debug.WrapFunc(pprof.Cmdline)(ctx)
		case "profile":
			return debug.WrapFunc(pprof.Profile)(ctx)
		case "symbol":
			return debug.WrapFunc(pprof.Symbol)(ctx)
		case "trace":
			return debug.WrapFunc(pprof.Trace)(ctx)
		case "allocs":
			return debug.Wrap(pprof.Handler("allocs"))(ctx)
		case "goroutine":
			return debug.Wrap(pprof.Handler("goroutine"))(ctx)
		case "heap":
			return debug.Wrap(pprof.Handler("heap"))(ctx)
		case "mutex":
			return debug.Wrap(pprof.Handler("mutex"))(ctx)
		case "threadcreate":
			return debug.Wrap(pprof.Handler("threadcreate"))(ctx)
		}
		return errors.New("name not found")
	})
}
