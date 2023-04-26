.PHONY: gen-go
gen-go:
	goctl api go --api ./api/app.api --dir ./api

.PHONY: protoc
protoc:
	goctl rpc protoc --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc ./rpc/app.proto

.PHONY: model
model:
	goctl model mysql ddl --database test -d "sql/model" --src "sql/*.sql" -c



.PHONY: swagger
swagger: genFile := ./swagger/app.swagger.json
swagger:
	goctl api plugin -plugin goctl-openapi3="openapi -filename $(genFile)" -api api/app.api -dir .
	sed -i '' s'/^  "components": {/  "components": {\n    "securitySchemes": {"bearerAuth": {"type": "http","scheme": "bearer","bearerFormat": "JWT"}},/' ./swagger/app.swagger.json
	sed -i '' s'/^  "components": {/  "security": [{"bearerAuth":[]}],\n  "components": {/' ./swagger/app.swagger.json
