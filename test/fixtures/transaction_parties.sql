insert into transaction_parties (
  uuid4,
  ebl_platform_identifier,
  legal_name,
  registration_number,
  location_of_registration,
  tax_reference,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
) VALUES (
    UNHEX(REPLACE('c2683eda-75fd-48fd-86da-92aa44381b0e','-','')),
    'BOLE',
    'Digital Container Shipping Association',
    '74567837',
    'NL',
    'NL859951480B01',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    '2023-03-07 12:12:12.000',
    '2023-04-07 12:12:12.000'
);
