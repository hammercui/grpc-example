>gRPC的使用示例


## 生成grpc
>protoc-gen-go使用1.3x版本
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```


执行生成命令
```
make proto
```

## FAQ

1 mustEmbedUnimplemented*** method appear in grpc-server #3794

参见[issues3794](https://github.com/grpc/grpc-go/issues/3794),解決方案
```
protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. api.proto
```
