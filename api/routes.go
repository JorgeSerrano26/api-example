package api

import (
	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	StoreID  uint64 `form:"store_id"`
	SellerID uint64 `form:"seller_id"`
}

func URLMapping(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiRouter := router.Group("/api")
	{
		apiRouter.GET("/test", TestHandler)
	}
}
