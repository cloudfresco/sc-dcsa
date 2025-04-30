insert into surrender_request_answers (
  uuid4,
  surrender_request_reference,
  action,
  comments,
  created_date_time,
  surrender_request_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
) VALUES (
    UNHEX(REPLACE('f69275f8-2d6e-49ed-a6d0-100c95cc0958','-','')),
    'Z12345',
    'SURR',
    'comments',
    '2023-03-07 12:12:12.000',
    1,
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    '2023-03-07 12:12:12.000',
    '2023-04-07 12:12:12.000'
);
