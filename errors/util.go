package errors

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MaxCode [0,1000]为系统错误, 业务错误从1000之后开始定义
const MaxCode = 1000

// Err2GrpcCode
// converts a standard Go error into its canonical code. Note that
// this is only used to translate the error returned by the server applications.
func Err2GrpcCode(err error) codes.Code {
	switch err {
	case nil:
		return codes.OK
	case io.EOF:
		return codes.OutOfRange
	case io.ErrClosedPipe, io.ErrNoProgress, io.ErrShortBuffer, io.ErrShortWrite, io.ErrUnexpectedEOF:
		return codes.FailedPrecondition
	case os.ErrInvalid:
		return codes.InvalidArgument
	case context.Canceled:
		return codes.Canceled
	case context.DeadlineExceeded:
		return codes.DeadlineExceeded
	}

	switch {
	case os.IsExist(err):
		return codes.AlreadyExists
	case os.IsNotExist(err):
		return codes.NotFound
	case os.IsPermission(err):
		return codes.PermissionDenied
	}
	return codes.Unknown
}

func Http2GrpcCode(code int32) codes.Code {
	switch code {
	case http.StatusOK:
		return codes.OK
	case http.StatusBadRequest:
		return codes.InvalidArgument
	case http.StatusRequestTimeout:
		return codes.DeadlineExceeded
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusConflict:
		return codes.AlreadyExists
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusPreconditionFailed:
		return codes.FailedPrecondition
	case http.StatusNotImplemented:
		return codes.Unimplemented
	case http.StatusInternalServerError:
		return codes.Internal
	case http.StatusServiceUnavailable:
		return codes.Unavailable
	}

	return codes.Unknown
}

func IsGrpcAcceptable(err error) bool {
	switch status.Code(err) {
	case codes.DeadlineExceeded, codes.Internal, codes.Unavailable, codes.DataLoss:
		return false
	default:
		return true
	}
}

func IsMemoryErr(err error) bool {
	if err == nil {
		return false
	}

	return strings.Contains(err.Error(), "invalid memory address or nil pointer dereference")
}

// FromError try to convert an error to *Error.
// It supports wrapped errors.
func FromError(err error) *Error {
	if err == nil {
		return nil
	}

	if se := new(Error); errors.As(err, &se) {
		return se
	}

	// grpc error
	gs, ok := status.FromError(err)
	if ok {
		ret := New("lava.grpc.error", int32(gs.Code()), gs.Message())
		for _, detail := range gs.Details() {
			switch d := detail.(type) {
			case *errdetails.ErrorInfo:
				ret.Reason = d.Reason
				return ret.WithMetadata(d.Metadata)
			}
		}
		return ret
	}

	return &Error{
		Code:     2,
		Reason:   "lava.unknown.error",
		Message:  err.Error(),
		Metadata: map[string]string{"detail": fmt.Sprintf("%v", err)},
	}
}
