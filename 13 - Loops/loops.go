package main

import (
	"fmt"
)


func main() {

	//i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }

	//fmt.Println(i)

	// for j := 0; j < 10; j ++ {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Caio", "Luciana", "Josano"}

	//se não quero imprimir o indice no lugar dele digito _ 
	for _, nome := range nomes {
		fmt.Println(indice, nome)
	}

	//se não quero imprimir o indice no lugar dele digito _ 
	for _, nome := range nomes {
		fmt.Println(nome)
	}	

	//imprime o codigo ascii da letra_ 
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
	}	
	
	//imprime a letra 
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}		

	usuario := map[string]string {
		"nome": "Josano",
		"Sobrenome": "Soares",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}	

}