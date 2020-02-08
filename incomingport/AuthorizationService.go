package incomingport

import "time"

type AuthorizationService interface {
	Authorize(IncomingAuthorizationRequest) IncomingAuthorizationResponse
}

type IncomingAuthorizationRequest struct {
	UUID              string
	AuthorizationCode string
}

type IncomingAuthorizationResponse struct {
	ResultCode   ResultCode
	ErrorMessage string
	CreatedTime  time.Time
}

type ResultCode int

const (
	Approved ResultCode = iota
	Error
)

func (rc ResultCode) String() string {
	return [...]string{"APPROVED", "ERROR"}[rc]
}
