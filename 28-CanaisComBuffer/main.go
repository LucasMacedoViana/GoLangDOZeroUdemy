package main

import "fmt"

func main()  {
	canal := make(chan string, 2) // dizendo a capacidade do meu canal
	canal <- "ola mundo"
	canal <- "Programando em GO!"

	mensagem :=  <- canal
	mensagem2 := <- canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
