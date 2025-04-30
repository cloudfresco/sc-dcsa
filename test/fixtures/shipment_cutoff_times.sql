insert into shipment_cutoff_times (
    shipment_id,
    cut_off_time_code,
    cut_off_time
) VALUES (
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'ABC123123123'),
    'AFD',
    DATE '2021-03-09'
);

insert into shipment_cutoff_times (
    shipment_id,
    cut_off_time_code,
    cut_off_time
) VALUES (
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'BR1239719871'),
    'DCO',
    DATE '2021-05-01'
);

insert into shipment_cutoff_times (
    shipment_id,
    cut_off_time_code,
    cut_off_time
) VALUES (
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'CR1239719872'),
    'ECP',
    DATE '2020-07-07'
);

insert into shipment_cutoff_times (
    shipment_id,
    cut_off_time_code,
    cut_off_time
) VALUES (
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'BR1239719971'),
    'EFC',
    DATE '2020-01-06'
);
