package middlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"log"
	"net/http"
)

//vai ser uma camada que vai ficar entre a requisição e a resposta
//é bastante utilizando quando temos alguma função que deva ser aplicada para todas as rotas
//é muito comum no middlewares termos uma especie de alinhamento de funções


//Logger escreve informações da requisição no terminal.
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w,r)
	}

}


//Autenticar verifica se o usuario fazendo a requisição está autenticado
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil{
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		proximaFuncao(w,r)
	}
}