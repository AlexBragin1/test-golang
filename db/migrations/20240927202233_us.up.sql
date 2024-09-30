CREATE TABLE IF NOT EXISTS users
(
    id         uuid             NOT NULL PRIMARY KEY,
    groups     bigint           NOT NULL,
    login      varchar(32)      NOT NULL,
	password   varchar(128)     NOT NULL,
	group      group NOT NULL,
    sessipon_start_at timestamptz      DEFAULT NULL,
	session_end_at timestamptz      DEFAULT NULL,
);