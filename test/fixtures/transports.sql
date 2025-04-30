insert into transports (
    uuid4,
    transport_reference,
    transport_name,
    load_transport_call_id,
    discharge_transport_call_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
) VALUES (
    UNHEX(REPLACE('561a5606-402e-11eb-b19a-0f3aa4962e1f','-','')), 
    'transport reference',
    'Transport name',
    1,
    1,
      'active',
      'auth0|66fd06d0bfea78a82bb42459',
      'auth0|66fd06d0bfea78a82bb42459',
      '2020-03-07 12:12:12.000',
      '2020-04-07 12:12:12.000'
), (
    UNHEX(REPLACE('561a5606-402e-11eb-b19a-0f3aa4962e2f','-','')), 
    'transport reference xx',
    'Transport name xx',
    1,
    1,
      'active',
      'auth0|66fd06d0bfea78a82bb42459',
      'auth0|66fd06d0bfea78a82bb42459',
      '2020-03-07 12:12:12.000',
      '2020-04-07 12:12:12.000'
), (
    UNHEX(REPLACE('561a5606-402e-11eb-b19a-0f3aa4962e3f','-','')), 
    'transport reference yy',
    'Transport name yy',
    1,
    1,
      'active',
      'auth0|66fd06d0bfea78a82bb42459',
      'auth0|66fd06d0bfea78a82bb42459',
      '2020-03-07 12:12:12.000',
      '2020-04-07 12:12:12.000'
);

insert into transports (
    uuid4,
    transport_reference,
    transport_name,
    load_transport_call_id,
    discharge_transport_call_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
) VALUES (
    UNHEX(REPLACE('6b14b74d-401a-4e66-a5ad-d3cd42953441','-','')), 
    'not another transport reference',
    'Transport Name in action',
    2,
    2,
      'active',
      'auth0|66fd06d0bfea78a82bb42459',
      'auth0|66fd06d0bfea78a82bb42459',
      '2020-03-07 12:12:12.000',
      '2020-04-07 12:12:12.000'
);

insert into transports (
    uuid4,
    transport_reference,
    transport_name,
    load_transport_call_id,
    discharge_transport_call_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
) VALUES (
    UNHEX(REPLACE('bda2e4ee-7cfa-4a07-91fd-e07b4ea9c0ad','-','')), 
    'transport reference',
    'Transport name (Singapore -> NYC)',
    3,
    3,
      'active',
      'auth0|66fd06d0bfea78a82bb42459',
      'auth0|66fd06d0bfea78a82bb42459',
      '2020-03-07 12:12:12.000',
      '2020-04-07 12:12:12.000'
);

insert into transports (
    uuid4,
    transport_reference,
    transport_name,
    load_transport_call_id,
    discharge_transport_call_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at
) VALUES (
    UNHEX(REPLACE('986cabf1-5bb8-4f74-94c6-57286fe769b9','-','')), 
    'transport reference',
    'Transport name (Singapore -> NYC)',
    4,
    4,
      'active',
      'auth0|66fd06d0bfea78a82bb42459',
      'auth0|66fd06d0bfea78a82bb42459',
      '2020-03-07 12:12:12.000',
      '2020-04-07 12:12:12.000'
);
