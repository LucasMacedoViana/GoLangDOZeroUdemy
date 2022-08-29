package main

import "fmt"

func main()  {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2" // determinou o tipo da variavel pelo valor dela, inferencia de tipo (nome tecnico)

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
		)
	 fmt.Println(variavel3 , variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "constate 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5 // invertendo as variaveis
	fmt.Println(variavel5, variavel6)
}
