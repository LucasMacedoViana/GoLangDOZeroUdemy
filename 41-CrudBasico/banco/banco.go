package banco

import (
	"database/sql"

	_ "github.com/lib/pq" // Driver de conexão com o postgres
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "user=postgres dbname=mesa password=12345678 host=localhost sslmode=disable"

	db, erro := sql.Open("postgres", stringConexao)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}