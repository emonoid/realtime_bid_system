
package main

import ( 
	"net/http" 

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
)
 
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func BidWebSocketHandler(redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookingID := c.Param("booking_id")
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		sub := redisClient.Subscribe(ctx, "bids_channel:"+bookingID)
		ch := sub.Channel()

		for msg := range ch {
			conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
		}
	}
}
