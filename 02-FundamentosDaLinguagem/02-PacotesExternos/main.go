package main

import (
	"modulo/auxiliar"
	"fmt"
	"github.com/badoux/checkmail"
)

func main()  {
	fmt.Println("Escrevendo do Arquivo main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("lucasmacedo9@hotmail.com")
	fmt.Println(erro)
}
