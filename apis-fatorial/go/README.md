# Compilation.hell Go

Este repo você irá encontrar exemplos de apis da forma mais simples desenvolvidas em Go que calcula Fatorial de um número.

Go é uma linguagem poderosa quando o cenário trata-se de apis, microservices, comunicação TCP, UPD e muito mais.

Para aprender mais sobre a lang existe hoje milhares de sites e blogs com conteúdos fantásticos. Aqui é o site oficial da Lang. [golang.org](https://golang.org)

Em Go você não precisa instalar nada no Server onde irá hospedar seu sistema em Go, basta levar ou subir o binário que você irá gerar.
Go é uma linguagem de programação Compilada.

Basta o simples comando abaixo para compilar seu programa em Go.

### Compilar

Você pode compilar seu programa e executar o binário

```bash

$ go build
$ ./framework-fiber

```

### Run

Você poderá simplesmente executar sem precisar compilar só para testar.

```bash

$ go run main.go

```

### Curl

Para fazer as requisições vocẽ poderá utilizar o cURL ou postman, o exemplo abaixo utilizando cURL.

```bash

$ curl -i -XGET localhost:8080/api/v1/fatorial/10

```