package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)
//CriarUsuario insere um usuario no bancode dados
func CriarUsuario(w http.ResponseWriter, r *http.Request)  {
	CorpoRequest,erro := ioutil.ReadAll(r.Body)
	if erro != nil{
		log.Fatal(erro)
	}
	var usuario modelos.Usuario
	if erro = json.Unmarshal(CorpoRequest, &usuario); erro != nil{
		log.Fatal(erro)
	}
	db,erro := banco.Conectar()
	if erro !=nil{
		log.Fatal(erro)
	}
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	repositorio.Criar(usuario)

}

//BuscarUsuarios busca todos os usarios salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Buscando todos os usuários!"))

}
//BuscarUsuario um usario salvo no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("BUscando um usuário!"))

}
//AtualizarUsuario altera as informaçÕes de um usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Atualizando Usuário!"))

}

//DeletarUsuario exclui as informações de um usuario no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("deletando um usuário!"))

}

