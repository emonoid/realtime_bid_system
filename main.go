
package main

import (
	"context" 

	"github.com/gin-gonic/gin" 
)

var ctx = context.Background()

func main() {
	r := gin.Default()
	redisClient := NewRedisClient()

	r.POST("/bid", func(c *gin.Context) {
		var bid Bid
		if err := c.BindJSON(&bid); err != nil {
			c.JSON(400, gin.H{"error": "invalid bid"})
			return
		}
		err :=  AddBid(redisClient, bid.BookingID, bid)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to save bid"})
			return
		}
		_ = PublishBid(redisClient, bid.BookingID, bid)
		c.JSON(200, gin.H{"status": "bid placed"})
	})

	r.GET("/bids/:booking_id", func(c *gin.Context) {
		bookingID := c.Param("booking_id")
		bids, err :=  GetBids(redisClient, bookingID)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to fetch bids"})
			return
		}
		c.JSON(200, bids)
	})

	r.GET("/ws/bids/:booking_id", BidWebSocketHandler(redisClient))

	r.Run(":8080")
}



 