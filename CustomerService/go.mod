module customerservice

go 1.22.7

toolchain go1.23.1

//protoc --go-grpc_out=. pbCustomer/customer.proto

require (
	golang.org/x/net v0.31.0
	google.golang.org/grpc v1.68.0
	google.golang.org/protobuf v1.35.2
)

require (
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241113202542-65e8d215514f // indirect
)
