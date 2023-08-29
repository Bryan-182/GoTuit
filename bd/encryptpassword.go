package bd

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	cost := 8 //Cantidad de encriptaciones que realiza el modulo de encriptacion. Entre más, alto más seguro, pero más costoso en cuanto a recursos y tiempo
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		return err.Error(), err
	}

	return string(bytes), nil
}
