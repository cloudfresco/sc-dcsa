# DCSA API's

## Booking API

- Booking request
- Booking update
- Booking confirmation
- Booking amendment
- Booking cancellation
- Booking rejection

### Use Cases

1. **Submit Booking request**  
Shipper to Carrier

2. **Request to update Booking request**  
Carrier to Shipper

3. **Submit updated Booking request**  
Shipper to Carrier

4. **Reject Booking request**  
Carrier to Shipper

5. **Confirm Booking request**  
Carrier to Shipper

6. **Request to amend confirmed Booking**  
Carrier to Shipper

7. **Request amendment to confirmed Booking**  
Shipper, Consignee or Endorsee to Carrier

8. **Process amendment to confirmed Booking**  
Carrier to Shipper, Consignee or Endorsee

9. **Cancel amendment to confirmed Booking**  
Shipper, Consignee or Endorsee to Carrier

10. **Decline Booking by Carrier**  
Carrier to Shipper

11. **Cancel Booking request by Shipper**  
Shipper to Carrier

12. **Confirm Booking completed**  
Carrier to Shipper

13. **Cancel confirmed Booking by Shipper**  
Shipper to Carrier

14. **Process Booking cancellation**  
Carrier to Shipper

### Resources

- DCSA BKG API: [https://app.swaggerhub.com/apis/dcsaorg/DCSA_BKG/2.0.2](https://app.swaggerhub.com/apis/dcsaorg/DCSA_BKG/2.0.2)  
- DCSA BKG API Docs: [https://app.swaggerhub.com/apis-docs/dcsaorg/DCSA_BKG/2.0.2](https://app.swaggerhub.com/apis-docs/dcsaorg/DCSA_BKG/2.0.2)

---

## EBL (Electronic Bill of Lading/Transport Document) API

- Shipping Instructions submission
- Shipping Instructions update
- Draft Transport Document publication
- Draft Transport Document update
- Draft Transport Document approval

### Use Cases

1. **Submit Shipping Instructions**  
Shipper to Carrier

2. **Request to update Shipping Instructions**  
Carrier to Shipper

3. **Submit updated Shipping Instructions**  
Shipper, Consignee or Endorsee to Carrier

4. **Process updated Shipping Instructions**  
Carrier to Shipper, Consignee or Endorsee

5. **Cancel update to Shipping Instructions**  
Shipper, Consignee or Endorsee to Carrier

6. **Publish Draft Transport Document**  
Carrier to Shipper

7. **Approve Draft Transport Document**  
Shipper to Carrier

8. **Issue Transport Document**  
Carrier to Shipper

9. **Request surrender Transport Document (amendment)**  
Shipper, Consignee or Endorsee to Carrier

10. **Process Transport Document surrender request (amendment)**  
Carrier to Shipper, Consignee or Endorsee

11. **Void original Transport Document and issue amended Transport Document**  
Carrier to Shipper, Consignee or Endorsee

12. **Request surrender Transport Document (delivery)**  
Shipper, Consignee or Endorsee to Carrier

13. **Process Transport Document surrender request (delivery)**  
Carrier to Shipper, Consignee or Endorsee

14. **Confirm Shipping Instructions completed**  

### Resources

- DCSA EBL API: [https://app.swaggerhub.com/apis/dcsaorg/DCSA_EBL/3.0.1](https://app.swaggerhub.com/apis/dcsaorg/DCSA_EBL/3.0.1)  
- DCSA EBL API Docs: [https://app.swaggerhub.com/apis-docs/dcsaorg/DCSA_EBL/3.0.1](https://app.swaggerhub.com/apis-docs/dcsaorg/DCSA_EBL/3.0.1)

---

## Track and Trace API

- ShipmentEvents
- TransportEvents
- EquipmentEvents

### Resources

- DCSA TNT API: [https://app.swaggerhub.com/apis/dcsaorg/DCSA_TNT/3.0.0-Beta-2](https://app.swaggerhub.com/apis/dcsaorg/DCSA_TNT/3.0.0-Beta-2)
- DCSA TNT API Docs: [https://app.swaggerhub.com/apis-docs/dcsaorg/DCSA_TNT/3.0.0-Beta-2](https://app.swaggerhub.com/apis-docs/dcsaorg/DCSA_TNT/3.0.0-Beta-2)
