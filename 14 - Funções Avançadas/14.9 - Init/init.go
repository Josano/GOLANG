package main

import "fmt"

var n int

func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println(n)
	fmt.Println("Executando a função main")
}