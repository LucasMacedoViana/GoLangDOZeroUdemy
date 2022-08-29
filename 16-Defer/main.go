package main

import "fmt"

func funcao1()  {
	fmt.Println("Executando a Função 1")
}

func funcao2()  {
	fmt.Println("Executando a Função 2")
}

func aluneEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("media calculada, o resultado sera retornado")
	fmt.Println( "Entrando na função para verificar se o aluno esta aprovado")
	media := (n1 + n2)/2
	if media >=6{

		return true
	} else {
		return false
	}
	
}
func main()  {
	defer funcao1()
	funcao2()
	fmt.Println(aluneEstaAprovado(7,8))
}