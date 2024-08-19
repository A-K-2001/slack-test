-- name: GetOrderAttributesByOrderID :one
SELECT * FROM med_order_attributes WHERE order_id = $1;

-- name: ListOrdersByUserID :many
SELECT * FROM med_order_attributes WHERE user_id = $1;

