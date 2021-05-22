CREATE TABLE IF NOT EXISTS comodity_image (
    comodity_id INTEGER NOT NULL,
    image_id INTEGER NOT NULL,

    FOREIGN KEY (comodity_id) REFERENCES comodity (id),
    FOREIGN KEY (image_id) REFERENCES image (id)
);