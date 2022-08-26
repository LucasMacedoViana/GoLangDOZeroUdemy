package main

import (
	"errors"
	"fmt"
)

func main()  {
	 var numero1 int16 = -100
	 fmt.Println(numero1)

	 var numero2 uint16 = 100
	 fmt.Println(numero2)

	 //allias
	 var numero3 rune = 12345 // igual ao int32, numeros que representam caracteres
	 fmt.Println(numero3)

	 var numero4 byte = 123 // igual ao uint8
	 fmt.Println(numero4)

	 var real1 float32= 123.45
	 fmt.Println(real1)

	 var real2 float64 = 5671231231209.43
	 fmt.Println(real2)

	 real3 := 123123.32
	 fmt.Println(real3)


	 var str1 string = "texto 1"
	 fmt.Println(str1)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	var texto string // string vazia
	var num int
	fmt.Println(texto , num)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro1 error
	fmt.Println(erro1)

	var erro2 error = errors.New("erro interno")
	fmt.Println(erro2)

}
