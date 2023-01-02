CREATE TABLE IF NOT EXISTS "user_reward"
(
    user_id    TEXT    NOT NULL,
    college_id TEXT    NOT NULL,
    reward_id  TEXT    NOT NULL,
    quantity   INTEGER NOT NULL,

    FOREIGN KEY (user_id) REFERENCES "user" (id) ON DELETE CASCADE,
    FOREIGN KEY (reward_id) REFERENCES "reward" (id) ON DELETE CASCADE
);

COMMENT ON TABLE "user_reward" IS 'Holds account users reward information';