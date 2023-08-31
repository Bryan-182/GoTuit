package bd

import (
	"context"

	"github.com/Bryan-182/GoTuit/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadTuits(Id string, page int64) ([]*models.RetrieveTuits, bool) {

	ctx := context.TODO()
	db := MongoConnection.Database(DataBaseName)
	col := db.Collection("tuits")

	var results []*models.RetrieveTuits

	condition := bson.M{
		"userid": Id,
	}

	opts := options.Find()

	//Opciones de busqueda
	opts.SetLimit(20)
	opts.SetSort(bson.D{{Key: "date", Value: -1}})
	opts.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, opts)

	if err != nil {
		return results, false
	}

	for cursor.Next(ctx) {
		var register models.RetrieveTuits
		err := cursor.Decode(&register) //Asigna automaticamente el cursor al registro

		if err != nil {
			return results, false
		}

		results = append(results, &register)
	}

	return results, true
}
