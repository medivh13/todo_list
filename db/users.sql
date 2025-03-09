CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,  -- Untuk menyimpan password hash (bcrypt)
    refresh_token TEXT,  -- Menyimpan refresh token terakhir
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
