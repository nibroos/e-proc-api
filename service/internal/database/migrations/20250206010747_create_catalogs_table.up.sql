BEGIN;

CREATE TABLE IF NOT EXISTS catalogs (
  id SERIAL PRIMARY KEY,
  customer_id INT NOT NULL,
  catalog_no TEXT NOT NULL,
  description TEXT,
  remark TEXT,
  is_active INT NOT NULL DEFAULT 1,
  created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  created_by_id INT,
  updated_by_id INT,
  deleted_by_id INT
);

COMMIT;