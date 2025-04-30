insert into seals (
    uuid4,
    utilized_transport_equipment_id,
    seal_number,
    seal_source_code,
    seal_type_code,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at
) VALUES (
     UNHEX(REPLACE('8976f0d6-25ff-45c8-97ac-5bd7ef24a3c1','-','')),    
     1,
     'SN123457',
     'CUS',
     'WIR',
     'active',
     'auth0|66fd06d0bfea78a82bb42459',
     'auth0|66fd06d0bfea78a82bb42459',
     '2020-03-07 12:12:12.000',
     '2020-04-07 12:12:12.000'
);
