insert into requested_equipments (
    uuid4,
    booking_id,
    shipment_id,
    requested_equipment_sizetype,
    requested_equipment_units,
    confirmed_equipment_sizetype,
    confirmed_equipment_units,
    is_shipper_owned,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at
    ) VALUES (
    UNHEX(REPLACE('1a595981-c2d9-46a9-a870-3086735b4529','-','')),
    10,
    15,
    '22GP',
    3,
    ' ',
    0,
    true,
'active',
'auth0|66fd06d0bfea78a82bb42459',
'auth0|66fd06d0bfea78a82bb42459',
'2020-03-07 12:12:12.000',
'2020-04-07 12:12:12.000'
);   
