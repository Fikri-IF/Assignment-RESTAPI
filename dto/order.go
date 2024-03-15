package dto

import (
	"kominfo-assignment-2/entity"
	"time"
)

type CreateOrderRequestDto struct {
	OrderedAt    time.Time              `json:"orderedAt"`
	CustomerName string                 `json:"customerName"`
	Items        []CreateItemRequestDto `json:"items"`
}

type UpdateOrderRequestDto struct {
	OrderedAt    time.Time              `json:"orderedAt"`
	CustomerName string                 `json:"customerName"`
	Items        []UpdateItemRequestDto `json:"items"`
}

type GetOrdersResponseDto struct {
	Orders []entity.Order `json:"orders"`
}
