INSERT INTO
  users (username, email, password, is_active, is_admin)
VALUES
  (
    'admin',
    'me@renodesper.com',
    uuid_generate_v4(),
    TRUE,
    TRUE
  )