package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Segunda-feira"
	case 2:
		return "Terça-feira"
	case 3:
		return "Quarta-feira"
	case 4:
		return "Quinta-feira"
	case 5:
		return "Sexta-feira"
	case 6:
		return "Sábado"
	case 7:
		return "Domingo"
	default:
		return "número invalido"															
	}
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Segunda-feira"
	case numero == 2:
		return "Terça-feira"
	case numero == 3:
		return "Quarta-feira"
	case numero == 4:
		return "Quinta-feira"
	case numero == 5:
		return "Sexta-feira"
	case numero == 6:
		return "Sábado"
	case numero == 7:
		return "Domingo"
	default:
		return "número invalido"															
	}
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(5)
	fmt.Println(dia)


}