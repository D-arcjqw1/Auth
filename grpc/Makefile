LOCAL_BIN := bin

install-deps:
	if not exist "$(LOCAL_BIN)" mkdir "$(LOCAL_BIN)"
	PowerShell -Command "$$env:GOBIN='$(CURDIR)/$(LOCAL_BIN)'; go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1"
	PowerShell -Command "$$env:GOBIN='$(CURDIR)/$(LOCAL_BIN)'; go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2"

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-note-api

generate-note-api:
	if not exist "pkg\user" mkdir "pkg\user"
	protoc --proto_path=api/user --go_out=pkg/user --go_opt=paths=source_relative --plugin=protoc-gen-go="$(LOCAL_BIN)/protoc-gen-go.exe" --go-grpc_out=pkg/user --go-grpc_opt=paths=source_relative --plugin=protoc-gen-go-grpc="$(LOCAL_BIN)/protoc-gen-go-grpc.exe" api/user/user.proto

build:
ifeq ($(OS),Windows_NT)
	PowerShell -Command "$$env:GOOS='linux'; $$env:GOARCH='amd64'; go build -o service_linux cmd/grpc_server/main.go"
else
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/grpc_server/main.go
endif

copy-to-server:
	"C:/Windows/System32/OpenSSH/scp.exe" service_linux root@5.159.102.94:/root/

docker-build-and-push:
	
	docker login -u token -p CRgAAAAAoPFXPJo1DjutuLgR-V6qnuaJe7gMmVTL cr.selcloud.ru
	
	
	docker buildx build --platform linux/amd64 --no-cache -t cr.selcloud.ru/olezhek28/test-server:v0.0.1 .

	
	docker push cr.selcloud.ru/olezhek28/test-server:v0.0.1