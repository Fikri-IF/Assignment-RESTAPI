package handler

import (
	"kominfo-assignment-2/infra/config"
	db "kominfo-assignment-2/infra/database"
	"kominfo-assignment-2/repository/order_repo/order_pg"
	"kominfo-assignment-2/service/order_service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	config.LoadAppConfig()
	db.InitiliazeDatabase()

	pgSql := db.GetDatabaseInstance()

	_ = pgSql

	orderRepo := order_pg.NewRepository(pgSql)

	_ = orderRepo

	orderService := order_service.NewService(orderRepo)

	oh := NewOrderHandler(orderService)

	_ = orderService

	route := gin.Default()

	route.POST("/orders", oh.CreateOrderWithItems)
	route.GET("/orders", oh.GetOrders)
	route.PUT("/orders/:orderId", oh.UpdateOrderWithItems)
	route.DELETE("/orders/:orderId", oh.DeleteOrder)

	route.Run(":8080")

	_ = route

}
