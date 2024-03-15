package handler

import (
	"kominfo-assignment-2/dto"
	"kominfo-assignment-2/service/order_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService order_service.Service
}

func NewOrderHandler(orderService order_service.Service) *orderHandler {
	return &orderHandler{
		orderService: orderService,
	}
}

func (o *orderHandler) CreateOrderWithItems(ctx *gin.Context) {
	var payload dto.CreateOrderRequestDto

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
			"message": "invalid request json",
		})
		return
	}

	if err := o.orderService.CreateOrderWithItems(payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]string{
		"message": "success",
	})
}

func (o *orderHandler) GetOrders(ctx *gin.Context) {
	orders, err := o.orderService.GetOrders()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (o *orderHandler) UpdateOrderWithItems(ctx *gin.Context) {
	var payload dto.UpdateOrderRequestDto

	orderId := ctx.Param("orderId")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
			"message": "invalid request json",
		})
		return
	}

	if err := o.orderService.UpdateOrderWithItems(payload, orderId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{
		"message": "update success",
	})
}

func (o *orderHandler) DeleteOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	if err := o.orderService.DeleteOrder(orderId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{
		"message": "delete success",
	})
}
