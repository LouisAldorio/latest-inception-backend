CREATE TABLE IF NOT EXISTS comodity (
    id INTEGER AUTO_INCREMENT PRIMARY KEY,
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

    FOREIGN KEY (user_id) REFERENCES user (id),
    FOREIGN KEY (category_id) REFERENCES category (id)
);