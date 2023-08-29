package bd

import (
	"context"

	"github.com/Bryan-182/GoTuit/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UserExist(email string) (models.User, bool, string) {
	ctx := context.TODO()

	db := MongoConnection.Database(DataBaseName)
	col := db.Collection("users") //Colleciones

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)

	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
