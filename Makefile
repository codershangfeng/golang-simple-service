# Proto file name 
PROTO?=""

.PHONY: go
go:
	protoc --go_out=plugins=grpc:.  $(PROTO)

.PHONY: tidy
tidy:
	cd src && go mod tidy