package bd

import (
	"context"

	"github.com/Bryan-182/GoTuit/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(u models.User, Id string) (bool, error) {
	ctx := context.TODO()

	db := MongoConnection.Database(DataBaseName)

	col := db.Collection("users")

	register := make(map[string]interface{})

	if len(u.Name) > 0 {
		register["name"] = u.Name
	}

	if len(u.Lastname) > 0 {
		register["lastname"] = u.Lastname
	}

	register["birthdate"] = u.Birthdate

	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}

	if len(u.Bio) > 0 {
		register["bio"] = u.Bio
	}

	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}

	if len(u.Location) > 0 {
		register["location"] = u.Location
	}

	if len(u.Website) > 0 {
		register["website"] = u.Website
	}

	//Creamos el registro de actualizacion para mongo
	updateStr := bson.M{
		"$set": register,
	}

	objId, _ := primitive.ObjectIDFromHex(Id)

	//Filtro (WHERE)
	filter := bson.M{"_id": bson.M{"$eq": objId}}

	_, err := col.UpdateOne(ctx, filter, updateStr)

	if err != nil {
		return false, err
	}

	return true, nil
}
