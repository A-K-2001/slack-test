-- -- name: GetSavedAddressByID :one
-- SELECT * FROM meds_savedAddress WHERE id = $1;

-- -- name: ListAddressesByUserID :many
-- SELECT * FROM meds_savedAddress WHERE user_id = $1;

-- -- name: GetAddressByUserAndLabel :one
-- SELECT * FROM meds_savedAddress WHERE user_id = $1 AND label = $2;