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
	collection = db.GetDB().Database("anino").Collection("board")
	bsonID, err := primitive.ObjectIDFromHex(score.BoardId)
	updateResult = collection.FindOneAndUpdate(context.TODO(),
		bson.M{"_id": bsonID},
		bson.D{
			{"$push",
				bson.D{
					{"entries", score},
				}},
		})
	updateResult = collection.FindOneAndUpdate(context.TODO(),
		bson.M{"_id": bsonID},
		bson.D{
			{"$sort",
				bson.D{
					{"score", -1},
				}},
		})

	// (1) Tasks To Do: use $each and $sort by score Issues: how to nest these?? #$%$%^
	// (2) Fix retrieve by creating struct compliant with specs (includes [name, rank], remove [board_id]) Issues: must finish (1)
	// (3) If rank is top 0 and if channel is not closed, close channel and previous go routine ends. If rank top 0 and channel is closed, run fake account go routine time.sleep for 5 *time.seconds
	score.Score = tempScoreVal
	return score, err
}
