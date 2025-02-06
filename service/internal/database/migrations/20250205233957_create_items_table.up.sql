BEGIN;

CREATE TABLE IF NOT EXISTS items (
  id SERIAL PRIMARY KEY,
  customer_id INT NOT NULL,
  item_group_id INT NOT NULL,
  item_sub_group_id INT NOT NULL,
  item_unit_id INT NOT NULL,
  name TEXT NOT NULL,
  tpb_code TEXT NOT NULL,
  specification TEXT,
  description TEXT,
  remark TEXT,
  price_sell INT NOT NULL,
  price_buy INT NOT NULL,
  minimum_stock INT NOT NULL,
  is_active INT NOT NULL DEFAULT 1,
  created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  created_by_id INT,
  updated_by_id INT,
  deleted_by_id INT
);

COMMIT;