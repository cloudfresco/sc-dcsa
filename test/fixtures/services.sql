INSERT INTO services (
    uuid4,
    carrier_id,
    carrier_service_code,
    carrier_service_name,
    universal_service_reference,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at
) VALUES (
     UNHEX(REPLACE('03482296-ef9c-11eb-9a03-0242ac131999','-','')),
     (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
      'A_CSC',
     'A_carrier_service_name',
     'SR00001D',
      'active',
     'auth0|66fd06d0bfea78a82bb42459',
      'auth0|66fd06d0bfea78a82bb42459',
      '2020-03-07 12:12:12.000',
      '2020-04-07 12:12:12.000'
);

INSERT INTO services (
    uuid4,
    carrier_id,
    carrier_service_code,
    carrier_service_name,
    universal_service_reference,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at
) VALUES (
     UNHEX(REPLACE('f65022f1-76e7-4cf2-8287-241cd7aed4de','-','')),
     (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
     'B_HLC',
     'B_carrier_service_name',
     'SR00002B',
      'active',
      'auth0|66fd06d0bfea78a82bb42459',
      'auth0|66fd06d0bfea78a82bb42459',
      '2020-03-07 12:12:12.000',
      '2020-04-07 12:12:12.000'
);

INSERT INTO services (
    uuid4,
    carrier_id,
    carrier_service_code,
    carrier_service_name,
    universal_service_reference,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at
) VALUES (
     UNHEX(REPLACE('f26ac90d-c89a-4bff-9fd3-35c134a3ec31','-','')),
     (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
     'B_HLC',
     'B_carrier_service_name_1',
     'SR00003H',
      'active',
     'auth0|66fd06d0bfea78a82bb42459',
      'auth0|66fd06d0bfea78a82bb42459',
      '2020-03-07 12:12:12.000',
      '2020-04-07 12:12:12.000'
);
