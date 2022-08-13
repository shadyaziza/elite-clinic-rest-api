

CREATE TABLE IF NOT EXISTS appointments(
    id bigserial UNIQUE NOT NULL PRIMARY,
    appointment_id uuid UNIQUE NOT NULL,
    date timestamptz UNIQUE NOT NULL,
    comment varchar(1020),
    patient_id uuid NOT NULL REFRENCES user.id ON UPDATE CASCADE ON DELETE CASCADE,
    doctor_id uuid NOT NULL REFRENCES user.id ON UPDATE CASCADE ON DELETE CASCADE

)