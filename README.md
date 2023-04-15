# 使用 goctl 创建项目

## 创建 api 服务

1. 编写 xx.api 文件
2. 运行 `goctl api go -api xx.api -dir . -style gozero` 自动生成项目目录
3. 运行 `go mod tidy` 自动寻找依赖，写到 go.mod 中，并进行依赖下载
4. 在 `internal/logic` 下编写具体的逻辑

或者：直接运行 `goctl api new api` 生成目录，再修改生成的 api 文件

## 创建 rpc 服务

1. 运行命令 `goctl rpc new rpc`，得到文件模板
2. 修改 proto 文件

## 测试

1. 构建 rpc 服务
2. 使用 http 的方式调用