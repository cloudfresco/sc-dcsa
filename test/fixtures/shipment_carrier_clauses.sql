insert into shipment_carrier_clauses (
    uuid4,
    carrier_clause_id,
    shipment_id
) VALUES (
    UNHEX(REPLACE('bfd1abc4-e54e-411d-971b-a4315d9cc4c7','-','')),
    1,
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'ABC123123123')
);

insert into shipment_carrier_clauses (
    uuid4,
    carrier_clause_id,
    shipment_id
) VALUES (
    UNHEX(REPLACE('16d50e1b-d672-4564-bee5-dd4760868e80','-','')),
    2,
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'BR1239719871')
);

insert into shipment_carrier_clauses (
    uuid4,
    carrier_clause_id,
    shipment_id
) VALUES (
    UNHEX(REPLACE('1b606f74-c831-4591-8951-55dcfe4ceb18','-','')),
    3,
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'CR1239719872')
);

insert into shipment_carrier_clauses (
    uuid4,
    carrier_clause_id,
    shipment_id
) VALUES (
    UNHEX(REPLACE('09e0f559-bc20-4686-9d80-0d3c00e2c620','-','')),
    4,
    (SELECT shipments.id FROM shipments WHERE carrier_booking_reference = 'BR1239719971')
);
