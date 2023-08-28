package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/Bryan-182/GoTuit/awsgo"
	"github.com/Bryan-182/GoTuit/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error) {
	var dataSecret models.Secret
	fmt.Println("> Pido secreto " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println(err.Error())
		return dataSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &dataSecret)
	fmt.Println("> Lectura de Secret OK " + secretName)
	return dataSecret, nil
}
