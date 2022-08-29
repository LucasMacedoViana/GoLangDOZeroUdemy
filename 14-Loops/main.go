package main

import (
	"fmt"
	"time"
)

func main()  {
	// primeira maneira
	i := 0
	for i < 10 {
		i++
		fmt.Printf("Incrementando i %v \n", i)
	}
		fmt.Println(i)
	// iniciando a variavel no proprio for
	for j := 0; j <10; j++{
		fmt.Printf("Incrementando j %v \n", j)
	}
	//for com a clausula range
	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes{
		fmt.Println(indice, nome)
	}
	for _, nome := range nomes{ // sem passar o indice
		fmt.Println( nome)
	}
	// Pegando cada letra de uma string
	for indice, letra := range "PALAVRA"{
		fmt.Println(indice, string(letra))
	}
	// usando o for em um MAP
	usuario := map[string]string{
		"nome": "Leonardo",
		"sobrenome":"Silva",
	}
	for chave, valor := range usuario{
		fmt.Println(chave,valor)
	}
	// loop infinido
	for {
		fmt.Println("Infinito")
		time.Sleep(time.Second)
	}
}
