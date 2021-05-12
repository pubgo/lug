package debug

import (
	"github.com/go-chi/chi/v5"
	"golang.org/x/net/trace"
)

func init() { On(traceRoute) }

func traceRoute(app *chi.Mux) {
	app.HandleFunc("/debug/requests", trace.Traces)
	app.HandleFunc("/debug/events", trace.Events)
}
