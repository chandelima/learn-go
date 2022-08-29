package main

import (
	"errors"
	"fmt"
)

func main() {
	// Tipos de números inteiros
	// int8, int16, int32, int64
	// int (apenas) usa a arquitetura do PC como base
	var numero int = 100000
	fmt.Println(numero)

	// tipos unsigned int
	// não tem sinal, logo sempre são positivos
	// uint8, uint16, uint32, uint64
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias int32 = rune
	var numero3 rune = 12456
	fmt.Println(numero3)

	// ailas uint8 = byte
	var numero4 byte = 123
	fmt.Println(numero4)

	// Tipos de números reais
	// float32, float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float32 = 12554546543.45
	fmt.Println(numeroReal2)

	numeroReal3 := 213123.56
	fmt.Println(numeroReal3)

	// Strings
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	// Valor Zero (valor que é atribuído a uma variáel quando não for inicializada)
	var texto int16
	fmt.Println(texto)

	// Boolean
	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool
	fmt.Println(booleano2)

	// Error (Em go o erro é um tipo, não há exception)
	var erro error = errors.New("O erro que eu quero imprimir.")
	fmt.Println(erro)
}
