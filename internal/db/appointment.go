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
	PatientID     int64
	DoctorID      int64
}

// CreateAppointment - helper to transform database types
// to service types, this can not be defined inside service
// layer since it will create circular dependancy
func CreateAppointment(appointmentRow AppointmentRow) appointment.Appointment {
	return appointment.Appointment{
		ID:            appointmentRow.ID,
		AppointmentID: appointmentRow.AppointmentID,
		Date:          appointmentRow.Date,
		Comment:       appointmentRow.Comment.String,
		PatientID:     appointmentRow.PatientID,
		DoctorID:      appointmentRow.DoctorID,
	}
}

const getOneQuery = `SELECT * FROM appointments WHERE id = $1`

func (db *Database) GetAppointment(ctx context.Context, id int) (appointment.Appointment, error) {
	var appointmentRow AppointmentRow
	row := db.Client.QueryRowContext(ctx, getOneQuery, id)
	err := row.Scan(
		&appointmentRow.ID,
		&appointmentRow.AppointmentID,
		&appointmentRow.Date,
		&appointmentRow.Comment,
		&appointmentRow.PatientID,
		&appointmentRow.DoctorID,
	)
	if err != nil {
		return appointment.Appointment{}, fmt.Errorf("error fetching appointment by id %w", err)
	}

	return CreateAppointment(appointmentRow), nil
}
