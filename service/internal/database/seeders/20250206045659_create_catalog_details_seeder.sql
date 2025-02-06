BEGIN;

INSERT INTO
  catalog_details (
    catalog_id,
    item_id,
    price,
    remark,
    created_at
  )
VALUES
  (
    1,
    1,
    1000,
    'Remark 1',
    CURRENT_TIMESTAMP
  ),
  (
    1,
    2,
    2000,
    'Remark 2',
    CURRENT_TIMESTAMP
  ),
  (
    1,
    3,
    3000,
    'Remark 3',
    CURRENT_TIMESTAMP
  ),
  (
    1,
    4,
    4000,
    'Remark 4',
    CURRENT_TIMESTAMP
  ),
  (
    2,
    1,
    5000,
    'Remark 5',
    CURRENT_TIMESTAMP
  ),
  (
    2,
    2,
    6000,
    'Remark 6',
    CURRENT_TIMESTAMP
  ),
  (
    3,
    3,
    7000,
    'Remark 7',
    CURRENT_TIMESTAMP
  ),
  (
    3,
    4,
    8000,
    'Remark 8',
    CURRENT_TIMESTAMP
  );

COMMIT;