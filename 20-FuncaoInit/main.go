package main

import "fmt"

var n int
func init()  {
	fmt.Println("Executando a função main")
	n = 10
}

func main()  {
	fmt.Println("Função main sendo executadoa")
	fmt.Println(n)
}
