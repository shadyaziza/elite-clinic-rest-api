-- name: GetAppointment :one
SELECT * FROM appointments WHERE patient_id = $1  LIMIT 1;