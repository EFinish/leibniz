generate:
	rm -rf ./gen/*
	protoc -I=. --go-grpc_out=./gen/ --go_out=./gen/ **/**/*.proto
	protoc-go-inject-tag -input="./gen/go/**/**/*.pb.go"

lint:
	protoc --lint_out=. *.proto

install:
	go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/ckaznocha/protoc-gen-lint \
		github.com/favadi/protoc-go-inject-tag
	sudo apt install protobuf-compiler
