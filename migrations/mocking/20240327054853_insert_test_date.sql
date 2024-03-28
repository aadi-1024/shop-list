-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO users (username, pass_hash) VALUES ('user', 'password'), ('user2', 'password');
INSERT INTO list_item (title, dsc, userid) VALUES ('abcd', 'item 1', 1), ('efgh', 'item 2', 1), ('ijkl', 'item 3', 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
