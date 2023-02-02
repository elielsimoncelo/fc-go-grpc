# fc-go-grpc

## Formas de comunicação

> Unária ou Unary

- Client: (1) Request
- Server: (1) Response

> Server streaming

- Client: (1) Request
- Server: (N) Response -> Response -> Response
- Você vai recebendo várias respostas, vai recebendo parcial

> Client streaming

- Cliente: (N) Request -> Request -> Request
- Server: (1) Response

> Bi directional streaming

- Cliente: (N) Request -> Request -> Request
- Server: (N) Response -> Response -> Response

## Configurando

> Instalar o compilador do protobuf

- <https://grpc.io/docs/protoc-installation>

> Instalar os plugins para o golang

- go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

> Baixar as dependências

- go mod tidy

> Criar os arquivos de gRPC

- protoc --go_out=. --go-grpc_out=. proto/course_category.proto
- go mod tidy

## Testando o projeto

> Instalando o cliente Evans

- <https://github.com/ktr0731/evans/blob/master/README.md>

> Executando o evans

- evans -r repl
- service CategoryService
- call CreateCategory
