package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome string
	Email string
}

func main()  {
	templates = template.Must(template.ParseGlob("*.html"))
	http.HandleFunc("/home",func (w http.ResponseWriter, r *http.Request) {

		u := usuario{"João", "J@J.com.br"}

		templates.ExecuteTemplate(w, "home.html", u )
	})
	fmt.Println("Iniciando o servidor...")
	log.Fatalln(http.ListenAndServe(":8000", nil))// criando o servidor e subindo na porta 8000

}
