BEGIN;

CREATE TABLE IF NOT EXISTS schedule (
    id INTEGER PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    dealed_unit INT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    comodity_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    FOREIGN KEY (comodity_id) REFERENCES comodity (id)
);

CREATE TABLE IF NOT EXISTS schedule_user (
    user_id INTEGER NOT NULL,
    schedule_id INTEGER NOT NULL,

    FOREIGN KEY (user_id) REFERENCES "user" (id),
    FOREIGN KEY (schedule_id) REFERENCES schedule (id)
);

COMMIT;