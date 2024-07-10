-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Insert predefined users
INSERT INTO users (id, name, surname, email, password, phone, created_at, updated_at)
VALUES 
(uuid_generate_v4(), 'John Doe', 'get', 'john@example.com', 'hashedpassword1', '1234567890', now(), now()),
(uuid_generate_v4(), 'Jane Doe', 'get', 'jane@example.com', 'hashedpassword2', '0987654321', now(), now())
ON CONFLICT (email) DO NOTHING;