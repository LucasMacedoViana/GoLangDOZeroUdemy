package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main()  {
	//string de conexão
	conexao := "user=postgres dbname=mesa password=12345678 host=localhost sslmode=disable"
	//abrindo a conexão
	db, erro := sql.Open("postgres", conexao)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil{
		log.Fatal(erro)
	}
	fmt.Println("conexão esta aberta")

	tabela, erro := db.Query("select * from usuario")
	if erro != nil{
		log.Fatal(erro)
	}
	defer tabela.Close()

	fmt.Println(tabela)
}
