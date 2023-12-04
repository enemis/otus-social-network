-- +goose Up
-- +goose StatementBegin
CREATE TABLE friends (
    "user_id" uuid NOT NULL,
    "friend_id" uuid NOT NULL,
    "created_at" timestamp NOT NULL,
    PRIMARY KEY (user_id, friend_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS friends;
-- +goose StatementEnd