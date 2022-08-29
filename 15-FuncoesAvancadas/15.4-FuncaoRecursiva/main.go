package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1{
		return posicao
	}
	return fibonacci(posicao -2) + fibonacci(posicao - 1)
}

func main()  {
	posicao := uint(100)
	fmt.Println(fibonacci(posicao))
	for i := uint(0); i< posicao ; i++{
		fmt.Println(fibonacci(i))
	}
}
