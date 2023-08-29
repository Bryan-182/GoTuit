package bd

import (
	"context"

	"github.com/Bryan-182/GoTuit/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUser(u models.User) (string, bool, error) {
	ctx := context.TODO()

	db := MongoConnection.Database(DataBaseName)
	col := db.Collection("users") //Colleciones

	//Ecriptar password
	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
