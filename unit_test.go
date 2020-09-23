package main

import (
	"anino_backend/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	total := Add(1, 2)
	assert.Equal(t, total, 2)

}

//notnil
func getDBconnection(t *testing.T) {

	client := db.GetDBconnection("mongodb+srv://aninoUser:3NmFiNQSN6VBvcz@cluster0.wftbh.mongodb.net/test?authSource=admin&replicaSet=atlas-a59at9-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true")
	assert.NotNil(t, client, "Not nil")
}

// https://godoc.org/github.com/stretchr/testify/assert
