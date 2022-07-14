
-- Extensions declaration
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


-- Types declaration
CREATE TYPE user_type AS ENUM ('doctor', 'patient');
CREATE TYPE gender AS ENUM ('male','female');

-- Table declaration
CREATE TABLE IF NOT EXISTS users (
    id uuid UNIQUE NOT NULL,
    type user_type NOT NULL DEFAULT ('patient'),
    phone_number varchar(50) NOT NULL UNIQUE,
    name varchar(50) NOT NULL,
    email varchar(100) UNIQUE NOT NULL,
    hashed_password VARCHAR NOT NULL,
    gender gender NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (NOW()),
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    phone_verified BOOLEAN NOT NULL DEFAULT FALSE,
    password_changed_at timestamptz NOT NULL DEFAULT ('0001-01-01 00:00:00Z')

);

-- Constraints
ALTER TABLE users ADD CONSTRAINT phone_number_constraint CHECK (phone_number ~ '^[+]201[0125][0-9]{8}$');


-- Indices
CREATE INDEX users_phone_idx ON users(phone_number);
CREATE INDEX users_gender_idx ON users(gender);
CREATE INDEX users_id_idx ON users(id);
CREATE INDEX users_type_idx ON users(type);


