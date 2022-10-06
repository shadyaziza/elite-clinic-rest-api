package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/shadyaziza/elite-clinic-rest-api/internal/appointment"
	"time"
)

type AppointmentRow struct {
	ID            int64
	AppointmentID uuid.UUID
	Date          time.Time
	Comment       sql.NullString
	//PatientID     int64
	//DoctorID      int64
}

// convertAppointmentRowToAppointment - helper to transform database types
// to service types, this can not be defined inside service
// layer since it will create circular dependancy
func convertAppointmentRowToAppointment(appointmentRow AppointmentRow) appointment.Appointment {
	return appointment.Appointment{
		ID:            appointmentRow.ID,
		AppointmentID: appointmentRow.AppointmentID,
		Date:          appointmentRow.Date,
		Comment:       appointmentRow.Comment.String,
		//PatientID:     appointmentRow.PatientID,
		//DoctorID:      appointmentRow.DoctorID,
	}
}

const selectOneQuery = `SELECT * FROM appointments WHERE id = $1`

func (db *Database) GetAppointment(ctx context.Context, id int) (appointment.Appointment, error) {
	var appointmentRow AppointmentRow
	row := db.Client.QueryRowContext(ctx, selectOneQuery, id)
	err := row.Scan(
		&appointmentRow.ID,
		&appointmentRow.AppointmentID,
		&appointmentRow.Date,
		&appointmentRow.Comment,
		//&appointmentRow.PatientID,
		//&appointmentRow.DoctorID,
	)
	if err != nil {
		return appointment.Appointment{}, fmt.Errorf("error fetching appointment by id %w", err)
	}

	return convertAppointmentRowToAppointment(appointmentRow), nil
}

const insertOneQuery = `INSERT INTO appointments (date,comment) VALUES ($1,$2) RETURNING *`

func (db *Database) PostAppointment(ctx context.Context,
	req appointment.CreateNewAppointmentRequest) (appointment.Appointment, error) {
	var appointmentRow AppointmentRow
	insertionRow := AppointmentRow{
		Date:    req.Date,
		Comment: sql.NullString{String: req.Comment, Valid: true},
	}
	rows := db.Client.QueryRowContext(ctx, insertOneQuery, insertionRow.Date, insertionRow.Comment)

	err := rows.Scan(
		&appointmentRow.ID,
		&appointmentRow.AppointmentID,
		&appointmentRow.Date,
		&appointmentRow.Comment)

	if err != nil {
		return appointment.Appointment{}, fmt.Errorf("error inserting new appointment %w", err)
	}

	return convertAppointmentRowToAppointment(appointmentRow), nil
}

const deleteOneQuery = `DELETE FROM appointments WHERE id = $1`

func (db *Database) DeleteAppointment(ctx context.Context, id int) error {
	_, err := db.Client.ExecContext(ctx, deleteOneQuery, id)
	if err != nil {
		return fmt.Errorf("error deleting an appointment %w", err)
	}
	return err
}

const updateQuery = `UPDATE appointments SET date = $1 , comment = $2 WHERE id=$3 RETURNING *`

func (db *Database) UpdateAppointment(ctx context.Context, req appointment.UpdateNewAppointmentRequest) (appointment.Appointment, error) {
	var appointmentRow AppointmentRow
	insertionRow := AppointmentRow{
		ID:      req.ID,
		Date:    req.Date,
		Comment: sql.NullString{String: req.Comment, Valid: true},
	}
	rows := db.Client.QueryRowContext(ctx, updateQuery, insertionRow.Date, insertionRow.Comment, insertionRow.ID)

	err := rows.Scan(
		&appointmentRow.ID,
		&appointmentRow.AppointmentID,
		&appointmentRow.Date,
		&appointmentRow.Comment)

	if err != nil {
		return appointment.Appointment{}, fmt.Errorf("error updating appointment %d %w", req.ID, err)
	}

	return convertAppointmentRowToAppointment(appointmentRow), nil
}
