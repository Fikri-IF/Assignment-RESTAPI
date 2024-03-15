package order_pg

import (
	"database/sql"
	"kominfo-assignment-2/entity"
	"kominfo-assignment-2/repository/order_repo"
)

type orderPG struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) order_repo.Repository {
	return &orderPG{
		db: db,
	}
}

func (o *orderPG) CreateOrderWithItems(order entity.Order, items []entity.Item) error {

	tx, err := o.db.Begin()

	if err != nil {
		return err
	}

	var orderId uint

	err = tx.QueryRow(
		create_new_order,
		order.OrderedAt,
		order.CustomerName,
	).Scan(&orderId)

	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range items {
		_, err = tx.Exec(
			create_new_item,
			item.ItemCode,
			item.Description,
			item.Quantity,
			orderId,
		)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (o *orderPG) GetOrders() ([]entity.Order, error) {
	rows, err := o.db.Query(get_all_orders)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orders := []entity.Order{}

	for rows.Next() {
		order := entity.Order{}
		err := rows.Scan(
			&order.OrderId,
			&order.OrderedAt,
			&order.CustomerName,
			&order.CreatedAt,
			&order.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (o *orderPG) UpdateOrderWithItems(order entity.Order, items []entity.Item, orderId string) error {
	tx, err := o.db.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec(
		update_order_by_id,
		order.OrderedAt,
		order.CustomerName,
		orderId,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range items {

		var itemId uint
		err = tx.QueryRow(check_item_exist, item.ItemId).Scan(&itemId)

		if err != nil {
			_, err = tx.Exec(
				create_new_item,
				item.ItemCode,
				item.Description,
				item.Quantity,
				orderId,
			)
		} else {
			_, err = tx.Exec(
				update_item_by_id,
				item.ItemCode,
				item.Description,
				item.Quantity,
				item.ItemId,
			)
		}

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (o *orderPG) DeleteOrder(orderId string) error {
	tx, err := o.db.Begin()

	if err != nil {
		return err
	}

	// Check if order exist
	var orderIdExist uint
	err = tx.QueryRow(check_order_exist, orderId).Scan(&orderIdExist)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(delete_order_by_id, orderId)

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
