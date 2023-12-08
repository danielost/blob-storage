-- +migrate Up

ALTER TABLE blobs ALTER COLUMN value TYPE jsonb USING value::jsonb;

-- +migrate Down

