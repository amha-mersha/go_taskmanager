package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/amha-mersha/go_taskmanager/data"
	"github.com/amha-mersha/go_taskmanager/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func GetTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, data.GetTasks())
}

func GetTaskByID(ctx *gin.Context) {
	taskID, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}
	task, _ := data.GetTaskByID(taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
	}
	ctx.JSON(http.StatusOK, task)
}

func UpdateTask(ctx *gin.Context) {
	taskID, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}
	var updatedTask models.Task
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
	}
	err = data.UpdateTask(taskID, updatedTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
	}
	ctx.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(ctx *gin.Context) {
	taskID, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}
	err = data.DeleteTask(taskID)
}

func PostTask(ctx *gin.Context) {
	var newTask models.Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
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
	}
	err = data.PostTask(newTask)
	ctx.JSON(http.StatusAccepted, newTask)
}
