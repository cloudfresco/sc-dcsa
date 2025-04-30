insert into equipment_events (
    event_id_s,
    event_classifier_code,
    event_created_date_time,
    event_date_time,
    equipment_event_type_code,
    transport_call_id,
    equipment_reference,
    empty_indicator_code
) VALUES (
    UNHEX(REPLACE('5e51e72c-d872-11ea-811c-0f8f10a32ea2','-','')),
    'ACT',
    '2003-05-02 21:02:44',
    '2003-05-03 21:02:44',
    'LOAD',
    7,
    'equipref3453',
    'EMPTY'
);

insert into equipment_events (
    event_id_s,
    event_classifier_code,
    event_created_date_time,
    event_date_time,
    equipment_event_type_code,
    transport_call_id,
    empty_indicator_code,
    equipment_reference
) VALUES (
    UNHEX(REPLACE('84db923d-2a19-4eb0-beb5-446c1ec57d34','-','')),
    'EST',
    '2021-01-09 14:12:56',
    '2019-11-12 07:41:00',
    'LOAD',
    7,
    'EMPTY',
    'APZU4812090'
);
