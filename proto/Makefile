generate:
	rm -rf ./gen/go/* && rm -rf ../fe/src/lib/proto/*
	protoc -I=. --go-grpc_out=./gen/ --go_out=./gen/ **/**/*.proto
	protoc-go-inject-tag -input="./gen/go/**/**/*.pb.go"
	npx protoc --ts_opt generate_dependencies --ts_out ../fe/src/lib/proto/ --proto_path ./ ./main.prot

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
	npm i

