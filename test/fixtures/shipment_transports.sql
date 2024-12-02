insert into shipment_transports (
    uuid4,
    shipment_id,
    transport_id,
    transport_plan_stage_sequence_number,
    transport_plan_stage_code,
    commercial_voyage_id,
    is_under_shippers_responsibility
) VALUES (
    UNHEX(REPLACE('a1c57fe2-3a38-4e8e-ae60-ac67550402a6','-','')),
     (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'BR1239719871'),
    1,
    1,
    'PRC',
    0,
    false
);

insert into shipment_transports (
    uuid4,
    shipment_id,
    transport_id,
    transport_plan_stage_sequence_number,
    transport_plan_stage_code,
    is_under_shippers_responsibility
) VALUES (
    UNHEX(REPLACE('ccf0fea3-460f-4008-9983-552d1b18ef5b','-','')),
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'ABC123123123'),
    2,
    1,
    'PRC',
    false
);

insert into shipment_transports (
    uuid4,
    shipment_id,
    transport_id,
    transport_plan_stage_sequence_number,
    transport_plan_stage_code,
    commercial_voyage_id,
    is_under_shippers_responsibility
) VALUES (
    UNHEX(REPLACE('46e44424-a3bc-4960-ac65-6c21ae8d2fb2','-','')),
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'DCR987876762'),
    4,
    1,
    'PRC',
    0,
    false
);
