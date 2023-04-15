.PHONY: gen-go
gen-go:
	goctl api go --api ./api/app.api --dir ./api

.PHONY: protoc
protoc:
	goctl rpc protoc --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc ./rpc/app.proto

.PHONY: model
model:
	goctl model mysql ddl --database test -d "sql/model" --src "sql/*.sql" -c