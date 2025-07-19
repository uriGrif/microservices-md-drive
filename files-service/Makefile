protoc:
	mkdir -p protogen/golang && cd proto && protoc --go_out=../protogen/golang --go_opt=paths=source_relative \
	--go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative \
	./*.proto

build:
	go build -o ./bin/server ./cmd/server/main.go