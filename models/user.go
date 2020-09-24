package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/jaadriano/anino_backend/db"
)

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name string             `json:"name"`
}

func (h User) GetByID(id string) (User, error) {
	collection := db.GetDB().Database("anino").Collection("user")
	bsonID, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": bsonID}
	var user User
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	return user, err
}

func (h User) PostUser(name string) (User, error) {
	user := User{Name: name}
	collection := db.GetDB().Database("anino").Collection("user")
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	user.ID = insertResult.InsertedID.(primitive.ObjectID)
	return user, err
}
