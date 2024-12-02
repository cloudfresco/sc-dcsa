insert into shipment_events (
    event_id_s,
    event_classifier_code,
    event_date_time,
    shipment_event_type_code,
    document_type_code,
    document_id,
    document_reference
) VALUES (
    UNHEX(REPLACE('784871e7-c9cd-4f59-8d88-2e033fa799a1','-','')),
    'ACT',
    '2020-07-15',
    'APPR',
    'BKG',
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'BR1239719971'),
    'BR1239719971'
);

insert into shipment_events (
    event_id_s,
    event_classifier_code,
    event_date_time,
    shipment_event_type_code,
    document_type_code,
    document_id,
    document_reference
) VALUES (
    UNHEX(REPLACE('e48f2bc0-c746-11ea-a3ff-db48243a89f4','-','')),
    'ACT',
    '2020-07-15 13:14:15',
    'APPR',
    'BKG',
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'BR1239719971'),
    'BR1239719971'
);

insert into shipment_events (
    event_id_s,
    event_classifier_code,
    event_date_time,
    shipment_event_type_code,
    document_type_code,
    document_id,
    document_reference
) VALUES (
    UNHEX(REPLACE('5e51e72c-d872-11ea-811c-0f8f10a32ea1','-','')),
    'ACT',
    '2003-05-03 21:02:44',
    'CONF',
    'BKG',
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'ABC123123123'),
    'ABC123123123'
);

insert into shipment_events (
   event_id_s,
   event_classifier_code,
   event_date_time,
   event_created_date_time,
   shipment_event_type_code,
   document_type_code,
   document_id,
   document_reference,
   reason
) VALUES (
   UNHEX(REPLACE('c448751d-3109-4fbf-a2ca-fcab6f5384ee','-','')),
   'ACT',
   '2021-01-08 13:22:53',
   '2021-01-08 13:22:53',
   'RECE',
   'SHI',
   1,
   (SELECT shipping_instructions.shipping_instruction_reference FROM shipping_instructions  WHERE id = 4),
   ''
), (
  UNHEX(REPLACE('aa3ad4a3-e7aa-4a55-b64c-3306157bcb1b','-','')),
  'ACT',
  '2021-01-08 17:22:53',
  '2021-01-08 17:22:53',
  'PENU',
  'SHI',
  1,
  (SELECT shipping_instructions.shipping_instruction_reference FROM shipping_instructions  WHERE id = 4),
  'Carrier Booking Reference present in both shipping instruction as well as cargo items.'
), (
  UNHEX(REPLACE('e2b6f94c-120f-42d9-b012-a7ad5b9faadb','-','')),
  'ACT',
  '2021-01-08 18:22:53',
  '2021-01-08 18:22:53',
  'DRFT',
  'SHI',
  1,
  (SELECT shipping_instructions.shipping_instruction_reference FROM shipping_instructions  WHERE id = 4),
  ''
), (
   UNHEX(REPLACE('ae58743d-68cf-4493-a263-40ec3d78ae15','-','')),
   'ACT',
   '2022-03-01 18:22:53',
   '2022-03-01 18:22:53',
   'RECE',
   'SHI',
   1,
    (SELECT shipping_instructions.shipping_instruction_reference FROM shipping_instructions  WHERE id = 6),
   ''
), (
   UNHEX(REPLACE('e93829d6-fb7b-40f1-ba94-d4fc01c87a38','-','')),
   'ACT',
   '2022-03-03 18:22:53',
   '2022-03-03 18:22:53',
   'DRFT',
   'SHI',
   2,
   (SELECT shipping_instructions.shipping_instruction_reference FROM shipping_instructions  WHERE id = 6),
   ''
), (
   UNHEX(REPLACE('e51cb663-2cdc-47a2-a48c-7de2fbb99670','-','')),
   'ACT',
   '2022-03-03 18:22:53',
   '2022-03-03 18:22:53',
   'DRFT',
   'TRD',
   (SELECT transport_documents.id FROM transport_documents WHERE transport_document_reference = '2b02401c-b2fb-5009'),
   '2b02401c-b2fb-5009',
   ''
), (
   UNHEX(REPLACE('82cddfe9-8e32-4bf7-8c10-0c7100de69bd','-','')),
   'ACT',
   '2022-03-05 13:56:12',
   '2022-03-05 13:56:12',
   'APPR',
   'TRD',
   (SELECT transport_documents.id FROM transport_documents WHERE transport_document_reference = '2b02401c-b2fb-5009'),
   '2b02401c-b2fb-5009',
   ''
 );
 
 insert into shipment_events (
   event_id_s,
   event_classifier_code,
   event_date_time,
   event_created_date_time,
   shipment_event_type_code,
   document_type_code,
   document_id,
   document_reference,
   reason
) VALUES (
   UNHEX(REPLACE('97eb7c09-571e-438f-8f65-ac6a29ba04e5','-','')),
   'ACT',
   '2021-01-08 13:22:53',
   '2021-01-08 13:22:53',
   'RECE',
   'CBR',
   1,
   'cbrr-b83765166707812c8ff4',
   ''
), (
   UNHEX(REPLACE('d7dde15f-5ddc-42ce-8103-9fa1c4da0bde','-','')),
   'ACT',
   '2021-01-08 13:22:53',
   '2021-01-08 13:22:53',
   'RECE',
   'BKG',
   (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'cbr-b83765166707812c8ff4'),
   'cbr-b83765166707812c8ff4',
   ''
), (
   UNHEX(REPLACE('8b654176-fe41-41fd-a457-a632d6811246','-','')),
   'ACT',
   '2021-01-08 13:22:53',
   '2021-01-08 13:22:53',
   'RECE',
   'SHI',
   1,
   'c144c6dff46b9fa67e65',
   ''
);
