
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS appointments(
    id bigserial UNIQUE NOT NULL PRIMARY KEY ,
    appointment_id uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    date timestamptz UNIQUE NOT NULL,
    comment varchar(1020)
--     patient_id bigserial NOT NULL REFERENCES  users(id) ON UPDATE CASCADE ON DELETE CASCADE,
--     doctor_id  bigserial NOT NULL REFERENCES  users(id) ON UPDATE CASCADE ON DELETE CASCADE

)