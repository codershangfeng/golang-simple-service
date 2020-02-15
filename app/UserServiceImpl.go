package app

import (
	incomingport "github.com/codershangfeng/grpcservice/incomingport"
)

type UserServiceImpl struct {
}

func (UserServiceImpl) register(incomingRegisterRequest *incomingport.IncomingRegisterRequest) (incomingport.IncomingRegisterResponse, error) {
	return incomingport.IncomingRegisterResponse{ResultCode: "APPROVED"}, nil
}
