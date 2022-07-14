

CREATE TABLE IF NOT EXISTS appointments(
    ID timestamptz UNIQUE NOT NULL,
    comment varchar(1020),
    patient_id uuid NOT NULL REFRENCES user.id ON UPDATE CASCADE ON DELETE CASCADE,
    doctor_id uuid NOT NULL REFRENCES user.id ON UPDATE CASCADE ON DELETE CASCADE

)