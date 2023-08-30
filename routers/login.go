package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Bryan-182/GoTuit/bd"
	"github.com/Bryan-182/GoTuit/jwt"
	"github.com/Bryan-182/GoTuit/models"

	"github.com/aws/aws-lambda-go/events"
)

func Login(ctx context.Context) models.ApiResponse {
	var t models.User
	var r models.ApiResponse
	r.Status = 400

	fmt.Println("Enter to login")

	body := ctx.Value(models.Key("body")).(string)

	err := json.Unmarshal([]byte(body), &t) //Parsea el JSON

	if err != nil {
		r.Message = "Datos incorrectos " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Debe especificar el email"
		fmt.Println(r.Message)
		return r
	}

	userData, exist := bd.TryLogin(t.Email, t.Password)

	if !exist {
		r.Message = "Datos incorrectos"
		fmt.Println(r.Message)
		return r
	}

	jwtKey, err := jwt.GenerateJWT(ctx, userData)

	if err != nil {
		r.Message = "Ocurrió un error al intentar generar el token: " + err.Error()
		return r
	}
	response := models.LoginResponse{
		Token: jwtKey,
	}

	token, err2 := json.Marshal(response) //Formatea de Go Struct a JSON

	if err2 != nil {
		r.Message = "Ocurrió un error al intentar fortear el token a jon: " + err2.Error()
		return r
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(time.Hour * 24),
	}

	cookieStr := cookie.String()

	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set.Cookie":                  cookieStr,
		},
	}

	r.Status = 200
	r.Message = string(token)
	r.CustomResp = res

	return r
}
