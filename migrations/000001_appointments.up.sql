CREATE TABLE IF NOT EXISTS appointments(
    ID timestamptz UNIQUE NOT NULL,
    comment Text,
    patient_id varchar NOT NULL,
    doctor_id uuid NOT NULL

)