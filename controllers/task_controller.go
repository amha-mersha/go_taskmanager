package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/amha-mersha/go_taskmanager/data"
	"github.com/amha-mersha/go_taskmanager/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var tasks = make(map[int64]models.Task)

func getTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, tasks)
}

func getTaskByID(ctx *gin.Context) {
	taskID, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}
	if _, exist := tasks[taskID]; !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": data.IDNotFound})
		return
	}
	ctx.JSON(http.StatusOK, tasks[taskID])
}

func udpateTask(ctx *gin.Context) {
	taskID, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}
	if _, exist := tasks[taskID]; !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": data.IDNotFound})
		return
	}
	updatedTask := models.Task{}
	if err = ctx.ShouldBindJSON(&updatedTask); err != nil {
		switch e := err.(type) {
		case *json.SyntaxError:
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": data.MalformedJSON})
		case *json.UnmarshalTypeError:
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": data.MismatchedFormat})
		case validator.ValidationErrors:
			missingRequireds := []string{}
			for _, fieldError := range e {
				missingRequireds = append(missingRequireds, fieldError.Error())
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": data.MissingRequireds, "Missing": missingRequireds})

		}
		ctx.JSON(http.StatusOK, tasks[taskID])
	}
}
func deleteTask(ctx *gin.Context) {

}

func postTask(ctx *gin.Context) {

}
