package app

import (
	"testing"
)

var service = AuthorizationServiceImpl{}

func TestHello(t *testing.T) {
	expected := "world"
	if got := service.Hello(); got != expected {
		t.Errorf("Hello() = %q, but expected = %v", got, expected)
	}
}
