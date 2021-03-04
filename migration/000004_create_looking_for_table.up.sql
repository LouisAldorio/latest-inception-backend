BEGIN;

CREATE TABLE IF NOT EXISTS looking_for (
    user_id INTEGER NOT NULL,
    item VARCHAR(50) NOT NULL,

    FOREIGN KEY (user_id) REFERENCES "user" (id)
);

COMMIT;