package data

import (
	"github.com/amha-mersha/go_taskmanager/models"
)

type TaskError struct {
	message string
}

func (err TaskError) Error() string {
	return err.message
}

const (
	IDNotFound        = "No item found with the specified ID."
	TaskAlreadyExists = "The task already exists in the database"
	MalformedJSON     = "Sent a malfomed JSON."
	MismatchedFormat  = "The task have a mismatched stucture."
	MissingRequireds  = "There are some missing required feilds."
)

var tasks = make(map[int64]models.Task)

func GetTasks() map[int64]models.Task {
	return tasks
}

func GetTaskByID(taskID int64) (models.Task, error) {
	if _, exist := tasks[taskID]; !exist {
		return models.Task{}, TaskError{message: IDNotFound}
	}
	return tasks[taskID], nil
}

func UpdateTask(taskID int64, updatedTask models.Task) error {
	if _, exist := tasks[taskID]; !exist {
		return TaskError{message: IDNotFound}
	}
	tasks[taskID] = updatedTask
	return nil
}

func DeleteTask(taskID int64) error {
	if _, exist := tasks[taskID]; !exist {
		return TaskError{message: IDNotFound}
	}
	delete(tasks, taskID)
	return nil
}

func PostTask(newTask models.Task) error {
	if _, exist := tasks[int64(newTask.ID)]; exist {
		return TaskError{message: TaskAlreadyExists}
	}
	tasks[int64(newTask.ID)] = newTask
	return nil
}
