/*
- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
- Solução: https://play.golang.org/p/GBZcnu0Ajp
*/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
}

func (p pessoa) demonstra() {
	fmt.Println("O nome completo da pessoa é:", p.nome, p.sobrenome)
	fmt.Println("E a idade dela é:", p.idade)
}

func main() {
	p1 := pessoa{"Alexandre", "Lima", 30}
	p1.demonstra()
}
