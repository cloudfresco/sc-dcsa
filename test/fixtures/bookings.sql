INSERT INTO `bookings` (uuid4,
    carrier_booking_request_reference,
    document_status,
    submission_date_time,
    receipt_type_at_origin,
    delivery_type_at_destination,
    cargo_movement_type_at_origin,
    cargo_movement_type_at_destination,
    created_at,
    service_contract_reference,
    payment_term_code,
    is_partial_load_allowed,
    is_export_declaration_required,
    export_declaration_reference,
    is_import_license_required,
    import_license_reference,
    is_destination_filing_required,
    inco_terms,
    expected_departure_date,
    transport_document_type_code,
    transport_document_reference,
    booking_channel_reference,
    communication_channel_code,
    is_equipment_substitution_allowed,
    vessel_id,
    declared_value_currency,
    declared_value,
    export_voyage_number,
    location_id,
    updated_at,
    invoice_payable_at,
    status_code,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('b521dbdb-a12b-48f5-b489-8594349731bf','-','')),
    'CARRIER_BOOKING_REQUEST_REFERENCE_01',
    'RECE',
    DATE '2020-03-07',
    'CY',
    'CFS',
    'FCL',
    'BB',
    DATE '2020-03-07',
    'SERVICE_CONTRACT_REFERENCE_01',
    'PRE',
    TRUE,
    TRUE,
    'EXPORT_DECLARATION_REFERENCE_01',
    FALSE,
    'IMPORT_LICENSE_REFERENCE_01',
    TRUE,
    'FCA',
    DATE '2020-03-07',
    'SWB',
    'TRANSPORT_DOC_REF_01',
    'BOOKING_CHA_REF_01',
    'EI',
    FALSE,
    (SELECT vessels.id FROM vessels WHERE vessel_imo_number = '9321483'),
    'EUR',
    1212,
    'CARRIER_VOYAGE_NUMBER_01',
    0,
    DATE '2021-12-09',
    'c703277f-84ca-4816-9ccf-fad8e202d3b6',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459'
);

INSERT INTO `bookings` (uuid4,
    carrier_booking_request_reference,
    document_status,
    submission_date_time,
    receipt_type_at_origin,
    delivery_type_at_destination,
    cargo_movement_type_at_origin,
    cargo_movement_type_at_destination,
    created_at,
    service_contract_reference,
    payment_term_code,
    is_partial_load_allowed,
    is_export_declaration_required,
    export_declaration_reference,
    is_import_license_required,
    import_license_reference,
    is_destination_filing_required,
    inco_terms,
    invoice_payable_at,
    expected_departure_date,
    transport_document_type_code,
    transport_document_reference,
    booking_channel_reference,
    communication_channel_code,
    is_equipment_substitution_allowed,
    vessel_id,
    declared_value_currency,
    declared_value,
    export_voyage_number,
    location_id,
    updated_at,
    status_code,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('9e7c5cfe-9725-4df4-b5fc-aec732762f3c','-','')),
    'CARRIER_BOOKING_REQUEST_REFERENCE_02',
    'RECE',
    DATE '2020-03-07',
    'CY',
    'CFS',
    'FCL',
    'BB',
    DATE '2020-03-07',
    'SERVICE_CONTRACT_REFERENCE_02',
    'PRE',
    TRUE,
    TRUE,
    'EXPORT_DECLARATION_REFERENCE_02',
    FALSE,
    'IMPORT_LICENSE_REFERENCE_02',
    TRUE,
    'FCA',
    '84bfcf2e-403b-11eb-bc4a-1fc4aa7d879d',
    DATE '2020-03-07',
    'SWB',
    'TRANSPORT_DOC_REF_02',
    'BOOKING_CHA_REF_02',
    'EI',
    FALSE,
    (SELECT vessels.id FROM vessels WHERE vessel_imo_number = '9321483'),
    'EUR',
    1212,
    'CARRIER_VOYAGE_NUMBER_02',
    8,
    DATE '2021-12-16',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459'
);

INSERT INTO `bookings` (uuid4,
    carrier_booking_request_reference,
    document_status,
    submission_date_time,
    receipt_type_at_origin,
    delivery_type_at_destination,
    cargo_movement_type_at_origin,
    cargo_movement_type_at_destination,
    created_at,
    service_contract_reference,
    payment_term_code,
    is_partial_load_allowed,
    is_export_declaration_required,
    export_declaration_reference,
    is_import_license_required,
    import_license_reference,
    is_destination_filing_required,
    inco_terms,
    expected_departure_date,
    transport_document_type_code,
    transport_document_reference,
    booking_channel_reference,
    communication_channel_code,
    is_equipment_substitution_allowed,
    vessel_id,
    declared_value_currency,
    declared_value,
    export_voyage_number,
    location_id,
    updated_at,
    status_code,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('61ee0dcf-0011-4cad-b33c-dfcca3c5fce4','-','')),
    'CARRIER_BOOKING_REQUEST_REFERENCE_03',
    'CONF',
    DATE '2020-03-07',
    'CY',
    'CFS',
    'FCL',
    'BB',
    DATE '2020-03-07',
    'SERVICE_CONTRACT_REFERENCE_03',
    'PRE',
    TRUE,
    TRUE,
    'EXPORT_DECLARATION_REFERENCE_03',
    FALSE,
    'IMPORT_LICENSE_REFERENCE_03',
    TRUE,
    'FCA',
    DATE '2020-03-07',
    'SWB',
    'TRANSPORT_DOC_REF_03',
    'BOOKING_CHA_REF_03',
    'EI',
    FALSE,
    (SELECT vessels.id FROM vessels WHERE vessel_imo_number = '9321483'),
     'EUR',
     1212,
    'CARRIER_VOYAGE_NUMBER_03',
    0,
    DATE '2021-12-09',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459'
);

INSERT INTO `bookings` (uuid4,
    carrier_booking_request_reference,
    document_status,
    submission_date_time,
    receipt_type_at_origin,
    delivery_type_at_destination,
    cargo_movement_type_at_origin,
    cargo_movement_type_at_destination,
    created_at,
    service_contract_reference,
    payment_term_code,
    is_partial_load_allowed,
    is_export_declaration_required,
    export_declaration_reference,
    is_import_license_required,
    import_license_reference,
    is_destination_filing_required,
    inco_terms,
    invoice_payable_at,
    expected_departure_date,
    transport_document_type_code,
    transport_document_reference,
    booking_channel_reference,
    communication_channel_code,
    is_equipment_substitution_allowed,
    vessel_id,
    declared_value_currency,
    declared_value,
    export_voyage_number,
    location_id,
    updated_at,
    status_code,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('f38f187a-4aa4-4352-b8f8-ed075dcb14d0','-','')),
    'CARRIER_BOOKING_REQUEST_REFERENCE_04',
    'CONF',
    DATE '2020-03-07',
    'CY',
    'CFS',
    'FCL',
    'BB',
    DATE '2020-03-07',
    'SERVICE_CONTRACT_REFERENCE_04',
    'PRE',
    TRUE,
    TRUE,
    'EXPORT_DECLARATION_REFERENCE_04',
    FALSE,
    'IMPORT_LICENSE_REFERENCE_04',
    TRUE,
    'FCA',
    '84bfcf2e-403b-11eb-bc4a-1fc4aa7d879d',
    DATE '2020-03-07',
    'SWB',
    'TRANSPORT_DOC_REF_04',
    'BOOKING_CHA_REF_04',
    'EI',
    FALSE,
    (SELECT vessels.id FROM vessels WHERE vessel_imo_number = '9321483'),
    'EUR',
    1212,
    'CARRIER_VOYAGE_NUMBER_04',
    8,
    DATE '2021-12-16',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459'
);


INSERT INTO `bookings` (uuid4,
    carrier_booking_request_reference,
    document_status,
    submission_date_time,
    receipt_type_at_origin,
    delivery_type_at_destination,
    cargo_movement_type_at_origin,
    cargo_movement_type_at_destination,
    created_at,
    service_contract_reference,
    payment_term_code,
    is_partial_load_allowed,
    is_export_declaration_required,
    export_declaration_reference,
    is_import_license_required,
    import_license_reference,
    is_destination_filing_required,
    inco_terms,
    expected_departure_date,
    transport_document_type_code,
    transport_document_reference,
    booking_channel_reference,
    communication_channel_code,
    is_equipment_substitution_allowed,
    vessel_id,
    declared_value_currency,
    declared_value,
    export_voyage_number,
    location_id,
    updated_at,
    status_code,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('0329b6ac-020d-4e1a-b819-c17dc8cdf959','-','')),
    'CARRIER_BOOKING_REQUEST_REFERENCE_05',
    'CONF',
    DATE '2020-03-07',
    'CY',
    'CFS',
    'FCL',
    'BB',
    DATE '2020-03-07',
    'SERVICE_CONTRACT_REFERENCE_05',
    'PRE',
    TRUE,
    TRUE,
    'EXPORT_DECLARATION_REFERENCE_05',
    FALSE,
    'IMPORT_LICENSE_REFERENCE_05',
    TRUE,
    'FCA',
    DATE '2020-03-07',
    'SWB',
    'TRANSPORT_DOC_REF_05',
    'BOOKING_CHA_REF_05',
    'EI',
    FALSE,
    (SELECT vessels.id FROM vessels WHERE vessel_imo_number = '9321483'),
    'EUR',
    1212,
    'CARRIER_VOYAGE_NUMBER_05',
    0,
    DATE '2021-12-09',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459'
);

INSERT INTO `bookings` (uuid4,
    carrier_booking_request_reference,
    document_status,
    submission_date_time,
    receipt_type_at_origin,
    delivery_type_at_destination,
    cargo_movement_type_at_origin,
    cargo_movement_type_at_destination,
    created_at,
    service_contract_reference,
    payment_term_code,
    is_partial_load_allowed,
    is_export_declaration_required,
    export_declaration_reference,
    is_import_license_required,
    import_license_reference,
    is_destination_filing_required,
    inco_terms,
    invoice_payable_at,
    expected_departure_date,
    transport_document_type_code,
    transport_document_reference,
    booking_channel_reference,
    communication_channel_code,
    is_equipment_substitution_allowed,
    vessel_id,
    declared_value_currency,
    declared_value,
    export_voyage_number,
    location_id,
    updated_at,
    status_code,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('24da3edf-bb66-4f82-9f45-f20ea33e041c','-','')),
    'CARRIER_BOOKING_REQUEST_REFERENCE_06',
    'CONF',
    DATE '2020-03-07',
    'CY',
    'CFS',
    'FCL',
    'BB',
    DATE '2020-03-07',
    'SERVICE_CONTRACT_REFERENCE_06',
    'PRE',
    TRUE,
    TRUE,
    'EXPORT_DECLARATION_REFERENCE_06',
    FALSE,
    'IMPORT_LICENSE_REFERENCE_06',
    TRUE,
    'FCA',
    '84bfcf2e-403b-11eb-bc4a-1fc4aa7d879d',
    DATE '2020-03-07',
    'SWB',
    'TRANSPORT_DOC_REF_06',
    'BOOKING_CHA_REF_06',
    'EI',
    FALSE,
    (SELECT vessels.id FROM vessels WHERE vessel_imo_number = '9321483'),
    'EUR',
    1212,
    'CARRIER_VOYAGE_NUMBER_06',
    8,
    DATE '2021-12-16',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459'
);

INSERT INTO `bookings` (uuid4,
    carrier_booking_request_reference,
    document_status,
    submission_date_time,
    receipt_type_at_origin,
    delivery_type_at_destination,
    cargo_movement_type_at_origin,
    cargo_movement_type_at_destination,
    created_at,
    service_contract_reference,
    payment_term_code,
    is_partial_load_allowed,
    is_export_declaration_required,
    export_declaration_reference,
    is_import_license_required,
    import_license_reference,
    is_destination_filing_required,
    inco_terms,
    expected_departure_date,
    transport_document_type_code,
    transport_document_reference,
    booking_channel_reference,
    communication_channel_code,
    is_equipment_substitution_allowed,
    vessel_id,
    declared_value_currency,
    declared_value,
    export_voyage_number,
    location_id,
    updated_at,
    invoice_payable_at,
    status_code,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('a169d494-d6dd-4334-b951-512e4e16f075','-','')),
    'KUBERNETES_IN_ACTION_01',
    'RECE',
    DATE '2020-03-07',
    'CY',
    'CFS',
    'FCL',
    'BB',
    DATE '2020-03-07',
    'SERVICE_CONTRACT_REFERENCE_01',
    'PRE',
    TRUE,
    TRUE,
    'EXPORT_DECLARATION_REFERENCE_01',
    FALSE,
    'IMPORT_LICENSE_REFERENCE_01',
    TRUE,
    'FCA',
    DATE '2020-03-07',
    'SWB',
    'TRANSPORT_DOC_REF_01',
    'BOOKING_CHA_REF_01',
    'EI',
    FALSE,
    (SELECT vessels.id FROM vessels WHERE vessel_imo_number = '9321483'),
    'EUR',
    1212,
    'CARRIER_VOYAGE_NUMBER_01',
    8,
    DATE '2021-12-09',
    'c703277f-84ca-4816-9ccf-fad8e202d3b6',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459'
), (
    UNHEX(REPLACE('59ede518-2224-4ecf-a0d0-4d641d365e1b','-','')),
    'KUBERNETES_IN_ACTION_02',
    'RECE',
    DATE '2020-03-07',
    'CY',
    'CFS',
    'FCL',
    'BB',
    DATE '2020-03-07',
    'SERVICE_CONTRACT_REFERENCE_01',
    'PRE',
    TRUE,
    TRUE,
    'EXPORT_DECLARATION_REFERENCE_01',
    FALSE,
    'IMPORT_LICENSE_REFERENCE_01',
    TRUE,
    'FCA',
    DATE '2020-03-07',
    'SWB',
    'TRANSPORT_DOC_REF_01',
    'BOOKING_CHA_REF_01',
    'EI',
    FALSE,
    (SELECT vessels.id FROM vessels WHERE vessel_imo_number = '9321483'),
    'EUR',
    1212,
    'CARRIER_VOYAGE_NUMBER_01',
    0,
    DATE '2021-12-09',
    'c703277f-84ca-4816-9ccf-fad8e202d3b6',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459'
);

INSERT INTO `bookings` (uuid4,
    carrier_booking_request_reference,
    document_status,
    submission_date_time,
    receipt_type_at_origin,
    delivery_type_at_destination,
    cargo_movement_type_at_origin,
    cargo_movement_type_at_destination,
    created_at,
    service_contract_reference,
    payment_term_code,
    is_partial_load_allowed,
    is_export_declaration_required,
    export_declaration_reference,
    is_import_license_required,
    import_license_reference,
    is_destination_filing_required,
    inco_terms,
    expected_departure_date,
    transport_document_type_code,
    transport_document_reference,
    booking_channel_reference,
    communication_channel_code,
    is_equipment_substitution_allowed,
    vessel_id,
    declared_value_currency,
    declared_value,
    export_voyage_number,
    location_id,
    updated_at,
    invoice_payable_at,
    status_code,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('66802442-4702-4464-9d61-d659fdb7e33c','-','')),
    'KUBERNETES_IN_ACTION_03',
    'CONF',
    DATE '2020-03-07',
    'CY',
    'CFS',
    'FCL',
    'BB',
    DATE '2020-03-07',
    'SERVICE_CONTRACT_REFERENCE_01',
    'PRE',
    TRUE,
    TRUE,
    'EXPORT_DECLARATION_REFERENCE_01',
    FALSE,
    'IMPORT_LICENSE_REFERENCE_01',
    TRUE,
    'FCA',
    DATE '2020-03-07',
    'SWB',
    'TRANSPORT_DOC_REF_01',
    'BOOKING_CHA_REF_01',
    'EI',
    FALSE,
    (SELECT vessels.id FROM vessels WHERE vessel_imo_number = '9321483'),
    'EUR',
    1212,
    'CARRIER_VOYAGE_NUMBER_01',
    0,
    DATE '2021-12-09',
    'c703277f-84ca-4816-9ccf-fad8e202d3b6',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459'
);

INSERT INTO `bookings` (uuid4,
    carrier_booking_request_reference,
    document_status,
    receipt_type_at_origin,
    delivery_type_at_destination,
    cargo_movement_type_at_origin,
    cargo_movement_type_at_destination,
    created_at,
    service_contract_reference,
    payment_term_code,
    is_partial_load_allowed,
    is_export_declaration_required,
    export_declaration_reference,
    is_import_license_required,
    import_license_reference,
    submission_date_time,
    is_ams_aci_filing_required,
    is_destination_filing_required,
    contract_quotation_reference,
    inco_terms,
    invoice_payable_at,
    expected_departure_date,
    transport_document_type_code,
    transport_document_reference,
    booking_channel_reference,
    communication_channel_code,
    is_equipment_substitution_allowed,
    vessel_id,
    declared_value_currency,
    declared_value,
    export_voyage_number,
    location_id,
    updated_at,
    status_code,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('a521dbdb-a12b-48f5-b489-8594349731bf','-','')),
    'ef223019-ff16-4870-be69-9dbaaaae9b11',
    'PENU',
    'CY',
    'CY',
    'FCL',
    'LCL',
    '2021-11-03 02:11:00.000',
    'Test',
     ' ',
     true,
     true,
     'Export declaration reference',
     true,
     'Import declaration reference',
     '2021-11-03 10:41:00.000',
     true,
     true,
     ' ',
     ' ',
     'c703277f-84ca-4816-9ccf-fad8e202d3b6',
     DATE '2020-03-07',
     ' ',
     ' ',
     ' ',
     'AO',
     true,
     0,
      'EUR',
      1212,
     ' ',
     0,
     DATE '2021-12-01',
     'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459');

INSERT INTO `bookings` (uuid4,
    carrier_booking_request_reference,
    document_status,
    receipt_type_at_origin,
    delivery_type_at_destination,
    cargo_movement_type_at_origin,
    cargo_movement_type_at_destination,
    created_at,
    service_contract_reference,
    payment_term_code,
    is_partial_load_allowed,
    is_export_declaration_required,
    export_declaration_reference,
    is_import_license_required,
    import_license_reference,
    submission_date_time,
    is_ams_aci_filing_required,
    is_destination_filing_required,
    contract_quotation_reference,
    inco_terms,
    invoice_payable_at,
    expected_departure_date,
    transport_document_type_code,
    transport_document_reference,
    booking_channel_reference,
    communication_channel_code,
    is_equipment_substitution_allowed,
    vessel_id,
    declared_value_currency,
    declared_value,
    export_voyage_number,
    location_id,
    updated_at,
    status_code,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('31629725-418b-41e1-9d10-521763c656c4','-','')),
    '52c2caa0-0137-44b7-9947-68687b3b4ae6',
    'PENC',
    'CY',
    'CY',
    'FCL',
    'LCL',
    '2021-11-03 02:11:00.000',
    'Test',
     ' ',
     true,
     true,
     'Export declaration reference',
     true,
     'Import declaration reference',
     '2021-11-03 10:41:00.000',
     true,
     true,
     ' ',
     ' ',
     'c703277f-84ca-4816-9ccf-fad8e202d3b6',
     DATE '2020-03-07',
     ' ',
     ' ',
     ' ',
     'AO',
     true,
     0,
     'EUR',
     1212,
     ' ',
     0,
     DATE '2021-12-01',
     'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459');

INSERT INTO `bookings` (uuid4,
    carrier_booking_request_reference,
    document_status,
    receipt_type_at_origin,
    delivery_type_at_destination,
    cargo_movement_type_at_origin,
    cargo_movement_type_at_destination,
    created_at,
    service_contract_reference,
    payment_term_code,
    is_partial_load_allowed,
    is_export_declaration_required,
    export_declaration_reference,
    is_import_license_required,
    import_license_reference,
    submission_date_time,
    is_ams_aci_filing_required,
    is_destination_filing_required,
    contract_quotation_reference,
    inco_terms,
    invoice_payable_at,
    expected_departure_date,
    transport_document_type_code,
    transport_document_reference,
    booking_channel_reference,
    communication_channel_code,
    is_equipment_substitution_allowed,
    vessel_id,
    declared_value_currency,
    declared_value,
    export_voyage_number,
    location_id,
    updated_at,
    status_code,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('8b78219e-d049-4c68-8d9e-f40bf9a85140','-','')),
    'a3a34f10-acc5-4e23-b52e-146f63458c90',
    'CONF',
    'CY',
    'CY',
    'FCL',
    'LCL',
    '2021-12-20 02:11:00.000',
    'Test',
     ' ',
     true,
     true,
     'Export declaration reference',
     true,
     'Import declaration reference',
     '2021-11-03 10:41:00.000',
     true,
     true,
     ' ',
     ' ',
     'c703277f-84ca-4816-9ccf-fad8e202d3b6',
     DATE '2020-03-07',
     ' ',
     ' ',
     ' ',
     'AO',
     true,
     0,
     'EUR',
     1212,
     ' ',
     0,
     DATE '2021-12-01',
     'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459');
     
  INSERT INTO `bookings` (uuid4,
    carrier_booking_request_reference,
    document_status,
    receipt_type_at_origin,
    delivery_type_at_destination,
    cargo_movement_type_at_origin,
    cargo_movement_type_at_destination,
    created_at,
    service_contract_reference,
    payment_term_code,
    is_partial_load_allowed,
    is_export_declaration_required,
    export_declaration_reference,
    is_import_license_required,
    import_license_reference,
    submission_date_time,
    is_ams_aci_filing_required,
    is_destination_filing_required,
    contract_quotation_reference,
    inco_terms,
    invoice_payable_at,
    expected_departure_date,
    transport_document_type_code,
    transport_document_reference,
    booking_channel_reference,
    communication_channel_code,
    is_equipment_substitution_allowed,
    vessel_id,
    declared_value_currency,
    declared_value,
    export_voyage_number,
    location_id,
    updated_at,
    status_code,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('b8376516-0c1c-4b6f-b51f-6707812c8ff4','-','')),
    'cbrr-b83765166707812c8ff4', /* carrier_booking_request_reference */
    'PENU', /* document_status */
    'CY', /* receipt_type_at_origin */
    'CY', /* delivery_type_at_destination */
    'FCL', /* cargo_movement_type_at_origin */
    'LCL', /* cargo_movement_type_at_destination */
    '2021-11-03 02:11:00.000', /* created_at */
    'Test', /* service_contract_reference */
     ' ', /* payment_term_code */
     true, /* is_partial_load_allowed */
     true, /* is_export_declaration_required */
     'Export declaration reference', /* export_declaration_reference */
     true, /* is_import_license_required */
     'Import declaration reference', /* import_license_reference */
     '2021-11-03 10:41:00.000', /* submission_date_time */
     true, /* is_ams_aci_filing_required */
     true, /* is_destination_filing_required */
     ' ', /* contract_quotation_reference */
     ' ', /* inco_terms */
     ' ', /* invoice_payable_at */
     DATE '2020-03-07', /* expected_departure_date */
     ' ', /* transport_document_type_code */
     ' ', /* transport_document_reference */
     ' ', /* booking_channel_reference */
     'AO', /* communication_channel_code */
     true, /* is_equipment_substitution_allowed */
     0, /* vessel_id */
     'EUR',
     1212,
     ' ', /* export_voyage_number */
     0, /* location_id */
     DATE '2021-12-01', /* updated_at */
     'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459'
);

INSERT INTO `bookings` (uuid4,
    carrier_booking_request_reference,
    document_status,
    submission_date_time,
    receipt_type_at_origin,
    delivery_type_at_destination,
    cargo_movement_type_at_origin,
    cargo_movement_type_at_destination,
    created_at,
    service_contract_reference,
    payment_term_code,
    is_partial_load_allowed,
    is_export_declaration_required,
    export_declaration_reference,
    is_import_license_required,
    import_license_reference,
    is_destination_filing_required,
    inco_terms,
    expected_departure_date,
    transport_document_type_code,
    transport_document_reference,
    booking_channel_reference,
    communication_channel_code,
    is_equipment_substitution_allowed,
    vessel_id,
    declared_value_currency,
    declared_value,
    export_voyage_number,
    updated_at,
    status_code,
    created_by_user_id,
    updated_by_user_id
) VALUES (
    UNHEX(REPLACE('37d0abab-bb7a-4da2-9b5f-ee1314ebe74e','-','')),
    'BR1239719971',
    'PENU',
    DATE '2020-03-07',
    'CY',
    'CFS',
    'FCL',
    'LCL',
    DATE '2020-03-07',
    'SERVICE_CONTRACT_REFERENCE_01',
    'PRE',
    TRUE,
    TRUE,
    'EXPORT_DECLARATION_REFERENCE_01',
    FALSE,
    'IMPORT_LICENSE_REFERENCE_01',
    TRUE,
    'FCA',
    DATE '2020-03-07',
    'SWB',
    'TRANSPORT_DOC_REF_01',
    'BOOKING_CHA_REF_01',
    'EI',
    FALSE,
    (SELECT vessels.id FROM vessels WHERE vessel_imo_number = '9321483'),
    'EUR',
    1212,
    'CARRIER_VOYAGE_NUMBER_01',
    DATE '2021-11-04',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459'
);

INSERT INTO `bookings` (uuid4,
    carrier_booking_request_reference,
    document_status,
    submission_date_time,
    receipt_type_at_origin,
    delivery_type_at_destination,
    cargo_movement_type_at_origin,
    cargo_movement_type_at_destination,
    created_at,
    service_contract_reference,
    payment_term_code,
    is_partial_load_allowed,
    is_export_declaration_required,
    export_declaration_reference,
    is_import_license_required,
    import_license_reference,
    is_destination_filing_required,
    inco_terms,
    expected_departure_date,
    transport_document_type_code,
    transport_document_reference,
    booking_channel_reference,
    communication_channel_code,
    is_equipment_substitution_allowed,
    vessel_id,
    declared_value_currency,
    declared_value,
    export_voyage_number,
    updated_at,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    expected_arrival_at_place_of_delivery_start_date,
    expected_arrival_at_place_of_delivery_end_date
) VALUES (
    UNHEX(REPLACE('0bbb347a-813b-448d-926e-dc68b4863693','-','')),
    'BR1239719872',
    'PENU',
    DATE '2020-04-15',
    'CY',
    'CFS',
    'FCL',
    'LCL',
    DATE '2020-04-15',
    'SERVICE_CONTRACT_REFERENCE_02',
    'PRE',
    TRUE,
    TRUE,
    'EXPORT_DECLARATION_REFERENCE_02',
    FALSE,
    'IMPORT_LICENSE_REFERENCE_02',
    TRUE,
    'FCA',
    DATE '2020-04-15',
    'SWB',
    'TRANSPORT_DOC_REF_02',
    'BOOKING_CHA_REF_02',
    'EI',
    FALSE,
    (SELECT vessels.id FROM vessels WHERE vessel_imo_number = '9321483'),
    'EUR',
    1212,
    'CARRIER_VOYAGE_NUMBER_02',
    DATE '2021-01-10',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    DATE '2020-04-16',
    DATE '2020-04-17'
);

INSERT INTO `bookings` (uuid4,
    carrier_booking_request_reference,
    document_status,
    submission_date_time,
    receipt_type_at_origin,
    delivery_type_at_destination,
    cargo_movement_type_at_origin,
    cargo_movement_type_at_destination,
    created_at,
    service_contract_reference,
    payment_term_code,
    is_partial_load_allowed,
    is_export_declaration_required,
    export_declaration_reference,
    is_import_license_required,
    import_license_reference,
    is_destination_filing_required,
    inco_terms,
    expected_departure_date,
    transport_document_type_code,
    transport_document_reference,
    booking_channel_reference,
    communication_channel_code,
    is_equipment_substitution_allowed,
    vessel_id,
    declared_value_currency,
    declared_value,
    export_voyage_number,
    updated_at,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    expected_arrival_at_place_of_delivery_start_date,
    expected_arrival_at_place_of_delivery_end_date
) VALUES (
    UNHEX(REPLACE('2b95bbb9-d2d4-4416-acfe-6207e181b5f4','-','')),
    'ABC123123123',
    'RECE',
    DATE '2020-03-10',
    'CY',
    'CFS',
    'FCL',
    'LCL',
    DATE '2020-03-10',
    'SERVICE_CONTRACT_REFERENCE_03',
    'PRE',
    TRUE,
    TRUE,
    'EXPORT_DECLARATION_REFERENCE_03',
    FALSE,
    'IMPORT_LICENSE_REFERENCE_03',
    TRUE,
    'FCA',
    DATE '2020-03-10',
    'SWB',
    'TRANSPORT_DOC_REF_03',
    'BOOKING_CHA_REF_03',
    'EI',
    FALSE,
    (SELECT vessels.id FROM vessels WHERE vessel_imo_number = '9321483'),
    'EUR',
    1212,
    'CARRIER_VOYAGE_NUMBER_03',
    DATE '2021-12-16',
    'active',
    'auth0|66fd06d0bfea78a82bb42459',
    'auth0|66fd06d0bfea78a82bb42459',
    DATE '2020-03-12',
    DATE '2020-03-13'
);
