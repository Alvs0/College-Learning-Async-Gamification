CREATE TABLE IF NOT EXISTS "college"
(
    id       TEXT NOT NULL,
    name     TEXT NOT NULL,
    address  TEXT NOT NULL,
    icon_url TEXT NOT NULL,

    PRIMARY KEY (id),
    UNIQUE (name)
);

COMMENT ON TABLE "college" IS 'Holds college information';