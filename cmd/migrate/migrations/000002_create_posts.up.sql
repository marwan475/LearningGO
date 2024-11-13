CREATE TABLE IF NOT EXISTS posts (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    userid bigint NOT NULL,
    content text NOT NULL,
    tags VARCHAR(100) [],
    createtimestamp timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updatetimestamp timestamp(0) with time zone NOT NULL DEFAULT NOW()
); 