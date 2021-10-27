package errno

import (
	"context"
	"fmt"
)

type Error struct {
	errno   int
	error   error
	message string

	record bool
	metric bool
}

const (
	Success        = 0
	InvalidRequest = 1
	UnknownError   = 3
	ServicePanic   = 4
)

var msg = map[int]string{
	Success:        "ok",
	InvalidRequest: "invalid request",
	UnknownError:   "unknown error",
	ServicePanic:   "service panic",
}

func (e Error) String() string {
	if e.error == nil {
		return fmt.Sprintf("coupon errno: %d, message:%s", e.errno, e.message)
	}
	return fmt.Sprintf("coupon errno: %d, message:%s, error:%v", e.errno, e.message, e.error)
}

func (e Error) Errno() int {
	return e.errno
}

func (e Error) Error() string {
	return e.String()
}

func (e Error) Metric() {
	if e.metric {
		return
	}
	e.metric = true
	// TODO metric
	fmt.Println("METRIC:", "coupon", "errno:", e.errno)
}

func (e Error) Record(ctx context.Context) Error {
	if e.record {
		return e
	}
	e.record = true
	// TODO LOG
	fmt.Println("LOG:", e)
	return e
}

func New(errno int, error error) Error {
	return Error{
		errno:   errno,
		error:   error,
		message: Msg(errno),
		record:  false,
		metric:  false,
	}
}

func Msg(errno int) string {
	if msg, exists := msg[errno]; exists {
		return msg
	}
	return "unknown errno"
}
