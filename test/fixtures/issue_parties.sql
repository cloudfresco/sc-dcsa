insert into issue_parties (
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
    UNHEX(REPLACE('7e33799e-9230-4c5e-b209-5c15efce1303','-','')),
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
