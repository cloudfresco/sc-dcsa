insert into transaction_party_supporting_codes (
  uuid4,
  transaction_party_id,
  party_code,
  party_code_list_provider,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
) VALUES (
    UNHEX(REPLACE('8839213f-8092-429e-8b7b-ca1c718cd140','-','')),
    1,
    '990052T8BM49ARSUDO55',
    'SPIU',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    '2023-03-07 12:12:12.000',
    '2023-04-07 12:12:12.000'
);
