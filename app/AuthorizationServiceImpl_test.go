package app

import (
	"testing"
	. "github.com/codershangfeng/golang-simple-serivce/incomingport"
)

var service = AuthorizationServiceImpl{}

func TestHello(t *testing.T) {
	expected := "world"
	if got := service.Hello(); got != expected {
		t.Errorf("Hello() = %q, but expected = %v", got, expected)
	}
}

func TestAuthorize(t *testing.T) {
	got := service.Authorize(new(IncomingAuthorizationRequest)); 

	t.Errorf("Test result: %v", got)
}

