insert into issuance_requests (
  uuid4,
  transport_document_reference,
  issuance_request_state,
  issue_to,
  ebl_visualization_id,
  transport_document_json,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
) VALUES (
    UNHEX(REPLACE('f40939a3-c05b-4a34-bc67-a8e3676cdc80','-','')),
    'HHL71800000',
    'DR',
    1,
    1,
    'SWB',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    '2023-03-07 12:12:12.000',
    '2023-04-07 12:12:12.000'
);
