package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var(
	//StringConexaoBanco é a string de conexão com o postgres
	StringConexaoBanco = ""

	//Porta onde a API vai estar rodando
	Porta = 0

	//Senha senha do banco de dados
	Senha = 0


	//SecretKey é a chave que vai ser usada para assinar o token
	SecretKey []byte
)

//Carregar vai inicializar as variaveis de ambiente
func Carregar()  {
	var erro error

	if erro = godotenv.Load(); erro!= nil{
		log.Fatal(erro)
	}
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil{
		Porta = 9000
	}

	Senha, erro = strconv.Atoi(os.Getenv("DB_SENHA"))
	if erro != nil{
		Senha = 12345678
	}
	StringConexaoBanco = fmt.Sprintf(
		"user=%s dbname=%s password=%v host=localhost sslmode=disable", os.Getenv("DB_USUARIO"), os.Getenv("DB_NOME"), Senha)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
