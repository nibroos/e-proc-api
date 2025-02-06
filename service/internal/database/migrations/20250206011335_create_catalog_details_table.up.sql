BEGIN;

CREATE TABLE IF NOT EXISTS catalog_details (
  id SERIAL PRIMARY KEY,
  catalog_id INT NOT NULL REFERENCES catalogs(id),
  item_id INT NOT NULL REFERENCES items(id),
  price_buy INT,
  price_sell INT,
  remark TEXT,
  created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone
);

COMMIT;