package modelos

import (
	"api/src/seguranca"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

//Usuario representa um usuario utilizando a rede social
type Usuario struct {
	ID 			uint64 `json:"id,omitempty"`
	Nome 		string `json:"nome,omitempty"`
	Nick	 	string `json:"nick,omitempty"`
	Email 		string `json:"email,omitempty"`
	Senha 		string `json:"senha,omitempty"`
	CriadoEm 	time.Time `json:"criado_em,omitempty"`
}

//Preparar vai chamar os dois metodos validar e formatar
func (usuario *Usuario) Preparar (etapa string) error  {
	if erro := usuario.validar(etapa); erro!= nil{
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil{
		return erro
	}
	return nil
}

//validar ai verificar se os campos nome, nick, email e senha estão preenchidos corretamente
func (usuario *Usuario) validar (etapa string) error  {
	if usuario.Nome ==""{
		return errors.New("O nome é obrigatorio e não pode estsar em branco!")
	}

	if usuario.Nick ==""{
		return errors.New("O nick é obrigatorio e não pode estsar em branco!")
	}

	if usuario.Email ==""{
		return errors.New("O e-mail é obrigatorio e não pode estsar em branco!")
	}
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil{
		return errors.New("O Email inserido é invalido")
	}

	if etapa == "cadastro" && usuario.Senha ==""{
		return errors.New("A senha é obrigatorio e não pode estsar em branco!")
	}

	return nil
}

// formatar vai retirar os espaços no começo e no final da string
func (usuario *Usuario) formatar(etapa string) error{
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro"{
		senhaHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil{
			return erro
		}
		usuario.Senha = string(senhaHash)
	}
	return nil
}