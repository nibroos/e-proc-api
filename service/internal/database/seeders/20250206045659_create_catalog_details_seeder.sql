BEGIN;

INSERT INTO
  catalog_details (
    catalog_id,
    item_id,
    price_buy,
    price_sell,
    remark,
    created_at
  )
VALUES
  (
    1,
    1,
    1000,
    2000,
    'Remark 1',
    CURRENT_TIMESTAMP
  ),
  (
    1,
    2,
    2000,
    3000,
    'Remark 2',
    CURRENT_TIMESTAMP
  ),
  (
    1,
    3,
    3000,
    4000,
    'Remark 3',
    CURRENT_TIMESTAMP
  ),
  (
    1,
    4,
    4000,
    5000,
    'Remark 4',
    CURRENT_TIMESTAMP
  ),
  (
    2,
    2,
    5000,
    6000,
    'Remark 5',
    CURRENT_TIMESTAMP
  ),
  (
    2,
    4,
    6000,
    7000,
    'Remark 6',
    CURRENT_TIMESTAMP
  ),
  (
    2,
    1,
    7000,
    8000,
    'Remark 7',
    CURRENT_TIMESTAMP
  ),
  (
    3,
    4,
    8000,
    9000,
    'Remark 8',
    CURRENT_TIMESTAMP
  ),
  (
    3,
    1,
    9000,
    10000,
    'Remark 9',
    CURRENT_TIMESTAMP
  ),
  (
    4,
    2,
    10000,
    11000,
    'Remark 10',
    CURRENT_TIMESTAMP
  );

COMMIT;