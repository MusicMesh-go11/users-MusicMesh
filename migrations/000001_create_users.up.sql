CREATE TABLE users (
        user_id uuid default gen_random_uuid() PRIMARY KEY,
        user_name VARCHAR(50) UNIQUE NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        deleted_at bigint default 0
);

CREATE TYPE user_role AS ENUM ('musician', 'listener', 'producer');

CREATE TABLE user_profiles (
        user_id uuid PRIMARY KEY REFERENCES users(user_id),
        full_name VARCHAR(100),
        bio TEXT,
        role user_role,
        location VARCHAR(100),
        avatar_url VARCHAR(255),
        website VARCHAR(255),
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        deleted_at bigint default 0
);