package main

import (
	"fmt"
)

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 = "lalala"
		variavel4 = "lelele"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável5", "Variável6"

	fmt.Println(variavel5, variavel6)

	const constante1 = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
	
}