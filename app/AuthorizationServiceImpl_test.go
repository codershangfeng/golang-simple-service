package app

import (
	. "github.com/codershangfeng/grpcservice/incomingport"
	"testing"
)

var service = AuthorizationServiceImpl{}

func TestHello(t *testing.T) {
	expected := "world"
	if got := service.Hello(); got != expected {
		t.Errorf("Hello() = %q, but expected = %v", got, expected)
	}
}

func TestAuthorize(t *testing.T) {
	if got := service.Authorize(*new(IncomingAuthorizationRequest)); got.ResultCode != Approved || got.ErrorMessage != "" {
		t.Errorf("Got = %v", got)
	}
}
