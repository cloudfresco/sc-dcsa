insert into shipment_locations (
    shipment_id,
    booking_id,
    location_id,
    shipment_location_type_code,
    event_date_time
) VALUES (
    1,
    1,
    2,
    'PRE',
    DATE '2020-03-07'
), (
    1,
    1,
    3,
    'POL',
    DATE '2020-03-07'
), (
    1,
    1,
    4,
    'POD',
    DATE '2020-03-07'
), (
    1,
    1,
    7,
    'PDE',
    DATE '2020-03-07'
),(
      7,
      1,
      7,
      'PDE',
      DATE '2020-03-07'
  );
  
  insert into shipment_locations (
    shipment_id,
    booking_id,
    location_id,
    shipment_location_type_code,
    displayed_name,
    event_date_time
) VALUES (
    0,
    7,
    9,
    'PRE',
    'HELLO!',
    DATE '2020-03-07'
),  (
    0,
    7,
    10,
    'POL',
    'HELLO!',
    DATE '2020-03-07'
), (
    0,
    7,
    11,
    'POD',
    'HELLO!',
    DATE '2020-03-07'
), (
    0,
    7,
    9,
    'PRE',
    'HELLO!',
    DATE '2020-03-07'
),  (
    0,
    8,
    10,
    'POL',
    'HELLO!',
    DATE '2020-03-07'
), (
    0,
    8,
    11,
    'POD',
    'HELLO!',
    DATE '2020-03-07'
);

insert into shipment_locations (
    shipment_id,
    booking_id,
    location_id,
    shipment_location_type_code,
    displayed_name,
    event_date_time
    ) VALUES (
    15,
    10,
    1,
    'POL',
    'Hamburg',
     DATE '2020-03-07');
