package modelos

import (
	"errors"
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
func (usuario *Usuario) Preparar () error  {
	if erro := usuario.validar(); erro!= nil{
		return erro
	}
	usuario.formatar()
	return nil
}


//validar ai verificar se os campos nome, nick, email e senha estão preenchidos corretamente
func (usuario *Usuario) validar () error  {
	if usuario.Nome ==""{
		return errors.New("O nome é obrigatorio e não pode estsar em branco!")
	}

	if usuario.Nick ==""{
		return errors.New("O nick é obrigatorio e não pode estsar em branco!")
	}

	if usuario.Email ==""{
		return errors.New("O e-mail é obrigatorio e não pode estsar em branco!")
	}

	if usuario.Senha ==""{
		return errors.New("A senha é obrigatorio e não pode estsar em branco!")
	}

	return nil
}

// formatar vai retirar os espaços no começo e no final da string
func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

}