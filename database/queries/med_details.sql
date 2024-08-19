-- name: GetMedicineByID :one
SELECT * FROM med_details WHERE id = $1;


-- name: GetMedicineBySlug :one
SELECT * FROM med_details WHERE external_slug = $1;


-- name: SearchMedicinesByPartialName :many
SELECT med.id , med.name  FROM med_details as med WHERE name ILIKE '%' || $1 || '%' ORDER BY name ;

