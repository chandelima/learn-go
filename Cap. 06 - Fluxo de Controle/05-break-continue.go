package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i == 3 {
			continue
			/*
				Diferentemente do break, que quebra todas as próxi-
				mas execuções do for, o continue quebra somente a
				execução do for atual após a sua declaração e dá se-
				quência as demais iterações
			*/
		}
		fmt.Println("i =", i)
	}
}
