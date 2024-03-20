-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE main_list (
    list_id INT,
    user_id INT,
    FOREIGN KEY (list_id) REFERENCES list_item(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE main_list;
-- +goose StatementEnd
