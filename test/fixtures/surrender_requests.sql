insert into surrender_requests (
  uuid4,
  surrender_request_reference,
  transport_document_reference,
  surrender_request_code,
  comments,
  surrender_requested_by,
  created_date_time,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
) VALUES (
    UNHEX(REPLACE('e40fc6f6-c8ce-4ef4-97c0-bf112b70d3f2','-','')),
    'Z12345',
    'string',
    'SREQ',
    'string',
    1,
    '2023-03-07 12:12:12.000',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    '2023-03-07 12:12:12.000',
    '2023-04-07 12:12:12.000'
);
