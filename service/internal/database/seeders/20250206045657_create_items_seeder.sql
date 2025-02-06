BEGIN;

INSERT INTO
  items (
    customer_id,
    item_group_id,
    item_sub_group_id,
    item_unit_id,
    name,
    tpb_code,
    specification,
    description,
    remark,
    price_sell,
    price_buy,
    minimum_stock,
    is_active,
    created_at
  )
VALUES
  (
    1,
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Office'
    ),
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Office tools'
    ),
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Pack'
    ),
    'Kertas HVS A4',
    'KRTS001',
    'Kertas HVS A4 80gsm',
    'Kertas HVS A4 80gsm',
    'Kertas HVS A4 80gsm',
    10000,
    8000,
    10,
    1,
    '2025-02-06 04:56:57'
  ),
  (
    1,
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Office'
    ),
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Office tools'
    ),
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Pack'
    ),
    'Kertas HVS A3',
    'KRTS002',
    'Kertas HVS A3 80gsm',
    'Kertas HVS A3 80gsm',
    'Kertas HVS A3 80gsm',
    20000,
    16000,
    10,
    1,
    '2025-02-06 04:56:57'
  ),
  (
    1,
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Office'
    ),
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Office tools'
    ),
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Pack'
    ),
    'Kertas HVS A2',
    'KRTS003',
    'Kertas HVS A2 80gsm',
    'Kertas HVS A2 80gsm',
    'Kertas HVS A2 80gsm',
    30000,
    24000,
    10,
    1,
    '2025-02-06 04:56:57'
  ),
  (
    1,
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Office'
    ),
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Office tools'
    ),
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Pack'
    ),
    'Kertas HVS A1',
    'KRTS004',
    'Kertas HVS A1 80gsm',
    'Kertas HVS A1 80gsm',
    'Kertas HVS A1 80gsm',
    40000,
    32000,
    10,
    1,
    '2025-02-06 04:56:57'
  ),
  (
    1,
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Office'
    ),
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Office tools'
    ),
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Pack'
    ),
    'Kertas HVS A0',
    'KRTS005',
    'Kertas HVS A0 80gsm',
    'Kertas HVS A0 80gsm',
    'Kertas HVS A0 80gsm',
    50000,
    40000,
    10,
    1,
    '2025-02-06 04:56:57'
  );

COMMIT;