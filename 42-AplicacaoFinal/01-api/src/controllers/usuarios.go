package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)
//CriarUsuario insere um usuario no bancode dados
func CriarUsuario(w http.ResponseWriter, r *http.Request)  {
	CorpoRequest,erro := ioutil.ReadAll(r.Body)
	if erro != nil{
		respostas.Erro(w, http.StatusUnprocessableEntity,erro)
		return
	}
	var usuario modelos.Usuario
	if erro = json.Unmarshal(CorpoRequest, &usuario); erro != nil{
		respostas.Erro(w, http.StatusBadRequest,erro)
		return
	}

	if erro = usuario.Preparar("cadastro"); erro!=nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	db,erro := banco.Conectar()
	if erro !=nil{
		respostas.Erro(w, http.StatusInternalServerError,erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID,erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError,erro)
		return
	}

	respostas.JSON(w,http.StatusCreated, usuario)
}

//BuscarUsuarios busca todos os usarios salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request)  {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db,erro := banco.Conectar()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.Buscar(nomeOuNick)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w,http.StatusOK, usuarios)

}

//BuscarUsuario um usario salvo no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request)  {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db,erro := banco.Conectar()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario, erro := repositorio.BuscarPorID(usuarioID)
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, usuario)


}

//AtualizarUsuario altera as informaçÕes de um usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request)  {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil{
		respostas.Erro(w,http.StatusUnprocessableEntity,erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil{
		respostas.Erro(w, http.StatusBadRequest,erro)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil{
		respostas.Erro(w,http.StatusBadRequest,erro)
		return
	}

	db,erro := banco.Conectar()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

//DeletarUsuario exclui as informações de um usuario no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request)  {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil{
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db,erro := banco.Conectar()
	if erro != nil{
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Deletar(usuarioID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

