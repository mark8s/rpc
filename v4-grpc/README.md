
## grpc

> http://liuqh.icu/2022/01/20/go/rpc/03-grpc-ru-men/

gRPC 是一个高性能、开源、通用的RPC框架，由Google推出，基于HTTP2协议标准设计开发，默认采用Protocol Buffers数据序列化协议，支持多种开发语言。gRPC提供了一种简单的方法来精确的定义服务，并且为客户端和服务端自动生成可靠的功能库。

### 前置准备

- 安装protoc
- 安装插件

验证安装

```shell
$ protoc --version
libprotoc 3.19.1

$ protoc-gen-go --version
protoc-gen-go v1.26.0

$ protoc-gen-go-grpc --version
protoc-gen-go-grpc 1.1.0
```

### 实现
- 编写proto文件
- 执行 `protoc --go-grpc_out=. --go_out=. hello.proto` 创建pb和grpc.pb文件
- 在生成的grpc文件中实现server接口，实现具体的方法logic
- 编写服务端和客户端代码


