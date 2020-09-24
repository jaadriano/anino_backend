package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jaadriano/anino_backend/models"
)

type BoardController struct{}

var boardModel = new(models.Board)
var scoreModel = new(models.Score)

func (u BoardController) RetrieveBoard(ctx *gin.Context) {
	board, err := boardModel.GetByID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Board not found", "error": err})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"board": board})
	return
}

func (u BoardController) AddBoard(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	var name string
	if len(query["name"]) == 1 {
		name = query["name"][0]
	} else {
		body, err := ioutil.ReadAll(ctx.Request.Body)
		board := models.Board{}
		err = json.Unmarshal(body, &board)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
			ctx.Abort()
			return
		}
		name = board.Name
	}
	if len(name) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		ctx.Abort()
		return
	}
	board, err := boardModel.PostBoard(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Board not found", "error": err})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"board": board})
	return
}

func (u BoardController) AddBoardScore(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	score := models.Score{
		BoardId:  ctx.Param("id"),
		ScoredAt: time.Now(),
		Score:    0,
		UserId:   ctx.Param("user_id")}
	if len(query["score_to_add"]) == 1 {
		scoreValue, _ := strconv.Atoi(query["score_to_add"][0])
		score.Score = scoreValue
	} else {
		body, err := ioutil.ReadAll(ctx.Request.Body)
		scoreVal := models.ScoreToAdd{}
		err = json.Unmarshal(body, &scoreVal)
		score.Score = scoreVal.Score
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
			ctx.Abort()
			return
		}
	}
	userExists, err := userModel.GetByID(score.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "User not found", "error": err})
		ctx.Abort()
		return
	}
	boardExists, err := boardModel.GetByID(score.BoardId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Board not found", "error": err})
		ctx.Abort()
		return
	}
	fmt.Sprintln(userExists, boardExists)
	scoreBoard, err := scoreModel.PostScore(score)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Score Post failed", "error": err})
		ctx.Abort()
		return
	}
	//append to board entries
	ctx.JSON(http.StatusOK, gin.H{"entry": scoreBoard})
	return
}
