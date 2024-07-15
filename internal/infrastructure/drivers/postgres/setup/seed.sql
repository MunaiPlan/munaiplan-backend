-- Insert predefined organizations
INSERT INTO organizations (id, name, email, phone, address, created_at, updated_at)
VALUES
(uuid_generate_v4(), 'KazMunaiGas', 'kazmunaigas@gmail.com', '1234567890', 'Kazakhstan', now(), now())
ON CONFLICT (email) WHERE deleted_at IS NULL DO NOTHING;