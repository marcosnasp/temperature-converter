## Estrutura do projeto

Estrutura do projeto, seguindo as convenções do Go e boas práticas para serviços gRPC.

```sh
temperature-converter/
├── proto/
│   ├── temperature.proto
│   └── gen/
│       └── go/
│           ├── temperature.pb.go
│           └── temperature_grpc.pb.go
├── cmd/
│   ├── server/
│   │   └── main.go
│   └── client/
│       └── main.go
├── internal/
│   └── service/
│       └── temperature.go
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

Comando para a geração dos arquivos temperature_grpc.pb.go, e temperature.pb.go:

```console
protoc --proto_path=proto \
       --go_out=. --go_opt=module=github.com/marcosnasp/temperature-converter \
       --go-grpc_out=. --go-grpc_opt=module=github.com/marcosnasp/temperature-converter \
       proto/temperature.proto
```

Para execução:

Execute o servidor:

```console
go run cmd/server/main.go
```
Execute o cliente:

```console
go run cmd/client/main.go 0 celsius kelvin
```




