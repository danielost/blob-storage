-- +migrate Up

CREATE TABLE users (
    id bigserial primary key,
    login varchar(56) NOT NULL,
    password text NOT NULL,
    created_at timestamp without time zone
);

ALTER TABLE blobs
ADD COLUMN owner_id bigint REFERENCES users(id);

ALTER TABLE blobs
  ADD CONSTRAINT ownerfk
  FOREIGN KEY (owner_id)
  REFERENCES users (id)
  ON UPDATE CASCADE ON DELETE CASCADE;

CREATE UNIQUE INDEX blobs_owner_constraint ON blobs USING btree (owner_id) WHERE owner_id IS NOT NULL;
CREATE UNIQUE INDEX user_login_constraint ON users USING btree (login) WHERE login IS NOT NULL;

-- +migrate Down

DROP TABLE IF EXISTS blobs;
DROP TABLE IF EXISTS users;
