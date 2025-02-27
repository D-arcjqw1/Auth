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
	if not exist "pkg\note_v1" mkdir "pkg\note_v1"
	protoc --proto_path=api/note_v1 --go_out=pkg/note_v1 --go_opt=paths=source_relative --plugin=protoc-gen-go="$(LOCAL_BIN)/protoc-gen-go.exe" --go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative --plugin=protoc-gen-go-grpc="$(LOCAL_BIN)/protoc-gen-go-grpc.exe" api/note_v1/note.proto
