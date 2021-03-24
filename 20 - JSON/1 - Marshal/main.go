package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)


type cachorro struct {
	Nome  string `json:"nome`
	Raca  string `json:"raca`
	Idade uint   `json:"idade`
}

func main() {
	c := cachorro{"Pipoca", "Vira lata caramelho", 1}
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome" : "Pipoca", 
		"raca" : "Vira lata caramelho",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJSON)

	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))	
}