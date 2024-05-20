DROP INDEX IF EXISTS idx_users_id;
DROP INDEX IF EXISTS idx_users_nip;
DROP INDEX IF EXISTS idx_users_name_lower;
DROP INDEX IF EXISTS idx_users_created_at;

DROP TABLE IF EXISTS users CASCADE;
