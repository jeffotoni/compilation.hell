# Compilation.hell

Este repo tem como objetivo mostrar algumas formas mais simples de desenvolver apis rEST em algumas linguagens de programação.
É somente uma inicitiva meramente didática, para que possamos criar apis em diversas linguagens de programação da forma mais simples possível.


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