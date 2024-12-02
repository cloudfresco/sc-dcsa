insert into endorsement_chain_links (
  uuid4,
  entry_order,
  action_date_time,
  actor,
  recipient,
  surrender_request_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
) VALUES (
    UNHEX(REPLACE('9d16c5f9-dd0a-444d-a9e4-4b252852cbf3','-','')),
    1,
    '2024-04-29 01:46:59.226',
    1,
    1,
    1,
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    '2023-03-07 12:12:12.000',
    '2023-04-07 12:12:12.000'
);
