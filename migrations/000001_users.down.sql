DROP INDEX IF EXISTS user_type_idx;
DROP INDEX IF EXISTS user_id_idx;
DROP INDEX IF EXISTS user_gender_idx;
DROP INDEX IF EXISTS user_phone_idx;

ALTER TABLE IF EXISTS users DROP CONSTRAINT IF EXISTS phone_number_constraints;

DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS gender;
DROP TYPE IF EXISTS user_type;