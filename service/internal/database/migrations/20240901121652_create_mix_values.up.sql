BEGIN;

CREATE TABLE IF NOT EXISTS mix_values (
  id SERIAL PRIMARY KEY,
  group_id INT NOT NULL,
  parent_id INT,
  name TEXT NOT NULL,
  num INT,
  description TEXT,
  status INT,
  options_json JSONB,
  created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  created_by_id INT,
  updated_by_id INT,
  deleted_by_id INT
);

COMMIT;