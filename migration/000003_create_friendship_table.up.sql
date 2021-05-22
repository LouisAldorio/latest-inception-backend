CREATE TABLE IF NOT EXISTS friendship (
    first_user_id INTEGER NOT NULL,
    second_user_id INTEGER NOT NULL,

    FOREIGN KEY (first_user_id) REFERENCES user (id),
    FOREIGN KEY (second_user_id) REFERENCES user (id)
);