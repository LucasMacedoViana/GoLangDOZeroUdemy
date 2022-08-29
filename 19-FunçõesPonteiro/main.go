package main

import "fmt"

// passando um parametro por copia
func inverterSinal(num int) int  {
	return num * -1
}
// passando uma referencia
func inverterSinalComPonteiro(num *int)  {
	*num = *num * -1
}

func main()  {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
