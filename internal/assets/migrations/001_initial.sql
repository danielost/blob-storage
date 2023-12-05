-- +migrate Up

CREATE TABLE blobs (
    id character(52) NOT NULL,
    type integer NOT NULL,
    value text NOT NULL,
    created_at timestamp without time zone
);

-- +migrate Down

DROP TABLE blobs
