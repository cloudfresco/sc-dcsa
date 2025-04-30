insert into consignment_items (
    uuid4,
    shipping_instruction_id,
    shipment_id,
    description_of_goods,
    hs_code,
    weight,
    weight_unit,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('1316ac3b-d3d3-4622-9a36-a862227bb385','-','')),
    2,
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'BR1239719871'),
    'Expensive Shoes',
    '411510',
    4000,
    'KGM','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
), (
    UNHEX(REPLACE('1316ac3b-d3d3-4622-9a36-a862227bb385','-','')),
     2,
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'BR1239719871'),
    'Massive Yacht',
    '720711',
    4000,
    'KGM','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
), (
    UNHEX(REPLACE('1316ac3b-d3d3-4622-9a36-a862227bb385','-','')),
    3,
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'bca68f1d3b804ff88aaa1e43055432f7'),
    'Leather Jackets',
    '411510',
    4000,
    'KGM','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
), (
    UNHEX(REPLACE('1316ac3b-d3d3-4622-9a36-a862227bb385','-','')),
    4,
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = '832deb4bd4ea4b728430b857c59bd057'),
    'Air ballons',
    '411510',
    4000,
    'KGM','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
),(
    UNHEX(REPLACE('1316ac3b-d3d3-4622-9a36-a862227bb385','-','')),
    5,
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = '994f0c2b590347ab86ad34cd1ffba505'),
    'Leather Jackets',
    '411510',
    4000,
    'KGM','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
),(
    UNHEX(REPLACE('1316ac3b-d3d3-4622-9a36-a862227bb385','-','')),
    6,
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = '02c965382f5a41feb9f19b24b5fe2906'),
    'Leather Jackets',
    '411510',
    4000,
    'KGM','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
),(
    UNHEX(REPLACE('1316ac3b-d3d3-4622-9a36-a862227bb385','-','')),
    7,
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'AR1239719871'),
    'Leather Jackets',
    '411510',
    4000,
    'KGM','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
);

insert into consignment_items (
    uuid4,
    shipping_instruction_id,
    shipment_id,
    description_of_goods,
    hs_code,
    weight,
    weight_unit,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('0e98eef4-6ebd-47eb-bd6e-d3878b341b7f','-','')),
    8,
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'E379021B7782'),
    'Expensive shoes',
    '411510',
    4000,
    'KGM','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
), (
    UNHEX(REPLACE('06c0e716-3128-4172-be09-7f82b7ec02ca','-','')),
    8,
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'E379021B7782'),
    'Slightly less expensive shoes',
    '411510',
    4000,
    'KGM','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
), (
    UNHEX(REPLACE('cf1798fe-9447-4ea8-a4a6-9515de751d5e','-','')),
    8,
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'A379021B7782'),
    'Even more expensive shoes',
    '411510',
    4000,
    'KGM','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
);

insert into consignment_items (
    uuid4,
    shipping_instruction_id,
    shipment_id,
    description_of_goods,
    hs_code,
    weight,
    weight_unit,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('5d943239-23fc-4d5c-ab70-a33a469f9e59','-','')),
    9,
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'D659FDB7E33C'),
    'Expensive shoes',
    '411510',
    4000,
    'KGM','active','2020-03-07 12:12:12.000','2020-04-07 12:12:12.000','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
);
