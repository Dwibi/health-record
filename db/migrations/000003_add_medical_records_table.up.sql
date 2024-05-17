CREATE TABLE IF NOT EXISTS medical_records (
    id SERIAL PRIMARY KEY,
    patient_id INT NOT NULL REFERENCES medical_patients(identity_number),
    symptoms VARCHAR(2000) NOT NULL,
    medications VARCHAR(2000) NOT NULL,
    created_by INT NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);