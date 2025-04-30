INSERT INTO vessel_schedules (
    uuid4,
    vessel_id,
    service_id,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at
) VALUES (
     UNHEX(REPLACE('43fb0eb7-b320-4932-8940-d727639983ac','-','')),
     (SELECT vessels.id FROM vessels WHERE vessel_imo_number = '9811000'),
     (SELECT services.id FROM services WHERE universal_service_reference = 'SR00002B'),
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000'
);

INSERT INTO vessel_schedules (
    uuid4,
    vessel_id,
    service_id,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at
) VALUES (
     UNHEX(REPLACE('74a87360-d02c-4d11-8a1e-0c68e8fb1cda','-','')),
     (SELECT vessels.id FROM vessels WHERE vessel_imo_number = '9136307'),
     (SELECT services.id FROM services WHERE universal_service_reference = 'SR00003H'),
      'active',
      'auth0|66fd06d0bfea78a82bb42459',
      'auth0|66fd06d0bfea78a82bb42459',
      '2020-03-07 12:12:12.000',
      '2020-04-07 12:12:12.000'
);
