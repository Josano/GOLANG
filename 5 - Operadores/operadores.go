package main

import (
	"fmt"
)

func main() {

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2
	
	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	var soma2 int16 = numero1 + numero2
	fmt.Println(soma2)

	//FIM ARITMETICOS

	//ATRIBUICAO
	var variavel1 string = "String"
	variavel2 := "String"

	fmt.Println(variavel1, variavel2)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//OPERADORES LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//OPERADORES UNARIOS
	numero := 10
	numero++
	numero += 15
	numero--
	numero -=5
	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)


}