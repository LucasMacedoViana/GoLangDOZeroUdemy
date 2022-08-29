package main

import (
	"fmt"
	"time"
)

func main()  {
	canal := make(chan string)

	go escrever("Olá mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")
	for mensagem := range canal{
		fmt.Println(mensagem)

	/*for{
		mensagem, aberto := <- canal // aqui o canal vai esperar receber um valor
		if !aberto{
			break
		}
		fmt.Println(mensagem)*/
	}
	fmt.Println("Fim do Programa")
}

func escrever(texto string, canal chan string)  {
	for i := 0; i < 5 ; i++{
		canal <- texto  //esta mandando um valor para dentro do canal
		time.Sleep(time.Second)
	}
	close(canal)
}
