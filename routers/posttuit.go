package routers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Bryan-182/GoTuit/bd"
	"github.com/Bryan-182/GoTuit/models"
)

func PostTuit(ctx context.Context, claim models.Claim) models.ApiResponse {
	var message models.Tuit

	var r models.ApiResponse
	r.Status = 400

	idUser := claim.ID.Hex()

	body := ctx.Value(models.Key("body")).(string)

	err := json.Unmarshal([]byte(body), &message)

	if err != nil {
		r.Message = "Ocurrió un error al decodificar el body " + err.Error()
		return r
	}

	register := models.SaveTuit{
		UserId:  idUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.SaveTuit(register)

	if err != nil {
		r.Message = "Ocurrió un error al insertar el registro " + err.Error()
		return r
	}

	if !status {
		r.Message = "Ocurrió un error al insertar el tuit"
		return r
	}

	r.Status = 200
	r.Message = "Tuit guardado Ok"
	return r
}
