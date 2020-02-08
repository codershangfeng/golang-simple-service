# Proto file name 
PROTO?=""

.PHONY: go
go:
	protoc --go_out=plugins=grpc:.  $(PROTO)

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	go fmt ./...