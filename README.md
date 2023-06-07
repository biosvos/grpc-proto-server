### 툴 설치

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
```

### build proto

```bash
protoc \
  --go_out=. \
  --go-grpc_out=. \
  --grpc-gateway_opt=logtostderr=true \
  --grpc-gateway_out=. \
  --grpc-gateway_opt=grpc_api_configuration=proto/service.yaml \
  proto/greeter.proto
```