DROP INDEX IF EXISTS idx_medical_records_identity_number_patient;
DROP INDEX IF EXISTS idx_medical_records_created_by;
DROP INDEX IF EXISTS idx_medical_records_created_at;

DROP TABLE IF EXISTS medical_records CASCADE;