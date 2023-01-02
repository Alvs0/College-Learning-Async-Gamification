CREATE TABLE IF NOT EXISTS "user_point"
(
    user_id    TEXT    NOT NULL,
    college_id TEXT    NOT NULL,
    point      INTEGER NOT NULL,

    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE
);

COMMENT ON TABLE "user_point" IS 'Holds account users point information';