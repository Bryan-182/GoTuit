package jwt

import (
	"errors"
	"strings"

	"github.com/Bryan-182/GoTuit/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var IdUser string

func TokenProcess(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	myKey := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("formato de token inválido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		//Rutina que chequea contra la BD
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("token inválido")
	}

	return &claims, true, string(""), err
}
