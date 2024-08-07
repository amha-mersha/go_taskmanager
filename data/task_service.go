package data

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/amha-mersha/go_taskmanager/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TaskError struct {
	message string
}

func (err TaskError) Error() string {
	return err.message
}

const (
	IDNotFound       = "No item found with the specified ID."
	MalformedJSON    = "Sent a malfomed JSON."
	MismatchedFormat = "The task have a mismatched stucture."
	MissingRequireds = "There are some missing required feilds."
)

var tasks = make(map[int64]models.Task)

func GetTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, tasks)
}

func GetTaskByID(ctx *gin.Context) {
	taskID, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}
	if _, exist := tasks[taskID]; !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": IDNotFound})
		return
	}
	ctx.JSON(http.StatusOK, tasks[taskID])
}

func UpdateTask(ctx *gin.Context) {
	taskID, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}
	if _, exist := tasks[taskID]; !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": IDNotFound})
		return
	}
	var updatedTask models.Task
	if err = ctx.ShouldBindJSON(&updatedTask); err != nil {
		switch e := err.(type) {
		case *json.SyntaxError:
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": MalformedJSON})
		case *json.UnmarshalTypeError:
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": MismatchedFormat})
		case validator.ValidationErrors:
			missingRequireds := []string{}
			for _, fieldError := range e {
				missingRequireds = append(missingRequireds, fieldError.Error())
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": MissingRequireds, "Missing": missingRequireds})
		}
	}
	tasks[taskID] = updatedTask
	ctx.JSON(http.StatusOK, tasks[taskID])
}
func DeleteTask(ctx *gin.Context) {
	taskID, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}
	if _, exist := tasks[taskID]; !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": IDNotFound})
		return
	}
	delete(tasks, taskID)
}

func PostTask(ctx *gin.Context) {
	var newTask models.Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		switch e := err.(type) {
		case *json.SyntaxError:
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": MalformedJSON})
		case *json.UnmarshalTypeError:
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": MismatchedFormat})
		case validator.ValidationErrors:
			missingRequireds := []string{}
			for _, fieldError := range e {
				missingRequireds = append(missingRequireds, fieldError.Error())
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": MissingRequireds, "Missing": missingRequireds})
		}
	}
	tasks[int64(newTask.ID)] = newTask
	ctx.JSON(http.StatusAccepted, newTask)
}
