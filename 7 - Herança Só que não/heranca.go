package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"João", "Pedro", 28, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Administração", "UNA"}
	fmt.Println(e1)
}