package requestID

import (
	"context"

	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pubgo/lava/pkg/httpx"
	"github.com/pubgo/lava/pkg/utils"
	"github.com/pubgo/lava/plugin"
	"github.com/pubgo/lava/service"
	"github.com/pubgo/lava/service/service_type"
)

const Name = "x-request-id"

func init() {
	plugin.RegisterMiddleware(Name, func(next service_type.HandlerFunc) service_type.HandlerFunc {
		return func(ctx context.Context, req service_type.Request, resp func(rsp service_type.Response) error) (gErr error) {
			defer func() {
				switch err := recover().(type) {
				case nil:
				case error:
					gErr = err
				default:
					gErr = status.Errorf(codes.Internal, "service=>%s, endpoint=>%s, msg=>%v", req.Service(), req.Endpoint(), err)
				}
			}()

			rid := utils.FirstNotEmpty(
				func() string { return getReqID(ctx) },
				func() string { return service.HeaderGet(req.Header(), Name) },
				func() string { return ksuid.New().String() },
			)

			req.Header().Set(httpx.HeaderXRequestID, rid)
			return next(WithReqID(ctx, rid), req, resp)
		}
	})
}
