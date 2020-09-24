package main

import (
	"github.com/jaadriano/anino_backend/config"
	"github.com/jaadriano/anino_backend/db"
	"github.com/jaadriano/anino_backend/server"
)

func main() {

	config.Init()
	server.Init()
	db.Init()
}
