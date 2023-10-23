-- +goose Up
-- +goose StatementBegin
CREATE INDEX CONCURRENTLY IF NOT EXISTS name_index ON users (lower((name)::text) varchar_pattern_ops);
CREATE INDEX CONCURRENTLY IF NOT EXISTS surname_index ON users (lower((surname)::text) varchar_pattern_ops);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
