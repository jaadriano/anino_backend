package models

import (
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/jaadriano/anino_backend/db"
)

type User struct {
	ID   string `bson:"_id" json:"_id"`
	Name string `json:"name"`
}

func (h User) GetByID(id string) (User, error) {
	collection := db.GetDB().Database("anino").Collection("user")
	bsonID, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": bsonID}
	var user User

	err = collection.FindOne(context.TODO(), filter).Decode(&user)

	return user, err
}

type UserPost struct {
	Name string `json:"name"`
}

func (h User) PostUser(name string) (User, error) {
	userPost := UserPost{Name: name}

	collection := db.GetDB().Database("anino").Collection("user")

	insertResult, err := collection.InsertOne(context.TODO(), userPost)

	if err != nil {
		log.Fatal(err)
	}

	id, _ := json.Marshal(insertResult.InsertedID)
	s := string(id)
	userId := s[1 : len(s)-1]

	user := User{ID: userId, Name: userPost.Name}
	return user, err
}
