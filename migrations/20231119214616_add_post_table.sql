-- +goose Up
-- +goose StatementBegin
CREATE TYPE post_status AS ENUM ('draft', 'published');
CREATE TABLE "posts" (
  "id" uuid NOT NULL,
  "title" character(255) NOT NULL,
  "post" text NOT NULL,
  "created_at" timestamp NOT NULL,
  "update_at" timestamp NOT NULL,
  "deleted_at" timestamp,
  "status" post_status NOT NULL default 'draft'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TYPE post_status;
DROP TABLE posts;
-- +goose StatementEnd
