package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado.")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado.")

	var mediaAprovacao float32 = 6.0
	mediaObtida := (n1 + n2) / 2

	if mediaObtida >= mediaAprovacao {
		return true
	}

	return false
}

func main() {
	defer funcao1() //DEFER = ADIAR (É executado por último, ou antes do return)
	funcao2()

	situacaoAluno := alunoEstaAprovado(7, 8)
	fmt.Println(situacaoAluno)
}
