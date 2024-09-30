CREATE TABLE IF NOT EXISTS variants
(
    id         uuid             NOT NULL PRIMARY KEY,
    test_user_id uuid NOT NULL,
    percent VARCHAR(15),
);