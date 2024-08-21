DO
$$
DECLARE
    org_id UUID;
BEGIN
    -- Ensure pgcrypto extension is available
    CREATE EXTENSION IF NOT EXISTS pgcrypto;

    -- Ensure the unique constraints exist
    CREATE UNIQUE INDEX IF NOT EXISTS idx_organizations_email ON organizations (email) WHERE deleted_at IS NULL;
    CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users (email) WHERE deleted_at IS NULL;
    CREATE UNIQUE INDEX IF NOT EXISTS idx_companies_name ON companies (name) WHERE deleted_at IS NULL;

    -- Insert Organization or fetch the existing one
    INSERT INTO organizations (id, name, email, phone, address, created_at, updated_at)
    VALUES (
        uuid_generate_v4(), 
        'KazMunaiGas', 
        'kazmunaigas@gmail.com', 
        '1234567890', 
        'Kazakhstan', 
        now(), 
        now()
    )
    ON CONFLICT (email) WHERE deleted_at IS NULL DO NOTHING;

    -- Fetch the organization ID
    SELECT id INTO org_id FROM organizations WHERE email = 'kazmunaigas@gmail.com' LIMIT 1;

    -- Insert Users
    INSERT INTO users (id, organization_id, name, surname, email, password, phone, created_at, updated_at)
    VALUES
    (uuid_generate_v4(), org_id, 'John', 'Doe', 'test1@gmail.com', crypt('password1', gen_salt('bf')), '1111111111', now(), now()),
    (uuid_generate_v4(), org_id, 'Jane', 'Smith', 'test2@gmail.com', crypt('password2', gen_salt('bf')), '2222222222', now(), now()),
    (uuid_generate_v4(), org_id, 'Alex', 'Johnson', 'test2johnson@gmail.com', crypt('password3', gen_salt('bf')), '3333333333', now(), now())
    ON CONFLICT (email) WHERE deleted_at IS NULL DO NOTHING;

    -- Insert Companies
    INSERT INTO companies (id, organization_id, name, division, "group", representative, address, phone, created_at, updated_at)
    VALUES
    (uuid_generate_v4(), org_id, 'KazMunaiGas Exploration', 'Exploration', 'A', 'John Doe', 'Almaty, Kazakhstan', '7771112233', now(), now()),
    (uuid_generate_v4(), org_id, 'KazMunaiGas Refining', 'Refining', 'B', 'Jane Smith', 'Astana, Kazakhstan', '7771112234', now(), now()),
    (uuid_generate_v4(), org_id, 'KazMunaiGas Production', 'Production', 'C', 'Alex Johnson', 'Atyrau, Kazakhstan', '7771112235', now(), now()),
    (uuid_generate_v4(), org_id, 'KazMunaiGas Marketing', 'Marketing', 'D', 'Jane Smith', 'Aktau, Kazakhstan', '7771112236', now(), now())
    ON CONFLICT (name) WHERE deleted_at IS NULL DO NOTHING;

END
$$;

