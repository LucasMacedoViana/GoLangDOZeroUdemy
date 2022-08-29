package main

import "fmt"

func main()  {
	func(){
		fmt.Println("Ola mnudo")
	}()

	// passando um parametro
	func(texto string){
		fmt.Println("Ola mnudo", texto)
	}("passando paramentro")

	// retorno da função anonima
	retorno := func(texto string) string{
		return fmt.Sprintf("recebido -> %s", texto)
	}("passando paramentro")
	fmt.Println(retorno)
}
