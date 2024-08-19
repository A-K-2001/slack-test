-- name: GetStockByPharmacyAndMedicine :one
SELECT * FROM med_stock WHERE pharmacy_id = $1 AND medicine_id = $2 AND batch = $3;


-- name: GetTotalStockForMedicine :one
SELECT SUM(stock) as total_stock FROM med_stock WHERE medicine_id = $1;