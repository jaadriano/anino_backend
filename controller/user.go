package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaadriano/anino_backend/models"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) RetrieveUser(ctx *gin.Context) {
	if ctx.Param("id") != "" {
		user, err := userModel.GetByID(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "User not found", "error": err})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"user": user})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	ctx.Abort()
	return
}

func (u UserController) AddUser(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	var name string
	if len(query["name"]) == 1 {
		name = query["name"][0]
	} else {
		body, err := ioutil.ReadAll(ctx.Request.Body)
		user := models.User{}
		err = json.Unmarshal(body, &user)
		name = user.Name
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
			ctx.Abort()
			return
		}
	}
	if len(name) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		ctx.Abort()
		return
	}
	user, err := userModel.PostUser(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "User not found", "error": err})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
	return
}
