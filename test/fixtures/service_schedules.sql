INSERT INTO service_schedules (
    uuid4,
    carrier_service_code,
    carrier_service_name,
    universal_service_reference,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at
) VALUES (
     UNHEX(REPLACE('d763f550-1f77-48e1-af53-4e038d604ad3','-','')),
     'B_HLC',
     'B_carrier_service_name',
     'SR00002B',
      'active',
      'auth0|66fd06d0bfea78a82bb42459',
      'auth0|66fd06d0bfea78a82bb42459',
      '2020-03-07 12:12:12.000',
      '2020-04-07 12:12:12.000'
);

INSERT INTO service_schedules (
    uuid4,
    carrier_service_code,
    carrier_service_name,
    universal_service_reference,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at
) VALUES (
     UNHEX(REPLACE('80138b28-b41d-42b9-b25f-7471f686c91b','-','')),
     'B_HLC',
     'B_carrier_service_name_1',
     'SR00003H',
      'active',
      'auth0|66fd06d0bfea78a82bb42459',
      'auth0|66fd06d0bfea78a82bb42459',
      '2020-03-07 12:12:12.000',
      '2020-04-07 12:12:12.000'
);
