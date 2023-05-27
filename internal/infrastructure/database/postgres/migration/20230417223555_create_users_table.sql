-- +goose Up
CREATE TABLE users (
                      id bigserial NOT NULL PRIMARY KEY,
                      name text NOT NULL,
                      lastname text NOT NULL,
                      login text NOT NULL UNIQUE,
                      email text NOT NULL UNIQUE,
                      password text NOT NULL,
                      created_at text NOT NULL,
                      updated_at text,
                      verified_at text
);

-- +goose Down
DROP TABLE users;
