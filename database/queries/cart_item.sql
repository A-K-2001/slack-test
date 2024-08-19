-- name: ListCartItemsByCartID :many
SELECT * FROM cart_item WHERE cart_id = $1;
