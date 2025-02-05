BEGIN;

INSERT INTO
  groups (
    id,
    name,
    description,
    status,
    options_json,
    created_by_id,
    updated_by_id
  )
VALUES
  (
    1,
    'roles',
    'Roles table for storing user data',
    1,
    '{}',
    1,
    1
  ),
  (
    2,
    'permissions',
    'Permissions table for storing user data',
    1,
    '{}',
    1,
    1
  ),
  (
    3,
    'users',
    'Users table for storing user data',
    1,
    '{}',
    1,
    1
  ),
  (
    4,
    'identifiers',
    'Identifier (ID)',
    1,
    '{}',
    1,
    1
  ),
  (
    5,
    'addresses',
    'Alamat',
    1,
    '{}',
    1,
    1
  ),
  (
    6,
    'contacts',
    'Contacts table for storing user data',
    1,
    '{}',
    1,
    1
  ),
  (
    7,
    'customer_types',
    'Customer types table for storing user data. Ex: Buyer, Agent, Vendor, Supplier, etc',
    1,
    '{}',
    1,
    1
  ),
  (
    8,
    'currencies',
    'Currencies table for storing user data. Ex: USD, EUR, etc',
    1,
    '{}',
    1,
    1
  ),
  (
    9,
    'item_groups',
    'Item groups table for storing user data. Ex: Office, Production, etc',
    1,
    '{}',
    1,
    1
  ),
  (
    10,
    'item_units',
    'Item units table for storing user data. Ex: Kg, Ltr, Pcs etc',
    1,
    '{}',
    1,
    1
  ),
  (
    11,
    'item_sub_groups',
    'Item sub groups table for storing user data. Ex: Finished goods, Semi-finished goods, Raw materials etc',
    1,
    '{}',
    1,
    1
  );

COMMIT;