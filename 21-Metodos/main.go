package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}
// todos os usuarios tem um metodo chamado salvar
// o "u" é uma variavel que podemos referenciar outros campos do mesmo usuario que chamou esse metodo

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuario %s no banco de dados \n", u.nome)

}
//não usamos o ponteiro pois so usamos a idade para fazer uma comparacão
func (u usuario) maiorIdade () bool {
	return u.idade >= 18
}

// usando o ponteiro
// como o idade agora é um ponteiro, isso vai refletir em todo o codigo
func (u *usuario) aniversairo (){
	u.idade ++
}

func main()  {
	usuario1 := usuario {"Lucas", 17}
	usuario1.aniversairo()
	fmt.Println(usuario1)
	usuario1.salvar()
	maioridade1 := usuario1.maiorIdade()
	fmt.Println(maioridade1)


	usuario2 := usuario{"Mateus", 26}
	usuario2.salvar()
	maioridade2 := usuario2.maiorIdade()
	fmt.Println(maioridade2)


}
