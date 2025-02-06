BEGIN;

INSERT INTO
  catalogs (
    customer_id,
    catalog_no,
    description,
    remark,
    is_active,
    created_at
  )
VALUES
  (
    1,
    'CAT-001',
    'Catalog 1',
    'Remark 1',
    1,
    CURRENT_TIMESTAMP
  ),
  (
    1,
    'CAT-002',
    'Catalog 2',
    'Remark 2',
    1,
    CURRENT_TIMESTAMP
  ),
  (
    2,
    'CAT-003',
    'Catalog 3',
    'Remark 3',
    1,
    CURRENT_TIMESTAMP
  ),
  (
    2,
    'CAT-004',
    'Catalog 4',
    'Remark 4',
    1,
    CURRENT_TIMESTAMP
  );

COMMIT;