package main

import "fmt"

func somar(n1 int8 , n2 int8) int8{
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8,int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao

}

func main()  {
	soma := somar(10, 30)
	fmt.Println(soma)
	fmt.Println("\n ")

	var f  = func() {
		fmt.Println("Função 1")
	}
	f()
	fmt.Println("\n ")


	var f2  = func(txt string) {
		fmt.Println(txt)
	}
	f2("Função 2")
	fmt.Println("\n ")

	var f3  = func(txt string) string{
		fmt.Println(txt)
		return txt
	}
	respota := f3("Função 3")
	fmt.Println(respota)
	fmt.Println("\n ")

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15) // usando os dois retornos
	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSoma, _  = calculosMatematicos(10, 15)  //usando somente um dos retornos
	fmt.Println(resultadoSoma)

	_ , resultadoSubtracao = calculosMatematicos(10, 15) //usando somente um dos retornos
	fmt.Println(resultadoSubtracao)

}
