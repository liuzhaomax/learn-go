# Microservices

1. RPC和gRPC
2. Protocol Buffer

## 1. RPC和gRPC

### 1.1 概念
本地调用远程服务器程序。传输采用序列化二进制，速度更快，体积更小。
协议一般使用TCP。可以打代码桩，通过桩函数直接调用服务函数，无需额外编程。

gRPC是一款go语言的RPC框架。传输协议使用HTTP2，速度更快。序列化采用ProtoBuf。

### 1.2 RPC四要素
+ server function
+ client function
+ server stub
+ client stub

### 1.3 gRPC的四种模式
+ 单一模式`unary`：客户端请求一次，服务端响应一次
+ 服务端流模式`server side stream`：客户端请求一次，服务端一直响应
+ 客户端流模式`client side stream`：客户端一直请求，服务端实时响应
+ 双向流模式`bidirectional stream`：客户端和服务端一直收发，用协程分开接收与发送，从而不会阻塞

## 2. Protocol Buffer

### 2.1 对比Json的优势
+ 跨语言
+ 跨平台
+ 传输速度快
+ 压缩体积小
+ 序列化二进制
+ 代码生成包含函数与数据结构

### 2.2 下载protobuf
>https://github.com/protocolbuffers/protobuf/releases

### 2.3 安装依赖
```shell
# protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go
# grpc
go get -u google.golang.org/grpc
# protobuf
go get -u google.golang.org/protobuf
```

### 2.4 官方文档
>http://github.com/protocolbuffers/protobuf

>http://developers.google.com/protocol-buffers

### 2.5 代码生成命令
```shell
protoc -I . --go_out=plugins=grpc:. *.proto
```
