package app

import (
	"github.com/codershangfeng/golang-simple-serivce/incomingport"
)

type AuthorizationServiceImpl struct {
}

func (service AuthorizationServiceImpl) Hello() string {
	return "world"
}

func (service AuthorizationServiceImpl) Authorize(incomingport.IncomingAuthorizationRequest) incomingport.IncomingAuthorizationResponse  {
	return new(incomingport.IncomingAuthorizationResponse)
}