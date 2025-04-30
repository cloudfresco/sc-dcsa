insert into operations_events (
    event_id_s,
    event_classifier_code,
    event_date_time,
    event_created_date_time,
    operations_event_type_code,
    transport_call_id,
    delay_reason_code,
    publisher_role,
    port_call_service_type_code,
    event_location,
    facility_type_code,
    publisher
) VALUES (
    UNHEX(REPLACE('a0993e06-a222-42ec-816f-ec1d775cfd10','-','')),
    'ACT',
    '2022-05-03 21:02:44', 
    '2022-05-08 13:22:53',
    'DEPA',
    10,
    'ANA',
    'TR',
    'BUNK',
    '06aca2f6-f1d0-48f8-ba46-9a3480adfd23',
    'BRTH',
    'be5bc290-7bac-48bb-a211-f3fa5a3ab3ae'
), (
    UNHEX(REPLACE('03482296-ef9c-11eb-9a03-0242ac130003','-','')),
    'EST',
    '2022-05-03 21:02:44', 
    '2022-02-08 13:22:53',
    'ARRI',
    9,
    'ANA',
    'CA',
    'SAFE',
    '6748a259-fb7e-4f27-9a88-3669e8b9c5f8',
    'BRTH',
    'be5bc290-7bac-48bb-a211-f3fa5a3ab3ae'
);

insert into operations_events (
    event_id_s,
    event_date_time,
    event_created_date_time,
    event_classifier_code,
    publisher,
    publisher_role,
    operations_event_type_code,
    event_location,
    transport_call_id,
    port_call_service_type_code,
    facility_type_code,
    delay_reason_code,
    vessel_position,
    remark,
    port_call_phase_type_code    
) VALUES (
    UNHEX(REPLACE('d330b6f5-edcb-4e9e-a09f-e98e91deba95','-','')),
    DATE '2022-03-07',
    '2022-05-03 21:02:44',
    'REQ',
    'c49ea2d6-3806-46c8-8490-294affc71286',
    'TR',
    'ARRI',
    'b4454ae5-dcd4-4955-8080-1f986aa5c6c3',
    1,
    '',
    'BRTH',
    '',
    '1d09e9e9-dba3-4de1-8ef8-3ab6d32dbb40',
    '',
    'INBD'
), (
    UNHEX(REPLACE('538312da-674c-4278-bf9f-10e2a7c018e3','-','')),
    DATE '2022-03-07', /* event_date_time */
    '2022-05-03 21:02:44',
    'PLN', /* event_classifier_code */
    '7bf6f428-58f0-4347-9ce8-d6be2f5d5745', /* publisher */
    'PLT', /* publisher_role */
    'STRT', /* operations_event_type_code */
    '06aca2f6-f1d0-48f8-ba46-9a3480adfd23', /* event_location */
    10, /* transport_call_id */
    'PILO', /* port_call_service_type_code */
    '', /* facility_type_code */
    'ANA', /* delay_reason_code */
    '', /* vessel_position */
    '', /* remark */
    'INBD' /* port_call_phase_type_code */
);
