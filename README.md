# Authorization-Service

gRPC auth service

## Generate new protobuf file

Проверить что есть модули для go

```bash
# Установка компилятора для стандартного Go-кода
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Установка компилятора для gRPC Go-кода
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

В каталоге `protos`.

```bash
protoc -I proto proto/sso/sso.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative
```

## Run app

Запуск приложения

```bash
go run cmd/sso/main.go --config=config/dev.yaml
```
