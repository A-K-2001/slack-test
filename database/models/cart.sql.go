// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: cart.sql

package models

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const getCartById = `-- name: GetCartById :one
SELECT cart_id, order_id, uniquer_floor, uniquer_round, address_id, cart_status, verification_status, agent_id, created_at, billed_at from cart WHERE cart_id = $1
`

func (q *Queries) GetCartById(ctx context.Context, cartID uuid.UUID) (Cart, error) {
	row := q.db.QueryRow(ctx, getCartById, cartID)
	var i Cart
	err := row.Scan(
		&i.CartID,
		&i.OrderID,
		&i.UniquerFloor,
		&i.UniquerRound,
		&i.AddressID,
		&i.CartStatus,
		&i.VerificationStatus,
		&i.AgentID,
		&i.CreatedAt,
		&i.BilledAt,
	)
	return i, err
}

const getCartByOrderId = `-- name: GetCartByOrderId :one
SELECT cart_id, order_id, uniquer_floor, uniquer_round, address_id, cart_status, verification_status, agent_id, created_at, billed_at from cart WHERE order_id = $1
`

func (q *Queries) GetCartByOrderId(ctx context.Context, orderID string) (Cart, error) {
	row := q.db.QueryRow(ctx, getCartByOrderId, orderID)
	var i Cart
	err := row.Scan(
		&i.CartID,
		&i.OrderID,
		&i.UniquerFloor,
		&i.UniquerRound,
		&i.AddressID,
		&i.CartStatus,
		&i.VerificationStatus,
		&i.AgentID,
		&i.CreatedAt,
		&i.BilledAt,
	)
	return i, err
}

const getCartWithCartItems = `-- name: GetCartWithCartItems :many
select c.cart_id, order_id, uniquer_floor, uniquer_round, address_id, cart_status, verification_status, agent_id, created_at, billed_at, id, cart_item.cart_id, item_type, item_name, discount, version, quantity, inventory_id, amount, cgst, igst, sgst from cart as c join cart_item on c.cart_id = cart_item.cart_id where  order_id= $1
`

type GetCartWithCartItemsRow struct {
	CartID             uuid.UUID          `db:"cart_id" json:"cart_id"`
	OrderID            string             `db:"order_id" json:"order_id"`
	UniquerFloor       *string            `db:"uniquer_floor" json:"uniquer_floor"`
	UniquerRound       *string            `db:"uniquer_round" json:"uniquer_round"`
	AddressID          *int32             `db:"address_id" json:"address_id"`
	CartStatus         string             `db:"cart_status" json:"cart_status"`
	VerificationStatus string             `db:"verification_status" json:"verification_status"`
	AgentID            *string            `db:"agent_id" json:"agent_id"`
	CreatedAt          pgtype.Timestamptz `db:"created_at" json:"created_at"`
	BilledAt           *int64             `db:"billed_at" json:"billed_at"`
	ID                 uuid.UUID          `db:"id" json:"id"`
	CartID_2           uuid.UUID          `db:"cart_id_2" json:"cart_id_2"`
	ItemType           string             `db:"item_type" json:"item_type"`
	ItemName           string             `db:"item_name" json:"item_name"`
	Discount           int32              `db:"discount" json:"discount"`
	Version            int32              `db:"version" json:"version"`
	Quantity           int32              `db:"quantity" json:"quantity"`
	InventoryID        *uuid.UUID         `db:"inventory_id" json:"inventory_id"`
	Amount             int32              `db:"amount" json:"amount"`
	Cgst               int32              `db:"cgst" json:"cgst"`
	Igst               int32              `db:"igst" json:"igst"`
	Sgst               int32              `db:"sgst" json:"sgst"`
}

func (q *Queries) GetCartWithCartItems(ctx context.Context, orderID string) ([]GetCartWithCartItemsRow, error) {
	rows, err := q.db.Query(ctx, getCartWithCartItems, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetCartWithCartItemsRow{}
	for rows.Next() {
		var i GetCartWithCartItemsRow
		if err := rows.Scan(
			&i.CartID,
			&i.OrderID,
			&i.UniquerFloor,
			&i.UniquerRound,
			&i.AddressID,
			&i.CartStatus,
			&i.VerificationStatus,
			&i.AgentID,
			&i.CreatedAt,
			&i.BilledAt,
			&i.ID,
			&i.CartID_2,
			&i.ItemType,
			&i.ItemName,
			&i.Discount,
			&i.Version,
			&i.Quantity,
			&i.InventoryID,
			&i.Amount,
			&i.Cgst,
			&i.Igst,
			&i.Sgst,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
