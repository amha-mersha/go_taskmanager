package route

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Run(port int) {
	router := gin.Default()
	router.Run("localhost:" + strconv.Itoa(port))
}
