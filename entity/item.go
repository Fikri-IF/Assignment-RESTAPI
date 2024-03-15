package entity

import "time"

type Item struct {
	ItemId      uint
	ItemCode    string
	Quantity    uint
	Description string
	OrderId     uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
