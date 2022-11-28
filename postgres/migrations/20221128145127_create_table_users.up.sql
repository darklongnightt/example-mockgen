BEGIN;
CREATE TABLE IF NOT EXISTS users(
  id uuid NOT NULL PRIMARY KEY,
  name varchar(255) NOT NULL,
  age integer NOT NULL,
  created_at timestamptz NOT NULL
);
COMMIT;
