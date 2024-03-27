-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE list_item (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    dsc VARCHAR(1024),
    userid INT,
    FOREIGN KEY (userid) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE list_item;
-- +goose StatementEnd
