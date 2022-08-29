package main

import "fmt"
// varios parametros
func soma(numeros ...int) int {
	total :=  0
	for _, numero := range numeros{
		total += numero
	}
	return total
}
// parametros fixos com parametros variaticos 
func escrever(texto string, numeros ...int)  {
	for _, numero := range numeros{
		fmt.Println(texto, numero)
	}
	
}
func main()  {
	totalDaSoma := soma(1,2,3,4,5,6,12,213,43,2314,23445,6567,578)
	fmt.Println(totalDaSoma)

	escrever("Ola mundo", 10,20,30,40,50)
}
