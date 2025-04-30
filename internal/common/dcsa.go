package common

// https://github.com/dcsaorg/DCSA-Information-Model/tree/master/datamodel/referencedata.d

// cargomovementtypes
const (

	// Full Container Load - The shipper/ consignee or its agent or subcontractor is responsible for stuffing or stripping the cargo into or out of the container and bears every responsibility and liability in such respect.
	CargoMovementFCL string = "FCL"

	// Less than Container Load - The carrier or its agent or subcontractor is responsible for stuffing or stripping the cargo into or out of the container and bears every responsibility and liability in such respect.
	CargoMovementLCL string = "LCL"

	// Break bulk 	Indicates that the carrier has received the cargo which is not containerised.
	CargoMovementBB string = "BB"
)

// codelistresponsibleagencycodes
const (
	// ISO,International Organization for Standardization
	CodeListResponsibleAgencyISO string = "ISO"
	// UN/ECE,United Nations Economic Commission for Europe
	CodeListResponsibleAgencyUNECE string = "UNECE"
	// LR,Lloyd's register of shipping
	CodeListResponsibleAgencyLLOYD string = "LLOYD"
	// BIC (Bureau International des Containeurs),The container industry's international organisation responsible for the issuance of container-related codes
	CodeListResponsibleAgencyBIC string = "BIC"
	// IMO (incl. IMDG),International Maritime Organisation
	CodeListResponsibleAgencyIMO string = "IMO"
	// "US, Standard Carrier Alpha Code (Motor)",Organisation maintaining the SCAC lists and transportation operating in North America
	CodeListResponsibleAgencySCAC string = "SCAC"
	// ITIGG,International Transport Implementation Guidelines Group
	CodeListResponsibleAgencyITIGG string = "ITIGG"
	// ITU,International Telecommunication Union
	CodeListResponsibleAgencyITU string = "ITU"
	// SMDG (Ship-planning Message Design Group),User Group for Shipping Lines and Container Terminals
	CodeListResponsibleAgencySMDG string = "SMDG"
	// EXIS,Exis Technologies Ltd.
	CodeListResponsibleAgencyEXIS string = "EXIS"
	// FMC,Federal Maritime Commission
	CodeListResponsibleAgencyFMC string = "FMC"
	// CBSA,Canada Border Services Agency
	CodeListResponsibleAgencyCBSA string = "CBSA"
	// zzz,Mutually defined,A code assigned within a code list to be used on an interim basis and as defined among trading partners until a precise code can be assigned to the code list
	CodeListResponsibleAgencyZZZ string = "ZZZ"
	// Digitial Container Shipping Association,Maintainer of the EBL Solution provider list
	CodeListResponsibleAgencyDCSA string = "DCSA"
)

// communicationchannelqualifier
const (
	// Electronic mail,Exchange of mail by electronic means
	CommunicationChannelQualifierEM string = "EM"
	// EDI transmission,Number identifying the service and service user
	CommunicationChannelQualifierEI string = "EI"
	// API,Uniform Resource Location
	CommunicationChannelQualifierAO string = "AO"
)

// cutofftimecodes
const (
	// Documentation cut-off,Document cut-off time for SI
	CutOffTimeDCO string = "DCO"
	// VGM cut-off,VGM cut-off time for submission
	CutOffTimeVGM string = "VGM"
	// FCL delivery cut-off,Latest deadline for delivering containers at the terminal gate
	CutOffTimeFCO string = "FCO"
	// LCL delivery cut-off,Latest deadline for delivering LCL cargo at the container freight station
	CutOffTimeLCO string = "LCO"
	// Empty container pick-up date and time,Time and date for shipper to pick-up empty container(s)
	CutOffTimeECP string = "ECP"
	// Earliest full-container delivery date,Earliest date where containers can be delivered at the terminal gate also called gate-opening
	CutOffTimeEFC string = "EFC"
	// AMS filing due date,Date when AMS filing should latest be done in the last port of call before visiting the first US port
	CutOffTimeAFD string = "AFD"
)

// documenttypecodes
const (
	// Booking
	DocumentTypeBKG string = "BKG"
	// Shipping Instructions
	DocumentTypeSHI string = "SHI"
	// Verified Gross Mass
	DocumentTypeVGM string = "VGM"
	// Shipment Release Message
	DocumentTypeSRM string = "SRM"
	// Transport document
	DocumentTypeTRD string = "TRD"
	// Arrival Notice
	DocumentTypeARN string = "ARN"
	// Cargo Survey
	DocumentTypeCAS string = "CAS"
	// Customs Inspection
	DocumentTypeCUS string = "CUS"
	// Dangerous Good Declaration
	DocumentTypeDGD string = "DGD"
	// Out Of Gauge
	DocumentTypeOOG string = "OOG"
	// Carrier Booking Request Reference
	DocumentTypeCBR string = "CBR"
)

// eblsolutionproviders
const (
	// TRAL,https://www.tradelens.com/
	EblsolutionProviderTRAL string = "TRAL"
	// Wave BL, https://wavebl.com/
	EblsolutionProviderWAVE string = "WAVE"
	// CargoX, https://cargox.io
	EblsolutionProviderCARX string = "CARX"
	// EssDocs, https://www.essdocs.com/
	EblsolutionProviderESSD string = "ESSD"
	// Bolero, https://www.bolero.net/
	EblsolutionProviderBOLE string = "BOLE"
	// EdoxOnline, https://web.edoxonline.com
	EblsolutionProviderEDOX string = "EDOX"
)

// emptyindicatorcodes
const (
	EmptyIndicatorEMPTY string = "EMPTY"
	EmptyIndicatorLADEN string = "LADEN"
)

// equipmenteventtypecodes
const (
	// Load,The action of lifting cargo or a container on board of the mode of transportation. Load is complete once the cargo or container has been lifted on board the mode of transport and secured.
	EquipmentEventTypeLOAD string = "LOAD"
	// Discharge,The action of lifting cargo or containers off a mode of transport. Discharge is the opposite of load.GTIN,Gate in,The action when a container is introduced into a controlled area like a port - or inland terminal. Gate in has been completed once the operator of the area is legally in possession of the container.
	EquipmentEventTypeDISC string = "DISC"
	// Gate in,The action when a container is introduced into a controlled area like a port - or inland terminal. Gate in has been completed once the operator of the area is legally in possession of the container.
	EquipmentEventTypeGTIN string = "GTIN"
	// Gate out,The action when a container is removed from a controlled area like a port – or inland terminal. Gate-out has been completed once the possession of the container has been transferred from the operator of the terminal to the entity who is picking up the container.
	EquipmentEventTypeGTOT string = "GTOT"
	// Stuffing,The process of loading the cargo in a container or in/onto another piece of equipment.
	EquipmentEventTypeSTUF string = "STUF"
	// STRP,Stripping,The action of unloading cargo from containers or equipment.
	EquipmentEventTypeSTRP string = "STRP"
	// Pick-up,The action of collecting the container at customer location.
	EquipmentEventTypePICK string = "PICK"
	// Drop-off, The action of delivering the container at customer location.
	EquipmentEventTypeDROP string = "DROP"
	// Inspected,Identifies that the seal on equipment has been inspected.
	EquipmentEventTypeINSP string = "INSP"
	// Resealed,Identifies that the equipment has been resealed after inspection
	EquipmentEventTypeRSEA string = "RSEA"
	// Removed,Identifies that a Seal has been removed from the equipment for inspection
	EquipmentEventTypeRMVD string = "RMVD"
)

// eventclassifiercodes
const (
	// Actual
	EventClassifierACT string = "ACT"
	// Planned
	EventClassifierPLN string = "PLN"
	// Estimated
	EventClassifierEST string = "EST"
	// Removed
	EventClassifierREQ string = "REQ"
)

// facilitytypes
const (
	// Border crossing,"Border crossing is the point at a border between two countries where people, transports or goods can cross. This may or may not include a customs checkpoint."
	FacilityTypeBOCR string = "BOCR"
	// Customer location,"Customer location is the premise of the customer, who can be either the shipper or the consignee."
	FacilityTypeCLOC string = "CLOC"
	// Container freight station,"Container freight station is a facility where LCL (Less Than Container Load) shipments are consolidated or dispersed, cargo is stuffed into containers prior to shipment, or cargo is stripped from containers prior to release to the consignee."
	FacilityTypeCOFS string = "COFS"
	// Container yard,"Deprecated, now called OFFD."
	FacilityTypeCOYA string = "COYA"
	// Off dock storage,"An interim storage facility where empty or full containers are stored in transit."DEPO,Depot,"Depot is a designated area where empty equipment is stored between use."
	FacilityTypeOFFD string = "OFFD"
	// Depot,"Depot is a designated area where empty equipment is stored between use."
	FacilityTypeDEPO string = "DEPO"
	// Inland terminal,"Inland terminal is a facility where containers are loaded, moved, or discharged. The inland terminal can be serviced by trucks, rail, and barges (at river terminals)."
	FacilityTypeINTE string = "INTE"
	// Port terminal,"Port terminal is a facility located adjacent to a waterway where containers are loaded, moved, or discharged onto/from sea-going vessels and barges. "
	FacilityTypePOTE string = "POTE"
	// Pilot boarding place,"The place where a pilot boards the vessel upon arrival at the port boundaries."
	FacilityTypePBPL string = "PBPL"
	// Berth,"A designated location in a port or harbour used for mooring vessels when they are not at sea."
	FacilityTypeBRTH string = "BRTH"
	// Ramp,"An inland container terminal (storing both full and empty containers) connected directly to a rail ramp where containers are loaded/discharged to/from a train."
	FacilityTypeRAMP string = "RAMP"
)

// incotermscodes
const (
	// Free Carrier,"The seller delivers the goods, cleared for export, at a named place (possibly including the seller's own premises)"
	IncotermsFCA string = "FCA"
	// Free on Board,Under FOB terms the seller bears all costs and risks up to the point the goods are loaded on board the vessel.
	IncotermsFOB string = "FOB"
)

// modeoftransportcodes
const (
	// Maritime transport,Transport of goods and/or persons is by sea.,VESSEL
	ModeOfTransportMaritime string = "Maritime transport"
	// Rail transport,Transport of goods and/or persons is by rail.,RAIL
	ModeOfTransportRail string = "Rail transport"
	// Road transport,Transport of goods and/or persons is by road.,TRUCK
	ModeOfTransportRoad string = "Road transport"
	// Inland water Transport,Transport of goods and/or persons is by inland water.,BARGE
	ModeOfTransportInland string = "Inland water"
)

// operationseventtypecodes
const (
	// Arrived
	OperationsEventARRI string = "ARRI"
	// Departed
	OperationsEventDEPA string = "DEPA"
	// Started
	OperationsEventSTRT string = "STRT"
	// Completed
	OperationsEventCMPL string = "CMPL"
)

// packagecodes
const (
	// Barrels, Wooden, bung type
	Package2C1 string = "2C1"
	// Boxes, Natural wood, ordinary
	Package4C1 string = "4C1"
	// Bags, Plastic film
	Package5H4 string = "5H4"
	// Bags, Paper, multiwall
	Package5M1 string = "5M1"
)

// partyfunctioncodes
const (
	// Original shipper.,The original supplier of the goods.
	PartyFunctionOS string = "OS"
	// Consignee.,Party to which goods are consigned.
	PartyFunctionCN string = "CN"
	// Freight payer on behalf of the consignor.,Freight payer is a third party acting on behalf of the consignor.
	PartyFunctionCOW string = "COW"
	// Freight payer on behalf of the consignee.,Freight payer is a third party acting on behalf of the consignee.
	PartyFunctionCOX string = "COX"
	// "Document/message issuer/sender",Issuer of a document and/or sender of a message.
	PartyFunctionMS string = "MS"
	// First Notify Party.,The first party which is to be notified.
	PartyFunctionN1 string = "N1"
	// Second Notify Party.,The second party which is to be notified.
	PartyFunctionN2 string = "N2"
	// Notify party,Party to be notified of arrival of goods.
	PartyFunctionNI string = "NI"
	// "Consignor's freight forwarder",Identification of freight forwarder giving services to the consignor (shipper).
	PartyFunctionDDR string = "DDR"
	// "Consignee's freight forwarder",Identification of freight forwarder giving services to the consignee.CA,Carrier operations,"The activity of discharging, shifting, loading, and lashing containers (both full and empty) as well as other cargo from/to a vessel during port stay. Normally quantified with a move-count and a number of moves per hour."
	PartyFunctionDDS string = "DDS"
	// Carrier local agent,
	PartyFunctionAG string = "AG"
	// Vessel,"A floating, sea going structure (mother vessels and feeder vessels) with either an internal or external mode of propulsion designed for the transport of cargo and/or passengers. Ocean vessels are uniquely identified by an IMO number consisting of 7 digits, or alternatively by their AIS signal with an MMSI number"
	PartyFunctionVSL string = "VSL"
	// Port Authorities
	PartyFunctionATH string = "ATH"
	// Pilot,The activity of conducting a vessel within restricted waters.
	PartyFunctionPLT string = "PLT"
	// Towage provider
	PartyFunctionTWG string = "TWG"
	// Lashing provider
	PartyFunctionLSH string = "LSH"
	// Bunkering service provider
	PartyFunctionBUK string = "BUK"
	// Terminal,"A facility for loading, moving or discharging containers. Terminals can be both inland terminals for trucks and rail or port terminals are accessed by vessels and these can contain multiple berths."EBL,EBL Solution Provider,"The solution provider for managing the electronic bill of lading."
	PartyFunctionTR string = "TR"
	// EBL Solution Provider,"The solution provider for managing the electronic bill of lading."
	PartyFunctionSCO string = "SCO"
	// Booking Agent,"Party acting as a booking office for transport and forwarding services."
	PartyFunctionBA string = "BA"
	// Carrier booking office (transportation office),"The carrier office responsible for processing the booking."
	PartyFunctionHE string = "HE"
)

// paymentterms
const (
	// PrePaid,Fee paid prior to transportation
	PaymentTermPRE string = "PRE"
	// Collect,Fee paid upon collection of the goods
	PaymentTermCOL string = "COL"
)

// portcallphasetype
const (
	// Inbound,Ship’s physical movement from approach to (anchor) berth
	PortCallPhaseTypeINBD string = "INBD"
	// Alongside,Time from First Line Secured till Last Line Released
	PortCallPhaseTypeALGS string = "ALGS"
	// Shifting,Ship’s physical movement from (anchor) berth to (anchor) berth
	PortCallPhaseTypeSHIF string = "SHIF"
	// Outbound,Ship’s physical movement from (anchor) berth to its next destination
	PortCallPhaseTypeOUTB string = "OUTB"
)

// publisherrole
const (
	// Carrier operations
	PublisherRoleCA string = "CA"
	// Carrier local agent
	PublisherRoleAG string = "AG"
	// Vessel
	PublisherRoleVSL string = "VSL"
	// Port Authorities
	PublisherRoleATH string = "ATH"
	// Pilot
	PublisherRolePLT string = "PLT"
	// Towage provider
	PublisherRoleTWG string = "TWG"
	// Lashing provider
	PublisherRoleLSH string = "LSH"
	// Bunkering service provider
	PublisherRoleBUK string = "BUK"
	// Terminal
	PublisherRoleTR string = "TR"
)

// receiptdeliverytypes
const (
	// Container yard (incl. rail ramp), "Where the carrier takes possession of a fully stuffed container delivered by the customer at the carrier or carrier's appointed supplier's facility or where a container is released to the customer by the carrier."
	ReceiptDeliveryTypeCY string = "CY"
	// Store Door, "Indicating that the carrier is taking possession of or delivers a fully stuffed container at the customer's appointed premises."
	ReceiptDeliveryTypeSD string = "SD"
	// Container Freight Station, "indicating that the carrier has received the cargo and is responsible for stuffing of the cargo within the container or the customer receives the cargo directly from the container freight station."
	ReceiptDeliveryTypeCFS string = "CFS"
)

// referencetypes
const (
	// Freight Forwarder’s Reference,Reference assigned to the shipment by the freight forwarder.
	ReferenceTypeFF string = "FF"
	// Shipper’s Reference,Reference assigned to the shipment by the shipper.
	ReferenceTypeSI string = "SI"
	// Purchase Order Reference,The PO reference that the shipper or freight forwarder received from the consignee and then shared with the carrier.
	ReferenceTypePO string = "PO"
	// Customer’s Reference,Reference assigned to the shipment by the customer.
	ReferenceTypeCR string = "CR"
	// Consignee’s Reference,Reference assigned to the shipment by the consignee.
	ReferenceTypeAAO string = "AAO"
	// Empty container release reference,Unique identifier to enable release of the container from a carrier nominated depot
	ReferenceTypeECR string = "ECR"
	// Customer shipment ID,Unique Shipment ID for the booking in the Shipper or Forwarder system. Used to identify the booking along with the Booking party.
	ReferenceTypeCSI string = "CSI"
	// Booking party reference number,A unique identifier provided by a booking party in the booking request.
	ReferenceTypeBPR string = "BPR"
	// Booking Request ID,The associated booking request ID provided by the shipper.
	ReferenceTypeBID string = "BID"
	// Equipment Reference,Reference to the equipment that is associated with document.
	ReferenceTypeEQ string = "EQ"
)

// sealsourcecodes
const (
	// Carrier
	SealSourceCAR string = "CAR"
	// Shipper
	SealSourceSHI string = "SHI"
	// Phytosanitary
	SealSourcePHY string = "PHY"
	// Veterinary
	SealSourceVET string = "VET"
	// Customs
	SealSourceCUS string = "CUS"
)

// sealtypecodes
const (
	// Keyless padlock
	SealTypeKLP string = "KLP"
	// Bolt
	SealTypeBLT string = "BLT"
	// Wire
	SealTypeWIR string = "WIR"
)

// shipmenteventtypecodes
const (
	// Received,Indicates that a document is received by the carrier or shipper
	ShipmentEventTypeRECE string = "RECE"
	// Drafted,Indicates that a document is in draft mode being updated by either the shipper or the carrier.
	ShipmentEventTypeDRFT string = "DRFT"
	// Pending Approval,Indicates that a document has been submitted by the carrier and is now awaiting approval by the shipper.
	ShipmentEventTypePENA string = "PENA"
	// Pending Update,Indicates that the carrier requested an update from the shipper which is not received yet.
	ShipmentEventTypePENU string = "PENU"
	// Pending Confirmation,Indicates that a document has been submitted by the shipper and is now awaiting approval by the carrier.
	ShipmentEventTypePENC string = "PENC"
	// Rejected,Indicates that a document has been rejected by the carrier.
	ShipmentEventTypeREJE string = "REJE"
	// Approved,Indicates that a document has been approved by the counterpart.
	ShipmentEventTypeAPPR string = "APPR"
	// Issued,Indicates that a document has been issued by the carrier.
	ShipmentEventTypeISSU string = "ISSU"
	// Surrendered,Indicates that a document has been surrendered by the customer to the carrier.
	ShipmentEventTypeSURR string = "SURR"
	// Submitted,Indicates that a document has been submitted by the customer to the carrier.
	ShipmentEventTypeSUBM string = "SUBM"
	// Void,Cancellation of an original document.
	ShipmentEventTypeVOID string = "VOID"
	// Confirmed,Indicates that the document is confirmed.
	ShipmentEventTypeCONF string = "CONF"
	// Requested,"A status indicator that can be used with a number of identifiers to denote that a certain activity, service or document has been requested by the carrier, customer or authorities. This status remains constant until the requested activity is  “Completed”."
	ShipmentEventTypeREQS string = "REQS"
	// Completed,"A status indicator that can be used with a number of activity identifiers to denote that a certain activity, service or document has been completed."
	ShipmentEventTypeCMPL string = "CMPL"
	// On Hold,"A status indicator that can be used with a number of activity identifiers to denote that a container or shipment has been placed on hold i.e. can’t  progress in the process."
	ShipmentEventTypeHOLD string = "HOLD"
	// Released,"A status indicator that can be used with a number of activity identifiers to denote that a container or shipment has been released i.e. allowed to move from depot or terminal by authorities or service provider."
	ShipmentEventTypeRELS string = "RELS"
	// Cancelled, "A status indicator to be used when the booking is cancelled by the Shipper"
	ShipmentEventTypeCANC string = "CANC"
)

// shipmentlocationtypes
const (
	// Place of Receipt,"The location where the cargo is handed over by the shipper, or his agent, to the shipping line. This indicates the point at which the shipping line takes on responsibility for carriage of the container."POL,Port of Loading,"The location where the cargo is loaded onto a first sea-going vessel for water transportation."
	ShipmentLocationPRE string = "PRE"
	// Port of Loading,"The location where the cargo is loaded onto a first sea-going vessel for water transportation."
	ShipmentLocationPOL string = "POL"
	// Port of Discharge,"The location where the cargo is discharged from the last sea-going vessel."
	ShipmentLocationPOD string = "POD"
	// Place of Delivery,"The location where the cargo is handed over to the consignee, or his agent, by the shipping line and where responsibility of the shipping line ceases."
	ShipmentLocationPDE string = "PDE"
	// Pre-carriage From,
	ShipmentLocationPCF string = "PCF"
	// Precarriage under shipper’s responsibility,"Place and mode of transportation for pre-carriage (e.g. truck, barge, rail) under shipper's responsibility"
	ShipmentLocationPSR string = "PSR"
	// Onward In-land Routing,"The location where the cargo is transported from port of discharge to consignee location on consignee's responsibility (merchant haulage)."
	ShipmentLocationOIR string = "OIR"
	// Depot release location
	ShipmentLocationDRL string = "DRL"
	// Origin of goods
	ShipmentLocationORI string = "ORI"
	// Container intermediate export stop off location,
	ShipmentLocationIEL string = "IEL"
	// Prohibited transshipment port
	ShipmentLocationPTP string = "PTP"
	// Requested transshipment port
	ShipmentLocationRTP string = "RTP"
	// Full container drop-off location
	ShipmentLocationFCD string = "FCD"
	// Empty container pick-up location
	ShipmentLocationECP string = "ECP"
)

// transportdocumenttypecodes
const (
	// Bill of Lading,"Contractual document issued to the shipper which confirms the carrier's receipt of the cargo, acknowledging goods being shipped or received for shipment and specifying the terms of delivery (as one of the evidences of the contract of carriage). The Bill of Lading is usually prepared based on shipping instructions, including cargo description, given by the shipper on forms issued by the carrier and is the title to the goods and can be a negotiable document."
	TransportDocumentTypeBOL string = "BOL"
	// Sea Waybill,"A separate specific transport document type which is non-negotiable, does not transfer title, but which evidences the contract of carriage and receipt of the goods. It must be issued to a named consignee and can be both in a physical or digital format. Goods can be released at destination without presenting the original sea waybill as proof of ownership."
	TransportDocumentTypeSWB string = "SWB"
)

// transporteventtypecodes
const (
	// Arrival
	TransportEventTypeARRI string = "ARRI"
	// Departure
	TransportEventTypeDEPA string = "DEPA"
)

// transportplanstagetypes
const (
	// Pre-Carriage,Transport leg occurring prior to the main transport leg.
	TransportPlanStageTypePRC string = "PRC"
	// Main Carriage Transport,"The main transport leg(s), happening on one or more main vessels."
	TransportPlanStageTypeMNC string = "MNC"
	// On-Carriage Transport,The transport leg occuring after the main leg to the final destination.
	TransportPlanStageTypeONC string = "ONC"
)

// unitofmeasures
const (
	// Kilogram
	UnitOfMeasuresKGM string = "KGM"
	// Pounds
	UnitOfMeasuresLBR string = "LBR"
	// Cubic Metres
	UnitOfMeasuresCBM string = "CBM"
	// Cubic Feet
	UnitOfMeasuresCFT string = "CFT"
	// Celsius
	UnitOfMeasuresCEL string = "CEL"
	// Fahrenheit
	UnitOfMeasuresFAH string = "FAH"
)

// valueaddedservicecodes
const (
	// Smart containers,Smart containers
	ValueAddedServiceSCON string = "SCON"
	// Cargo insurance,Cargo insurance
	ValueAddedServiceCINS string = "CINS"
	// Smart IoT devices,Smart IoT devices
	ValueAddedServiceSIOT string = "SIOT"
	// Customs declaration,Customs Declaration
	ValueAddedServiceCDECL string = "CDECL"
	// Shipping guarantee,Shipping guarantee
	ValueAddedServiceSGUAR string = "SGUAR"
	// Upfront payment,Upfront payment
	ValueAddedServiceUPPY string = "UPPY"
)

// vesselsharingagreementtypes
const (
	// Vessel Sharing Agreement,An agreement between two or more carriers agreeing to share vessel capacity in specific trades using a specified number of vessels.
	VesselSharingAgreementTypeVSA string = "VSA"
	// Slot Charter Agreement,An agreement between two carriers to sell or exchange a specific number of TEU slots in one or more trades.
	VesselSharingAgreementTypeSCA string = "SCA"
)
