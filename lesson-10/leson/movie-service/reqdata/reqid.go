package reqdata

import (
	"context"
)

type contextKey string

var RequestIDContextKey = "X-Request-ID"

const RequestIDHeader = "X-Request-ID"

func GetRequestID(ctx context.Context) string {
	ridRaw, _ := ctx.Value(RequestIDContextKey).(string)
	return string(ridRaw)
}
