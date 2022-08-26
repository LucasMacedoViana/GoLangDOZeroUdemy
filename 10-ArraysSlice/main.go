package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var array1 [5] string
	array1[0] = "Posição 1"
	fmt.Println(array1)


	array2 := [5]string{"Posição1","Posição2","Posição3","Posição4","Posição5"}
	fmt.Println(array2)

	array3 := [...] int{1,2,3,4,5} // o array vai declarar que tem o tamanho de 5 opr conta da quantidade de itens passados
	fmt.Println(array3)
	fmt.Println("")

	slice := []int{ 1,2,3,4,5,6,7,8}
	fmt.Println(slice)
	fmt.Println("")

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))
	fmt.Println("")

	slice = append(slice, 18) // adicionando ao final do slice
	fmt.Println(slice)
	fmt.Println("")

	slice2 := array2[1:3] // criando uma fatia do array, primeiro indice é inclusivo e o segundo indice é ignorado
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)
	fmt.Println("")


	// Array Interno
	slice3 := make([]float32, 10, 11) // inicialmente essa função cria um array de 11 posições e retira as 10 primeiras posições para o slice
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 12)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) // capacidade

	slice4 := make([] float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) //tamanho
	fmt.Println(cap(slice4)) // capacidade

	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) //tamanho
	fmt.Println(cap(slice4)) // capacidade

}
