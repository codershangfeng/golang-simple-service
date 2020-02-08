package app

import (
	"time"

	. "github.com/codershangfeng/grpcservice/incomingport"
)

type AuthorizationServiceImpl struct {
}

func (service AuthorizationServiceImpl) Hello() string {
	return "world"
}

func (service AuthorizationServiceImpl) Authorize(IncomingAuthorizationRequest) IncomingAuthorizationResponse {
	return IncomingAuthorizationResponse{ResultCode: Approved, ErrorMessage: "", CreatedTime: time.Now()}
}
