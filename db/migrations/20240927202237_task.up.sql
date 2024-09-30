CREATE TABLE IF NOT EXISTS task
(
    id  uuid   NOT NULL PRIMARY KEY,
    variant_id uuid NOT NULL,    
    description VARCHAR(70) NOT NULL,
    correct_answer VARCHAR(10),
    options VARCHAR(80) NOT NULL
);