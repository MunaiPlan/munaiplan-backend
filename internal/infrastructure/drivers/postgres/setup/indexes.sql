CREATE UNIQUE INDEX IF NOT EXISTS idx_organizations_email ON organizations (email) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users (email) WHERE deleted_at IS NULL;
