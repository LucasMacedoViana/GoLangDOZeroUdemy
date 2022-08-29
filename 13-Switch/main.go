package main

import "fmt"

func diaDaSemana(numero int)string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "terça-Feira"
	case 4:
		return "quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sabado"
	default:
		return "digite um numero de 1 a 7"
	}
}

func diaDaSemana2(num int)string  {

	switch {
	case num == 1:
		return "Domingo"
	case num == 2:
		return "Segunda-Feira"
	case num == 3:
		return "terça-Feira"
	case num == 4:
		return "quarta-Feira"
	case num == 5:
		return "Quinta-Feira"
	case num == 6:
		return "Sexta-feira"
	case num == 7:
		return "Sabado"
	default:
		return "digite um numero de 1 a 7"
	}
}
func diaDaSemana3(num int)string  {
	var diaDaSemana string
	switch {
	case num == 1:
		diaDaSemana = "Domingo"
	case num == 2:
		diaDaSemana = "Segunda-Feira"
	case num == 3:
		diaDaSemana = "terça-Feira"
	case num == 4:
		diaDaSemana = "quarta-Feira"
	case num == 5:
		diaDaSemana = "Quinta-Feira"
	case num == 6:
		diaDaSemana = "Sexta-feira"
	case num == 7:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "digite um numero de 1 a 7"
	}

	return diaDaSemana
}

func diaDaSemana4(num int)string  {
	var diaDaSemana string
	switch {
	case num == 1:
		diaDaSemana = "Domingo"
		fallthrough
	case num == 2:
		diaDaSemana = "Segunda-Feira"
	case num == 3:
		diaDaSemana = "terça-Feira"
	case num == 4:
		diaDaSemana = "quarta-Feira"
	case num == 5:
		diaDaSemana = "Quinta-Feira"
	case num == 6:
		diaDaSemana = "Sexta-feira"
	case num == 7:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "digite um numero de 1 a 7"
	}

	return diaDaSemana
}
func main()  {
	dia := diaDaSemana(1)
	fmt.Println(dia)
	fmt.Println("-----------------")
	dia2 := diaDaSemana2(4)
	fmt.Println(dia2)
	fmt.Println("-----------------")
	dia3 := diaDaSemana3(7)
	fmt.Println(dia3)
	fmt.Println("-----------------")
	dia4 := diaDaSemana4(1)
	fmt.Println(dia4)
	fmt.Println("-----------------")
}
