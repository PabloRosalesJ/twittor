package bd

import (
	"context"
	"time"

	"github.com/PabloRosalesJ/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StoreUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConnection.Database("twittor")
	collection := db.Collection("users")

	u.Password, _ = BcryptPwd(u.Password)

	res, err := collection.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjId, _ := res.InsertedID.(primitive.ObjectID)
	return ObjId.String(), true, nil
}

func ExistUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConnection.Database("twittor")
	collection := db.Collection("users")

	where := bson.M{"email": email}
	var user models.User

	err := collection.FindOne(ctx, where).Decode(&user)
	ID := user.ID.Hex()
	if err != nil {
		return user, false, ID
	}
	return user, true, ID
}
