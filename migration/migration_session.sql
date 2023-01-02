CREATE TABLE IF NOT EXISTS "session"
(
    id         TEXT NOT NULL,
    college_id TEXT NOT NULL,
    name       TEXT NOT NULL,
    link       TEXT NOT NULL,
    image_url  TEXT NOT NULL,

    PRIMARY KEY (id),
    UNIQUE (name)
);

COMMENT ON TABLE "session" IS 'Holds session information';