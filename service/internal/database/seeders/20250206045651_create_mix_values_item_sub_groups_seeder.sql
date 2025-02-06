BEGIN;

INSERT INTO
  mix_values (
    group_id,
    parent_id,
    name,
    description,
    status,
    created_at,
    updated_at
  )
VALUES
  (
    (
      SELECT
        id
      FROM
        groups
      WHERE
        name = 'item_sub_groups'
    ),
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Production'
    ),
    'Finished goods',
    'Finished goods',
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  ),
  (
    (
      SELECT
        id
      FROM
        groups
      WHERE
        name = 'item_sub_groups'
    ),
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Production'
    ),
    'Semi-finished goods',
    'Semi-finished goods',
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  ),
  (
    (
      SELECT
        id
      FROM
        groups
      WHERE
        name = 'item_sub_groups'
    ),
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Production'
    ),
    'Raw materials',
    'Raw materials',
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  ),
  (
    (
      SELECT
        id
      FROM
        groups
      WHERE
        name = 'item_sub_groups'
    ),
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Production'
    ),
    'Packaging materials',
    'Packaging materials',
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  ),
  (
    (
      SELECT
        id
      FROM
        groups
      WHERE
        name = 'item_sub_groups'
    ),
    (
      SELECT
        id
      FROM
        mix_values
      WHERE
        name = 'Office'
    ),
    'Office tools',
    'Office tools',
    1,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
  );

COMMIT;