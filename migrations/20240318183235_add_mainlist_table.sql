-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE main_list (
    list_id INT FOREIGN KEY REFERENCES list_item.id,
    user_id INT FOREIGN KEY REFERENCES user.id
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE main_list;
-- +goose StatementEnd
