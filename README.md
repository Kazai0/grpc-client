# Título do Projeto

Descrição curta do projeto.

## Tabela de Conteúdo

- [Instruções de Uso](#instruções-de-uso)
- [Recursos e Funcionalidades](#recursos-e-funcionalidades)
- [Autores](#autores)

## Instruções de Uso

- Verifique a instação passo do protoc na doc: https://grpc.io/docs/languages/go/quickstart/

- Após instalar e settar variavéis de ambiente, utilize o comando na raiz do projeto protoc --go_out=. --go-grpc_out=. proto/user_message.proto

- Se o arquivo proto for gerado com sucesso, entre no diretorio cmd/server e rode o comando go run server.go

- após o terceiro passo entre no diretorio cmd/server e rode o comando go run server.go

- faço uma chamada dentro do postman ou i insomnia e utilize esse curl de exeplo

curl --location --request GET 'http://localhost:3000/user' \
--header 'Content-Type: application/json' \
--data-raw '{
        "name": "Jean",
        "email": "jean@example.com"
}'

## Recursos e Funcionalidades

- Lista de recursos e funcionalidades principais.

## Autores

- Jean Franzen 