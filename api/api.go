package api

import "github.com/gin-gonic/gin"

// Handlers for the API
func TestHandler(c *gin.Context) {
	var params QueryParams



	if err := c.ShouldBind(&params); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid query parameters",
		})
		return
	}

	c.JSON(200, gin.H{
		"storeId":  params.StoreID,
		"sellerId": params.SellerID,
	})
}
