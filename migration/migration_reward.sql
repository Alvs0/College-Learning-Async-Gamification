CREATE TABLE IF NOT EXISTS "reward"
(
    id              TEXT    NOT NULL,
    college_id      TEXT    NOT NULL,
    name            TEXT    NOT NULL,
    description     TEXT    NOT NULL,
    quantity        INTEGER NOT NULL,
    required_points INTEGER NOT NULL,
    minimal_level   INTEGER NOT NULL,
    image_url       TEXT    NOT NULL,
    is_active       BOOLEAN NOT NULL,

    PRIMARY KEY (id)
);

COMMENT ON TABLE "reward" IS 'Holds reward information';