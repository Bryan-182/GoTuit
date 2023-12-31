package handlers

import (
	"context"
	"fmt"

	"github.com/Bryan-182/GoTuit/jwt"
	"github.com/Bryan-182/GoTuit/models"
	"github.com/Bryan-182/GoTuit/routers"
	"github.com/aws/aws-lambda-go/events"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.ApiResponse {
	fmt.Println("Procesar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var r models.ApiResponse

	r.Status = 400

	isOk, statusCode, message, claim := validAuth(ctx, request)

	if !isOk {
		r.Status = statusCode
		r.Message = message
		return r
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "register":
			return routers.Register(ctx)
		case "login":
			return routers.Login(ctx)
		case "posttuit":
			return routers.PostTuit(ctx, claim)
		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		case "profile":
			return routers.Profile(request)
		case "viewtuits":
			return routers.ViewTuits(request)
		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		case "updateuser":
			return routers.UpdateUser(ctx, claim)
		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		case "deletetuit":
			return routers.DeleteTuit(request, claim)
		}
	}

	r.Message = "Method Invalid"
	return r
}

func validAuth(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)

	if path == "register" || path == "login" || path == "getAvatar" || path == "getBanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]

	if len(token) == 0 {
		return false, 401, "Token requerido", models.Claim{}
	}

	claim, allOk, msg, err := jwt.TokenProcess(token, ctx.Value(models.Key("jwtsign")).(string))

	if !allOk {
		if err != nil {
			fmt.Println("Error en el token " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error en el token " + msg)
			return false, 401, msg, models.Claim{}
		}
	}

	fmt.Println("Token OK")

	return true, 200, msg, *claim
}
