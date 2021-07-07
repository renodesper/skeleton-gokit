CREATE TABLE IF NOT EXISTS "verification" (
  id uuid DEFAULT uuid_generate_v4(),
  user_id uuid NOT NULL,
  type VARCHAR NOT NULL,
  token VARCHAR NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT FALSE,
  expired_at TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
  updated_at TIMESTAMP DEFAULT NULL,
  PRIMARY KEY (id)
);