package route

import (
	"strconv"

	"github.com/amha-mersha/go_taskmanager/data"
	"github.com/gin-gonic/gin"
)

func Run(port int) {
	router := gin.Default()

	router.GET("/api/v1/tasks", data.GetTasks)
	router.GET("/api/v1/tasks/:id", data.GetTaskByID)
	router.POST("/api/v1/tasks", data.PostTask)
	router.PUT("/api/v1/tasks/:id", data.UpdateTask)
	router.DELETE("/api/v1/tasks/:id", data.DeleteTask)

	router.Run("localhost:" + strconv.Itoa(port))
}
