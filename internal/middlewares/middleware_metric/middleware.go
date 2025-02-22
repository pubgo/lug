package middleware_metric

import (
	"context"
	"strings"
	"time"

	"github.com/pubgo/funk/generic"
	"github.com/uber-go/tally/v4"

	"github.com/pubgo/lava/core/metrics"
	"github.com/pubgo/lava/lava"
)

// grpc metric
// ref: https://github.com/grpc-ecosystem/go-grpc-middleware/blob/v2/providers/openmetrics/server_metrics.go
//		https://github.com/grpc-ecosystem/go-grpc-middleware/blob/v2/providers/openmetrics/server_options.go

var requestDurationBucket = tally.DurationBuckets{100 * time.Millisecond, 300 * time.Millisecond, 1200 * time.Millisecond, 5000 * time.Millisecond, 10000 * time.Millisecond}

// Total number of rpc call started on the server.
func grpcServerRpcCallTotal(m metrics.Metric, method string) {
	m.Tagged(metrics.Tags{"method": method}).Counter("grpc_server_rpc_total").Inc(1)
}

// Total number of rpc call failed.
func grpcServerRpcErrTotal(m metrics.Metric, method string) {
	m.Tagged(metrics.Tags{"method": method}).Counter("grpc_server_rpc_failed_total").Inc(1)
}

// Histogram of response latency (seconds) of gRPC that had been application-level handled by the server.
func grpcServerHandlingSecondsCount(m metrics.Metric, method string, val time.Duration) {
	m.Tagged(metrics.Tags{"method": method}).
		Histogram("grpc_server_handling_seconds_count", requestDurationBucket).
		RecordDuration(val)
}

func New(m metrics.Metric) *MetricMiddleware {
	return &MetricMiddleware{m: m}
}

var _ lava.Middleware = (*MetricMiddleware)(nil)

type MetricMiddleware struct {
	m metrics.Metric
}

func (m MetricMiddleware) String() string { return "metric" }

func (m MetricMiddleware) Middleware(next lava.HandlerFunc) lava.HandlerFunc {
	return func(ctx context.Context, req lava.Request) (rsp lava.Response, gErr error) {
		if req.Kind() == "http" || strings.Contains(req.Operation(), " ") {
			return next(ctx, req)
		}

		now := time.Now()
		grpcServerRpcCallTotal(m.m, req.Operation())

		defer func() {
			if !generic.IsNil(gErr) {
				grpcServerRpcErrTotal(m.m, req.Operation())
			}

			grpcServerHandlingSecondsCount(m.m, req.Operation(), time.Since(now))
		}()

		return next(ctx, req)
	}
}
