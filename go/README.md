# Compilation.hell Go

Este repo você irá encontrar alguns exemplos de apis feitas na linguagem de programação Go. Go é uma linguagem poderosa quando o cenário trata-se de apis, microservices, comunicação TCP, UPD e muito mais.

Para aprender mais sobre a lang existe hoje milhares de sites e blogs com conteúdos fantásticos. Aqui é o site oficial da Lang. [golang.org](https://golang.org)

Em Go você não precisa instalar nada no Server onde irá hospedar seu sistema em Go, basta levar ou subir o binário que você irá gerar.
Go é uma linguagem de programação Compilada.

Basta o simples comando abaixo para compilar seu programa em Go.
```bash

$ go build

```

## Passos para instalar

Para instalar em ambiente Linux basta fazer download no site oficial e depois rodar o comando.
[Go Download](https://dl.google.com/go/go1.15.2.linux-amd64.tar.gz)

Instalando A lang Go
```bash

$ tar -C /usr/local -xzf go1.15.2.linux-amd64.tar.gz

```

Aqui iremos definir em seu PATH para que possa reconhecer Go
```bash

$ export PATH=$PATH:/usr/local/go/bin

```

Voila, prontinho linguagem instalada e prontinha para brincarmos.
```bash

$ go version

```

Em Go conseguimos visualizar **4 Pilares** bem Claros:
```bash

- Simplicidade
- Legibilidade
- Produtividade
- Retrocompatibilidade

```

Neste repo você irá encontrar Go usando de sua stand library default e deixamos também alguns frameworks como fiber e atreugo ambos extremamentes otimizados para alto desempenho.

## Resumo da Linguagem de programação Go 

Go foi projetado pelo Google em 2007 para melhorar a produtividade de programação em uma era de multicore, rede de máquinas e grandes bases de código. Os designers queriam abordar críticas de outras línguagens em uso no Google, mas manter suas características úteis. Os criadores Rob Pike, Ken Thompson e Robert Griesemer mantiveram a sintaxe de Go semelhante ao C. No final de 2008 Russ Cox juntou-se a equipe e ajudou a mudar a linguagem e as bibliotecas de protótipo para realidade.

A linguagem Go foi lançada em 2009 com propósito de facilitar a resolução de problemas quando o assunto é desenvolvimento em camadas de rede, escalabilidade, desempenho, produtividade e o mais importante, concorrência. O próprio Rob Pike declarou que “Go foi projetado para tratar de um conjunto de problemas de engenharia de software a que estávamos expostos na construção de grandes softwares de servidor”.


Um simples exemplo de uma api rEST, escutando na porta 8080

```go
/// curl localhost:8080/hello
//////////////////////////////////////

package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":8080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello, welcome to the world, Go!"))
		}))
}
```
