package order_pg

const create_new_order = `
	INSERT INTO "orders"
	(ordered_at, customer_name)
	VALUES ($1, $2)
	RETURNING order_id
`

const create_new_item = `
	INSERT INTO "items"
	(item_code, description, quantity, order_id)
	VALUES($1, $2, $3, $4)
`

const get_all_orders = `
	SELECT order_id, ordered_at, customer_name, created_at, updated_at
	FROM "orders"
`

const update_order_by_id = `
	UPDATE "orders"
	SET ordered_at = $1, customer_name = $2, updated_at = now()
	WHERE order_id = $3
`

const update_item_by_id = `
	UPDATE "items"
	SET item_code = $1, description = $2, quantity = $3, updated_at = now()
	WHERE item_id = $4
`

const delete_order_by_id = `
	DELETE FROM "orders"
	WHERE order_id = $1
`

const check_item_exist = `
	SELECT item_id
	FROM "items"
	WHERE item_id = $1
`

const check_order_exist = `
	SELECT order_id
	FROM "orders"
	WHERE order_id = $1
`
