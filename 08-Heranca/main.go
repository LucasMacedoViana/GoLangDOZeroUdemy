package main

import "fmt"

type pessoa struct {
	nome 		string
	sobrenome 	string
	idade 		uint8
	altura 		uint8
}

type estudante struct {
	pessoa
	curso 		string
	faculdade 	string
	matricula	uint
}
func main()  {
	p1 := pessoa{"João", "Pedro", 30, 179}
	fmt.Println(p1)
	e1 := estudante{p1, "Engenharia", "usp", 00013 }
	fmt.Println(e1)
	fmt.Println(e1.nome)
	fmt.Println(e1.sobrenome)
}
