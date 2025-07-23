package main

import (
	"encoding/json" 
	"time"

	"github.com/go-redis/redis/v8"
)


func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func AddBid(client *redis.Client, bookingID string, bid Bid) error {
	key := "bids:" + bookingID
	bidJSON, err := json.Marshal(bid)
	if err != nil {
		return err
	}
	err = client.RPush(ctx, key, bidJSON).Err()
	if err == nil {
		client.Expire(ctx, key, time.Hour) // optional expiration
	}
	return err
}

func GetBids(client *redis.Client, bookingID string) ([] Bid, error) {
	key := "bids:" + bookingID
	bidStrings, err := client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	var bids []Bid
	for _, s := range bidStrings {
		var bid Bid
		if err := json.Unmarshal([]byte(s), &bid); err == nil {
			bids = append(bids, bid)
		}
	}
	return bids, nil
}

func PublishBid(client *redis.Client, bookingID string, bid Bid) error {
	channel := "bids_channel:" + bookingID
	msg, err := json.Marshal(bid)
	if err != nil {
		return err
	}
	return client.Publish(ctx, channel, msg).Err()
}
