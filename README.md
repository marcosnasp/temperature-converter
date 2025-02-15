## Estrutura do projeto

Estrutura do projeto, seguindo as convenções do Go e boas práticas para serviços gRPC.

temperature-converter/
├── proto/                       # Pasta para arquivos Protobuf
│   ├── temperature.proto        # Definição do serviço e mensagens
│   └── gen/                     # Códigos gerados pelo protoc
│       └── go/                  # Arquivos Go gerados (não versionados)
│           ├── temperature.pb.go
│           └── temperature_grpc.pb.go
├── cmd/                         # Pasta para executáveis
│   ├── server/                  # Servidor gRPC
│   │   └── main.go
│   └── client/                 # Cliente gRPC
│       └── main.go
├── internal/                    # Código interno do projeto
│   └── service/                # Implementação do serviço
│       └── temperature.go
├── go.mod                      # Módulo Go
├── go.sum                      # Dependências do módulo
├── Makefile                    # Automação de tarefas (opcional)
└── README.md                   # Documentação do projeto

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




