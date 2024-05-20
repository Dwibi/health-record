DROP INDEX IF EXISTS idx_medical_patients_identity_number;
DROP INDEX IF EXISTS idx_medical_patients_name_lower;
DROP INDEX IF EXISTS idx_medical_patients_phone_number;
DROP INDEX IF EXISTS idx_medical_patients_created_at;

DROP TABLE IF EXISTS medical_patients CASCADE;
DROP TYPE gender_enum CASCADE;

