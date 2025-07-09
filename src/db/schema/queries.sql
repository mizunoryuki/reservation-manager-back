-- name: CreateUser :exec
INSERT INTO users (email, name, password_hash, role)
VALUES (?, ?, ?, 'general');

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = ?;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = ?;

-- name: CreateRefreshToken :exec
INSERT INTO refresh_tokens (user_id, token, expires_at)
VALUES (?, ?, ?)
ON DUPLICATE KEY UPDATE token = VALUES(token), expires_at = VALUES(expires_at), created_at = CURRENT_TIMESTAMP;

-- name: GetRefreshTokenByUserID :one
SELECT * FROM refresh_tokens
WHERE user_id = ?;

-- name: DeleteRefreshTokenByUserID :exec
DELETE FROM refresh_tokens
WHERE user_id = ?;

-- name: CreateStore :exec
INSERT INTO stores (name, address, business_start_time, business_end_time, details)
VALUES (?, ?, ?, ?, ?);

-- name: GetAllStores :many
SELECT * FROM stores;

-- name: GetStoreByID :one
SELECT * FROM stores
WHERE id = ?;

-- name: DeleteStore :exec
DELETE FROM stores
WHERE id = ?;

-- name: UpdateStore :exec
UPDATE stores
SET name = ?, address = ?, business_start_time = ?, business_end_time = ?, details = ?
WHERE id = ?;

-- name: CreateReservation :exec
INSERT INTO reservations (user_id, store_id, visit_date)
VALUES (?, ?, ?);

-- name: GetAllReservations :many
SELECT * FROM reservations
ORDER BY visit_date DESC;

-- name: GetReservationByID :one
SELECT * FROM reservations
WHERE id = ?;

-- name: GetReservationsByUser :many
SELECT * FROM reservations
WHERE user_id = ?
ORDER BY visit_date DESC;

-- name: GetReservationsByStoreAndDate :many
SELECT * FROM reservations
WHERE store_id = ? AND DATE(visit_date) = ?;

-- name: CancelReservation :exec
DELETE FROM reservations
WHERE id = ? AND user_id = ?;

-- name: DeleteReservationAsAdmin :exec
DELETE FROM reservations
WHERE id = ?;