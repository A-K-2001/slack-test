-- name: GetCartById :one
SELECT * from cart WHERE cart_id = $1;

-- name: GetCartByOrderId :one
SELECT * from cart WHERE order_id = $1;

-- name: GetCartWithCartItems :many
select * from cart as c join cart_item on c.cart_id = cart_item.cart_id where  order_id= $1;