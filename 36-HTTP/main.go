package main

import (
	"log"
	"net/http"
)

func usuario (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar Pagina de Usuarios"))
}

func home (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pagina Home"))
}

func main()  {

	http.HandleFunc("/home",home) // normalmente as rotas são declaradas antes.
	http.HandleFunc("/usuario",usuario)

	log.Fatalln(http.ListenAndServe(":8000", nil))// criando o servidor e subindo na porta 8000
	
}
