BEGIN;

INSERT INTO
  customers (
    customer_type_id,
    name,
    email,
    phone,
    address,
    pic,
    is_active,
    user_id,
    created_at,
    updated_at
  )
VALUES
  (
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Vendor'
    ),
    'PT. Vendor A',
    'email@gmail.com',
    '08123456789',
    'Jl. Vendor A No. 1',
    'Vendor A PIC',
    1,
    2,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  ),
  (
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Buyer'
    ),
    'PT. Buyer B',
    'buyer@gmail.com',
    '08123456789',
    'Jl. Buyer B No. 1',
    'Buyer B PIC',
    1,
    3,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  ),
  (
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Agent'
    ),
    'PT. Agent C',
    'agent@gmail.com',
    '08123456789',
    'Jl. Agent C No. 1',
    'Agent C PIC',
    1,
    4,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  ),
  (
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Supplier'
    ),
    'PT. Supplier D',
    'sup@gmail.com',
    '08123456789',
    'Jl. Supplier D No. 1',
    'Supplier D PIC',
    1,
    5,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  );

COMMIT;