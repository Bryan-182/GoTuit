package routers

import (
	"github.com/Bryan-182/GoTuit/bd"
	"github.com/Bryan-182/GoTuit/models"
	"github.com/aws/aws-lambda-go/events"
)

func DeleteTuit(request events.APIGatewayProxyRequest, claim models.Claim) models.ApiResponse {
	var r models.ApiResponse
	r.Status = 400

	ID := request.QueryStringParameters["id"]

	if len(ID) < 1 {
		r.Message = "El parametro ID es obligatorio"
		return r
	}

	err := bd.DeleteTuit(ID, claim.ID.Hex())

	if err != nil {
		r.Message = "Ocurrió un error al intentar borrar el tuit " + err.Error()
		return r
	}

	r.Message = "Tuit eliminado OK"
	r.Status = 200
	return r
}
