-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id    bigint NOT NULL,
    state varchar(30)
);

create unique index users_id_uindex
    on users (id);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
