insert into shipping_instructions (
    uuid4,
    shipping_instruction_reference,
    document_status,
    is_shipped_onboard_type,
    number_of_copies,
    number_of_originals,
    is_electronic,
    is_to_order,
    are_charges_displayed_on_originals,
    are_charges_displayed_on_copies,
    created_date_time,
    updated_date_time
) VALUES (
    UNHEX(REPLACE('01670315-a51f-4a11-b947-ce8e245128eb','-','')),  
    'SI_REF_1',
    'RECE',
    TRUE,
    2,
    4,
    TRUE,
    TRUE,
    TRUE,
    FALSE,
    DATE '2021-12-24',
    DATE '2021-12-31'
);

/**
 * Data used in integration tests - Do not modify - make your own data
 */
insert into shipping_instructions (
    uuid4,
    shipping_instruction_reference,
    document_status,
    is_shipped_onboard_type,
    number_of_copies,
    number_of_originals,
    is_electronic,
    is_to_order,
    are_charges_displayed_on_originals,
    are_charges_displayed_on_copies,
    created_date_time,
    updated_date_time
) VALUES (
    UNHEX(REPLACE('9d5965a5-9e2f-4c78-b8cb-fbb7095e13a0','-','')),  
    'SI_REF_2',
    'APPR',
    TRUE,
    2,
    4,
    TRUE,
    TRUE,
    TRUE,
    FALSE,
    DATE '2022-01-24',
    DATE '2022-01-31'
),(
    UNHEX(REPLACE('877ce0f8-3126-45f5-b22e-2d1d27d42d85','-','')),  
    'SI_REF_3',
    'RECE',
    TRUE,
    2,
    4,
    TRUE,
    TRUE,
    TRUE,
    FALSE,
    DATE '2022-02-01',
    DATE '2022-02-07'
),(
    UNHEX(REPLACE('770f11e5-aae2-4ae4-b27e-0c689ed2e333','-','')),  
    'SI_REF_4',
    'RECE',
    TRUE,
    2,
    4,
    TRUE,
    TRUE,
    TRUE,
    FALSE,
    DATE '2021-02-08',
    DATE '2021-02-09'
),(
    UNHEX(REPLACE('cb6354c9-1ceb-452c-aed0-3cb25a04647a','-','')),  
    'SI_REF_5',
    'PENU',
    TRUE,
    2,
    4,
    TRUE,
    TRUE,
    TRUE,
    FALSE,
    DATE '2021-02-08',
    DATE '2021-02-09'
),(
    UNHEX(REPLACE('8fbb78cc-e7c6-4e17-9a23-24dc3ad0378d','-','')),  
    'SI_REF_6',
    'APPR',
    TRUE,
    2,
    4,
    TRUE,
    TRUE,
    TRUE,
    FALSE,
    DATE '2022-03-01',
    DATE '2022-03-07'
),(
      UNHEX(REPLACE('9fbb78cc-e7c6-4e17-9a23-24dc3ad0378d','-','')),  
      'SI_REF_7',
      'APPR',
      TRUE,
      2,
      4,
      TRUE,
      TRUE,
      TRUE,
      FALSE,
      DATE '2022-03-01',
      DATE '2022-03-07'
);

insert into shipping_instructions (
    uuid4,
    shipping_instruction_reference,
    document_status,
    is_shipped_onboard_type,
    number_of_copies,
    number_of_originals,
    is_electronic,
    is_to_order,
    are_charges_displayed_on_originals,
    are_charges_displayed_on_copies,
    created_date_time,
    updated_date_time
) VALUES (
    UNHEX(REPLACE('a1c7b95d-3004-40a5-bae1-e379021b7782','-','')),  
    'SI_REF_9',
    'RECE',
    TRUE,
    2,
    4,
    TRUE,
    TRUE,
    TRUE,
    FALSE,
    DATE '2021-12-24',
    DATE '2021-12-31'
);

insert into shipping_instructions (
    uuid4,
    shipping_instruction_reference,
    document_status,
    is_shipped_onboard_type,
    number_of_copies,
    number_of_originals,
    is_electronic,
    is_to_order,
    are_charges_displayed_on_originals,
    are_charges_displayed_on_copies,
    created_date_time,
    updated_date_time
) VALUES (
    UNHEX(REPLACE('2c337fcc-2814-42b3-a752-f1847e74cba7','-','')),  
    'SI_REF_10',
    'DRFT',
    TRUE,
    2,
    4,
    TRUE,
    TRUE,
    TRUE,
    FALSE,
    DATE '2021-12-24',
    DATE '2021-12-31'
);

insert into shipping_instructions (
    uuid4,
    shipping_instruction_reference,
    document_status,
    is_shipped_onboard_type,
    number_of_copies,
    number_of_originals,
    is_electronic,
    is_to_order,
    are_charges_displayed_on_originals,
    are_charges_displayed_on_copies,
    created_date_time,
    updated_date_time
) VALUES (
    UNHEX(REPLACE('c144c6df-440e-4065-8430-f46b9fa67e65','-','')),  
    'c144c6dff46b9fa67e65',
    'RECE',
    TRUE,
    2,
    4,
    TRUE,
    TRUE,
    TRUE,
    FALSE,
    DATE '2021-12-24',
    DATE '2021-12-31'
);
