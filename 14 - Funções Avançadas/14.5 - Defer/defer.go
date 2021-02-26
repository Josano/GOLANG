package main

import "fmt"

func funcao1() {
	fmt.Println("executando a função 1")
}

func alunoEstaoAprovado(n1, n2 float32) bool {

	defer fmt.Println("Média calculada, Resultado será retornado.")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
} 

func funcao2() {
	fmt.Println("executando a função 2")
}

func main() {
	fmt.Println(alunoEstaoAprovado(7,8))
}