INSERT INTO
  "user" (username, email, password, is_active, is_admin, created_from)
VALUES
  (
    'admin',
    'boy.arriezona@gmail.com',
    uuid_generate_v4(),
    TRUE,
    TRUE,
    'DBMigration'
  )