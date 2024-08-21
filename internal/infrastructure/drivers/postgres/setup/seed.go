package setup

var Data = `
INSERT INTO users (id, organization_id, name, surname, email, password, phone, created_at, updated_at)
VALUES 
(uuid_generate_v4(), 'your-organization-id', 'John', 'Doe', 'john@example.com', $1, '1234567890', now(), now()),
(uuid_generate_v4(), 'your-organization-id', 'Jane', 'Smith', 'jane@example.com', $2, '0987654321', now(), now()),
(uuid_generate_v4(), 'your-organization-id', 'Alice', 'Johnson', 'alice@example.com', $3, '1122334455', now(), now())
ON CONFLICT (email) WHERE deleted_at IS NULL DO NOTHING;

INSERT INTO companies (id, organization_id, name, division, "group", representative, address, phone, created_at, updated_at)
VALUES 
(uuid_generate_v4(), 'your-organization-id', 'Company A', 'Division A', 'Group A', 'John Rep', 'Address A', '1111111111', now(), now()),
(uuid_generate_v4(), 'your-organization-id', 'Company B', 'Division B', 'Group B', 'Jane Rep', 'Address B', '2222222222', now(), now()),
(uuid_generate_v4(), 'your-organization-id', 'Company C', 'Division C', 'Group C', 'Alice Rep', 'Address C', '3333333333', now(), now()),
(uuid_generate_v4(), 'your-organization-id', 'Company D', 'Division D', 'Group D', 'Bob Rep', 'Address D', '4444444444', now(), now())
ON CONFLICT (name) WHERE deleted_at IS NULL DO NOTHING;
`
