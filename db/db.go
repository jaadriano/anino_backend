package db

import (
	"fmt"

	"github.com/jaadriano/anino_backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init() {
	// server.Use(middleware.Logger(), gindump.Dump())
	var err error
	client, err = mongo.Connect(nil, options.Client().ApplyURI(config.GetConfig().Database))
	if err != nil {
		fmt.Println(err)
	}

}

func GetDB() *mongo.Client {
	return client
}
