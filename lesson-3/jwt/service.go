package swagger

import (
	"github.com/azomio/courses/lesson4/pkg/grpc/user"
)

type Service struct {
	AuthGRPCClient user.UserClient
}
