  insert into timestamps (
      uuid4,
      event_type_code,
      event_classifier_code,
      delay_reason_code,
      change_remark,
      event_date_time
  ) VALUES (
      UNHEX(REPLACE('4f6a0367-2e32-46ee-bc4f-a2f3adc03f66','-','')),
      'ARRI',
      'ACT',
      'ANA',
      'Authorities not available',
      '2020-03-07 12:12:12.000'
  );

  insert into timestamps (
      uuid4,
      event_type_code,
      event_classifier_code,
      delay_reason_code,
      change_remark,
      event_date_time
  ) VALUES (
      UNHEX(REPLACE('44dd23c4-5478-4356-a328-42165b95d129','-','')),
      'TTA',
      'PLN',
      'WEA',
      'Bad weather',
      '2020-03-07 12:12:12.000'
  );
