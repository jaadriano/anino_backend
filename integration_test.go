package main

import (
	"anino_backend/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

//notnil
func TestDBConnection(t *testing.T) {

	client := db.GetDBconnection("mongodb+srv://aninoUser:3NmFiNQSN6VBvcz@cluster0.wftbh.mongodb.net/test?authSource=admin&replicaSet=atlas-a59at9-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true")
	assert.NotNil(t, client, "Not nil")
}

//notnil
func TestPostUser(t *testing.T) {

	client := db.GetDBconnection("mongodb+srv://aninoUser:3NmFiNQSN6VBvcz@cluster0.wftbh.mongodb.net/test?authSource=admin&replicaSet=atlas-a59at9-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true")
	assert.NotNil(t, client, "Not nil")
}

//notnil
func TestGetUserByID(t *testing.T) {

	client := db.GetDBconnection("mongodb+srv://aninoUser:3NmFiNQSN6VBvcz@cluster0.wftbh.mongodb.net/test?authSource=admin&replicaSet=atlas-a59at9-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true")
	assert.NotNil(t, client, "Not nil")
}

// https://godoc.org/github.com/stretchr/testify/assert
