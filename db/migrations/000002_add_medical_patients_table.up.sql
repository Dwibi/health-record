CREATE TYPE gender_enum AS ENUM ('male', 'female');

CREATE TABLE IF NOT EXISTS medical_patients (
    id SERIAL PRIMARY KEY,
    identity_number CHAR(16) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    name VARCHAR(255) NOT NULL,
    birth_date DATE NOT NULL,
    gender gender_enum NOT NULL,
    identity_card_scan_img VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
);