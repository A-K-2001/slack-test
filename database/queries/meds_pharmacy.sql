-- name: GetPharmacyByID :one
SELECT * FROM meds_pharmacy WHERE id = $1;

-- name: ListActivePharmacies :many
SELECT * FROM meds_pharmacy WHERE is_active = true;

-- name: GetPharmacyByDLNumber :one
SELECT * FROM meds_pharmacy WHERE dl_number = $1;

-- name: ListPharmaciesByShippingPartner :many
SELECT * FROM meds_pharmacy WHERE shipping_partner = $1;