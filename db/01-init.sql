CREATE TABLE people (
  id         SERIAL PRIMARY KEY,
  fullname   VARCHAR NOT NULL,
  age        SMALLINT NOT NULL,
  email      VARCHAR NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
