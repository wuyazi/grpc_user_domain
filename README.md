> goctl rpc protoc --proto_path=.:$GOPATH/src --go_out=./user_domain --go_opt=paths=source_relative --go-grpc_out=./user_domain --go-grpc_opt=paths=source_relative --zrpc_out=. --style=go_zero ./grpc_user_domain.proto

> protoc --proto_path=.:$GOPATH/src --go_out=./user_domain --go_opt=paths=source_relative --go-grpc_out=./user_domain --go-grpc_opt=paths=source_relative ./event.proto

