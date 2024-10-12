DO
$$
BEGIN
    -- Ensure pgcrypto extension is available
    CREATE EXTENSION IF NOT EXISTS pgcrypto;

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
    DECLARE
        org_id UUID;
    BEGIN
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

        -- Insert Strings
        BEGIN

            -- Insert Sections
            INSERT INTO library_sections (id, description, manufacturer, type, body_od, body_id, avg_joint_length, 
                                  stabilizer_length, stabilizer_od, stabilizer_id, weight, material, grade, class, 
                                  friction_coefficient, min_yield_strength, created_at, updated_at)
            VALUES 
                -- Add rows based on your data, 24 rows in total
                (uuid_generate_v4(), 'Drill Pipe section', 'ZMS', 'Drill Pipe', 127, 108.61, 9.144, 0.433, 152.4, 82.55, 32.62, 'CS_API 5D/7', 'G', 2, 0.25, 105, now(), now()),
                (uuid_generate_v4(), 'Drill Pipe section', 'ZMS', 'Drill Pipe', 127, 108.61, 9.144, 0.433, 152.4, 82.55, 32.62, 'CS_API 5D/7', 'G', 2, 0.25, 105, now(), now()),
                (uuid_generate_v4(), 'Drill Pipe section', 'ZMS', 'Drill Pipe', 127, 108.61, 9.144, 0.433, 152.4, 82.55, 32.62, 'CS_API 5D/7', 'G', 2, 0.3, 105, now(), now()),
                (uuid_generate_v4(), 'Drill Pipe section', 'ZMS', 'Drill Pipe', 127, 108.61, 9.144, 0.433, 152.4, 82.55, 32.62, 'CS_API 5D/7', 'G', 2, 0.3, 105, now(), now()),
                (uuid_generate_v4(), 'Drill Pipe section', 'ZMS', 'Drill Pipe', 127, 108.61, 9.144, 0.433, 152.4, 82.55, 32.62, 'CS_API 5D/7', 'G', 2, 0.3, 105, now(), now()),
                (uuid_generate_v4(), 'Drill Pipe section', 'ZMS', 'Drill Pipe', 127, 108.61, 9.144, 0.433, 152.4, 82.55, 32.62, 'CS_API 5D/7', 'G', 2, 0.3, 105, now(), now()),
                (uuid_generate_v4(), 'Drill Pipe section', 'ZMS', 'Drill Pipe', 127, 108.61, 9.144, 0.433, 152.4, 82.55, 32.62, 'CS_API 5D/7', 'G', 2, 0.3, 105, now(), now()),
                (uuid_generate_v4(), 'Heavy Weight section', 'ZMS', 'Heavy Weight', 127, 76.2, 9.14, 1.219, 165.1, 76.2, 73.13, 'CS_1340 MOD', '1340 MOD', NULL, 0.3, 55, now(), now()),
                (uuid_generate_v4(), 'Heavy Weight section', 'ZMS', 'Heavy Weight', 127, 76.2, 9.14, 1.219, 165.1, 76.2, 73.13, 'CS_1340 MOD', '1340 MOD', NULL, 0.3, 55, now(), now()),
                (uuid_generate_v4(), 'Jar section', 'ZMS', 'Jar', 165.1, 69.85, 10.058, NULL, NULL, NULL, 136.6, 'CS_API 5D/7', '4145H MOD', NULL, 0.3, 110, now(), now()),
                (uuid_generate_v4(), 'Heavy Weight section', 'ZMS', 'Heavy Weight', 127, 76.2, 9.14, 1.219, 165.1, 76.2, 73.13, 'CS_1340 MOD', '1340 MOD', NULL, 0.3, 55, now(), now()),
                (uuid_generate_v4(), 'Heavy Weight section', 'ZMS', 'Heavy Weight', 127, 76.2, 9.14, 1.219, 165.1, 76.2, 73.13, 'CS_1340 MOD', '1340 MOD', NULL, 0.3, 55, now(), now()),
                (uuid_generate_v4(), 'Heavy Weight section', 'ZMS', 'Heavy Weight', 127, 76.2, 9.14, 1.219, 165.1, 76.2, 73.13, 'CS_1340 MOD', '1340 MOD', NULL, 0.3, 55, now(), now()),
                (uuid_generate_v4(), 'Heavy Weight section', 'ZMS', 'Heavy Weight', 127, 76.2, 9.14, 1.219, 165.1, 76.2, 73.13, 'CS_1340 MOD', '1340 MOD', NULL, 0.3, 55, now(), now()),
                (uuid_generate_v4(), 'Drill Pipe section', 'ZMS', 'Drill Pipe', 127, 108.61, 9.144, 0.433, 152.4, 82.55, 32.62, 'CS_API 5D/7', 'G', 2, 0.3, 105, now(), now()),
                (uuid_generate_v4(), 'Drill Pipe section', 'ZMS', 'Drill Pipe', 127, 108.61, 9.144, 0.433, 152.4, 82.55, 32.62, 'CS_API 5D/7', 'G', 2, 0.3, 105, now(), now()),
                (uuid_generate_v4(), 'Drill Pipe section', 'ZMS', 'Drill Pipe', 127, 108.61, 9.144, 0.433, 152.4, 82.55, 32.62, 'CS_API 5D/7', 'G', 2, 0.3, 105, now(), now()),
                (uuid_generate_v4(), 'Drill Pipe section', 'ZMS', 'Drill Pipe', 127, 108.61, 9.144, 0.433, 152.4, 82.55, 32.62, 'CS_API 5D/7', 'G', 2, 0.3, 105, now(), now()),
                (uuid_generate_v4(), 'Drill Pipe section', 'ZMS', 'Drill Pipe', 127, 108.61, 9.144, 0.433, 152.4, 82.55, 32.62, 'CS_API 5D/7', 'G', 2, 0.3, 105, now(), now()),
                (uuid_generate_v4(), 'Drill Pipe section', 'ZMS', 'Drill Pipe', 127, 108.61, 9.144, 0.433, 152.4, 82.55, 32.62, 'CS_API 5D/7', 'G', 2, 0.3, 105, now(), now()),
                (uuid_generate_v4(), 'Drill Pipe section', 'ZMS', 'Drill Pipe', 127, 108.61, 9.144, 0.433, 152.4, 82.55, 32.62, 'CS_API 5D/7', 'G', 2, 0.3, 105, now(), now()),
                (uuid_generate_v4(), 'MWD section', 'ZMS', 'MWD', 172, 83, 10.4, NULL, NULL, NULL, 149.77, 'SS_15-15LC', '15-15LC MOD (1)', NULL, 0.3, 110, now(), now()),
                (uuid_generate_v4(), 'Mud Motor section', 'ZMS', 'Mud Motor', 171.45, 76.2, 9.71, NULL, NULL, NULL, 103.53, 'CS_API 5D/7', '4145H MOD', NULL, 0.3, 110, now(), now()),
                (uuid_generate_v4(), 'Bit section', 'ZMS', 'Bit', 776, 215.9, NULL, 0.4, NULL, NULL, 100, NULL, NULL, NULL, NULL, NULL, now(), now());
        END;
    END;
END
$$;
