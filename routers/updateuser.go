package routers

import (
	"context"
	"encoding/json"

	"github.com/Bryan-182/GoTuit/bd"
	"github.com/Bryan-182/GoTuit/models"
)

func UpdateUser(ctx context.Context, claim models.Claim) models.ApiResponse {
	var r models.ApiResponse
	r.Status = 400

	var t models.User

	body := ctx.Value(models.Key("body")).(string)

	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		r.Message = "Datos incorrectos " + err.Error()
	}

	status, err := bd.UpdateUser(t, claim.ID.Hex())

	if err != nil {
		r.Message = "Ocurrió un error al intentar modificar el usuario " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se logró modificar el usuario"
		return r
	}

	r.Status = 200
	r.Message = "Modificación de perfil OK"
	return r
}
