package bd

import (
	"context"

	"github.com/Bryan-182/GoTuit/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveTuit(t models.SaveTuit) (string, bool, error) {
	ctx := context.TODO()

	db := MongoConnection.Database(DataBaseName)
	col := db.Collection("tuits") //Colleciones

	register := bson.M{
		"userid":  t.UserId,
		"message": t.Message,
		"date":    t.Date,
	}

	result, err := col.InsertOne(ctx, register)

	if err != nil {
		return "", false, err
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)

	return objId.String(), true, nil
}
