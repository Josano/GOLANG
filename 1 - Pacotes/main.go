package main

import (
	"modulo/auxiliar"
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	
	erro := checkmail.ValidateFormat("josano@hotmail.com")
	fmt.Println(erro)
}