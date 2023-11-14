## grpc-demo

一个简单的grpc示例

需要安装：
- protoc
- protoc-gen-go
- protoc-gen-go-grpc

**生成.pb.go文件**

```sh
protoc --go_out=. ./proto/*.proto
```

**生成_grpc.pb.go文件**

```sh
protoc --go-grpc_out=. ./proto/*.proto
```