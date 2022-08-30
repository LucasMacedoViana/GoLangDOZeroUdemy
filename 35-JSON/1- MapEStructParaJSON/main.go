package main

import (
	"bytes" // pacote para converter o slice de bytes do JSON
	"encoding/json"
	"fmt"
	"log"

)

type cachorro struct {
	Nome string `json:"nome"` // falando qual é a chave que vai ser representada quando for convertida para json
	Raca string `json:"raca"`
	idade uint	`json:"idade"`
}
func main()  {
	c := cachorro{"Rex,", "Dalmata", 3 }
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil{
		log.Fatalln(erro)
	}
	fmt.Println(cachorroEmJSON) // Assim ele vai retornar um slice de byte ou uint8
	fmt.Println(bytes.NewBuffer(cachorroEmJSON)) // assim ele devolve corretamente o JSON

	c2 := map[string]string{
		"nome": "Bento",
		"raca": "Golden",
	}
	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil{
		log.Fatalln(erro)
	}
	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
