-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE user(
    id SERIAL PRIMARY KEY,
    username VARCHAR(16),
    pass_hash VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE user;
-- +goose StatementEnd
