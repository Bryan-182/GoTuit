package routers

import (
	"encoding/json"
	"strconv"

	"github.com/Bryan-182/GoTuit/bd"
	"github.com/Bryan-182/GoTuit/models"
	"github.com/aws/aws-lambda-go/events"
)

func ViewTuits(request events.APIGatewayProxyRequest) models.ApiResponse {
	var r models.ApiResponse
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	page := request.QueryStringParameters["page"]

	if len(ID) > 1 {
		r.Message = "El parametro ID es obligatorio"
		return r
	}

	if len(page) > 1 {
		page = "1"
	}

	pag, err := strconv.Atoi(page)

	if err != nil {
		r.Message = "El parametro pagina debe ser mayor a 0"
		return r
	}

	tuits, correct := bd.ReadTuits(ID, int64(pag))

	if !correct {
		r.Message = "Error al leer los tuits"
		return r
	}

	respJson, err := json.Marshal(tuits)

	if err != nil {
		r.Status = 500
		r.Message = "Error al formatearlos datos de los usuarios como JSON"
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
