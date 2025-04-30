insert into issuance_request_responses (
  uuid4,
  transport_document_reference,
  issuance_response_code,
  reason,
  created_date_time,
  issuance_request_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
) VALUES (
    UNHEX(REPLACE('311ca76d-56a9-42af-94a9-6602af9ed683','-','')),
    'HHL71800000',
    'ISSU',
    'null',
    '2023-03-07 12:12:12.000',
    1,
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    '2023-03-07 12:12:12.000',
    '2023-04-07 12:12:12.000'
);
