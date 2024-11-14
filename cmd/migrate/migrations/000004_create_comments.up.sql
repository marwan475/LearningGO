CREATE TABLE IF NOT EXISTS comments (
    id bigserial PRIMARY KEY,
    postid bigserial NOT NULL,
    userid bigserial NOT NULL,
    content TEXT NOT NULL,
    createtimestamp timestamp(0) with time zone NOT NULL DEFAULT NOW()
);