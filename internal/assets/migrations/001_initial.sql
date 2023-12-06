-- +migrate Up

CREATE TABLE users (
    id bigserial primary key,
    login varchar(56) NOT NULL,
    password text NOT NULL,
    created_at timestamp without time zone
);

CREATE TABLE blobs (
    id bigserial primary key,
    value text NOT NULL,
    created_at timestamp without time zone,
    owner_id bigint not null references users (id) on delete cascade
);

CREATE INDEX blobs_owner_constraint ON blobs USING btree (owner_id) WHERE owner_id IS NOT NULL;
CREATE INDEX user_login_constraint ON users USING btree (login) WHERE login IS NOT NULL;

-- +migrate Down

DROP TABLE IF EXISTS blobs;
DROP TABLE IF EXISTS users;
