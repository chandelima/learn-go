package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido."
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		println(diaDaSemana)
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
		println(diaDaSemana)
	case numero == 3:
		diaDaSemana = "Terça-Feira"
		println(diaDaSemana)
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
		println(diaDaSemana)
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
		println(diaDaSemana)
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
		println(diaDaSemana)
	case numero == 7:
		diaDaSemana = "Sábado"
		println(diaDaSemana)
	default:
		diaDaSemana = "Número inválido"
		println(diaDaSemana)
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	//dia := diaDaSemana(4)
	diaDaSemana2(1)
	//fmt.Println(dia)
	//fmt.Println(dia2)
}
