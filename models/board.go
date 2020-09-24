package models

import (
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/jaadriano/anino_backend/db"
)

type Board struct {
	ID    string  `bson:"_id" json:"_id"`
	Name  string  `json:"name"`
	Entry []Score `json:"entries"`
}

func (b Board) GetByID(id string) (Board, error) {
	collection := db.GetDB().Database("anino").Collection("board")
	bsonID, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": bsonID}
	var board Board
	err = collection.FindOne(context.TODO(), filter).Decode(&board)
	return board, err
}

type BoardPost struct {
	Name  string  `json:"name"`
	Entry []Score `json:"entries"`
}

func (b Board) PostBoard(name string) (Board, error) {
	boardPost := BoardPost{Name: name}

	collection := db.GetDB().Database("anino").Collection("board")

	insertResult, err := collection.InsertOne(context.TODO(), boardPost)

	if err != nil {
		log.Fatal(err)
	}

	id, _ := json.Marshal(insertResult.InsertedID)
	s := string(id)
	boardId := s[1 : len(s)-1]

	board := Board{ID: boardId, Name: boardPost.Name}
	return board, err
}

func (b Board) PutBoard(name string) (Board, error) {
	boardPost := BoardPost{Name: name}

	collection := db.GetDB().Database("anino").Collection("board")

	insertResult, err := collection.InsertOne(context.TODO(), boardPost)

	if err != nil {
		log.Fatal(err)
	}

	id, _ := json.Marshal(insertResult.InsertedID)
	s := string(id)
	boardId := s[1 : len(s)-1]

	board := Board{ID: boardId, Name: boardPost.Name}
	return board, err
}
