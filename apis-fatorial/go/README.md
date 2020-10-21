# Compilation.hell Go

Este repo você irá encontrar exemplos de apis da forma mais simples desenvolvidas em Go que calcula Fatorial de um número.

Go é uma linguagem poderosa quando o cenário trata-se de apis, microservices, comunicação TCP, UPD e muito mais.

Para aprender mais sobre a lang existe hoje milhares de sites e blogs com conteúdos fantásticos. Aqui é o site oficial da Lang. [golang.org](https://golang.org)

Em Go você não precisa instalar nada no Server onde irá hospedar seu sistema em Go, basta levar ou subir o binário que você irá gerar.
Go é uma linguagem de programação Compilada.

Basta o simples comando abaixo para compilar seu programa em Go ou pode utilizar o Makefile para facilitar ainda mais.

Você poderá utilizar o docker e desejar abaixo o manual bonitinho para facilitar.

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

### Utilizando Docker

Vamos criar sua imagem

```bash

$ make docker

```

Agora basta executar seu container e prontinho!

```bash

$ docker run --rm --name fatorial -p 8080:8080 jeffotoni/fatorial
 ┌───────────────────────────────────────────────────┐ 
 │                    Fiber v2.1.0                   │ 
 │               http://127.0.0.1:8080               │ 
 │                                                   │ 
 │ Handlers ............. 2  Threads ............. 4 │ 
 │ Prefork ....... Disabled  PID ................. 1 │ 
 └───────────────────────────────────────────────────┘ 

```

### Curl

Para fazer as requisições vocẽ poderá utilizar o cURL ou postman, o exemplo abaixo utilizando cURL.

```bash

$ curl -i -XGET localhost:8080/api/v1/fatorial/10

```