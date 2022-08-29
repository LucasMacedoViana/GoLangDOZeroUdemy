package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // quantas goroutines vao estar trabalhando ao mesmo tempo
	// as funções anonimas vao estar rodando concorrentemente
	go func() {
		escrever("Ola mundo!") //goroutine
		waitGroup.Done() //-1 do add
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()
	waitGroup.Wait() //esperar a contagem das gorotines chegar em zero

}

func escrever(texto string)  {
	for i := 0; i < 5 ; i++{
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}