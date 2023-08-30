package jwt

import (
	"context"
	"time"

	"github.com/Bryan-182/GoTuit/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(ctx context.Context, t models.User) (string, error) {
	jwtSign := ctx.Value(models.Key("jwtsign")).(string)

	myKey := []byte(jwtSign)

	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastname":  t.Lastname,
		"birthdate": t.Birthdate,
		"bio":       t.Bio,
		"location":  t.Location,
		"website":   t.Website,
		"_id":       t.ID.Hex(),                            //Convierte el id a texto
		"exp":       time.Now().Add(time.Hour * 24).Unix(), //Agrega 24 horas y la convierte en foramto unix
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
