package main

import "fmt"

func main()  {

	num := 10

	if num > 15{
		fmt.Println("é maior que 15")
	} else{
		fmt.Println("é menor do que 15")
	}
	fmt.Println("")

	// if init
	if num2 := -5 ; num2 > 0 {
		fmt.Println("numero é maior que zero")
	}else {
		fmt.Println("é menor que zero")
	}
	fmt.Println("")

	if num2 := num ; num2 > 0 {
		fmt.Println("numero é maior que zero")
	}else {
		fmt.Println("é menor que zero")
	}
	
}
