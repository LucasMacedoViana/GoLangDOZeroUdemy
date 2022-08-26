package main

import "fmt"

type usuario struct {
	nome 		string
	idade 		uint8
	endereco 	endereco
}
type endereco struct {
	rua 	string
	numero	uint8
}
func main()  {
	var u usuario
	u.nome = "Lucas"
	u.idade = 30
	fmt.Println(u)

	endereco1 := endereco{"rua dos bobos", 0 }
	u2 := usuario{"Mateus", 26, endereco1}
	fmt.Println(u2)

	u3 := usuario{nome:"Davi"} // sem colocar o campo  idade
	fmt.Println(u3)

	u4 := usuario{idade:12} // sem colocar o campo nome
	fmt.Println(u4)



}
