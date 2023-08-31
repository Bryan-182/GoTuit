package bd

import (
	"context"

	"github.com/Bryan-182/GoTuit/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(ID string) (models.User, error) {
	ctx := context.TODO()

	db := MongoConnection.Database(DataBaseName)
	col := db.Collection("users")

	var profile models.User

	objectID, _ := primitive.ObjectIDFromHex(ID) //Recibe Sting y convierte en primitive

	condition := bson.M{
		"_id": objectID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile) //Puntero a variable perfil

	profile.Password = ""

	if err != nil {
		return profile, err
	}

	return profile, nil
}
