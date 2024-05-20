CREATE TABLE IF NOT EXISTS medical_records (
    id SERIAL PRIMARY KEY,
    identity_number_patient CHAR(16) NOT NULL REFERENCES medical_patients(identity_number),
    symptoms VARCHAR(2000) NOT NULL,
    medications VARCHAR(2000) NOT NULL,
    created_by INT NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_medical_records_identity_number_patient ON medical_records (identity_number_patient);
CREATE INDEX IF NOT EXISTS idx_medical_records_created_by ON medical_records (created_by);
CREATE INDEX IF NOT EXISTS idx_medical_records_created_at ON medical_records (created_at);