package models

import (
	"context"
	"log"
	"time"

	"github.com/jaadriano/anino_backend/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScoreToAdd struct {
	Score int `json:"score_to_add"`
}
type Score struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	BoardId  string             `bson:"board_id" json:"board_id"`
	Score    int                `json:"score"`
	ScoredAt time.Time          `json:"scored_at"`
	UserId   string             `bson:"user_id" json:"user_id"`
}

func (b Score) PostScore(score Score) (Score, error) {
	filter := bson.M{
		"board_id": score.BoardId,
		"user_id":  score.UserId}
	update := bson.D{
		{"$set",
			bson.D{
				{"score", score.Score},
			}},
	}
	collection := db.GetDB().Database("anino").Collection("score")
	updateResult := collection.FindOneAndUpdate(context.TODO(), filter, update)
	tempScoreId := score.ID
	tempScoreVal := score.Score
	updateResult.Decode(&score)
	if tempScoreId == score.ID {
		insertResult, err := collection.InsertOne(context.TODO(), score)
		if err != nil {
			log.Fatal(err)
		}
		score.ID = insertResult.InsertedID.(primitive.ObjectID)
	}

	// //find board where id = board id
	// collection = db.GetDB().Database("anino").Collection("board")
	// filter = bson.M{
	// 	"_id": scorePost.BoardId}
	// //push
	// updateResult = collection.FindOneAndUpdate(context.TODO(), filter, update)
	// //sort
	score.Score = tempScoreVal
	return score, nil
}
