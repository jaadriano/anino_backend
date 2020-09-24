package models

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jaadriano/anino_backend/db"
	"go.mongodb.org/mongo-driver/bson"
)

type ScoreToAdd struct {
	Score int `json:"score_to_add"`
}
type Score struct {
	ID       string    `bson:"_id" json:"_id"`
	BoardId  string    `bson:"board_id" json:"board_id"`
	Score    int       `json:"score"`
	ScoredAt time.Time `json:"scored_at"`
	UserId   string    `bson:"user_id" json:"user_id"`
}

type ScorePost struct {
	BoardId  string    `bson:"board_id" json:"board_id"`
	Score    int       `json:"score"`
	ScoredAt time.Time `json:"scored_at"`
	UserId   string    `bson:"user_id" json:"user_id"`
}

func (b Score) PostScore(scorePost ScorePost) (Score, error) {
	collection := db.GetDB().Database("anino").Collection("score")
	filter := bson.M{"board_id": scorePost.BoardId, "user_id": scorePost.UserId}
	var board Board
	var scoreId string
	err := collection.FindOne(context.TODO(), filter).Decode(&board)
	if err != nil {
		insertResult, err := collection.InsertOne(context.TODO(), scorePost)
		if err != nil {
			log.Fatal(err)
		}
		id, _ := json.Marshal(insertResult.InsertedID)
		s := string(id)
		scoreId = s[1 : len(s)-1]
	} else {
		fmt.Println(board.ID)
	}
	score := Score{
		ID:       scoreId,
		BoardId:  scorePost.BoardId,
		Score:    scorePost.Score,
		ScoredAt: scorePost.ScoredAt,
		UserId:   scorePost.UserId}
	err = nil
	return score, err
}
