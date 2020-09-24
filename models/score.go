package models

import (
	"context"
	"encoding/json"
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
	update := bson.D{
		{"$set", bson.D{
			{"score", scorePost.Score},
		}},
	}
	score := Score{
		ID:       scorePost.BoardId,
		BoardId:  scorePost.BoardId,
		Score:    scorePost.Score,
		ScoredAt: scorePost.ScoredAt,
		UserId:   scorePost.UserId}
	updateResult := collection.FindOneAndUpdate(context.TODO(), filter, update)
	updateResult.Decode(&score)
	if score.ID == score.BoardId {
		insertResult, err := collection.InsertOne(context.TODO(), scorePost)
		if err != nil {
			log.Fatal(err)
		}
		id, _ := json.Marshal(insertResult.InsertedID)
		s := string(id)
		scoreId := s[1 : len(s)-1]
		score.ID = scoreId
	}
	return score, nil
}
