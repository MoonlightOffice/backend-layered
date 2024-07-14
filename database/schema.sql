CREATE DATABASE giants;
\c giants;

CREATE TABLE users (
  user_id TEXT NOT NULL,
  email TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,

  PRIMARY KEY (user_id)
);
