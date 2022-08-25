package main

import "fmt"

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)

	/*
		- O defer é chamado por último. Ou antes da linha de
		  encerramento da função "}", ou antes de um return;
		- Se vc tiver vários defers, o que é chamado primero
		  é executado por último.
	*/
}
