CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY,
    username varchar(255) NOT NULL,
    email citext UNIQUE NOT NULL,
    password bytea NOT NULL,
    createtimestamp timestamp(0) with time zone NOT NULL DEFAULT NOW()
);