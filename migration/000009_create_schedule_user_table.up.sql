CREATE TABLE IF NOT EXISTS schedule_user (
    user_id INTEGER NOT NULL,
    schedule_id INTEGER NOT NULL,

    FOREIGN KEY (user_id) REFERENCES user (id),
    FOREIGN KEY (schedule_id) REFERENCES schedule (id)
);