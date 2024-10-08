// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: meds_pharmacy.sql

package models

import (
	"context"

	"github.com/google/uuid"
)

const getPharmacyByDLNumber = `-- name: GetPharmacyByDLNumber :one
SELECT id, name, latitude, longitude, address, phone, is_active, printer_id, dl_number, fssai_number, gst_number, pan_number, shipping_partner, pref_otp_enabled_delivery, dummy_cart_id FROM meds_pharmacy WHERE dl_number = $1
`

func (q *Queries) GetPharmacyByDLNumber(ctx context.Context, dlNumber *string) (MedsPharmacy, error) {
	row := q.db.QueryRow(ctx, getPharmacyByDLNumber, dlNumber)
	var i MedsPharmacy
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Latitude,
		&i.Longitude,
		&i.Address,
		&i.Phone,
		&i.IsActive,
		&i.PrinterID,
		&i.DlNumber,
		&i.FssaiNumber,
		&i.GstNumber,
		&i.PanNumber,
		&i.ShippingPartner,
		&i.PrefOtpEnabledDelivery,
		&i.DummyCartID,
	)
	return i, err
}

const getPharmacyByID = `-- name: GetPharmacyByID :one
SELECT id, name, latitude, longitude, address, phone, is_active, printer_id, dl_number, fssai_number, gst_number, pan_number, shipping_partner, pref_otp_enabled_delivery, dummy_cart_id FROM meds_pharmacy WHERE id = $1
`

func (q *Queries) GetPharmacyByID(ctx context.Context, id uuid.UUID) (MedsPharmacy, error) {
	row := q.db.QueryRow(ctx, getPharmacyByID, id)
	var i MedsPharmacy
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Latitude,
		&i.Longitude,
		&i.Address,
		&i.Phone,
		&i.IsActive,
		&i.PrinterID,
		&i.DlNumber,
		&i.FssaiNumber,
		&i.GstNumber,
		&i.PanNumber,
		&i.ShippingPartner,
		&i.PrefOtpEnabledDelivery,
		&i.DummyCartID,
	)
	return i, err
}

const listActivePharmacies = `-- name: ListActivePharmacies :many
SELECT id, name, latitude, longitude, address, phone, is_active, printer_id, dl_number, fssai_number, gst_number, pan_number, shipping_partner, pref_otp_enabled_delivery, dummy_cart_id FROM meds_pharmacy WHERE is_active = true
`

func (q *Queries) ListActivePharmacies(ctx context.Context) ([]MedsPharmacy, error) {
	rows, err := q.db.Query(ctx, listActivePharmacies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []MedsPharmacy{}
	for rows.Next() {
		var i MedsPharmacy
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Latitude,
			&i.Longitude,
			&i.Address,
			&i.Phone,
			&i.IsActive,
			&i.PrinterID,
			&i.DlNumber,
			&i.FssaiNumber,
			&i.GstNumber,
			&i.PanNumber,
			&i.ShippingPartner,
			&i.PrefOtpEnabledDelivery,
			&i.DummyCartID,
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

const listPharmaciesByShippingPartner = `-- name: ListPharmaciesByShippingPartner :many
SELECT id, name, latitude, longitude, address, phone, is_active, printer_id, dl_number, fssai_number, gst_number, pan_number, shipping_partner, pref_otp_enabled_delivery, dummy_cart_id FROM meds_pharmacy WHERE shipping_partner = $1
`

func (q *Queries) ListPharmaciesByShippingPartner(ctx context.Context, shippingPartner string) ([]MedsPharmacy, error) {
	rows, err := q.db.Query(ctx, listPharmaciesByShippingPartner, shippingPartner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []MedsPharmacy{}
	for rows.Next() {
		var i MedsPharmacy
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Latitude,
			&i.Longitude,
			&i.Address,
			&i.Phone,
			&i.IsActive,
			&i.PrinterID,
			&i.DlNumber,
			&i.FssaiNumber,
			&i.GstNumber,
			&i.PanNumber,
			&i.ShippingPartner,
			&i.PrefOtpEnabledDelivery,
			&i.DummyCartID,
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
