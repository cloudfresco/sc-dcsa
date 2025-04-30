insert into point_to_point_routings (
    uuid4,
    sequence_number,
    place_of_receipt_id,
    place_of_delivery_id,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at
) VALUES (
    UNHEX(REPLACE('df9d6bcc-f983-4d8a-ae62-37ba5d572c4c','-','')),
    10,
    1,
    1,
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000'
);

insert into point_to_point_routings (
    uuid4,
    sequence_number,
    place_of_receipt_id,
    place_of_delivery_id,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at
) VALUES (
    UNHEX(REPLACE('48230edb-2b87-4f26-a964-15095fefd7a4','-','')),
    20,
    2,
    2,
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000'
);
