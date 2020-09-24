package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/jaadriano/anino_backend/db"
)

type Board struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name    string             `json:"name"`
	Entries []Score            `json:"entries"`
}

func (b Board) GetByID(id string) (Board, error) {
	collection := db.GetDB().Database("anino").Collection("board")
	bsonID, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": bsonID}
	var board Board
	err = collection.FindOne(context.TODO(), filter).Decode(&board)
	return board, err
}

func (b Board) PostBoard(name string) (Board, error) {
	board := Board{Name: name}
	collection := db.GetDB().Database("anino").Collection("board")
	insertResult, err := collection.InsertOne(context.TODO(), board)
	if err != nil {
		log.Fatal(err)
	}
	board.ID = insertResult.InsertedID.(primitive.ObjectID)
	return board, err
}

// func (b Board) PutBoard(name string) (Board, error) {
// 	board := Board{Name: name}
// 	collection := db.GetDB().Database("anino").Collection("board")
// 	insertResult, err := collection.InsertOne(context.TODO(), boardPost)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	id, _ := json.Marshal(insertResult.InsertedID)
// 	s := string(id)
// 	boardId := s[1 : len(s)-1]

// 	board := Board{ID: boardId, Name: boardPost.Name}
// 	return board, err
// }
