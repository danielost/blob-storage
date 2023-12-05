-- +migrate Up

CREATE TABLE blobs (
    id bigserial primary key,
    value text NOT NULL,
    created_at timestamp without time zone
    deleted_at timestamp without time zone
);

-- +migrate Down

DROP TABLE blobs;

