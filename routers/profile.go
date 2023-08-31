package routers

import (
	"encoding/json"
	"fmt"

	"github.com/Bryan-182/GoTuit/bd"
	"github.com/Bryan-182/GoTuit/models"
	"github.com/aws/aws-lambda-go/events"
)

func Profile(request events.APIGatewayProxyRequest) models.ApiResponse {
	var r models.ApiResponse
	r.Status = 400

	fmt.Println("Enter to profile")

	ID := request.QueryStringParameters["id"]

	if len(ID) < 1 {
		r.Message = "El parámetro ID es obligatorio"
		fmt.Println(r.Message)
		return r
	}

	profile, err := bd.SearchProfile(ID)

	if err != nil {
		r.Message = "Ocurrió un error al intentar buscar el registro " + err.Error()
		return r
	}

	respJson, err := json.Marshal(profile)

	if err != nil {
		r.Status = 500 //Error de back, no de solicitud
		r.Message = "Error al foramtear los datos del usuario a JSON " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
