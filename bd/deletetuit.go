package bd

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTuit(ID string, UserId string) error {
	ctx := context.TODO()
	bd := MongoConnection.Database(DataBaseName)
	col := bd.Collection("tuits")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":    objID,
		"userid": UserId,
	}

	_, err := col.DeleteOne(ctx, condition)

	return err
}
