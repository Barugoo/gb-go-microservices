package reqdata

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type contextKey string

var RequestIDContextKey = "X-Request-ID"

const RequestIDHeader = "X-Request-ID"

func GetRequestID(ctx context.Context) string {
	ridRaw, ok := ctx.Value(RequestIDContextKey).(string)
	if !ok {
		md, _ := metadata.FromIncomingContext(ctx)
		if ridSli := md.Get(RequestIDHeader); len(ridSli) > 0 {
			return ridSli[0]
		}
	}
	return string(ridRaw)
}
