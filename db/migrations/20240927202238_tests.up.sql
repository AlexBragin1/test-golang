CREATE TABLE IF NOT EXISTS tests
(
    id         uuid   NOT NULL PRIMARY KEY,
    user_id    uuid   NOT NULL,     
    variant_id uuid NOT NULL,
    start_at timestamptz      DEFAULT NULL,
    end_at   timestamptz      DEFAULT NULLS,
    percent VARCHAR(15),
);