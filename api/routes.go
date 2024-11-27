package api

import (
	"github.com/JorgeSerrano26/api-example/api/clients/cat"
	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	StoreID  uint64 `form:"store_id"`
	SellerID uint64 `form:"seller_id"`
}

func URLMapping(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/origin", func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		c.JSON(200, gin.H{
			"host": origin,
		})
	})

	router.GET("/fact", func(c *gin.Context) {
		catFact, err := cat.Instance.GetFact()

		if err != nil {
			c.JSON(500, gin.H{})
		}

		c.JSON(200, catFact)
	})

	apiRouter := router.Group("/api")
	{
		apiRouter.GET("/test", TestHandler)
	}
}
