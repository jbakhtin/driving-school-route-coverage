-- +goose Up
CREATE TABLE routes (
                       id bigserial NOT NULL PRIMARY KEY,
                       user_id bigint NOT NULL,
                       name text,
                       description text,
                       linestring GEOMETRY(LineString, 4326),
                       created_at text NOT NULL,
                       updated_at text
);

ALTER TABLE routes
    ADD CONSTRAINT user_id
        FOREIGN KEY (user_id) REFERENCES users (id)


-- +goose Down
DROP TABLE routes;
