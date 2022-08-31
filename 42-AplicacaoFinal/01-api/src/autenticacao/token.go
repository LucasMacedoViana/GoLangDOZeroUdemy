package autenticacao

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

//CriarToken retorna um token assinado com as permissões do usuario
func CriarToken(usuarioId uint64) (string, error)  {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour*6).Unix()
	permissoes["usuarioId"] = usuarioId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SigningString() //secret
}
