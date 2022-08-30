package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"` // falando qual é a chave que vai ser representada quando for convertida para json
	Raca string `json:"raca"` // caso nao queira que aparece algum campo, basta colocar a tag com "-" e ai não vai ser identificado
	Idade uint	`json:"idade"`
}
func main()  {
	cachorroEmJSON := `{"nome":"Bento","raca":"Golden","idade":4}`

	// convertendo para struct
	var c cachorro
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil{
		log.Fatalln(erro)
	} // tem que passar com o ponteiro pois queremos que a variavel fique alterada
	// convertendo a string para o slice de byte []byte()
	fmt.Println(c)

	// convertendo para MAP
	cachorro2EmJSON := `{"nome":"Sushi","raca":"Chiuaua"}`
	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil{
		log.Fatalln(erro)
	}
	fmt.Println(c2)
}
