insert into ebl_visualizations (
  uuid4,
  name,
  content,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
) VALUES (
    UNHEX(REPLACE('1270cfd4-4402-45c5-997b-63cc92ed0a9d','-','')),
    'Carrier rendered copy of the EBL.pdf',
    'string',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    '2023-03-07 12:12:12.000',
    '2023-04-07 12:12:12.000'
);
