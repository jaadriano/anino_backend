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

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"JP Adriano Anino Exam": "API",
		})
	})

	user := new(controller.UserController)
	board := new(controller.BoardController)

	router.GET("/user/:id", user.RetrieveUser)
	router.POST("/user/", user.AddUser)

	router.POST("/admin/leaderboard/", board.AddBoard)
	router.GET("/leaderboard/:id", board.RetrieveBoard)
	router.PUT("/leaderboard/:id/user/:user_id/add_score", board.AddBoardScore)

	return router
}
