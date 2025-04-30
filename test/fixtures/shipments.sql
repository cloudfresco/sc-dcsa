insert into shipments (uuid4,
    carrier_id,
    booking_id,
    carrier_booking_reference,
    terms_and_conditions,
    confirmation_datetime,
    updated_date_time,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('372e1778-e853-4745-9470-169d00935315','-','')),
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'CARRIER_BOOKING_REQUEST_REFERENCE_01'),
    'BR1239719871',
    'TERMS AND CONDITIONS!',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
);

insert into shipments (uuid4,
    carrier_id,
    booking_id,
    carrier_booking_reference,
    terms_and_conditions,
    confirmation_datetime,
    updated_date_time,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('34390ee7-ca05-47dc-9740-3929511dbbef','-','')),
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'CARRIER_BOOKING_REQUEST_REFERENCE_02'),
    'CR1239719872',
    'TERMS AND CONDITIONS!',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
);

/**
 * Data used in integration tests - Do not modify - make your own data
 */
insert into shipments (uuid4,
    carrier_id,
    booking_id,
    carrier_booking_reference,
    terms_and_conditions,
    confirmation_datetime,
    updated_date_time,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('00e87010-9b54-4fb3-b444-00280da44d89','-','')),
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'CARRIER_BOOKING_REQUEST_REFERENCE_01'),
    'bca68f1d3b804ff88aaa1e43055432f7',
    'TERMS AND CONDITIONS!',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
),(
    UNHEX(REPLACE('23e09959-23b2-470a-81ba-0fb8bb22a562','-','')),
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'CARRIER_BOOKING_REQUEST_REFERENCE_01'),
    '832deb4bd4ea4b728430b857c59bd057',
    'TERMS AND CONDITIONS!',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
),(
    UNHEX(REPLACE('d8078144-80b5-4d8f-9cd3-148cfe9da1c7','-','')),
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'CARRIER_BOOKING_REQUEST_REFERENCE_01'),
    '994f0c2b590347ab86ad34cd1ffba505',
    'TERMS AND CONDITIONS!',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
),(
    UNHEX(REPLACE('58089407-445c-4a5d-983b-9ba1d2fccb07','-','')),
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'CARRIER_BOOKING_REQUEST_REFERENCE_01'),
    '02c965382f5a41feb9f19b24b5fe2906',
    'TERMS AND CONDITIONS!',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
),( 
    UNHEX(REPLACE('6a2bcacf-4ee6-4723-bd9b-b45765670cd7','-','')),
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'CARRIER_BOOKING_REQUEST_REFERENCE_01'),
    'AR1239719871',
    'TERMS AND CONDITIONS!',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
);

insert into shipments (uuid4,
    carrier_id,
    booking_id,
    carrier_booking_reference,
    terms_and_conditions,
    confirmation_datetime,
    updated_date_time,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('a1dbbafc-904a-44c0-9ffb-fdc550a3d768','-','')),
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'CARRIER_BOOKING_REQUEST_REFERENCE_03'),
    '43f615138efc4d3286b36402405f851b',
    'TERMS AND CONDITIONS!',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
),(
    UNHEX(REPLACE('ca60809a-f214-4e12-9966-6089cf8c863d','-','')),
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'CARRIER_BOOKING_REQUEST_REFERENCE_04'),
    'e8e9d64172934a40aec82e4308cdf97a',
    'TERMS AND CONDITIONS!',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
),(
    UNHEX(REPLACE('6227e254-7ed1-4b33-bdfe-1b22c8e10629','-','')),
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'CARRIER_BOOKING_REQUEST_REFERENCE_05'),
    '6fe84758a4cc471fb5eb-4de63ddadc41',
    'TERMS AND CONDITIONS!',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
),
(
    UNHEX(REPLACE('b69ccaa0-e942-4854-a92f-47fbe4e558e0','-','')),
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'CARRIER_BOOKING_REQUEST_REFERENCE_06'),
    '5dc92988f48a420495b786c224efce7d',
    'TERMS AND CONDITIONS!',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
);

insert into shipments (uuid4,
    carrier_id,
    booking_id,
    carrier_booking_reference,
    terms_and_conditions,
    confirmation_datetime,
    updated_date_time,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('b3f0a6f3-d5bd-4ed0-81ab-0177672052e2','-','')),
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'KUBERNETES_IN_ACTION_01'),
    'E379021B7782',
    'TERMS AND CONDITIONS!',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
), (
    UNHEX(REPLACE('f09ae946-c25f-4447-9290-d5002ced691c','-','')),
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'KUBERNETES_IN_ACTION_02'),
    'A379021B7782',
    'TERMS AND CONDITIONS!',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
);

insert into shipments (uuid4,
    carrier_id,
    booking_id,
    carrier_booking_reference,
    terms_and_conditions,
    confirmation_datetime,
    updated_date_time,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('4c082102-c6d2-4f5f-b6c3-b4d24a5432b4','-','')),
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'KUBERNETES_IN_ACTION_03'),
    'D659FDB7E33C',
    'TERMS AND CONDITIONS!',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
);

insert into shipments (uuid4,
    booking_id,
    carrier_id,
    carrier_booking_reference,
    terms_and_conditions,
    confirmation_datetime,
    updated_date_time,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
    ) VALUES (
    UNHEX(REPLACE('2037d2fa-3d30-43f5-b399-b85b24e8df29','-','')),
    10,
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'HLC'),
    'DCR987876762',
    'TERMS AND CONDITIONS!',
    '2021-12-12 12:12:12.000',
    '2021-12-12 12:12:12.000','active','2021-12-12 12:12:12.000','2021-12-12 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459');

insert into shipments (
    uuid4,
    booking_id,
    carrier_id,
    carrier_booking_reference,
    terms_and_conditions,
    confirmation_datetime,
    updated_date_time,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
    ) VALUES (
    UNHEX(REPLACE('dd4c2e5e-4a0e-43fa-b4ea-c41fb71b4fd1','-','')),
    12,
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'HLC'),
    'C501576CD94F',
    'TERMS AND CONDITIONS!',
    '2022-02-02 02:22:22.000',
    '2022-03-03 12:12:12.000','active','2022-02-02 12:12:12.000','2021-12-12 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459');
    
 insert into shipments (uuid4,
    carrier_id,
    booking_id,
    carrier_booking_reference,
    terms_and_conditions,
    confirmation_datetime,
    updated_date_time,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('20fa6b72-ea9a-424b-bc6b-ef1c70f1d02c','-','')),
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
    13,
    'cbr-b83765166707812c8ff4',
    'TERMS AND CONDITIONS!',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
);

insert into shipments (uuid4,
    carrier_id,
    booking_id,
    carrier_booking_reference,
    terms_and_conditions,
    confirmation_datetime,
    updated_date_time,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('1316ac3b-d3d3-4622-9a36-a862227bb385','-','')),
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'CARRIER_BOOKING_REQUEST_REFERENCE_01'),
    'BR1239719971',
    'TERMS AND CONDITIONS!',
    '2021-12-12 12:12:12.000',
    '2021-12-12 12:12:12.000','active','2021-12-12 12:12:12.000','2021-12-12 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
);

insert into shipments (uuid4,
    carrier_id,
    booking_id,
    carrier_booking_reference,
    terms_and_conditions,
    confirmation_datetime,
    updated_date_time,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('70ddcc63-58c2-4c20-ae6b-62508f071d89','-','')),
    (SELECT carriers.id FROM carriers WHERE smdg_code = 'MSK'),
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'CARRIER_BOOKING_REQUEST_REFERENCE_02'),
    'ABC123123123',
    'TERMS AND CONDITIONS!',
    '2021-12-12 12:12:12.000',
    '2021-12-12 12:12:12.000','active','2021-12-12 12:12:12.000','2021-12-12 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
);
