package incomingadapter

import (
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.String(200, "APPROVED")
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResiterResponse struct {
	ResultCode   string
	ErrorMessage string
}
