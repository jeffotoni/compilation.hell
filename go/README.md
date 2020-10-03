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



