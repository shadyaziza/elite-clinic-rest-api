CREATE TABLE IF NOT EXISTS appointments(
    ID timestamptz UNIQUE NOT NULL,
    comment Text,
    patientID uuid NOT NULL,
    doctorID uuid NOT NULL

)