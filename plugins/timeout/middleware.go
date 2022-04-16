package timeout

import (
	"context"
	"github.com/pubgo/lava/abc"
	"net/http"
	"time"

	"github.com/pubgo/lava/consts"
	"github.com/pubgo/lava/errors"
	"github.com/pubgo/lava/pkg/httpx"
	"github.com/pubgo/lava/plugin"
)

const Name = "middleware.timeout"

func init() {
	plugin.RegisterMiddleware(Name, func(next abc.HandlerFunc) abc.HandlerFunc {
		var defaultTimeout = consts.DefaultTimeout
		return func(ctx context.Context, req abc.Request, resp func(rsp abc.Response) error) error {
			// 过滤 websocket 请求
			// 过滤 stream
			if httpx.IsWebsocket(http.Header(req.Header())) || req.Stream() {
				return next(ctx, req, resp)
			}

			// 从header中获取超时设置
			//	key: x-request-timeout
			if t := abc.HeaderGet(req.Header(), "X-REQUEST-TIMEOUT"); t != "" {
				var dur, err = time.ParseDuration(t)
				if dur != 0 && err == nil {
					defaultTimeout = dur
				}
			}

			if _, ok := ctx.Deadline(); !ok {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, defaultTimeout)
				defer cancel()
			}

			var err error
			var done = make(chan struct{})
			go func() {
				defer func() {
					switch c := recover().(type) {
					case nil:
					case error:
						err = c
					default:
						err = errors.Internal("timeout", "service=>%s, endpoint=>%s, msg=>%v", req.Service(), req.Endpoint(), err)
					}
					close(done)
				}()

				err = next(ctx, req, resp)
			}()

			select {
			case <-ctx.Done():
				return errors.DeadlineExceeded("timeout", req.Endpoint())
			case <-done:
				return err
			}
		}
	})
}
