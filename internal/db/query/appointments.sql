-- name: GetAppointment :one
SELECT * FROM appointments WHERE id = $1  LIMIT 1;