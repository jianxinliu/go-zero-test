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

# 总结

## http api

1. 使用命令直接创建所有文件
2. 编辑 api 文件，定义自己业务上的接口
3. 重新使用命令根据更新后的 api 文件生成新的代码
4. 在生成 logic 文件中编写具体的业务逻辑
5. 如果有添加其他 service（rpc 等），或者是 model ,则在 `servicecontext.go` 文件中将此添加到上下文里，方便各个 logic 取用
6. 如果需要新增接口，则更新 api 文件，再使用命令生成对应的 handler, logic 等文件。运行的脚本可以写在 makefile 里，方便每次运行

## rpc

1. 使用命令直接生成所有文件
2. 编辑 proto 文件定义接口
3. 在 logic 文件中写业务逻辑。只需要关注 logic 即可
4. rpc 本身会启动一个 rpcServer, 用于提供服务。也会提供一个程序包 client 供其他调用者调用，比较容易混淆
5. proto 生成的文件可以不用管