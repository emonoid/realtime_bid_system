package main

type Bid struct {
	ID           string `json:"id"`
	BookingID    string `json:"booking_id"`
	BidAmount    int    `json:"bid_amount"`
	DriverID     int64  `json:"driver_id"`
	DriverName   string `json:"driver_name"`
	DriverRating int    `json:"driver_rating"`
	DriverMobile string `json:"driver_mobile"`
	CarID        int64  `json:"car_id"`
	CarType      string `json:"car_type"`
	CarImage     string `json:"car_image"`
}