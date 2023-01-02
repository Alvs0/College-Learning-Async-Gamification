CREATE TABLE IF NOT EXISTS "user"
(
    id                TEXT    NOT NULL,
    college_id        TEXT    NOT NULL,
    name              TEXT    NOT NULL,
    email             TEXT    NOT NULL,
    phone_number      TEXT    NOT NULL,
    birth_date        TEXT    NOT NULL,
    profile_image_url TEXT    NOT NULL,
    password          TEXT    NOT NULL,
    is_admin          BOOLEAN NOT NULL DEFAULT FALSE,

    PRIMARY KEY (id),
    UNIQUE (email)
);

COMMENT ON TABLE "user" IS 'Holds account user information';