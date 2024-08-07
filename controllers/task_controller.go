package controllers

import (
	"bufio"
	"fmt"
	"os"
)

var availableCommands = map[string]string{
	"Add a task to the list":       "ADD",
	"Remove a task from the list":  "REMOVE",
	"Mark a task as completed":     "COMPLETE",
	"Mark a task as not completed": "PENDING",
	"View all tasks":               "VIEW_ALL",
	"View completed tasks":         "VIEW_COMPLETED",
	"View pending tasks":           "VIEW_PENDING",
	"Quit out of the system":       "QUIT",
}

var CommandtoHandler = map[string]func(models.Library) error{
	"ADD":            handleAddTask,
	"REMOVE":         handleRemoveTask,
	"COMPLETE":       handleCompleteTask,
	"PENDING":        handlePendingTask,
	"VIEW_ALL":       handleViewAllTasks,
	"VIEW_COMPLETED": handleViewCompletedTasks,
	"VIEW_PENDING":   handleViewPendingTasks,
	"QUIT":           handleQuit,
}
var reader = bufio.NewReader(os.Stdin)

func Run() {
	fmt.Println("Hello there, this is a library managment system.")
	fmt.Println("Here are the available commands in this library:")
	maxDescWidth := 0
	for _, cmd := range availableCommands {
		if maxDescWidth < len(cmd) {
			maxDescWidth = len(cmd)
		}
	}

	for desc, cmd := range availableCommands {
		fmt.Printf("\tTo %-*v : %v.\n", maxDescWidth, desc, cmd)
	}
	fmt.Println("Enter your command after >")
}
