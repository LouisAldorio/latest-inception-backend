BEGIN;

CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    role VARCHAR(15) NOT NULL,
    whatsapp VARCHAR(15) NOT NULL,
    password VARCHAR(255) NOT NULL,
    avatar INTEGER,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,

    FOREIGN KEY (avatar) REFERENCES image (id)
); 

COMMIT;