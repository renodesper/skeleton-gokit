CREATE TABLE IF NOT EXISTS users (
  id uuid DEFAULT uuid_generate_v4(),
  username VARCHAR NOT NULL,
  email VARCHAR NOT NULL,
  password VARCHAR NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT FALSE,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  is_admin BOOLEAN NOT NULL DEFAULT FALSE,
  access_token VARCHAR,
  refresh_token VARCHAR,
  expired_at TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
  updated_at TIMESTAMP DEFAULT NULL,
  PRIMARY KEY (id)
);

-- Used when we implement get by cursor
-- CREATE INDEX idx_user_pagination ON users (created_at, id);