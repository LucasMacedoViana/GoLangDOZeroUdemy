package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}
//NovoRepositorioDeUsuarios cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

//Criar insere um usuario no banco de dados
func (repositorio Usuarios) Criar (usuario modelos.Usuario)(uint64,error)  {
	statemente, erro := repositorio.db.Prepare(
		"insert into usuarios(nome, nick, email, senha) values ($1, $2, $3, $4)")
	if erro != nil {
		return 0,erro
	}
	defer statemente.Close()
	resultado, erro := statemente.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0,erro
	}
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0,erro
	}
	return uint64(ultimoIDInserido), nil

}

//Buscar traz todos os usuarios que atendem o filtro de nome ou nick
func (repositorio Usuarios) Buscar (nomeOuNick string)([] modelos.Usuario, error)  {
	nomeOuNick = fmt.Sprintf("%%%s%%",nomeOuNick)
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome like $1 or nick like $2",
		nomeOuNick, nomeOuNick)
	if erro != nil{
		return nil, erro
	}
	defer linhas.Close()
	var usuarios[]modelos.Usuario
	for linhas.Next(){
		var usuario modelos.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
			); erro !=nil{
			return nil,erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

//BuscarPorID traz um usuario do banco de dados
func (repositorio Usuarios) BuscarPorID (ID uint64)(modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = $1", ID)

	if erro != nil{
		return modelos.Usuario{},erro
	}
	defer linhas.Close()
	var usuario modelos.Usuario
	if linhas.Next(){
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
			); erro!= nil{
			return modelos.Usuario{},erro
		}
	}
	return usuario, nil
}

//Atualizar altera as informações de um usuario no banco de dados
func (repositorio Usuarios) Atualizar (ID uint64, usuario modelos.Usuario) error{
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = $1, nick = $2, email = $3 where id = $4 ")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, ID); erro != nil{
		return erro
	}
	return nil
}

//Deletar apaga as informações de um usuario no banco de dados
func (repositorio Usuarios) Deletar (ID uint64) error{
	statement, erro := repositorio.db.Prepare(
		"delete from usuarios where id = $1")
	if erro != nil{
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil{
		return erro
	}
	return nil

}