CREATE TABLE IF NOT EXISTS cameras (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    year integer NOT NULL,
    manufacturer text NOT NULL,
    model text NOT NULL,
    details text NOT NULL,
    version integer NOT NULL DEFAULT 1
);
