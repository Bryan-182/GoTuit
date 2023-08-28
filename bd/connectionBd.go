package bd

import (
	"context"
	"fmt"

	"github.com/Bryan-182/GoTuit/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnection *mongo.Client
var DataBaseName string

func ConnectDB(ctx context.Context) error {
	user := ctx.Value(models.Key("username")).(string)
	pasword := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)

	connectionStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, pasword, host)

	var clientOptions = options.Client().ApplyURI(connectionStr)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = client.Ping(ctx, nil) //Hago ping para verificar conexion

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("> Conexion exitosa con la BD")

	MongoConnection = client

	DataBaseName = ctx.Value(models.Key("database")).(string)

	return nil
}

func DBConnected() bool {
	err := MongoConnection.Ping(context.TODO(), nil)

	return err == nil
}
