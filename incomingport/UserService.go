package incomingport

// UserService handle user request, should be implemented by app layer
type UserService interface {
	register(incomingRegisterRequest *IncomingRegisterRequest) (IncomingRegisterResponse, error)
}

type IncomingRegisterRequest struct {
	Username string
	Password string
}

type IncomingRegisterResponse struct {
	ResultCode string
}
