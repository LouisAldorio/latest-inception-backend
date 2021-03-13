BEGIN;

CREATE TABLE IF NOT EXISTS comodity (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    unit_price DECIMAL(15, 2) NOT NULL,
    unit_type VARCHAR(15) NOT NULL,
    min_purchase INTEGER NOT NULL,
    description TEXT,
    user_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES "user" (id)
    FOREIGN KEY (category_id) REFERENCES category (id)
);

CREATE TABLE IF NOT EXISTS comodity_image (
    comodity_id INTEGER NOT NULL,
    image_id INTEGER NOT NULL,

    FOREIGN KEY (comodity_id) REFERENCES comodity (id),
    FOREIGN KEY (image_id) REFERENCES image (id)
);

COMMIT;