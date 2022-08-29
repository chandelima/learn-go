package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá mundo") //goroutine
	escrever("Tô pra endoidar com esse erro!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
