INSERT INTO `parties` (
   uuid4,
    party_name,
    address_id,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('4e448f26-4035-11eb-a49d-7f9eb9bc8dd9','-','')),
    'Malwart',
    1, 
    'active','2019-07-23 10:04:26','2019-07-23 10:04:26','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
), (
    UNHEX(REPLACE('8dd9a4c4-4039-11eb-8770-0b2b19847fab','-','')),
    'Malwart Düsseldorf', 
    2,
    'active','2019-07-23 10:04:26','2019-07-23 10:04:26','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
), (
     UNHEX(REPLACE('9dd9a4c4-4039-11eb-8770-0b2b19847fab','-','')),
     'Malwart Lyngy', 
     3,
     'active','2019-07-23 10:04:26','2019-07-23 10:04:26','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
 );

INSERT INTO `parties` (
    uuid4,
    party_name,
    address_id,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('499918a2-d12d-4df6-840c-dd92357002df','-','')),
    'FTL International', 
    4,
    'active','2019-07-23 10:04:26','2019-07-23 10:04:26','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
);


INSERT INTO `parties` (
   uuid4,
    party_name,
    address_id,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('8e463a84-0a2d-47cd-9332-51e6cb36b635','-','')),
    'Superdæk Albertslund', 
    5,
    'active','2019-07-23 10:04:26','2019-07-23 10:04:26','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
);

INSERT INTO `parties` (
   uuid4,
    party_name,
    address_id,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id 
) VALUES (
    UNHEX(REPLACE('c49ea2d6-3806-46c8-8490-294affc71286','-','')),
    'FDM Quality Control', 
    6,
    'active','2019-07-23 10:04:26','2019-07-23 10:04:26','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
);

INSERT INTO `parties` (
    uuid4,
    party_name,
    tax_reference1,
    tax_reference2,
    public_key,
    address_id,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('7bf6f428-58f0-4347-9ce8-d6be2f5d5745','-','')),
    'Hapag Lloyd',
    'CVR-25645774',
    'CVR-25645774',
    'eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IkFzaW',
    7, 
    'active','2019-07-23 10:04:26','2019-07-23 10:04:26','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
);

INSERT INTO  `parties` (
    uuid4,
    party_name,
    tax_reference1,
    tax_reference2,
    public_key,
    address_id,
    status_code,
    created_at,
    updated_at,
    created_by_user_id,
    updated_by_user_id
) VALUES ( 
    UNHEX(REPLACE('be5bc290-7bac-48bb-a211-f3fa5a3ab3ae','-','')),
    'Asseco Denmark',
    'CVR-25645774',
    'CVR-25645774',
    'eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IkFzaW',
    8, 
    'active','2019-07-23 10:04:26','2019-07-23 10:04:26','auth0|66fd06d0bfea78a82bb42459','auth0|66fd06d0bfea78a82bb42459'
);
