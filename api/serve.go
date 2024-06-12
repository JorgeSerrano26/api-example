package api

import (
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()

	URLMapping(router)

	if err := router.Run(":8080"); err != nil {
		print(err)
	}
}
