insert into transport_events (
    event_id_s,
    event_classifier_code,
    event_date_time,
    transport_event_type_code,
    transport_call_id,
    delay_reason_code,
    change_remark
) VALUES (
    UNHEX(REPLACE('5e51e72c-d872-11ea-811c-0f8f10a32ea3','-','')),
    'ACT',
    '2003-05-03 21:02:44',
    'DEPA',
    7,
    'ANA',
    'Authorities not available'
);

insert into transport_events (
    event_id_s,
    event_classifier_code,
    event_created_date_time,
    event_date_time,
    transport_event_type_code,
    transport_call_id,
    delay_reason_code,
    change_remark
) VALUES (
    UNHEX(REPLACE('84db923d-2a19-4eb0-beb5-446c1ec57d34','-','')),
    'ACT',
    '2021-01-09 14:12:56',
    '2019-11-12 07:41:00',
    'ARRI',
    7,
    'WEA',
    'Bad weather'
);

insert into transport_events (
    event_id_s,
    event_classifier_code,
    event_date_time,
    event_created_date_time,
    transport_event_type_code,
    transport_call_id,
    delay_reason_code,
    change_remark
) VALUES (
    UNHEX(REPLACE('704c41d8-8718-42fc-9051-88e25c5b1770','-','')),
    'EST',
    '2003-05-03 21:02:44',
    '2003-05-01 21:02:44',
    'DEPA',
    8,
    'ANA',
    'Authorities not available'
);

insert into transport_events (
    event_id_s,
    event_classifier_code,
    event_date_time,
    event_created_date_time,
    transport_event_type_code,
    transport_call_id,
    delay_reason_code,
    change_remark
) VALUES (
    UNHEX(REPLACE('25dee589-5c80-42b9-935e-5ae7f7c0193e','-','')),
    'ACT',
    '2003-05-03 21:02:44',
    '2003-05-03 21:02:44',
    'DEPA',
    8,
    'ANA',
    'Authorities not available'
);

insert into transport_events (
  event_id_s,
  event_classifier_code,
  event_created_date_time,
  event_date_time,
  transport_event_type_code,
  transport_call_id,
  delay_reason_code,
  change_remark
) VALUES (
  UNHEX(REPLACE('2968b966-ee81-46ba-af87-0c5031c641f3','-','')),
  'PLN',
  '2021-11-28 14:12:56',
  '2021-12-01 07:41:00',
  'ARRI',
  3,
  'WEA',
  'Bad weather'
), (
    UNHEX(REPLACE('2968b966-ee81-46ba-af87-0c5031c641f2','-','')),
    'PLN',
    '2021-11-28 14:12:56',
    '2021-12-01 07:41:00',
    'DEPA',
    4,
    'WEA',
    'Bad weather'
);

insert into transport_events (
    event_id_s,
    event_classifier_code,
    event_created_date_time,
    event_date_time,
    transport_event_type_code,
    transport_call_id,
    delay_reason_code,
    change_remark
) VALUES (
    UNHEX(REPLACE('2968b966-ee81-46ba-af87-0c5031c641f4','-','')),
    'PLN',
    '2021-11-28 14:12:56',
    '2021-12-01 07:41:00',
    'ARRI',
    1,
    'WEA',
    'Bad weather'
), (
    UNHEX(REPLACE('9d5d0824-b228-4ea8-b2cb-4ebd8da76e15','-','')),
    'PLN',
    '2021-11-29 14:12:56',
    '2021-12-03 07:41:00',
    'DEPA',
    6,
    'WEA',
    'Bad weather'
);
