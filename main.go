package main

import (
	"context"
	"os"
	"strings"

	"github.com/Bryan-182/GoTuit/awsgo"
	"github.com/Bryan-182/GoTuit/bd"
	"github.com/Bryan-182/GoTuit/handlers"
	"github.com/Bryan-182/GoTuit/models"
	"github.com/Bryan-182/GoTuit/secretmanager"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(RunLambda)
}

func RunLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse
	awsgo.InitAWS()

	if !ValidParameters() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno, deben incluir 'SecretName', 'BucketName', 'UrlPrefix'",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}

		return res, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))

	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura de 'SecretName' " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}

		return res, nil
	}

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("username"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtsign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketname"), os.Getenv("BucketName"))

	path := strings.Replace(request.PathParameters["gotuit"], os.Getenv("UrlPrefix"), "", -1)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)

	//Chequeo conexion a DB o la conecto
	err = bd.ConnectDB(awsgo.Ctx)

	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error al intentar conectarse a la BD " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}

		return res, nil
	}

	resApi := handlers.Handlers(awsgo.Ctx, request)
	if resApi.CustomResp == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: resApi.Status,
			Body:       resApi.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}

		return res, nil
	} else {
		return resApi.CustomResp, nil
	}
}

func ValidParameters() bool {
	_, hasParameters := os.LookupEnv("SecretName")

	if !hasParameters {
		return hasParameters
	}

	_, hasParameters = os.LookupEnv("BucketName")

	if !hasParameters {
		return hasParameters
	}

	_, hasParameters = os.LookupEnv("UrlPrefix")

	if !hasParameters {
		return hasParameters
	}

	return hasParameters
}
