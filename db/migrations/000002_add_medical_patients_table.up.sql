DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_enum') THEN
        CREATE TYPE gender_enum AS ENUM ('male', 'female');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS medical_patients (
    id SERIAL PRIMARY KEY,
    identity_number CHAR(16) NOT NULL UNIQUE,
    phone_number VARCHAR(20) NOT NULL,
    name VARCHAR(255) NOT NULL,
    birth_date DATE NOT NULL,
    gender gender_enum NOT NULL,
    identity_card_scan_img VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_medical_patients_identity_number ON medical_patients (identity_number);
CREATE INDEX IF NOT EXISTS idx_medical_patients_name_lower ON medical_patients (LOWER(name));
CREATE INDEX IF NOT EXISTS idx_medical_patients_phone_number ON medical_patients (phone_number);
CREATE INDEX IF NOT EXISTS idx_medical_patients_created_at ON medical_patients (created_at);