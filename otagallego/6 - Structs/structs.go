package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u1 usuario
	u1.nome = "Davi"
	u1.idade = 21
	fmt.Println(u1)

	u1e := endereco{"Rua dos Bobos", 0}
	u1.endereco = u1e
	fmt.Println(u1)
}
