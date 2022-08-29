package main

import "fmt"

func main()  {
	usuario := map[string]string{
		"nome":			"Pedro",
		"sobrenome":	"Silva",

	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome":{
			"primeiro": "Lucas",
			"ultimo":"Viana",
		},
		"curso":{
			"nome": "Engenharia",
			"tipo": "Civil",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome") //deletando a chave
	fmt.Println(usuario2)
	usuario2["signo"] = map[string]string{
		"tipo": "gemeos",
	}
	fmt.Println(usuario2)
}

