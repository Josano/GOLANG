package main

import (
	"fmt"
)

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo structs")

	//con variavel
	var u usuario
	u.nome = "Josano"
	u.idade = 43

	fmt.Println(u.nome, u.idade)
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua Francisco", 0}

	//com inferência de tipo
	usuario2 := usuario{"Caio", 12, enderecoExemplo}
	fmt.Println(usuario2)

	//passando apenas um valor
	usuario3 := usuario{nome: "Luciana"}
	fmt.Println(usuario3)	
}