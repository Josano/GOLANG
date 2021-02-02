package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Arrays e Slices")

	var array1[5] string
	array1[0] = "Posicao 1"
	fmt.Println(array1)

	array2 := [5] string{"Possicao 1", "Possicao 2", "Possicao 3", "Possicao 4", "Possicao 5"}
	fmt.Println(array2)

	//nesse formato o array vai ter o tamanho de itens que receber, não é um tamanho dinâmico
	array3 := [...] int {1, 2, 3, 4, 5} 
	fmt.Println(array3)

	//Slice tem tamanho dinâmico mas tem limitação de tipo
	slice := [] int {10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)	

	array2[1] = "Posicao Alterada"
	fmt.Println(slice2)

	//Arrays Internos
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade

	

}