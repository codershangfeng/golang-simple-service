package app

import (
	incomingport "github.com/codershangfeng/grpcservice/incomingport"
	"testing"
)

var service = UserServiceImpl{}

func TestUserRegisterShouldSuccess(t *testing.T) {
	want := incomingport.IncomingRegisterResponse{ResultCode: "APPROVED"}
	if got, _ := service.register(&incomingport.IncomingRegisterRequest{}); got != want {
		t.Errorf("Got = %v, but want = %v", got, want)
	}
}
