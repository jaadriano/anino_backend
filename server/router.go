package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jaadriano/anino_backend/controller"
	"github.com/jaadriano/anino_backend/db"
)

func NewRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Recovery()) //, middleware.Logger(), gindump.Dump())

	db.Init()
	client := db.GetDB()
	if client == nil {
		router.GET("/", func(ctx *gin.Context) {
			ctx.JSON(500, gin.H{
				"MongoDB Atlass connection failed": "error",
			})
		})
	}

	user := new(controller.UserController)

	router.GET("/user/:id", user.RetrieveUser)
	router.POST("/user/", user.AddUser)

	return router
}
