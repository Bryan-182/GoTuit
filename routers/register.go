package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Bryan-182/GoTuit/bd"
	"github.com/Bryan-182/GoTuit/models"
)

func Register(ctx context.Context) models.ApiResponse {
	var t models.User
	var r models.ApiResponse
	r.Status = 400

	fmt.Println("Enter to register")

	body := ctx.Value(models.Key("body")).(string)

	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Debe especificar el email"
		fmt.Println(r.Message)
		return r
	}

	if len(t.Password) < 6 {
		r.Message = "Debe especificar una contraseña de al menos 6 caracteres"
		fmt.Println(r.Message)
		return r
	}

	_, exist, _ := bd.UserExist(t.Email)

	if exist {
		r.Message = "Usuario ya existe"
		fmt.Println(r.Message)
		return r
	}

	_, status, err := bd.RegisterUser(t)

	if err != nil {
		r.Message = "Ocurrio un error al intentar registrar el usuario " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	if !status {
		r.Message = "No se ha logrado insetrar el registro del usuario"
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "Registro Ok"
	fmt.Println(r.Message)
	return r
}
