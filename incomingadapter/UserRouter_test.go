package incomingadapter

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	Register(c)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "APPROVED", w.Body.String())
}
