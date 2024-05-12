# 实用工具

jsonlint 检查json格式
> https://jsonlint.com/

wtools JSON escape等操作
> https://wtools.io/json-escape-unescape

Open API 可视化工具
> https://editor.swagger.io/

查看jwt
> https://jwt.io/

Base64编解码
> https://base64.us/

RSA加解密
> https://www.bejson.com/enc/rsa/#google_vignette

缩小字体文件
> https://blog.csdn.net/qq_32228265/article/details/120525130

图片压缩
> https://tinypng.com/

Github徽记
> https://shields.io/

React-Markdown
> https://blog.csdn.net/m0_57707788/article/details/133156951

pb code gen
> Download the protoc bin file.
> https://github.com/protocolbuffers/protobuf/releases
```shell
# protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go
# grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

go get -u google.golang.org/protobuf
go get -u google.golang.org/grpc
```
```shell
protoc -I . --go_out=plugins=grpc:. *.proto
```

wire
```shell
# 安装
go install github.com/google/wire/main/wire@latest
go get -u github.com/google/wire/main/wire@v0.5.0
# 生成
cd internal/app
go run github.com/google/wire/main/wire
# 生成
wire
```

mysql：
```shell
mysql -u root -p
```

代码覆盖率检查：
```shell
# 查看pwd下所有go文件代码覆盖率
go test -cover
# 查看pwd下所有go文件代码覆盖率，并输出覆盖率报告文件unit_test.out
go test -cover -coverprofile=unit_test.out
# 用html方式读取报告文件unit_test.out，可查看具体是哪段代码没有覆盖
go tool cover -html=unit_test.out
```
