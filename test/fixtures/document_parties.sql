INSERT INTO document_parties (
    uuid4,
    party_id,
    shipment_id,
    party_function,
    is_to_be_notified,
    booking_id
) VALUES (
    UNHEX(REPLACE('8aafdf32-ec6c-4cd9-a096-98aecd7d58eb','-','')),
    8,
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'BR1239719871'),
    'OS',
    true,
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'CARRIER_BOOKING_REQUEST_REFERENCE_01')
), (
    UNHEX(REPLACE('d7bbd1d4-f650-4f64-aa2b-c6a0be666296','-','')),
    9,
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'BR1239719871'),
    'CN',
    true,
    (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'CARRIER_BOOKING_REQUEST_REFERENCE_01')
), (
      UNHEX(REPLACE('b5c103cc-4964-4ada-923c-5f318b2e4ca0','-','')),
      10,
      (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'AR1239719871'),
      'CN',
      true,
      (SELECT bookings.id FROM bookings WHERE carrier_booking_request_reference = 'CARRIER_BOOKING_REQUEST_REFERENCE_01')
  );

INSERT INTO document_parties (
    uuid4,
    party_id,
    shipping_instruction_id,
    shipment_id,
    party_function,
    is_to_be_notified,
    booking_id
    ) VALUES (
    UNHEX(REPLACE('c678ce03-3859-4db3-a23f-d7c3f998fd0a','-','')),
    14,
    0,
    8,
    'DDS',
    true,
    10);
