# learn-go

protobuf
```shell
# protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go
# grpc
go get -u google.golang.org/grpc
# protobuf
go get -u google.golang.org/protobuf
```

```shell
protoc -I . --go_out=plugins=grpc:. *.proto
```

wire
```shell
# 安装
go install github.com/google/wire/cmd/wire@latest
go get github.com/google/wire/cmd/wire
# 生成
cd internal/app
go run github.com/google/wire/cmd/wire
```