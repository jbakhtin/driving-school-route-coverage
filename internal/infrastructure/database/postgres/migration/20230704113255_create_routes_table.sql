-- +goose Up
CREATE TABLE routes (
                       id bigserial NOT NULL PRIMARY KEY,
                       line GEOMETRY(LineString, 4326),
                       created_at text NOT NULL,
                       updated_at text
);

-- +goose Down
DROP TABLE routes;
