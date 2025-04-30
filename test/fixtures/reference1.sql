insert into reference1 (
    reference_type_code,
    reference_value,
    shipment_id,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at
) VALUES (
    'CR',
    'AB-123743CR',
    1,
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000'
), (
    'PO',
    'PO0027',
    1,
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000'
), (
    'CR',
    'BC-346267CR',
    2,
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000'
), (
    'PO',
    'PO0028',
    2,
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    '2020-03-07 12:12:12.000',
    '2020-04-07 12:12:12.000'
);

insert into reference1 (
    reference_type_code,
    reference_value,
    shipment_id,
    shipping_instruction_id,
    booking_id,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at
    ) VALUES (
    'FF',
    'test',
    15,
    1,
    10,
'active',
'auth0|66fd06d0bfea78a82bb42459',
'auth0|66fd06d0bfea78a82bb42459',
'2020-03-07 12:12:12.000',
'2020-04-07 12:12:12.000');
    
insert into reference1 (
    reference_type_code,
    reference_value,
    shipment_id,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at
) VALUES (
    'FF',
    'string',
    19,
'active',
'auth0|66fd06d0bfea78a82bb42459',
'auth0|66fd06d0bfea78a82bb42459',
'2020-03-07 12:12:12.000',
'2020-04-07 12:12:12.000'
);   
