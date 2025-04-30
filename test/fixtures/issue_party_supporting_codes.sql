insert into issue_party_supporting_codes (
  uuid4,
  issue_party_id,
  party_code,
  party_code_list_provider,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
) VALUES (
    UNHEX(REPLACE('1be42473-a5a8-4dcf-abfe-53fc0ed01ba8','-','')),
    1,
    '529900T8BM49AURSDO55',
    'EPIU',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    '2023-03-07 12:12:12.000',
    '2023-04-07 12:12:12.000'
);
