package main

import "fmt"

func main()  {
	// ARITIMETICOS
	soma := 1+2
	subtracao := 2-1
	divisao := 10/5
	multiplicacao := 5*2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)
	fmt.Println("")

	// ATRIBUIÇÃO
	var variavel1 string = "String1"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	fmt.Println("")

	// OPERADORES RELACIONAIS
	fmt.Println(1>2)
	fmt.Println(1>=2)
	fmt.Println(1==2)
	fmt.Println(1<2)
	fmt.Println(1<=2)
	fmt.Println(1!=2)
	fmt.Println("")

	// OPERADORES LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	fmt.Println("")

	//OPERADORES UNARIOS
	num := 10
	num ++ // numero = numero + 1
	fmt.Println(num)
	num += 15 // numero = numero + 15
	fmt.Println(num)
	num -- // numero = numero - 1
	fmt.Println(num)
	num -= 20// numero = numero - 20
	fmt.Println(num)
	num *= 3 // numero = numero * 3
	fmt.Println(num)
	fmt.Println("")


}