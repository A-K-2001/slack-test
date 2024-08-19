-- name: GetBatchByMedicineAndBatch :one
SELECT * FROM med_batch WHERE medicine_id = $1 AND batch = $2;

-- name: ListBatchesByMedicineID :many
SELECT * FROM med_batch WHERE medicine_id = $1;

-- name: ListExpiredBatches :many
SELECT * FROM med_batch WHERE expiry_date < CURRENT_DATE;