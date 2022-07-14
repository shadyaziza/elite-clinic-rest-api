CREATE TABLE IF NOT EXISTS appointments(
    ID timestamptz,
    comment Text,
    patientID uuid,
    doctorID uuid

)