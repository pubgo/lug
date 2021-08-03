// Code generated by protoc-gen-lug. DO NOT EDIT.
// source: example/proto/login/code.proto

package login

import (
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	fb "github.com/pubgo/lug/builder/fiber"
	"github.com/pubgo/lug/pkg/gutil"
	"github.com/pubgo/lug/plugins/grpcc"
	"github.com/pubgo/lug/xgen"
	"github.com/pubgo/xerror"
)

var _ = strings.Trim
var _ = utils.UnsafeString
var _ fiber.Router = nil
var _ = gutil.MapFormByTag
var _ = fb.Cfg{}

func GetCodeClient(srv string, opts ...func(cfg *grpcc.Cfg)) func(func(cli CodeClient)) error {
	client := grpcc.GetClient(srv, opts...)
	return func(fn func(cli CodeClient)) (err error) {
		defer xerror.RespErr(&err)

		c, err := client.Get()
		if err != nil {
			return xerror.WrapF(err, "srv: %s", srv)
		}

		fn(&codeClient{c})
		return
	}
}

func init() {
	var mthList []xgen.GrpcRestHandler

	mthList = append(mthList, xgen.GrpcRestHandler{
		Service:      "login.Code",
		Name:         "SendCode",
		Method:       "POST",
		Path:         "/user/code/send-code",
		ClientStream: "False" == "True",
		ServerStream: "False" == "True",
		DefaultUrl:   "False" == "True",
	})

	mthList = append(mthList, xgen.GrpcRestHandler{
		Service:      "login.Code",
		Name:         "Verify",
		Method:       "POST",
		Path:         "/user/code/verify",
		ClientStream: "False" == "True",
		ServerStream: "False" == "True",
		DefaultUrl:   "False" == "True",
	})

	mthList = append(mthList, xgen.GrpcRestHandler{
		Service:      "login.Code",
		Name:         "IsCheckImageCode",
		Method:       "POST",
		Path:         "/user/code/is-check-image-code",
		ClientStream: "False" == "True",
		ServerStream: "False" == "True",
		DefaultUrl:   "False" == "True",
	})

	mthList = append(mthList, xgen.GrpcRestHandler{
		Service:      "login.Code",
		Name:         "VerifyImageCode",
		Method:       "POST",
		Path:         "/user/code/verify-image-code",
		ClientStream: "False" == "True",
		ServerStream: "False" == "True",
		DefaultUrl:   "False" == "True",
	})

	mthList = append(mthList, xgen.GrpcRestHandler{
		Service:      "login.Code",
		Name:         "GetSendStatus",
		Method:       "POST",
		Path:         "/user/code/get-send-status",
		ClientStream: "False" == "True",
		ServerStream: "False" == "True",
		DefaultUrl:   "False" == "True",
	})

	xgen.Add(reflect.ValueOf(RegisterCodeServer), mthList)
	xgen.Add(reflect.ValueOf(RegisterCodeRestServer), nil)
	xgen.Add(reflect.ValueOf(RegisterCodeHandler), nil)
}

func RegisterCodeRestServer(app fiber.Router, server CodeServer) {
	xerror.Assert(app == nil || server == nil, "app is nil or server is nil")

	// restful
	app.Add("POST", "/user/code/send-code", func(ctx *fiber.Ctx) error {
		var req = new(SendCodeRequest)
		if err := ctx.BodyParser(req); err != nil {
			return xerror.Wrap(err)
		}

		var resp, err = server.SendCode(ctx.UserContext(), req)
		if err != nil {
			return xerror.Wrap(err)
		}

		return xerror.Wrap(ctx.JSON(resp))
	})

	// restful
	app.Add("POST", "/user/code/verify", func(ctx *fiber.Ctx) error {
		var req = new(VerifyRequest)
		if err := ctx.BodyParser(req); err != nil {
			return xerror.Wrap(err)
		}

		var resp, err = server.Verify(ctx.UserContext(), req)
		if err != nil {
			return xerror.Wrap(err)
		}

		return xerror.Wrap(ctx.JSON(resp))
	})

	// restful
	app.Add("POST", "/user/code/is-check-image-code", func(ctx *fiber.Ctx) error {
		var req = new(IsCheckImageCodeRequest)
		if err := ctx.BodyParser(req); err != nil {
			return xerror.Wrap(err)
		}

		var resp, err = server.IsCheckImageCode(ctx.UserContext(), req)
		if err != nil {
			return xerror.Wrap(err)
		}

		return xerror.Wrap(ctx.JSON(resp))
	})

	// restful
	app.Add("POST", "/user/code/verify-image-code", func(ctx *fiber.Ctx) error {
		var req = new(VerifyImageCodeRequest)
		if err := ctx.BodyParser(req); err != nil {
			return xerror.Wrap(err)
		}

		var resp, err = server.VerifyImageCode(ctx.UserContext(), req)
		if err != nil {
			return xerror.Wrap(err)
		}

		return xerror.Wrap(ctx.JSON(resp))
	})

	// restful
	app.Add("POST", "/user/code/get-send-status", func(ctx *fiber.Ctx) error {
		var req = new(GetSendStatusRequest)
		if err := ctx.BodyParser(req); err != nil {
			return xerror.Wrap(err)
		}

		var resp, err = server.GetSendStatus(ctx.UserContext(), req)
		if err != nil {
			return xerror.Wrap(err)
		}

		return xerror.Wrap(ctx.JSON(resp))
	})

}
