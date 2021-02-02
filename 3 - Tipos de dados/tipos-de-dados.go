package main

import (
	"errors"
	"fmt"
)

func main(){
	var numero int16 = 100
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias
	//INT32 = rune
	var numero3 rune = 12345
	fmt.Println(numero3)

	//BYTE = uint8
	var numero4 byte = 123
	fmt.Println(numero4)	

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)
	
	var numeroReal2 float64 = 12300000000.45
	fmt.Println(numeroReal2)	

	var str string = "Texto"
	fmt.Println(str)

	str2 :=  "Texto"
	fmt.Println(str2)	

	//devolve o numero do caracter na tabela ascii
	char :=  'B'
	fmt.Println(char)	

	//quando a variavel está vazio o compilador retorna o valor inicial do tipo string <vazio> int 0
	var texto string
	fmt.Println(texto)

	var num int
	fmt.Println(num)

	//o default do bool é false se não passar valor
	var boolenado1 bool = false
	fmt.Println(boolenado1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}