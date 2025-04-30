/*!999999\- enable the sandbox mode */ 
-- MariaDB dump 10.19  Distrib 10.6.18-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: scfdcsadb1
-- ------------------------------------------------------
-- Server version	10.6.18-MariaDB-0ubuntu0.22.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `addresses`
--

DROP TABLE IF EXISTS `addresses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `addresses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `name1` varchar(100) DEFAULT '',
  `street` varchar(100) DEFAULT '',
  `street_number` varchar(50) DEFAULT '',
  `floor1` varchar(50) DEFAULT '',
  `postal_code` varchar(10) DEFAULT '',
  `city` varchar(65) DEFAULT '',
  `state_region` varchar(65) DEFAULT '',
  `country_name` varchar(75) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `addresses`
--

LOCK TABLES `addresses` WRITE;
/*!40000 ALTER TABLE `addresses` DISABLE KEYS */;
/*!40000 ALTER TABLE `addresses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `booking_cores`
--

DROP TABLE IF EXISTS `booking_cores`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `booking_cores` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `receipt_type_at_origin` varchar(3) DEFAULT '',
  `delivery_type_at_destination` varchar(3) DEFAULT '',
  `cargo_movement_type_at_origin` varchar(3) DEFAULT '',
  `cargo_movement_type_at_destination` varchar(3) DEFAULT '',
  `service_contract_reference` varchar(30) DEFAULT '',
  `vessel_name` varchar(50) DEFAULT '',
  `carrier_export_voyage_number` varchar(50) DEFAULT '',
  `universal_export_voyage_reference` varchar(100) DEFAULT '',
  `declared_value` double DEFAULT 0,
  `delivery_value_currency` varchar(3) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `booking_cores`
--

LOCK TABLES `booking_cores` WRITE;
/*!40000 ALTER TABLE `booking_cores` DISABLE KEYS */;
/*!40000 ALTER TABLE `booking_cores` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `booking_ref_statuses`
--

DROP TABLE IF EXISTS `booking_ref_statuses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `booking_ref_statuses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `carrier_booking_request_reference` varchar(100) DEFAULT '',
  `document_status` varchar(4) DEFAULT '',
  `booking_request_created_date_time` datetime DEFAULT current_timestamp(),
  `booking_request_updated_date_time` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `booking_ref_statuses`
--

LOCK TABLES `booking_ref_statuses` WRITE;
/*!40000 ALTER TABLE `booking_ref_statuses` DISABLE KEYS */;
/*!40000 ALTER TABLE `booking_ref_statuses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `booking_requests`
--

DROP TABLE IF EXISTS `booking_requests`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `booking_requests` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `receipt_type_at_origin` varchar(3) DEFAULT '',
  `delivery_type_at_destination` varchar(3) DEFAULT '',
  `cargo_movement_type_at_origin` varchar(3) DEFAULT '',
  `cargo_movement_type_at_destination` varchar(3) DEFAULT '',
  `service_contract_reference` varchar(30) DEFAULT '',
  `vessel_name` varchar(50) DEFAULT '',
  `carrier_export_voyage_number` varchar(50) DEFAULT '',
  `universal_export_voyage_reference` varchar(100) DEFAULT '',
  `declared_value` double DEFAULT 0,
  `delivery_value_currency` varchar(3) DEFAULT '',
  `payment_term_code` varchar(3) DEFAULT '',
  `is_partial_load_allowed` tinyint(1) DEFAULT 0,
  `is_export_declaration_required` tinyint(1) DEFAULT 0,
  `export_declaration_reference` varchar(35) DEFAULT '',
  `is_import_license_required` tinyint(1) DEFAULT 0,
  `import_license_reference` varchar(35) DEFAULT '',
  `is_ams_aci_filing_required` tinyint(1) DEFAULT 0,
  `is_destination_filing_required` tinyint(1) DEFAULT 0,
  `contract_quotation_reference` varchar(35) DEFAULT '',
  `transport_document_type_code` varchar(50) DEFAULT '',
  `transport_document_reference` varchar(100) DEFAULT '',
  `booking_channel_reference` varchar(100) DEFAULT '',
  `inco_terms` varchar(50) DEFAULT '',
  `communication_channel_code` varchar(50) DEFAULT '',
  `is_equipment_substitution_allowed` tinyint(1) DEFAULT 0,
  `vessel_imo_number` varchar(50) DEFAULT '',
  `pre_carriage_mode_of_transport_code` varchar(50) DEFAULT '',
  `submission_date_time` datetime DEFAULT current_timestamp(),
  `expected_departure_date` datetime DEFAULT current_timestamp(),
  `expected_arrival_at_place_of_delivery_start_date` datetime DEFAULT current_timestamp(),
  `expected_arrival_at_place_of_delivery_end_date` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `booking_requests`
--

LOCK TABLES `booking_requests` WRITE;
/*!40000 ALTER TABLE `booking_requests` DISABLE KEYS */;
/*!40000 ALTER TABLE `booking_requests` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `booking_responses`
--

DROP TABLE IF EXISTS `booking_responses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `booking_responses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `carrier_booking_request_reference` varchar(100) DEFAULT '',
  `document_status` varchar(4) DEFAULT '',
  `payment_term_code` varchar(3) DEFAULT '',
  `is_partial_load_allowed` tinyint(1) DEFAULT 0,
  `is_export_declaration_required` tinyint(1) DEFAULT 0,
  `export_declaration_reference` varchar(35) DEFAULT '',
  `is_import_license_required` tinyint(1) DEFAULT 0,
  `import_license_reference` varchar(35) DEFAULT '',
  `is_ams_aci_filing_required` tinyint(1) DEFAULT 0,
  `is_destination_filing_required` tinyint(1) DEFAULT 0,
  `contract_quotation_reference` varchar(35) DEFAULT '',
  `transport_document_type_code` varchar(50) DEFAULT '',
  `transport_document_reference` varchar(100) DEFAULT '',
  `booking_channel_reference` varchar(100) DEFAULT '',
  `inco_terms` varchar(50) DEFAULT '',
  `communication_channel_code` varchar(50) DEFAULT '',
  `is_equipment_substitution_allowed` tinyint(1) DEFAULT 0,
  `vessel_imo_number` varchar(50) DEFAULT '',
  `booking_request_created_date_time` datetime DEFAULT current_timestamp(),
  `booking_request_updated_date_time` datetime DEFAULT current_timestamp(),
  `expected_departure_date` datetime DEFAULT current_timestamp(),
  `expected_arrival_at_place_of_delivery_start_date` datetime DEFAULT current_timestamp(),
  `expected_arrival_at_place_of_delivery_end_date` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `booking_responses`
--

LOCK TABLES `booking_responses` WRITE;
/*!40000 ALTER TABLE `booking_responses` DISABLE KEYS */;
/*!40000 ALTER TABLE `booking_responses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `booking_shallow_cores`
--

DROP TABLE IF EXISTS `booking_shallow_cores`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `booking_shallow_cores` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `receipt_type_at_origin` varchar(3) DEFAULT '',
  `delivery_type_at_destination` varchar(3) DEFAULT '',
  `cargo_movement_type_at_origin` varchar(3) DEFAULT '',
  `cargo_movement_type_at_destination` varchar(3) DEFAULT '',
  `service_contract_reference` varchar(30) DEFAULT '',
  `vessel_name` varchar(50) DEFAULT '',
  `carrier_export_voyage_number` varchar(50) DEFAULT '',
  `universal_export_voyage_reference` varchar(100) DEFAULT '',
  `declared_value` double DEFAULT 0,
  `delivery_value_currency` varchar(3) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `booking_shallow_cores`
--

LOCK TABLES `booking_shallow_cores` WRITE;
/*!40000 ALTER TABLE `booking_shallow_cores` DISABLE KEYS */;
/*!40000 ALTER TABLE `booking_shallow_cores` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `booking_shallows`
--

DROP TABLE IF EXISTS `booking_shallows`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `booking_shallows` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `receipt_type_at_origin` varchar(3) DEFAULT '',
  `delivery_type_at_destination` varchar(3) DEFAULT '',
  `cargo_movement_type_at_origin` varchar(3) DEFAULT '',
  `cargo_movement_type_at_destination` varchar(3) DEFAULT '',
  `service_contract_reference` varchar(30) DEFAULT '',
  `vessel_name` varchar(50) DEFAULT '',
  `carrier_export_voyage_number` varchar(50) DEFAULT '',
  `universal_export_voyage_reference` varchar(100) DEFAULT '',
  `declared_value` double DEFAULT 0,
  `delivery_value_currency` varchar(3) DEFAULT '',
  `payment_term_code` varchar(3) DEFAULT '',
  `is_partial_load_allowed` tinyint(1) DEFAULT 0,
  `is_export_declaration_required` tinyint(1) DEFAULT 0,
  `export_declaration_reference` varchar(35) DEFAULT '',
  `is_import_license_required` tinyint(1) DEFAULT 0,
  `import_license_reference` varchar(35) DEFAULT '',
  `is_ams_aci_filing_required` tinyint(1) DEFAULT 0,
  `is_destination_filing_required` tinyint(1) DEFAULT 0,
  `contract_quotation_reference` varchar(35) DEFAULT '',
  `transport_document_type_code` varchar(50) DEFAULT '',
  `transport_document_reference` varchar(100) DEFAULT '',
  `booking_channel_reference` varchar(100) DEFAULT '',
  `inco_terms` varchar(50) DEFAULT '',
  `communication_channel_code` varchar(50) DEFAULT '',
  `is_equipment_substitution_allowed` tinyint(1) DEFAULT 0,
  `vessel_imo_number` varchar(50) DEFAULT '',
  `pre_carriage_mode_of_transport_code` varchar(50) DEFAULT '',
  `expected_departure_date` datetime DEFAULT current_timestamp(),
  `expected_arrival_at_place_of_delivery_start_date` datetime DEFAULT current_timestamp(),
  `expected_arrival_at_place_of_delivery_end_date` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `booking_shallows`
--

LOCK TABLES `booking_shallows` WRITE;
/*!40000 ALTER TABLE `booking_shallows` DISABLE KEYS */;
/*!40000 ALTER TABLE `booking_shallows` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `booking_summaries`
--

DROP TABLE IF EXISTS `booking_summaries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `booking_summaries` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `carrier_booking_request_reference` varchar(100) DEFAULT '',
  `document_status` varchar(4) DEFAULT '',
  `receipt_type_at_origin` varchar(3) DEFAULT '',
  `delivery_type_at_destination` varchar(3) DEFAULT '',
  `cargo_movement_type_at_origin` varchar(3) DEFAULT '',
  `cargo_movement_type_at_destination` varchar(3) DEFAULT '',
  `service_contract_reference` varchar(30) DEFAULT '',
  `vessel_name` varchar(50) DEFAULT '',
  `carrier_export_voyage_number` varchar(50) DEFAULT '',
  `universal_export_voyage_reference` varchar(100) DEFAULT '',
  `declared_value` double DEFAULT 0,
  `delivery_value_currency` varchar(3) DEFAULT '',
  `payment_term_code` varchar(3) DEFAULT '',
  `is_partial_load_allowed` tinyint(1) DEFAULT 0,
  `is_export_declaration_required` tinyint(1) DEFAULT 0,
  `export_declaration_reference` varchar(35) DEFAULT '',
  `is_import_license_required` tinyint(1) DEFAULT 0,
  `import_license_reference` varchar(35) DEFAULT '',
  `is_ams_aci_filing_required` tinyint(1) DEFAULT 0,
  `is_destination_filing_required` tinyint(1) DEFAULT 0,
  `contract_quotation_reference` varchar(35) DEFAULT '',
  `transport_document_type_code` varchar(50) DEFAULT '',
  `transport_document_reference` varchar(100) DEFAULT '',
  `booking_channel_reference` varchar(100) DEFAULT '',
  `inco_terms` varchar(50) DEFAULT '',
  `communication_channel_code` varchar(50) DEFAULT '',
  `is_equipment_substitution_allowed` tinyint(1) DEFAULT 0,
  `vessel_imo_number` varchar(50) DEFAULT '',
  `pre_carriage_mode_of_transport_code` varchar(50) DEFAULT '',
  `booking_request_created_date_time` datetime DEFAULT current_timestamp(),
  `booking_request_updated_date_time` datetime DEFAULT current_timestamp(),
  `submission_date_time` datetime DEFAULT current_timestamp(),
  `expected_departure_date` datetime DEFAULT current_timestamp(),
  `expected_arrival_at_place_of_delivery_start_date` datetime DEFAULT current_timestamp(),
  `expected_arrival_at_place_of_delivery_end_date` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `booking_summaries`
--

LOCK TABLES `booking_summaries` WRITE;
/*!40000 ALTER TABLE `booking_summaries` DISABLE KEYS */;
/*!40000 ALTER TABLE `booking_summaries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bookings`
--

DROP TABLE IF EXISTS `bookings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bookings` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `carrier_booking_request_reference` varchar(100) DEFAULT '',
  `document_status` varchar(4) DEFAULT '',
  `receipt_type_at_origin` varchar(3) DEFAULT '',
  `delivery_type_at_destination` varchar(3) DEFAULT '',
  `cargo_movement_type_at_origin` varchar(3) DEFAULT '',
  `cargo_movement_type_at_destination` varchar(3) DEFAULT '',
  `service_contract_reference` varchar(30) DEFAULT '',
  `payment_term_code` varchar(3) DEFAULT '',
  `is_partial_load_allowed` tinyint(1) DEFAULT 0,
  `is_export_declaration_required` tinyint(1) DEFAULT 0,
  `export_declaration_reference` varchar(35) DEFAULT '',
  `is_import_license_required` tinyint(1) DEFAULT 0,
  `import_license_reference` varchar(35) DEFAULT '',
  `is_ams_aci_filing_required` tinyint(1) DEFAULT 0,
  `is_destination_filing_required` tinyint(1) DEFAULT 0,
  `contract_quotation_reference` varchar(35) DEFAULT '',
  `transport_document_type_code` varchar(50) DEFAULT '',
  `transport_document_reference` varchar(100) DEFAULT '',
  `booking_channel_reference` varchar(100) DEFAULT '',
  `inco_terms` varchar(50) DEFAULT '',
  `communication_channel_code` varchar(50) DEFAULT '',
  `is_equipment_substitution_allowed` tinyint(1) DEFAULT 0,
  `vessel_name` varchar(50) DEFAULT '',
  `vessel_imo_number` varchar(50) DEFAULT '',
  `export_voyage_number` varchar(50) DEFAULT '',
  `pre_carriage_mode_of_transport_code` varchar(50) DEFAULT '',
  `vessel_id` int(10) unsigned DEFAULT 0,
  `declared_value_currency_code` varchar(3) DEFAULT '',
  `declared_value` double DEFAULT 0,
  `voyage_id` int(10) unsigned DEFAULT 0,
  `location_id` int(10) unsigned DEFAULT 0,
  `invoice_payable_at` varchar(50) DEFAULT '',
  `submission_date_time` datetime DEFAULT current_timestamp(),
  `expected_departure_date` datetime DEFAULT current_timestamp(),
  `expected_arrival_at_place_of_delivery_start_date` datetime DEFAULT current_timestamp(),
  `expected_arrival_at_place_of_delivery_end_date` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bookings`
--

LOCK TABLES `bookings` WRITE;
/*!40000 ALTER TABLE `bookings` DISABLE KEYS */;
/*!40000 ALTER TABLE `bookings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cargo_items`
--

DROP TABLE IF EXISTS `cargo_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cargo_items` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `consignment_item_id` int(10) unsigned DEFAULT 0,
  `weight` double DEFAULT 0,
  `volume` double DEFAULT 0,
  `weight_unit` varchar(10) DEFAULT '',
  `volume_unit` varchar(10) DEFAULT '',
  `number_of_packages` int(10) unsigned DEFAULT 0,
  `package_code` varchar(3) DEFAULT '',
  `utilized_transport_equipment_id` int(10) unsigned DEFAULT 0,
  `package_name_on_bl` varchar(50) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cargo_items`
--

LOCK TABLES `cargo_items` WRITE;
/*!40000 ALTER TABLE `cargo_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `cargo_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cargo_line_items`
--

DROP TABLE IF EXISTS `cargo_line_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cargo_line_items` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `cargo_item_id` int(10) unsigned DEFAULT 0,
  `shipping_marks` varchar(100) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cargo_line_items`
--

LOCK TABLES `cargo_line_items` WRITE;
/*!40000 ALTER TABLE `cargo_line_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `cargo_line_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cargo_movement_types`
--

DROP TABLE IF EXISTS `cargo_movement_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cargo_movement_types` (
  `cargo_movement_type_code` varchar(3) DEFAULT '',
  `cargo_movement_type_name` varchar(50) DEFAULT '',
  `cargo_movement_type_description` varchar(300) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cargo_movement_types`
--

LOCK TABLES `cargo_movement_types` WRITE;
/*!40000 ALTER TABLE `cargo_movement_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `cargo_movement_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `carrier_clauses`
--

DROP TABLE IF EXISTS `carrier_clauses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `carrier_clauses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `clause_content` varchar(300) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `carrier_clauses`
--

LOCK TABLES `carrier_clauses` WRITE;
/*!40000 ALTER TABLE `carrier_clauses` DISABLE KEYS */;
/*!40000 ALTER TABLE `carrier_clauses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `carriers`
--

DROP TABLE IF EXISTS `carriers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `carriers` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `carrier_name` varchar(100) DEFAULT '',
  `smdg_code` varchar(3) DEFAULT '',
  `nmfta_code` varchar(4) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `carriers`
--

LOCK TABLES `carriers` WRITE;
/*!40000 ALTER TABLE `carriers` DISABLE KEYS */;
/*!40000 ALTER TABLE `carriers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `charges`
--

DROP TABLE IF EXISTS `charges`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `charges` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `transport_document_id` int(10) unsigned DEFAULT 0,
  `shipment_id` int(10) unsigned DEFAULT 0,
  `charge_type` varchar(50) DEFAULT '',
  `currency_amount` double DEFAULT 0,
  `currency_code` varchar(50) DEFAULT '',
  `payment_term_code` varchar(50) DEFAULT '',
  `calculation_basis` varchar(50) DEFAULT '',
  `unit_price` double DEFAULT 0,
  `quantity` double DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `charges`
--

LOCK TABLES `charges` WRITE;
/*!40000 ALTER TABLE `charges` DISABLE KEYS */;
/*!40000 ALTER TABLE `charges` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `code_list_responsible_agencies`
--

DROP TABLE IF EXISTS `code_list_responsible_agencies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `code_list_responsible_agencies` (
  `dcsa_responsible_agency_code` varchar(5) DEFAULT '',
  `code_list_responsible_agency_code` varchar(3) DEFAULT '',
  `code_list_responsible_agency_name` varchar(100) DEFAULT '',
  `code_list_responsible_agency_description` varchar(300) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `code_list_responsible_agencies`
--

LOCK TABLES `code_list_responsible_agencies` WRITE;
/*!40000 ALTER TABLE `code_list_responsible_agencies` DISABLE KEYS */;
/*!40000 ALTER TABLE `code_list_responsible_agencies` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `commodities`
--

DROP TABLE IF EXISTS `commodities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `commodities` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `booking_id` int(10) unsigned DEFAULT 0,
  `commodity_type` varchar(550) DEFAULT '',
  `hs_code` varchar(50) DEFAULT '',
  `cargo_gross_weight` double DEFAULT 0,
  `cargo_gross_weight_unit` varchar(3) DEFAULT '',
  `cargo_gross_volume` double DEFAULT 0,
  `cargo_gross_volume_unit` varchar(3) DEFAULT '',
  `number_of_packages` int(10) unsigned DEFAULT 0,
  `export_license_issue_date` datetime DEFAULT current_timestamp(),
  `export_license_expiry_date` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `commodities`
--

LOCK TABLES `commodities` WRITE;
/*!40000 ALTER TABLE `commodities` DISABLE KEYS */;
/*!40000 ALTER TABLE `commodities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `communication_channel_qualifiers`
--

DROP TABLE IF EXISTS `communication_channel_qualifiers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `communication_channel_qualifiers` (
  `communication_channel_qualifier_code` varchar(2) DEFAULT '',
  `communication_channel_qualifier_name` varchar(100) DEFAULT '',
  `communication_channel_qualifier_description` varchar(250) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `communication_channel_qualifiers`
--

LOCK TABLES `communication_channel_qualifiers` WRITE;
/*!40000 ALTER TABLE `communication_channel_qualifiers` DISABLE KEYS */;
/*!40000 ALTER TABLE `communication_channel_qualifiers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `confirmed_equipments`
--

DROP TABLE IF EXISTS `confirmed_equipments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `confirmed_equipments` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `confirmed_equipment_sizetype` varchar(4) DEFAULT '',
  `confirmed_equipment_units` int(11) DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `confirmed_equipments`
--

LOCK TABLES `confirmed_equipments` WRITE;
/*!40000 ALTER TABLE `confirmed_equipments` DISABLE KEYS */;
/*!40000 ALTER TABLE `confirmed_equipments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `consignment_items`
--

DROP TABLE IF EXISTS `consignment_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `consignment_items` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `description_of_goods` varchar(100) DEFAULT '',
  `hs_code` varchar(10) DEFAULT '',
  `shipping_instruction_id` int(10) unsigned DEFAULT 0,
  `weight` double DEFAULT 0,
  `volume` double DEFAULT 0,
  `weight_unit` varchar(10) DEFAULT '',
  `volume_unit` varchar(10) DEFAULT '',
  `shipment_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `consignment_items`
--

LOCK TABLES `consignment_items` WRITE;
/*!40000 ALTER TABLE `consignment_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `consignment_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `countries`
--

DROP TABLE IF EXISTS `countries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `countries` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `country_code` varchar(2) DEFAULT '',
  `country_name` varchar(75) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `countries`
--

LOCK TABLES `countries` WRITE;
/*!40000 ALTER TABLE `countries` DISABLE KEYS */;
/*!40000 ALTER TABLE `countries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customs_reference_types`
--

DROP TABLE IF EXISTS `customs_reference_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customs_reference_types` (
  `customs_reference_type_code` varchar(50) DEFAULT '',
  `customs_reference_country_code` varchar(2) DEFAULT '',
  `customs_reference_type_name` varchar(50) DEFAULT '',
  `customs_reference_type_description` varchar(400) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customs_reference_types`
--

LOCK TABLES `customs_reference_types` WRITE;
/*!40000 ALTER TABLE `customs_reference_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `customs_reference_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cut_off_times`
--

DROP TABLE IF EXISTS `cut_off_times`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cut_off_times` (
  `cut_off_time_code` varchar(3) DEFAULT '',
  `cut_off_time_name` varchar(100) DEFAULT '',
  `cut_off_time_description` varchar(250) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cut_off_times`
--

LOCK TABLES `cut_off_times` WRITE;
/*!40000 ALTER TABLE `cut_off_times` DISABLE KEYS */;
/*!40000 ALTER TABLE `cut_off_times` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `displayed_addresses`
--

DROP TABLE IF EXISTS `displayed_addresses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `displayed_addresses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `document_party_id` int(10) unsigned DEFAULT 0,
  `address_line_number` int(10) unsigned DEFAULT 0,
  `address_line_text` varchar(250) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `displayed_addresses`
--

LOCK TABLES `displayed_addresses` WRITE;
/*!40000 ALTER TABLE `displayed_addresses` DISABLE KEYS */;
/*!40000 ALTER TABLE `displayed_addresses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `document_parties`
--

DROP TABLE IF EXISTS `document_parties`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `document_parties` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `party_id` int(10) unsigned DEFAULT 0,
  `shipping_instruction_id` int(10) unsigned DEFAULT 0,
  `shipment_id` int(10) unsigned DEFAULT 0,
  `party_function` varchar(3) DEFAULT '',
  `is_to_be_notified` tinyint(1) DEFAULT 0,
  `booking_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `document_parties`
--

LOCK TABLES `document_parties` WRITE;
/*!40000 ALTER TABLE `document_parties` DISABLE KEYS */;
/*!40000 ALTER TABLE `document_parties` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `document_types`
--

DROP TABLE IF EXISTS `document_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `document_types` (
  `document_type_code` varchar(3) DEFAULT '',
  `document_type_name` varchar(100) DEFAULT '',
  `document_type_description` varchar(250) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `document_types`
--

LOCK TABLES `document_types` WRITE;
/*!40000 ALTER TABLE `document_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `document_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ebl_solution_provider_types`
--

DROP TABLE IF EXISTS `ebl_solution_provider_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ebl_solution_provider_types` (
  `ebl_solution_provider_name` varchar(50) DEFAULT '',
  `ebl_solution_provider_code` varchar(5) DEFAULT '',
  `ebl_solution_provider_url` varchar(100) DEFAULT '',
  `ebl_solution_provider_description` varchar(250) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ebl_solution_provider_types`
--

LOCK TABLES `ebl_solution_provider_types` WRITE;
/*!40000 ALTER TABLE `ebl_solution_provider_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `ebl_solution_provider_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ebl_visualizations`
--

DROP TABLE IF EXISTS `ebl_visualizations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ebl_visualizations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `name` varchar(100) DEFAULT '',
  `content` varchar(100) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ebl_visualizations`
--

LOCK TABLES `ebl_visualizations` WRITE;
/*!40000 ALTER TABLE `ebl_visualizations` DISABLE KEYS */;
/*!40000 ALTER TABLE `ebl_visualizations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `empty_indicators`
--

DROP TABLE IF EXISTS `empty_indicators`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `empty_indicators` (
  `empty_indicator_code` varchar(5) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `empty_indicators`
--

LOCK TABLES `empty_indicators` WRITE;
/*!40000 ALTER TABLE `empty_indicators` DISABLE KEYS */;
/*!40000 ALTER TABLE `empty_indicators` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `endorsement_chain_links`
--

DROP TABLE IF EXISTS `endorsement_chain_links`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `endorsement_chain_links` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `entry_order` int(11) DEFAULT 0,
  `action_date_time` datetime DEFAULT current_timestamp(),
  `actor` int(10) unsigned DEFAULT 0,
  `recipient` int(10) unsigned DEFAULT 0,
  `surrender_request_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `endorsement_chain_links`
--

LOCK TABLES `endorsement_chain_links` WRITE;
/*!40000 ALTER TABLE `endorsement_chain_links` DISABLE KEYS */;
/*!40000 ALTER TABLE `endorsement_chain_links` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `equipment`
--

DROP TABLE IF EXISTS `equipment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `equipment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `equipment_reference` varchar(15) DEFAULT '',
  `iso_equipment_code` varchar(4) DEFAULT '',
  `tare_weight` double DEFAULT 0,
  `weight_unit` varchar(3) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `equipment`
--

LOCK TABLES `equipment` WRITE;
/*!40000 ALTER TABLE `equipment` DISABLE KEYS */;
/*!40000 ALTER TABLE `equipment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `equipment_event_types`
--

DROP TABLE IF EXISTS `equipment_event_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `equipment_event_types` (
  `equipment_event_type_code` varchar(4) DEFAULT '',
  `equipment_event_type_name` varchar(30) DEFAULT '',
  `equipment_event_type_description` varchar(300) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `equipment_event_types`
--

LOCK TABLES `equipment_event_types` WRITE;
/*!40000 ALTER TABLE `equipment_event_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `equipment_event_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `equipment_events`
--

DROP TABLE IF EXISTS `equipment_events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `equipment_events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `event_id_s` binary(16) DEFAULT NULL,
  `event_classifier_code` varchar(3) DEFAULT '',
  `equipment_event_type_code` varchar(4) DEFAULT '',
  `equipment_reference` varchar(15) DEFAULT '',
  `empty_indicator_code` varchar(5) DEFAULT '',
  `transport_call_id` int(10) unsigned DEFAULT 0,
  `event_location` varchar(100) DEFAULT '',
  `event_created_date_time` datetime DEFAULT current_timestamp(),
  `event_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `equipment_events`
--

LOCK TABLES `equipment_events` WRITE;
/*!40000 ALTER TABLE `equipment_events` DISABLE KEYS */;
/*!40000 ALTER TABLE `equipment_events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `event_cache_queue_deads`
--

DROP TABLE IF EXISTS `event_cache_queue_deads`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `event_cache_queue_deads` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `event_id` binary(16) DEFAULT NULL,
  `event_type` varchar(16) DEFAULT '',
  `failure_reason_type` varchar(200) DEFAULT '',
  `failure_reason_message` varchar(200) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `event_cache_queue_deads`
--

LOCK TABLES `event_cache_queue_deads` WRITE;
/*!40000 ALTER TABLE `event_cache_queue_deads` DISABLE KEYS */;
/*!40000 ALTER TABLE `event_cache_queue_deads` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `event_cache_queues`
--

DROP TABLE IF EXISTS `event_cache_queues`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `event_cache_queues` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `event_id` binary(16) DEFAULT NULL,
  `event_type` varchar(16) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `event_cache_queues`
--

LOCK TABLES `event_cache_queues` WRITE;
/*!40000 ALTER TABLE `event_cache_queues` DISABLE KEYS */;
/*!40000 ALTER TABLE `event_cache_queues` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `event_caches`
--

DROP TABLE IF EXISTS `event_caches`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `event_caches` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `event_id` binary(16) DEFAULT NULL,
  `event_type` varchar(16) DEFAULT '',
  `content` varchar(200) DEFAULT '',
  `document_references` varchar(200) DEFAULT '',
  `event_created_date_time` datetime DEFAULT current_timestamp(),
  `event_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `event_caches`
--

LOCK TABLES `event_caches` WRITE;
/*!40000 ALTER TABLE `event_caches` DISABLE KEYS */;
/*!40000 ALTER TABLE `event_caches` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `event_classifiers`
--

DROP TABLE IF EXISTS `event_classifiers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `event_classifiers` (
  `event_classifier_code` varchar(3) DEFAULT '',
  `event_classifier_name` varchar(30) DEFAULT '',
  `event_classifier_description` varchar(350) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `event_classifiers`
--

LOCK TABLES `event_classifiers` WRITE;
/*!40000 ALTER TABLE `event_classifiers` DISABLE KEYS */;
/*!40000 ALTER TABLE `event_classifiers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `event_subscription_document_type_codes`
--

DROP TABLE IF EXISTS `event_subscription_document_type_codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `event_subscription_document_type_codes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `subscription_id` int(10) unsigned DEFAULT 0,
  `document_type_code` varchar(4) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `event_subscription_document_type_codes`
--

LOCK TABLES `event_subscription_document_type_codes` WRITE;
/*!40000 ALTER TABLE `event_subscription_document_type_codes` DISABLE KEYS */;
/*!40000 ALTER TABLE `event_subscription_document_type_codes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `event_subscription_equipment_event_type_codes`
--

DROP TABLE IF EXISTS `event_subscription_equipment_event_type_codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `event_subscription_equipment_event_type_codes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `subscription_id` int(10) unsigned DEFAULT 0,
  `equipment_event_type_code` varchar(4) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `event_subscription_equipment_event_type_codes`
--

LOCK TABLES `event_subscription_equipment_event_type_codes` WRITE;
/*!40000 ALTER TABLE `event_subscription_equipment_event_type_codes` DISABLE KEYS */;
/*!40000 ALTER TABLE `event_subscription_equipment_event_type_codes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `event_subscription_event_types`
--

DROP TABLE IF EXISTS `event_subscription_event_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `event_subscription_event_types` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `subscription_id` binary(16) DEFAULT NULL,
  `event_type` varchar(40) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `event_subscription_event_types`
--

LOCK TABLES `event_subscription_event_types` WRITE;
/*!40000 ALTER TABLE `event_subscription_event_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `event_subscription_event_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `event_subscription_shipment_event_type_codes`
--

DROP TABLE IF EXISTS `event_subscription_shipment_event_type_codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `event_subscription_shipment_event_type_codes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `subscription_id` int(10) unsigned DEFAULT 0,
  `shipment_event_type_code` varchar(4) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `event_subscription_shipment_event_type_codes`
--

LOCK TABLES `event_subscription_shipment_event_type_codes` WRITE;
/*!40000 ALTER TABLE `event_subscription_shipment_event_type_codes` DISABLE KEYS */;
/*!40000 ALTER TABLE `event_subscription_shipment_event_type_codes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `event_subscription_transport_event_type_codes`
--

DROP TABLE IF EXISTS `event_subscription_transport_event_type_codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `event_subscription_transport_event_type_codes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `subscription_id` int(10) unsigned DEFAULT 0,
  `transport_event_type_code` varchar(4) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `event_subscription_transport_event_type_codes`
--

LOCK TABLES `event_subscription_transport_event_type_codes` WRITE;
/*!40000 ALTER TABLE `event_subscription_transport_event_type_codes` DISABLE KEYS */;
/*!40000 ALTER TABLE `event_subscription_transport_event_type_codes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `event_subscriptions`
--

DROP TABLE IF EXISTS `event_subscriptions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `event_subscriptions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `subscription_id` binary(16) DEFAULT NULL,
  `callback_url` varchar(100) DEFAULT '',
  `document_reference` varchar(100) DEFAULT '',
  `equipment_reference` varchar(100) DEFAULT '',
  `transport_call_reference` varchar(100) DEFAULT '',
  `vessel_imo_number` varchar(7) DEFAULT '',
  `carrier_export_voyage_number` varchar(50) DEFAULT '',
  `universal_export_voyage_reference` varchar(100) DEFAULT '',
  `carrier_service_code` varchar(11) DEFAULT '',
  `universal_service_reference` varchar(100) DEFAULT '',
  `un_location_code` varchar(5) DEFAULT '',
  `secret` varchar(100) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `event_subscriptions`
--

LOCK TABLES `event_subscriptions` WRITE;
/*!40000 ALTER TABLE `event_subscriptions` DISABLE KEYS */;
/*!40000 ALTER TABLE `event_subscriptions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `facilities`
--

DROP TABLE IF EXISTS `facilities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `facilities` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `facility_name` varchar(100) DEFAULT '',
  `un_location_code` varchar(5) DEFAULT '',
  `facility_bic_code` varchar(4) DEFAULT '',
  `facility_smdg_code` varchar(6) DEFAULT '',
  `location_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `facilities`
--

LOCK TABLES `facilities` WRITE;
/*!40000 ALTER TABLE `facilities` DISABLE KEYS */;
/*!40000 ALTER TABLE `facilities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `facility_types`
--

DROP TABLE IF EXISTS `facility_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `facility_types` (
  `facility_type_code` varchar(4) DEFAULT '',
  `facility_type_name` varchar(100) DEFAULT '',
  `facility_type_description` varchar(250) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `facility_types`
--

LOCK TABLES `facility_types` WRITE;
/*!40000 ALTER TABLE `facility_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `facility_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `general_reference_types`
--

DROP TABLE IF EXISTS `general_reference_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `general_reference_types` (
  `general_reference_type_code` varchar(3) DEFAULT '',
  `general_reference_type_name` varchar(100) DEFAULT '',
  `general_reference_type_description` varchar(400) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `general_reference_types`
--

LOCK TABLES `general_reference_types` WRITE;
/*!40000 ALTER TABLE `general_reference_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `general_reference_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `hs_codes`
--

DROP TABLE IF EXISTS `hs_codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `hs_codes` (
  `hs_code` varchar(10) DEFAULT '',
  `hs_code_description` varchar(250) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `hs_codes`
--

LOCK TABLES `hs_codes` WRITE;
/*!40000 ALTER TABLE `hs_codes` DISABLE KEYS */;
/*!40000 ALTER TABLE `hs_codes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `identifying_codes`
--

DROP TABLE IF EXISTS `identifying_codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `identifying_codes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `dcsa_responsible_agency_code` varchar(4) DEFAULT '',
  `party_code` varchar(100) DEFAULT '',
  `code_list_name` varchar(100) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `identifying_codes`
--

LOCK TABLES `identifying_codes` WRITE;
/*!40000 ALTER TABLE `identifying_codes` DISABLE KEYS */;
/*!40000 ALTER TABLE `identifying_codes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `incoterms`
--

DROP TABLE IF EXISTS `incoterms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `incoterms` (
  `incoterms_code` varchar(3) DEFAULT '',
  `incoterms_name` varchar(100) DEFAULT '',
  `incoterms_description` varchar(250) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `incoterms`
--

LOCK TABLES `incoterms` WRITE;
/*!40000 ALTER TABLE `incoterms` DISABLE KEYS */;
/*!40000 ALTER TABLE `incoterms` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `iso_equipment_codes`
--

DROP TABLE IF EXISTS `iso_equipment_codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `iso_equipment_codes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `iso_equipment_code` varchar(4) DEFAULT '',
  `iso_equipment_name` varchar(35) DEFAULT '',
  `iso_equipment_size_code` varchar(2) DEFAULT '',
  `iso_equipment_type_code_a` varchar(2) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `iso_equipment_codes`
--

LOCK TABLES `iso_equipment_codes` WRITE;
/*!40000 ALTER TABLE `iso_equipment_codes` DISABLE KEYS */;
/*!40000 ALTER TABLE `iso_equipment_codes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `issuance_request_responses`
--

DROP TABLE IF EXISTS `issuance_request_responses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `issuance_request_responses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `transport_document_reference` varchar(20) DEFAULT '',
  `issuance_response_code` varchar(4) DEFAULT '',
  `reason` varchar(100) DEFAULT '',
  `created_date_time` datetime DEFAULT current_timestamp(),
  `issuance_request_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_issuance_request_responses_issuance_response_code` (`issuance_response_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `issuance_request_responses`
--

LOCK TABLES `issuance_request_responses` WRITE;
/*!40000 ALTER TABLE `issuance_request_responses` DISABLE KEYS */;
/*!40000 ALTER TABLE `issuance_request_responses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `issuance_requests`
--

DROP TABLE IF EXISTS `issuance_requests`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `issuance_requests` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `transport_document_reference` varchar(20) DEFAULT '',
  `issuance_request_state` varchar(4) DEFAULT '',
  `issue_to` int(10) unsigned DEFAULT 0,
  `ebl_visualization_id` int(10) unsigned DEFAULT 0,
  `transport_document_json` varchar(100) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_issuance_requests_issuance_request_state` (`issuance_request_state`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `issuance_requests`
--

LOCK TABLES `issuance_requests` WRITE;
/*!40000 ALTER TABLE `issuance_requests` DISABLE KEYS */;
/*!40000 ALTER TABLE `issuance_requests` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `issue_parties`
--

DROP TABLE IF EXISTS `issue_parties`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `issue_parties` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `ebl_platform_identifier` varchar(100) DEFAULT '',
  `legal_name` varchar(100) DEFAULT '',
  `registration_number` varchar(100) DEFAULT '',
  `location_of_registration` varchar(2) DEFAULT '',
  `tax_reference` varchar(100) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `issue_parties`
--

LOCK TABLES `issue_parties` WRITE;
/*!40000 ALTER TABLE `issue_parties` DISABLE KEYS */;
/*!40000 ALTER TABLE `issue_parties` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `issue_party_supporting_codes`
--

DROP TABLE IF EXISTS `issue_party_supporting_codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `issue_party_supporting_codes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `issue_party_id` int(10) unsigned DEFAULT 0,
  `party_code` varchar(100) DEFAULT '',
  `party_code_list_provider` varchar(4) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `issue_party_supporting_codes`
--

LOCK TABLES `issue_party_supporting_codes` WRITE;
/*!40000 ALTER TABLE `issue_party_supporting_codes` DISABLE KEYS */;
/*!40000 ALTER TABLE `issue_party_supporting_codes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `legs`
--

DROP TABLE IF EXISTS `legs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `legs` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `sequence_number` int(11) DEFAULT 0,
  `mode_of_transport` varchar(50) DEFAULT '',
  `vessel_operator_smdg_liner_code` varchar(50) DEFAULT '',
  `vessel_imo_number` varchar(50) DEFAULT '',
  `vessel_name` varchar(50) DEFAULT '',
  `carrier_service_name` varchar(50) DEFAULT '',
  `universal_service_reference` varchar(50) DEFAULT '',
  `carrier_service_code` varchar(50) DEFAULT '',
  `universal_import_voyage_reference` varchar(50) DEFAULT '',
  `universal_export_voyage_reference` varchar(50) DEFAULT '',
  `carrier_import_voyage_number` varchar(50) DEFAULT '',
  `carrier_export_voyage_number` varchar(50) DEFAULT '',
  `departure_id` int(10) unsigned DEFAULT 0,
  `arrival_id` int(10) unsigned DEFAULT 0,
  `point_to_point_routing_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `legs`
--

LOCK TABLES `legs` WRITE;
/*!40000 ALTER TABLE `legs` DISABLE KEYS */;
/*!40000 ALTER TABLE `legs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `locations`
--

DROP TABLE IF EXISTS `locations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `locations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `location_name` varchar(100) DEFAULT '',
  `latitude` varchar(10) DEFAULT '',
  `longitude` varchar(11) DEFAULT '',
  `un_location_code` varchar(5) DEFAULT '',
  `address_id` int(10) unsigned DEFAULT 0,
  `facility_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `locations`
--

LOCK TABLES `locations` WRITE;
/*!40000 ALTER TABLE `locations` DISABLE KEYS */;
/*!40000 ALTER TABLE `locations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message_routing_rules`
--

DROP TABLE IF EXISTS `message_routing_rules`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `message_routing_rules` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `api_url` varchar(255) DEFAULT '',
  `login_type` varchar(8) DEFAULT '',
  `login_information` varchar(80) DEFAULT '',
  `vessel_imo_number` varchar(255) DEFAULT '',
  `publisher_role` varchar(3) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message_routing_rules`
--

LOCK TABLES `message_routing_rules` WRITE;
/*!40000 ALTER TABLE `message_routing_rules` DISABLE KEYS */;
/*!40000 ALTER TABLE `message_routing_rules` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mode_of_transports`
--

DROP TABLE IF EXISTS `mode_of_transports`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mode_of_transports` (
  `mode_of_transport_code` varchar(3) DEFAULT '',
  `mode_of_transport_name` varchar(100) DEFAULT '',
  `mode_of_transport_description` varchar(250) DEFAULT '',
  `dcsa_transport_type` varchar(50) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mode_of_transports`
--

LOCK TABLES `mode_of_transports` WRITE;
/*!40000 ALTER TABLE `mode_of_transports` DISABLE KEYS */;
/*!40000 ALTER TABLE `mode_of_transports` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `negotiation_cycles`
--

DROP TABLE IF EXISTS `negotiation_cycles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `negotiation_cycles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cycle_key` varchar(80) DEFAULT '',
  `cycle_name` varchar(80) DEFAULT '',
  `display_order` int(11) DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `negotiation_cycles`
--

LOCK TABLES `negotiation_cycles` WRITE;
/*!40000 ALTER TABLE `negotiation_cycles` DISABLE KEYS */;
/*!40000 ALTER TABLE `negotiation_cycles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `notification_end_points`
--

DROP TABLE IF EXISTS `notification_end_points`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `notification_end_points` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `end_point_id` binary(16) DEFAULT NULL,
  `subscription_id` varchar(100) DEFAULT '',
  `secret` varbinary(255) DEFAULT NULL,
  `endpoint_reference` varchar(100) DEFAULT '',
  `managed_endpoint` tinyint(1) DEFAULT 0,
  `subscription_url` varchar(500) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `notification_end_points`
--

LOCK TABLES `notification_end_points` WRITE;
/*!40000 ALTER TABLE `notification_end_points` DISABLE KEYS */;
/*!40000 ALTER TABLE `notification_end_points` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `operations_event_types`
--

DROP TABLE IF EXISTS `operations_event_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `operations_event_types` (
  `operations_event_type_code` varchar(4) DEFAULT '',
  `operations_event_type_name` varchar(30) DEFAULT '',
  `operations_event_type_description` varchar(250) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `operations_event_types`
--

LOCK TABLES `operations_event_types` WRITE;
/*!40000 ALTER TABLE `operations_event_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `operations_event_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `operations_events`
--

DROP TABLE IF EXISTS `operations_events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `operations_events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `event_id_s` binary(16) DEFAULT NULL,
  `event_classifier_code` varchar(3) DEFAULT '',
  `publisher` varchar(100) DEFAULT '',
  `publisher_role` varchar(3) DEFAULT '',
  `operations_event_type_code` varchar(4) DEFAULT '',
  `event_location` varchar(100) DEFAULT '',
  `transport_call_id` int(10) unsigned DEFAULT 0,
  `port_call_service_type_code` varchar(4) DEFAULT '',
  `facility_type_code` varchar(4) DEFAULT '',
  `delay_reason_code` varchar(4) DEFAULT '',
  `vessel_position` varchar(100) DEFAULT '',
  `remark` varchar(500) DEFAULT '',
  `port_call_phase_type_code` varchar(4) DEFAULT '',
  `vessel_draft` double DEFAULT 0,
  `vessel_draft_unit` varchar(3) DEFAULT '',
  `miles_remaining_to_destination` double DEFAULT 0,
  `event_created_date_time` datetime DEFAULT current_timestamp(),
  `event_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `operations_events`
--

LOCK TABLES `operations_events` WRITE;
/*!40000 ALTER TABLE `operations_events` DISABLE KEYS */;
/*!40000 ALTER TABLE `operations_events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ops_event_timestamp_definitions`
--

DROP TABLE IF EXISTS `ops_event_timestamp_definitions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ops_event_timestamp_definitions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `event_id` binary(16) DEFAULT NULL,
  `timestamp_definition` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ops_event_timestamp_definitions`
--

LOCK TABLES `ops_event_timestamp_definitions` WRITE;
/*!40000 ALTER TABLE `ops_event_timestamp_definitions` DISABLE KEYS */;
/*!40000 ALTER TABLE `ops_event_timestamp_definitions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `outbox_messages`
--

DROP TABLE IF EXISTS `outbox_messages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `outbox_messages` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `message_routing_rule_id` int(10) unsigned DEFAULT 0,
  `payload` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `outbox_messages`
--

LOCK TABLES `outbox_messages` WRITE;
/*!40000 ALTER TABLE `outbox_messages` DISABLE KEYS */;
/*!40000 ALTER TABLE `outbox_messages` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `outgoing_event_queue_deads`
--

DROP TABLE IF EXISTS `outgoing_event_queue_deads`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `outgoing_event_queue_deads` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `event_id` binary(16) DEFAULT NULL,
  `delivery_id` int(10) unsigned DEFAULT 0,
  `subscription_id` int(10) unsigned DEFAULT 0,
  `payload` varchar(80) DEFAULT '',
  `failure_reason_type` varchar(200) DEFAULT '',
  `failure_reason_message` varchar(200) DEFAULT '',
  `enqueued_at_date_time` datetime DEFAULT current_timestamp(),
  `last_failed_at_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `outgoing_event_queue_deads`
--

LOCK TABLES `outgoing_event_queue_deads` WRITE;
/*!40000 ALTER TABLE `outgoing_event_queue_deads` DISABLE KEYS */;
/*!40000 ALTER TABLE `outgoing_event_queue_deads` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `outgoing_event_queues`
--

DROP TABLE IF EXISTS `outgoing_event_queues`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `outgoing_event_queues` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `event_id` binary(16) DEFAULT NULL,
  `delivery_id` int(10) unsigned DEFAULT 0,
  `subscription_id` int(10) unsigned DEFAULT 0,
  `payload` varchar(80) DEFAULT '',
  `enqueued_at_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `outgoing_event_queues`
--

LOCK TABLES `outgoing_event_queues` WRITE;
/*!40000 ALTER TABLE `outgoing_event_queues` DISABLE KEYS */;
/*!40000 ALTER TABLE `outgoing_event_queues` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `package_codes`
--

DROP TABLE IF EXISTS `package_codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `package_codes` (
  `package_code` varchar(3) DEFAULT '',
  `package_code_description` varchar(50) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `package_codes`
--

LOCK TABLES `package_codes` WRITE;
/*!40000 ALTER TABLE `package_codes` DISABLE KEYS */;
/*!40000 ALTER TABLE `package_codes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `parties`
--

DROP TABLE IF EXISTS `parties`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `parties` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `party_name` varchar(100) DEFAULT '',
  `tax_reference1` varchar(20) DEFAULT '',
  `tax_reference2` varchar(20) DEFAULT '',
  `public_key` varchar(500) DEFAULT '',
  `address_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `parties`
--

LOCK TABLES `parties` WRITE;
/*!40000 ALTER TABLE `parties` DISABLE KEYS */;
/*!40000 ALTER TABLE `parties` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `party_contact_details`
--

DROP TABLE IF EXISTS `party_contact_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `party_contact_details` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `party_id` int(10) unsigned DEFAULT 0,
  `name` varchar(255) NOT NULL,
  `email` varchar(100) DEFAULT '',
  `phone` varchar(30) DEFAULT '',
  `url` varchar(100) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `party_contact_details`
--

LOCK TABLES `party_contact_details` WRITE;
/*!40000 ALTER TABLE `party_contact_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `party_contact_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `party_functions`
--

DROP TABLE IF EXISTS `party_functions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `party_functions` (
  `party_function_code` varchar(3) DEFAULT '',
  `party_function_name` varchar(100) DEFAULT '',
  `party_function_description` varchar(350) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `party_functions`
--

LOCK TABLES `party_functions` WRITE;
/*!40000 ALTER TABLE `party_functions` DISABLE KEYS */;
/*!40000 ALTER TABLE `party_functions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `party_identifying_codes`
--

DROP TABLE IF EXISTS `party_identifying_codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `party_identifying_codes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `dcsa_responsible_agency_code` varchar(5) DEFAULT '',
  `party_id` int(10) unsigned DEFAULT 0,
  `code_list_name` varchar(100) DEFAULT '',
  `party_code` varchar(100) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `party_identifying_codes`
--

LOCK TABLES `party_identifying_codes` WRITE;
/*!40000 ALTER TABLE `party_identifying_codes` DISABLE KEYS */;
/*!40000 ALTER TABLE `party_identifying_codes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment_term_types`
--

DROP TABLE IF EXISTS `payment_term_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payment_term_types` (
  `payment_term_code` varchar(3) DEFAULT '',
  `payment_term_name` varchar(100) DEFAULT '',
  `payment_term_description` varchar(250) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment_term_types`
--

LOCK TABLES `payment_term_types` WRITE;
/*!40000 ALTER TABLE `payment_term_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `payment_term_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pending_email_notification_deads`
--

DROP TABLE IF EXISTS `pending_email_notification_deads`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pending_email_notification_deads` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `event_id` binary(16) DEFAULT NULL,
  `template_name` varchar(80) DEFAULT '',
  `failure_reason_type` varchar(200) DEFAULT '',
  `failure_reason_message` varchar(200) DEFAULT '',
  `enqueued_at_date_time` datetime DEFAULT current_timestamp(),
  `last_failed_at_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pending_email_notification_deads`
--

LOCK TABLES `pending_email_notification_deads` WRITE;
/*!40000 ALTER TABLE `pending_email_notification_deads` DISABLE KEYS */;
/*!40000 ALTER TABLE `pending_email_notification_deads` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pending_email_notifications`
--

DROP TABLE IF EXISTS `pending_email_notifications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pending_email_notifications` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `event_id` binary(16) DEFAULT NULL,
  `template_name` varchar(80) DEFAULT '',
  `enqueued_at_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pending_email_notifications`
--

LOCK TABLES `pending_email_notifications` WRITE;
/*!40000 ALTER TABLE `pending_email_notifications` DISABLE KEYS */;
/*!40000 ALTER TABLE `pending_email_notifications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pending_event_queues`
--

DROP TABLE IF EXISTS `pending_event_queues`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pending_event_queues` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `delivery_id` binary(16) DEFAULT NULL,
  `subscription_id` binary(16) DEFAULT NULL,
  `event_id` binary(16) DEFAULT NULL,
  `payload` varchar(35) DEFAULT '',
  `last_error_message` varchar(35) DEFAULT '',
  `retry_count` int(10) unsigned DEFAULT 0,
  `enqueued_at_date_time` datetime DEFAULT current_timestamp(),
  `last_attempt_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pending_event_queues`
--

LOCK TABLES `pending_event_queues` WRITE;
/*!40000 ALTER TABLE `pending_event_queues` DISABLE KEYS */;
/*!40000 ALTER TABLE `pending_event_queues` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `point_to_point_routings`
--

DROP TABLE IF EXISTS `point_to_point_routings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `point_to_point_routings` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `sequence_number` int(11) DEFAULT 0,
  `place_of_receipt_id` int(10) unsigned DEFAULT 0,
  `place_of_delivery_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `point_to_point_routings`
--

LOCK TABLES `point_to_point_routings` WRITE;
/*!40000 ALTER TABLE `point_to_point_routings` DISABLE KEYS */;
/*!40000 ALTER TABLE `point_to_point_routings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `port_call_parts`
--

DROP TABLE IF EXISTS `port_call_parts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `port_call_parts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `port_call_part` varchar(100) DEFAULT '',
  `display_order` int(11) DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `port_call_parts`
--

LOCK TABLES `port_call_parts` WRITE;
/*!40000 ALTER TABLE `port_call_parts` DISABLE KEYS */;
/*!40000 ALTER TABLE `port_call_parts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `port_call_phase_types`
--

DROP TABLE IF EXISTS `port_call_phase_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `port_call_phase_types` (
  `port_call_phase_type_code` varchar(4) DEFAULT '',
  `port_call_phase_type_name` varchar(30) DEFAULT '',
  `port_call_phase_type_description` varchar(250) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `port_call_phase_types`
--

LOCK TABLES `port_call_phase_types` WRITE;
/*!40000 ALTER TABLE `port_call_phase_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `port_call_phase_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `port_call_service_types`
--

DROP TABLE IF EXISTS `port_call_service_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `port_call_service_types` (
  `port_call_service_type_code` varchar(4) DEFAULT '',
  `port_call_service_type_name` varchar(30) DEFAULT '',
  `port_call_service_type_description` varchar(250) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `port_call_service_types`
--

LOCK TABLES `port_call_service_types` WRITE;
/*!40000 ALTER TABLE `port_call_service_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `port_call_service_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `port_call_status_types`
--

DROP TABLE IF EXISTS `port_call_status_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `port_call_status_types` (
  `port_call_status_type_code` varchar(4) DEFAULT '',
  `port_call_status_type_name` varchar(30) DEFAULT '',
  `port_call_status_type_description` varchar(250) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `port_call_status_types`
--

LOCK TABLES `port_call_status_types` WRITE;
/*!40000 ALTER TABLE `port_call_status_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `port_call_status_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `port_time_zones`
--

DROP TABLE IF EXISTS `port_time_zones`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `port_time_zones` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `un_location_code` varchar(4) DEFAULT '',
  `iana_timezone` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `port_time_zones`
--

LOCK TABLES `port_time_zones` WRITE;
/*!40000 ALTER TABLE `port_time_zones` DISABLE KEYS */;
/*!40000 ALTER TABLE `port_time_zones` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `publisher_patterns`
--

DROP TABLE IF EXISTS `publisher_patterns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `publisher_patterns` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pattern_id` varchar(80) DEFAULT '',
  `publisher_role` varchar(3) DEFAULT '',
  `primary_receiver` varchar(3) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `publisher_patterns`
--

LOCK TABLES `publisher_patterns` WRITE;
/*!40000 ALTER TABLE `publisher_patterns` DISABLE KEYS */;
/*!40000 ALTER TABLE `publisher_patterns` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `receipt_delivery_types`
--

DROP TABLE IF EXISTS `receipt_delivery_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `receipt_delivery_types` (
  `receipt_delivery_type_code` varchar(3) DEFAULT '',
  `receipt_delivery_type_name` varchar(50) DEFAULT '',
  `receipt_delivery_type_description` varchar(300) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `receipt_delivery_types`
--

LOCK TABLES `receipt_delivery_types` WRITE;
/*!40000 ALTER TABLE `receipt_delivery_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `receipt_delivery_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reference1`
--

DROP TABLE IF EXISTS `reference1`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `reference1` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `reference_type_code` varchar(3) DEFAULT '',
  `reference_value` varchar(100) DEFAULT '',
  `shipment_id` int(10) unsigned DEFAULT 0,
  `shipping_instruction_id` int(10) unsigned DEFAULT 0,
  `booking_id` int(10) unsigned DEFAULT 0,
  `consignment_item_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reference1`
--

LOCK TABLES `reference1` WRITE;
/*!40000 ALTER TABLE `reference1` DISABLE KEYS */;
/*!40000 ALTER TABLE `reference1` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reference_types`
--

DROP TABLE IF EXISTS `reference_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `reference_types` (
  `reference_type_code` varchar(3) DEFAULT '',
  `reference_type_name` varchar(100) DEFAULT '',
  `reference_type_description` varchar(400) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reference_types`
--

LOCK TABLES `reference_types` WRITE;
/*!40000 ALTER TABLE `reference_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `reference_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `requested_equipments`
--

DROP TABLE IF EXISTS `requested_equipments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `requested_equipments` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `booking_id` int(10) unsigned DEFAULT 0,
  `shipment_id` int(10) unsigned DEFAULT 0,
  `requested_equipment_sizetype` varchar(4) DEFAULT '',
  `requested_equipment_units` int(11) DEFAULT 0,
  `confirmed_equipment_sizetype` varchar(4) DEFAULT '',
  `confirmed_equipment_units` int(11) DEFAULT 0,
  `is_shipper_owned` tinyint(1) DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `requested_equipments`
--

LOCK TABLES `requested_equipments` WRITE;
/*!40000 ALTER TABLE `requested_equipments` DISABLE KEYS */;
/*!40000 ALTER TABLE `requested_equipments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `seal_sources`
--

DROP TABLE IF EXISTS `seal_sources`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `seal_sources` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `seal_source_code` varchar(5) DEFAULT '',
  `seal_source_description` varchar(50) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `seal_sources`
--

LOCK TABLES `seal_sources` WRITE;
/*!40000 ALTER TABLE `seal_sources` DISABLE KEYS */;
/*!40000 ALTER TABLE `seal_sources` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `seal_types`
--

DROP TABLE IF EXISTS `seal_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `seal_types` (
  `seal_type_code` varchar(5) DEFAULT '',
  `seal_type_description` varchar(50) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `seal_types`
--

LOCK TABLES `seal_types` WRITE;
/*!40000 ALTER TABLE `seal_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `seal_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `seals`
--

DROP TABLE IF EXISTS `seals`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `seals` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `utilized_transport_equipment_id` int(10) unsigned DEFAULT 0,
  `seal_number` varchar(50) DEFAULT '',
  `seal_source_code` varchar(50) DEFAULT '',
  `seal_type_code` varchar(50) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `seals`
--

LOCK TABLES `seals` WRITE;
/*!40000 ALTER TABLE `seals` DISABLE KEYS */;
/*!40000 ALTER TABLE `seals` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_schedules`
--

DROP TABLE IF EXISTS `service_schedules`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `service_schedules` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `carrier_service_name` varchar(50) DEFAULT '',
  `carrier_service_code` varchar(5) DEFAULT '',
  `universal_service_reference` varchar(8) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_schedules`
--

LOCK TABLES `service_schedules` WRITE;
/*!40000 ALTER TABLE `service_schedules` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_schedules` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `services`
--

DROP TABLE IF EXISTS `services`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `services` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `carrier_id` int(10) unsigned DEFAULT 0,
  `carrier_service_name` varchar(50) DEFAULT '',
  `carrier_service_code` varchar(5) DEFAULT '',
  `tradelane_id` varchar(8) DEFAULT '',
  `universal_service_reference` varchar(8) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `services`
--

LOCK TABLES `services` WRITE;
/*!40000 ALTER TABLE `services` DISABLE KEYS */;
/*!40000 ALTER TABLE `services` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shipment_carrier_clauses`
--

DROP TABLE IF EXISTS `shipment_carrier_clauses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shipment_carrier_clauses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `carrier_clause_id` int(10) unsigned DEFAULT 0,
  `shipment_id` int(10) unsigned DEFAULT 0,
  `transport_document_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shipment_carrier_clauses`
--

LOCK TABLES `shipment_carrier_clauses` WRITE;
/*!40000 ALTER TABLE `shipment_carrier_clauses` DISABLE KEYS */;
/*!40000 ALTER TABLE `shipment_carrier_clauses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shipment_cutoff_times`
--

DROP TABLE IF EXISTS `shipment_cutoff_times`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shipment_cutoff_times` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `shipment_id` int(10) unsigned DEFAULT 0,
  `cut_off_time_code` varchar(3) DEFAULT '',
  `cut_off_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shipment_cutoff_times`
--

LOCK TABLES `shipment_cutoff_times` WRITE;
/*!40000 ALTER TABLE `shipment_cutoff_times` DISABLE KEYS */;
/*!40000 ALTER TABLE `shipment_cutoff_times` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shipment_event_types`
--

DROP TABLE IF EXISTS `shipment_event_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shipment_event_types` (
  `shipment_event_type_code` varchar(4) DEFAULT '',
  `shipment_event_type_name` varchar(30) DEFAULT '',
  `shipment_event_type_description` varchar(350) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shipment_event_types`
--

LOCK TABLES `shipment_event_types` WRITE;
/*!40000 ALTER TABLE `shipment_event_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `shipment_event_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shipment_events`
--

DROP TABLE IF EXISTS `shipment_events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shipment_events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `event_id_s` binary(16) DEFAULT NULL,
  `event_classifier_code` varchar(3) DEFAULT '',
  `shipment_event_type_code` varchar(4) DEFAULT '',
  `document_type_code` varchar(3) DEFAULT '',
  `document_id` int(10) unsigned DEFAULT 0,
  `document_reference` varchar(100) DEFAULT '',
  `reason` varchar(250) DEFAULT '',
  `event_created_date_time` datetime DEFAULT current_timestamp(),
  `event_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shipment_events`
--

LOCK TABLES `shipment_events` WRITE;
/*!40000 ALTER TABLE `shipment_events` DISABLE KEYS */;
/*!40000 ALTER TABLE `shipment_events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shipment_location_types`
--

DROP TABLE IF EXISTS `shipment_location_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shipment_location_types` (
  `shipment_location_type_code` varchar(3) DEFAULT '',
  `shipment_location_type_name` varchar(50) DEFAULT '',
  `shipment_location_type_description` varchar(250) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shipment_location_types`
--

LOCK TABLES `shipment_location_types` WRITE;
/*!40000 ALTER TABLE `shipment_location_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `shipment_location_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shipment_locations`
--

DROP TABLE IF EXISTS `shipment_locations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shipment_locations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `shipment_id` int(10) unsigned DEFAULT 0,
  `booking_id` int(10) unsigned DEFAULT 0,
  `location_id` int(10) unsigned DEFAULT 0,
  `shipment_location_type_code` varchar(3) DEFAULT '',
  `displayed_name` varchar(250) DEFAULT '',
  `event_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shipment_locations`
--

LOCK TABLES `shipment_locations` WRITE;
/*!40000 ALTER TABLE `shipment_locations` DISABLE KEYS */;
/*!40000 ALTER TABLE `shipment_locations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shipment_summaries`
--

DROP TABLE IF EXISTS `shipment_summaries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shipment_summaries` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `carrier_booking_reference` varchar(35) DEFAULT '',
  `terms_and_conditions` varchar(50) DEFAULT '',
  `carrier_booking_request_reference` varchar(100) DEFAULT '',
  `booking_document_status` varchar(50) DEFAULT '',
  `shipment_created_date_time` datetime DEFAULT current_timestamp(),
  `shipment_updated_date_time` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shipment_summaries`
--

LOCK TABLES `shipment_summaries` WRITE;
/*!40000 ALTER TABLE `shipment_summaries` DISABLE KEYS */;
/*!40000 ALTER TABLE `shipment_summaries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shipment_summary_cores`
--

DROP TABLE IF EXISTS `shipment_summary_cores`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shipment_summary_cores` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `carrier_booking_request_reference` varchar(100) DEFAULT '',
  `booking_document_status` varchar(50) DEFAULT '',
  `terms_and_conditions` varchar(50) DEFAULT '',
  `shipment_created_date_time` datetime DEFAULT current_timestamp(),
  `shipment_updated_date_time` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shipment_summary_cores`
--

LOCK TABLES `shipment_summary_cores` WRITE;
/*!40000 ALTER TABLE `shipment_summary_cores` DISABLE KEYS */;
/*!40000 ALTER TABLE `shipment_summary_cores` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shipment_transports`
--

DROP TABLE IF EXISTS `shipment_transports`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shipment_transports` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `shipment_id` int(10) unsigned DEFAULT 0,
  `transport_id` int(10) unsigned DEFAULT 0,
  `transport_plan_stage_sequence_number` int(10) unsigned DEFAULT 0,
  `transport_plan_stage_code` varchar(3) DEFAULT '',
  `commercial_voyage_id` int(10) unsigned DEFAULT 0,
  `is_under_shippers_responsibility` tinyint(1) DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shipment_transports`
--

LOCK TABLES `shipment_transports` WRITE;
/*!40000 ALTER TABLE `shipment_transports` DISABLE KEYS */;
/*!40000 ALTER TABLE `shipment_transports` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shipments`
--

DROP TABLE IF EXISTS `shipments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shipments` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `booking_id` int(10) unsigned DEFAULT 0,
  `carrier_id` int(10) unsigned DEFAULT 0,
  `carrier_booking_reference` varchar(35) DEFAULT '',
  `terms_and_conditions` varchar(50) DEFAULT '',
  `confirmation_datetime` datetime DEFAULT current_timestamp(),
  `updated_date_time` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shipments`
--

LOCK TABLES `shipments` WRITE;
/*!40000 ALTER TABLE `shipments` DISABLE KEYS */;
/*!40000 ALTER TABLE `shipments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shipping_instruction_ref_statuses`
--

DROP TABLE IF EXISTS `shipping_instruction_ref_statuses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shipping_instruction_ref_statuses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `shipping_instruction_reference` varchar(100) DEFAULT '',
  `document_status` varchar(4) DEFAULT '',
  `shipping_instruction_created_date_time` datetime DEFAULT current_timestamp(),
  `shipping_instruction_updated_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shipping_instruction_ref_statuses`
--

LOCK TABLES `shipping_instruction_ref_statuses` WRITE;
/*!40000 ALTER TABLE `shipping_instruction_ref_statuses` DISABLE KEYS */;
/*!40000 ALTER TABLE `shipping_instruction_ref_statuses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shipping_instruction_requests`
--

DROP TABLE IF EXISTS `shipping_instruction_requests`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shipping_instruction_requests` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `amend_to_transport_document` varchar(100) DEFAULT '',
  `transport_document_type_code` varchar(3) DEFAULT '',
  `is_shipped_onboard_type` tinyint(1) DEFAULT 0,
  `number_of_copies` int(10) unsigned DEFAULT 0,
  `number_of_originals` int(10) unsigned DEFAULT 0,
  `is_electronic` tinyint(1) DEFAULT 0,
  `is_to_order` tinyint(1) DEFAULT 0,
  `are_charges_displayed_on_originals` tinyint(1) DEFAULT 0,
  `are_charges_displayed_on_copies` tinyint(1) DEFAULT 0,
  `displayed_name_for_place_of_receipt` varchar(250) DEFAULT '',
  `displayed_name_for_port_of_load` varchar(250) DEFAULT '',
  `displayed_name_for_port_of_discharge` varchar(250) DEFAULT '',
  `displayed_name_for_place_of_delivery` varchar(250) DEFAULT '',
  `carrier_booking_reference` varchar(100) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shipping_instruction_requests`
--

LOCK TABLES `shipping_instruction_requests` WRITE;
/*!40000 ALTER TABLE `shipping_instruction_requests` DISABLE KEYS */;
/*!40000 ALTER TABLE `shipping_instruction_requests` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shipping_instruction_responses`
--

DROP TABLE IF EXISTS `shipping_instruction_responses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shipping_instruction_responses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `shipping_instruction_reference` varchar(100) DEFAULT '',
  `document_status` varchar(4) DEFAULT '',
  `amend_to_transport_document` varchar(100) DEFAULT '',
  `transport_document_type_code` varchar(3) DEFAULT '',
  `is_shipped_onboard_type` tinyint(1) DEFAULT 0,
  `number_of_copies` int(10) unsigned DEFAULT 0,
  `number_of_originals` int(10) unsigned DEFAULT 0,
  `is_electronic` tinyint(1) DEFAULT 0,
  `is_to_order` tinyint(1) DEFAULT 0,
  `are_charges_displayed_on_originals` tinyint(1) DEFAULT 0,
  `are_charges_displayed_on_copies` tinyint(1) DEFAULT 0,
  `displayed_name_for_place_of_receipt` varchar(250) DEFAULT '',
  `displayed_name_for_port_of_load` varchar(250) DEFAULT '',
  `displayed_name_for_port_of_discharge` varchar(250) DEFAULT '',
  `displayed_name_for_place_of_delivery` varchar(250) DEFAULT '',
  `carrier_booking_reference` varchar(100) DEFAULT '',
  `shipping_instruction_created_date_time` datetime DEFAULT current_timestamp(),
  `shipping_instruction_updated_date_time` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shipping_instruction_responses`
--

LOCK TABLES `shipping_instruction_responses` WRITE;
/*!40000 ALTER TABLE `shipping_instruction_responses` DISABLE KEYS */;
/*!40000 ALTER TABLE `shipping_instruction_responses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shipping_instruction_shallows`
--

DROP TABLE IF EXISTS `shipping_instruction_shallows`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shipping_instruction_shallows` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `amend_to_transport_document` varchar(100) DEFAULT '',
  `transport_document_type_code` varchar(3) DEFAULT '',
  `is_shipped_onboard_type` tinyint(1) DEFAULT 0,
  `number_of_copies` int(10) unsigned DEFAULT 0,
  `number_of_originals` int(10) unsigned DEFAULT 0,
  `is_electronic` tinyint(1) DEFAULT 0,
  `is_to_order` tinyint(1) DEFAULT 0,
  `are_charges_displayed_on_originals` tinyint(1) DEFAULT 0,
  `are_charges_displayed_on_copies` tinyint(1) DEFAULT 0,
  `displayed_name_for_place_of_receipt` varchar(250) DEFAULT '',
  `displayed_name_for_port_of_load` varchar(250) DEFAULT '',
  `displayed_name_for_port_of_discharge` varchar(250) DEFAULT '',
  `displayed_name_for_place_of_delivery` varchar(250) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shipping_instruction_shallows`
--

LOCK TABLES `shipping_instruction_shallows` WRITE;
/*!40000 ALTER TABLE `shipping_instruction_shallows` DISABLE KEYS */;
/*!40000 ALTER TABLE `shipping_instruction_shallows` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shipping_instruction_summaries`
--

DROP TABLE IF EXISTS `shipping_instruction_summaries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shipping_instruction_summaries` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `shipping_instruction_reference` varchar(100) DEFAULT '',
  `document_status` varchar(4) DEFAULT '',
  `amend_to_transport_document` varchar(100) DEFAULT '',
  `transport_document_type_code` varchar(3) DEFAULT '',
  `is_shipped_onboard_type` tinyint(1) DEFAULT 0,
  `number_of_copies` int(10) unsigned DEFAULT 0,
  `number_of_originals` int(10) unsigned DEFAULT 0,
  `is_electronic` tinyint(1) DEFAULT 0,
  `is_to_order` tinyint(1) DEFAULT 0,
  `are_charges_displayed_on_originals` tinyint(1) DEFAULT 0,
  `are_charges_displayed_on_copies` tinyint(1) DEFAULT 0,
  `displayed_name_for_place_of_receipt` varchar(250) DEFAULT '',
  `displayed_name_for_port_of_load` varchar(250) DEFAULT '',
  `displayed_name_for_port_of_discharge` varchar(250) DEFAULT '',
  `displayed_name_for_place_of_delivery` varchar(250) DEFAULT '',
  `shipping_instruction_created_date_time` datetime DEFAULT current_timestamp(),
  `shipping_instruction_updated_date_time` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shipping_instruction_summaries`
--

LOCK TABLES `shipping_instruction_summaries` WRITE;
/*!40000 ALTER TABLE `shipping_instruction_summaries` DISABLE KEYS */;
/*!40000 ALTER TABLE `shipping_instruction_summaries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `shipping_instructions`
--

DROP TABLE IF EXISTS `shipping_instructions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `shipping_instructions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `shipping_instruction_reference` varchar(100) DEFAULT '',
  `document_status` varchar(4) DEFAULT '',
  `is_shipped_onboard_type` tinyint(1) DEFAULT 0,
  `number_of_copies` int(10) unsigned DEFAULT 0,
  `number_of_originals` int(10) unsigned DEFAULT 0,
  `is_electronic` tinyint(1) DEFAULT 0,
  `is_to_order` tinyint(1) DEFAULT 0,
  `are_charges_displayed_on_originals` tinyint(1) DEFAULT 0,
  `are_charges_displayed_on_copies` tinyint(1) DEFAULT 0,
  `location_id` int(10) unsigned DEFAULT 0,
  `transport_document_type_code` varchar(3) DEFAULT '',
  `displayed_name_for_place_of_receipt` varchar(250) DEFAULT '',
  `displayed_name_for_port_of_load` varchar(250) DEFAULT '',
  `displayed_name_for_port_of_discharge` varchar(250) DEFAULT '',
  `displayed_name_for_place_of_delivery` varchar(250) DEFAULT '',
  `amend_to_transport_document` varchar(100) DEFAULT '',
  `created_date_time` datetime DEFAULT current_timestamp(),
  `updated_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `shipping_instructions`
--

LOCK TABLES `shipping_instructions` WRITE;
/*!40000 ALTER TABLE `shipping_instructions` DISABLE KEYS */;
/*!40000 ALTER TABLE `shipping_instructions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `smdg_delay_reasons`
--

DROP TABLE IF EXISTS `smdg_delay_reasons`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `smdg_delay_reasons` (
  `delay_reason_code` varchar(3) DEFAULT '',
  `delay_reason_name` varchar(100) DEFAULT '',
  `delay_reason_description` varchar(250) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `smdg_delay_reasons`
--

LOCK TABLES `smdg_delay_reasons` WRITE;
/*!40000 ALTER TABLE `smdg_delay_reasons` DISABLE KEYS */;
/*!40000 ALTER TABLE `smdg_delay_reasons` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `subscription_bodies`
--

DROP TABLE IF EXISTS `subscription_bodies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subscription_bodies` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `callback_url` varchar(100) DEFAULT '',
  `carrier_booking_reference` varchar(35) DEFAULT '',
  `transport_document_reference` varchar(20) DEFAULT '',
  `transport_call_id` varchar(100) DEFAULT '',
  `vessel_imo_number` varchar(7) DEFAULT '',
  `export_voyage_number` varchar(50) DEFAULT '',
  `carrier_service_code` varchar(5) DEFAULT '',
  `un_location_code` varchar(5) DEFAULT '',
  `equipment_reference` varchar(15) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `subscription_bodies`
--

LOCK TABLES `subscription_bodies` WRITE;
/*!40000 ALTER TABLE `subscription_bodies` DISABLE KEYS */;
/*!40000 ALTER TABLE `subscription_bodies` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `subscription_body_with_secrets`
--

DROP TABLE IF EXISTS `subscription_body_with_secrets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subscription_body_with_secrets` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `callback_url` varchar(100) DEFAULT '',
  `carrier_booking_reference` varchar(35) DEFAULT '',
  `transport_document_reference` varchar(20) DEFAULT '',
  `transport_call_id` varchar(100) DEFAULT '',
  `vessel_imo_number` varchar(7) DEFAULT '',
  `export_voyage_number` varchar(50) DEFAULT '',
  `carrier_service_code` varchar(5) DEFAULT '',
  `un_location_code` varchar(5) DEFAULT '',
  `equipment_reference` varchar(15) DEFAULT '',
  `secret` varchar(100) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `subscription_body_with_secrets`
--

LOCK TABLES `subscription_body_with_secrets` WRITE;
/*!40000 ALTER TABLE `subscription_body_with_secrets` DISABLE KEYS */;
/*!40000 ALTER TABLE `subscription_body_with_secrets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `subscriptions`
--

DROP TABLE IF EXISTS `subscriptions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subscriptions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `subscription_id` varchar(100) DEFAULT '',
  `callback_url` varchar(100) DEFAULT '',
  `carrier_booking_reference` varchar(35) DEFAULT '',
  `transport_document_reference` varchar(20) DEFAULT '',
  `transport_call_id` varchar(100) DEFAULT '',
  `vessel_imo_number` varchar(7) DEFAULT '',
  `export_voyage_number` varchar(50) DEFAULT '',
  `carrier_service_code` varchar(5) DEFAULT '',
  `un_location_code` varchar(5) DEFAULT '',
  `equipment_reference` varchar(15) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `subscriptions`
--

LOCK TABLES `subscriptions` WRITE;
/*!40000 ALTER TABLE `subscriptions` DISABLE KEYS */;
/*!40000 ALTER TABLE `subscriptions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `surrender_request_answers`
--

DROP TABLE IF EXISTS `surrender_request_answers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `surrender_request_answers` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `surrender_request_reference` varchar(100) DEFAULT '',
  `action` varchar(4) DEFAULT '',
  `comments` varchar(255) DEFAULT '',
  `created_date_time` datetime DEFAULT current_timestamp(),
  `surrender_request_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_surrender_request_answers_action` (`action`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `surrender_request_answers`
--

LOCK TABLES `surrender_request_answers` WRITE;
/*!40000 ALTER TABLE `surrender_request_answers` DISABLE KEYS */;
/*!40000 ALTER TABLE `surrender_request_answers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `surrender_requests`
--

DROP TABLE IF EXISTS `surrender_requests`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `surrender_requests` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `surrender_request_reference` varchar(100) DEFAULT '',
  `transport_document_reference` varchar(20) DEFAULT '',
  `surrender_request_code` varchar(4) DEFAULT '',
  `comments` varchar(255) DEFAULT '',
  `surrender_requested_by` int(10) unsigned DEFAULT 0,
  `created_date_time` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_surrender_requests_surrender_request_reference` (`surrender_request_reference`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `surrender_requests`
--

LOCK TABLES `surrender_requests` WRITE;
/*!40000 ALTER TABLE `surrender_requests` DISABLE KEYS */;
/*!40000 ALTER TABLE `surrender_requests` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tax_and_legal_reference_types`
--

DROP TABLE IF EXISTS `tax_and_legal_reference_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tax_and_legal_reference_types` (
  `tax_and_legal_reference_type_code` varchar(50) DEFAULT '',
  `tax_and_legal_country_code` varchar(2) DEFAULT '',
  `tax_and_legal_reference_type_name` varchar(50) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tax_and_legal_reference_types`
--

LOCK TABLES `tax_and_legal_reference_types` WRITE;
/*!40000 ALTER TABLE `tax_and_legal_reference_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `tax_and_legal_reference_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `timestamp_definition_publisher_patterns`
--

DROP TABLE IF EXISTS `timestamp_definition_publisher_patterns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `timestamp_definition_publisher_patterns` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `timestamp_id` varchar(80) DEFAULT '',
  `pattern_id` varchar(80) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `timestamp_definition_publisher_patterns`
--

LOCK TABLES `timestamp_definition_publisher_patterns` WRITE;
/*!40000 ALTER TABLE `timestamp_definition_publisher_patterns` DISABLE KEYS */;
/*!40000 ALTER TABLE `timestamp_definition_publisher_patterns` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `timestamp_definitions`
--

DROP TABLE IF EXISTS `timestamp_definitions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `timestamp_definitions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `timestamp_id` varchar(80) DEFAULT '',
  `timestamp_type_name` varchar(80) DEFAULT '',
  `event_classifier_code` varchar(3) DEFAULT '',
  `operations_event_type_code` varchar(4) DEFAULT '',
  `port_call_phase_type_code` varchar(4) DEFAULT '',
  `port_call_service_type_code` varchar(4) DEFAULT '',
  `facility_type_code` varchar(4) DEFAULT '',
  `port_call_part` varchar(100) DEFAULT '',
  `event_location_requirement` varchar(10) DEFAULT '',
  `is_terminal_needed` tinyint(1) DEFAULT 0,
  `is_vessel_draft_relevant` tinyint(1) DEFAULT 0,
  `vessel_position_requirement` varchar(10) DEFAULT '',
  `is_miles_to_destination_relevant` tinyint(1) DEFAULT 0,
  `provided_in_standard` varchar(80) DEFAULT '',
  `accept_timestamp_definition` varchar(80) DEFAULT '',
  `reject_timestamp_definition` varchar(80) DEFAULT '',
  `negotiation_cycle` varchar(50) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `timestamp_definitions`
--

LOCK TABLES `timestamp_definitions` WRITE;
/*!40000 ALTER TABLE `timestamp_definitions` DISABLE KEYS */;
/*!40000 ALTER TABLE `timestamp_definitions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `timestamp_notification_deads`
--

DROP TABLE IF EXISTS `timestamp_notification_deads`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `timestamp_notification_deads` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `message_routing_rule_id` int(10) unsigned DEFAULT 0,
  `payload` varchar(80) DEFAULT '',
  `latest_delivery_attempted_datetime` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `timestamp_notification_deads`
--

LOCK TABLES `timestamp_notification_deads` WRITE;
/*!40000 ALTER TABLE `timestamp_notification_deads` DISABLE KEYS */;
/*!40000 ALTER TABLE `timestamp_notification_deads` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `timestamps`
--

DROP TABLE IF EXISTS `timestamps`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `timestamps` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `event_type_code` varchar(7) DEFAULT '',
  `event_classifier_code` varchar(35) DEFAULT '',
  `delay_reason_code` varchar(3) DEFAULT '',
  `change_remark` varchar(250) DEFAULT '',
  `event_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `timestamps`
--

LOCK TABLES `timestamps` WRITE;
/*!40000 ALTER TABLE `timestamps` DISABLE KEYS */;
/*!40000 ALTER TABLE `timestamps` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tradelanes`
--

DROP TABLE IF EXISTS `tradelanes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tradelanes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `tradelane_name` varchar(150) DEFAULT '',
  `vessel_sharing_agreement_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tradelanes`
--

LOCK TABLES `tradelanes` WRITE;
/*!40000 ALTER TABLE `tradelanes` DISABLE KEYS */;
/*!40000 ALTER TABLE `tradelanes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transaction_parties`
--

DROP TABLE IF EXISTS `transaction_parties`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transaction_parties` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `ebl_platform_identifier` varchar(100) DEFAULT '',
  `legal_name` varchar(100) DEFAULT '',
  `registration_number` varchar(100) DEFAULT '',
  `location_of_registration` varchar(2) DEFAULT '',
  `tax_reference` varchar(100) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaction_parties`
--

LOCK TABLES `transaction_parties` WRITE;
/*!40000 ALTER TABLE `transaction_parties` DISABLE KEYS */;
/*!40000 ALTER TABLE `transaction_parties` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transaction_party_supporting_codes`
--

DROP TABLE IF EXISTS `transaction_party_supporting_codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transaction_party_supporting_codes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `transaction_party_id` int(10) unsigned DEFAULT 0,
  `party_code` varchar(100) DEFAULT '',
  `party_code_list_provider` varchar(4) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaction_party_supporting_codes`
--

LOCK TABLES `transaction_party_supporting_codes` WRITE;
/*!40000 ALTER TABLE `transaction_party_supporting_codes` DISABLE KEYS */;
/*!40000 ALTER TABLE `transaction_party_supporting_codes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transport_call_jit_port_visits`
--

DROP TABLE IF EXISTS `transport_call_jit_port_visits`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transport_call_jit_port_visits` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `port_visit_id` int(10) unsigned DEFAULT 0,
  `transport_call_id` int(10) unsigned DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transport_call_jit_port_visits`
--

LOCK TABLES `transport_call_jit_port_visits` WRITE;
/*!40000 ALTER TABLE `transport_call_jit_port_visits` DISABLE KEYS */;
/*!40000 ALTER TABLE `transport_call_jit_port_visits` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transport_calls`
--

DROP TABLE IF EXISTS `transport_calls`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transport_calls` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `transport_call_reference` varchar(100) DEFAULT '',
  `transport_call_sequence_number` int(10) unsigned DEFAULT 0,
  `facility_id` int(10) unsigned DEFAULT 0,
  `facility_type_code` varchar(4) DEFAULT '',
  `other_facility` varchar(50) DEFAULT '',
  `location_id` int(10) unsigned DEFAULT 0,
  `mode_of_transport_code` varchar(50) DEFAULT '',
  `vessel_id` int(10) unsigned DEFAULT 0,
  `import_voyage_id` int(10) unsigned DEFAULT 0,
  `export_voyage_id` int(10) unsigned DEFAULT 0,
  `port_call_status_code` varchar(4) DEFAULT '',
  `port_visit_reference` varchar(50) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transport_calls`
--

LOCK TABLES `transport_calls` WRITE;
/*!40000 ALTER TABLE `transport_calls` DISABLE KEYS */;
/*!40000 ALTER TABLE `transport_calls` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transport_document_ref_statuses`
--

DROP TABLE IF EXISTS `transport_document_ref_statuses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transport_document_ref_statuses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `transport_document_reference` varchar(20) DEFAULT '',
  `document_status` varchar(4) DEFAULT '',
  `transport_document_created_date_time` datetime DEFAULT current_timestamp(),
  `transport_document_updated_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transport_document_ref_statuses`
--

LOCK TABLES `transport_document_ref_statuses` WRITE;
/*!40000 ALTER TABLE `transport_document_ref_statuses` DISABLE KEYS */;
/*!40000 ALTER TABLE `transport_document_ref_statuses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transport_document_roots`
--

DROP TABLE IF EXISTS `transport_document_roots`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transport_document_roots` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `transport_document_reference` varchar(20) DEFAULT '',
  `number_of_originals` int(10) unsigned DEFAULT 0,
  `carrier_code` varchar(4) DEFAULT '',
  `carrier_code_list_provider` varchar(4) DEFAULT '',
  `number_of_rider_pages` int(11) DEFAULT 0,
  `transport_document_created_date_time` datetime DEFAULT current_timestamp(),
  `transport_document_updated_date_time` datetime DEFAULT current_timestamp(),
  `issue_date` datetime DEFAULT current_timestamp(),
  `shipped_onboard_date` datetime DEFAULT current_timestamp(),
  `received_for_shipment_date` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transport_document_roots`
--

LOCK TABLES `transport_document_roots` WRITE;
/*!40000 ALTER TABLE `transport_document_roots` DISABLE KEYS */;
/*!40000 ALTER TABLE `transport_document_roots` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transport_document_summaries`
--

DROP TABLE IF EXISTS `transport_document_summaries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transport_document_summaries` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `transport_document_reference` varchar(20) DEFAULT '',
  `number_of_originals` int(10) unsigned DEFAULT 0,
  `carrier_code` varchar(4) DEFAULT '',
  `carrier_code_list_provider` varchar(4) DEFAULT '',
  `number_of_rider_pages` int(11) DEFAULT 0,
  `shipping_instruction_reference` varchar(100) DEFAULT '',
  `document_status` varchar(4) DEFAULT '',
  `transport_document_created_date_time` datetime DEFAULT current_timestamp(),
  `transport_document_updated_date_time` datetime DEFAULT current_timestamp(),
  `issue_date` datetime DEFAULT current_timestamp(),
  `shipped_onboard_date` datetime DEFAULT current_timestamp(),
  `received_for_shipment_date` datetime DEFAULT current_timestamp(),
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transport_document_summaries`
--

LOCK TABLES `transport_document_summaries` WRITE;
/*!40000 ALTER TABLE `transport_document_summaries` DISABLE KEYS */;
/*!40000 ALTER TABLE `transport_document_summaries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transport_document_types`
--

DROP TABLE IF EXISTS `transport_document_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transport_document_types` (
  `transport_document_type_code` varchar(3) DEFAULT '',
  `transport_document_type_name` varchar(20) DEFAULT '',
  `transport_document_type_description` varchar(500) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transport_document_types`
--

LOCK TABLES `transport_document_types` WRITE;
/*!40000 ALTER TABLE `transport_document_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `transport_document_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transport_documents`
--

DROP TABLE IF EXISTS `transport_documents`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transport_documents` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `transport_document_reference` varchar(20) DEFAULT '',
  `location_id` int(10) unsigned DEFAULT 0,
  `number_of_originals` int(10) unsigned DEFAULT 0,
  `carrier_id` int(10) unsigned DEFAULT 0,
  `shipping_instruction_id` int(10) unsigned DEFAULT 0,
  `declared_value_currency` varchar(3) DEFAULT '',
  `declared_value` double DEFAULT 0,
  `number_of_rider_pages` int(11) DEFAULT 0,
  `issuing_party` varchar(100) DEFAULT '',
  `issue_date` datetime DEFAULT current_timestamp(),
  `shipped_onboard_date` datetime DEFAULT current_timestamp(),
  `received_for_shipment_date` datetime DEFAULT current_timestamp(),
  `created_date_time` datetime DEFAULT current_timestamp(),
  `updated_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transport_documents`
--

LOCK TABLES `transport_documents` WRITE;
/*!40000 ALTER TABLE `transport_documents` DISABLE KEYS */;
/*!40000 ALTER TABLE `transport_documents` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transport_event_types`
--

DROP TABLE IF EXISTS `transport_event_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transport_event_types` (
  `transport_event_type_code` varchar(4) DEFAULT '',
  `transport_event_type_name` varchar(30) DEFAULT '',
  `transport_event_type_description` varchar(250) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transport_event_types`
--

LOCK TABLES `transport_event_types` WRITE;
/*!40000 ALTER TABLE `transport_event_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `transport_event_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transport_events`
--

DROP TABLE IF EXISTS `transport_events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transport_events` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `event_id_s` binary(16) DEFAULT NULL,
  `event_classifier_code` varchar(3) DEFAULT '',
  `transport_event_type_code` varchar(4) DEFAULT '',
  `delay_reason_code` varchar(4) DEFAULT '',
  `change_remark` varchar(250) DEFAULT '',
  `transport_call_id` int(10) unsigned DEFAULT 0,
  `event_created_date_time` datetime DEFAULT current_timestamp(),
  `event_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transport_events`
--

LOCK TABLES `transport_events` WRITE;
/*!40000 ALTER TABLE `transport_events` DISABLE KEYS */;
/*!40000 ALTER TABLE `transport_events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transport_plan_stage_types`
--

DROP TABLE IF EXISTS `transport_plan_stage_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transport_plan_stage_types` (
  `transport_plan_stage_code` varchar(3) DEFAULT '',
  `transport_plan_stage_name` varchar(100) DEFAULT '',
  `transport_plan_stage_description` varchar(250) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transport_plan_stage_types`
--

LOCK TABLES `transport_plan_stage_types` WRITE;
/*!40000 ALTER TABLE `transport_plan_stage_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `transport_plan_stage_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transports`
--

DROP TABLE IF EXISTS `transports`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transports` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `transport_reference` varchar(50) DEFAULT '',
  `transport_name` varchar(100) DEFAULT '',
  `load_transport_call_id` int(10) unsigned DEFAULT 0,
  `discharge_transport_call_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transports`
--

LOCK TABLES `transports` WRITE;
/*!40000 ALTER TABLE `transports` DISABLE KEYS */;
/*!40000 ALTER TABLE `transports` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `un_locations`
--

DROP TABLE IF EXISTS `un_locations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `un_locations` (
  `un_location_code` varchar(5) DEFAULT '',
  `un_location_name` varchar(100) DEFAULT '',
  `location_code` varchar(3) DEFAULT '',
  `country_code` varchar(2) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `un_locations`
--

LOCK TABLES `un_locations` WRITE;
/*!40000 ALTER TABLE `un_locations` DISABLE KEYS */;
/*!40000 ALTER TABLE `un_locations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `unit_of_measures`
--

DROP TABLE IF EXISTS `unit_of_measures`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `unit_of_measures` (
  `unit_of_measure_code` varchar(3) DEFAULT '',
  `unit_of_measure_description` varchar(50) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `unit_of_measures`
--

LOCK TABLES `unit_of_measures` WRITE;
/*!40000 ALTER TABLE `unit_of_measures` DISABLE KEYS */;
/*!40000 ALTER TABLE `unit_of_measures` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `unmapped_event_queues`
--

DROP TABLE IF EXISTS `unmapped_event_queues`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `unmapped_event_queues` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `event_id` binary(16) DEFAULT NULL,
  `enqueued_at_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `unmapped_event_queues`
--

LOCK TABLES `unmapped_event_queues` WRITE;
/*!40000 ALTER TABLE `unmapped_event_queues` DISABLE KEYS */;
/*!40000 ALTER TABLE `unmapped_event_queues` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `utilized_transport_equipments`
--

DROP TABLE IF EXISTS `utilized_transport_equipments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `utilized_transport_equipments` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `equipment_reference` varchar(15) DEFAULT '',
  `cargo_gross_weight` double DEFAULT 0,
  `cargo_gross_weight_unit` varchar(3) DEFAULT '',
  `is_shipper_owned` tinyint(1) DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `utilized_transport_equipments`
--

LOCK TABLES `utilized_transport_equipments` WRITE;
/*!40000 ALTER TABLE `utilized_transport_equipments` DISABLE KEYS */;
/*!40000 ALTER TABLE `utilized_transport_equipments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `value_added_service_requests`
--

DROP TABLE IF EXISTS `value_added_service_requests`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `value_added_service_requests` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `booking_id` int(10) unsigned DEFAULT 0,
  `value_added_service_code` varchar(5) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `value_added_service_requests`
--

LOCK TABLES `value_added_service_requests` WRITE;
/*!40000 ALTER TABLE `value_added_service_requests` DISABLE KEYS */;
/*!40000 ALTER TABLE `value_added_service_requests` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `value_added_services`
--

DROP TABLE IF EXISTS `value_added_services`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `value_added_services` (
  `value_added_service_code` varchar(5) DEFAULT '',
  `value_added_service_name` varchar(100) DEFAULT '',
  `value_added_service_description` varchar(200) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `value_added_services`
--

LOCK TABLES `value_added_services` WRITE;
/*!40000 ALTER TABLE `value_added_services` DISABLE KEYS */;
/*!40000 ALTER TABLE `value_added_services` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `vessel_schedule_terminal_visits`
--

DROP TABLE IF EXISTS `vessel_schedule_terminal_visits`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vessel_schedule_terminal_visits` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `vessel_schedule_id` int(10) unsigned DEFAULT 0,
  `actual_arrival_event_id` int(10) unsigned DEFAULT 0,
  `planned_arrival_event_id` int(10) unsigned DEFAULT 0,
  `estimated_arrival_event_id` int(10) unsigned DEFAULT 0,
  `actual_departure_event_id` int(10) unsigned DEFAULT 0,
  `planned_departure_event_id` int(10) unsigned DEFAULT 0,
  `estimated_departure_event_id` int(10) unsigned DEFAULT 0,
  `port_call_status_event_id` int(10) unsigned DEFAULT 0,
  `transport_call_sequence` int(11) DEFAULT 0,
  `created_date_time` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `vessel_schedule_terminal_visits`
--

LOCK TABLES `vessel_schedule_terminal_visits` WRITE;
/*!40000 ALTER TABLE `vessel_schedule_terminal_visits` DISABLE KEYS */;
/*!40000 ALTER TABLE `vessel_schedule_terminal_visits` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `vessel_schedules`
--

DROP TABLE IF EXISTS `vessel_schedules`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vessel_schedules` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `vessel_id` int(10) unsigned DEFAULT 0,
  `service_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `vessel_schedules`
--

LOCK TABLES `vessel_schedules` WRITE;
/*!40000 ALTER TABLE `vessel_schedules` DISABLE KEYS */;
/*!40000 ALTER TABLE `vessel_schedules` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `vessel_sharing_agreement_types`
--

DROP TABLE IF EXISTS `vessel_sharing_agreement_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vessel_sharing_agreement_types` (
  `vessel_sharing_agreement_type_code` varchar(3) DEFAULT '',
  `vessel_sharing_agreement_type_name` varchar(50) DEFAULT '',
  `vessel_sharing_agreement_type_description` varchar(250) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `vessel_sharing_agreement_types`
--

LOCK TABLES `vessel_sharing_agreement_types` WRITE;
/*!40000 ALTER TABLE `vessel_sharing_agreement_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `vessel_sharing_agreement_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `vessel_sharing_agreements`
--

DROP TABLE IF EXISTS `vessel_sharing_agreements`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vessel_sharing_agreements` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `vessel_sharing_agreement_name` varchar(50) DEFAULT '',
  `vessel_sharing_agreement_type_code` varchar(10) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `vessel_sharing_agreements`
--

LOCK TABLES `vessel_sharing_agreements` WRITE;
/*!40000 ALTER TABLE `vessel_sharing_agreements` DISABLE KEYS */;
/*!40000 ALTER TABLE `vessel_sharing_agreements` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `vessel_types`
--

DROP TABLE IF EXISTS `vessel_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vessel_types` (
  `vessel_type_code` varchar(4) DEFAULT '',
  `vessel_type_name` varchar(100) DEFAULT '',
  `unece_concatenated_means_of_transport_code` varchar(4) DEFAULT '',
  `vessel_type_description` varchar(100) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `vessel_types`
--

LOCK TABLES `vessel_types` WRITE;
/*!40000 ALTER TABLE `vessel_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `vessel_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `vessels`
--

DROP TABLE IF EXISTS `vessels`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vessels` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `vessel_imo_number` varchar(7) DEFAULT '',
  `vessel_name` varchar(35) DEFAULT '',
  `vessel_flag` varchar(2) DEFAULT '',
  `vessel_call_sign` varchar(10) DEFAULT '',
  `is_dummy_vessel` tinyint(1) DEFAULT 0,
  `vessel_operator_carrier_code` varchar(50) DEFAULT '',
  `vessel_operator_carrier_code_list_provider` varchar(50) DEFAULT '',
  `vessel_length` double DEFAULT 0,
  `vessel_width` double DEFAULT 0,
  `dimension_unit` varchar(10) DEFAULT '',
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `vessels`
--

LOCK TABLES `vessels` WRITE;
/*!40000 ALTER TABLE `vessels` DISABLE KEYS */;
/*!40000 ALTER TABLE `vessels` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `voyages`
--

DROP TABLE IF EXISTS `voyages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `voyages` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid4` binary(16) DEFAULT NULL,
  `carrier_voyage_number` varchar(50) DEFAULT '',
  `universal_voyage_reference` varchar(50) DEFAULT '',
  `service_id` int(10) unsigned DEFAULT 0,
  `status_code` varchar(50) DEFAULT 'active',
  `created_by_user_id` varchar(50) DEFAULT 'active',
  `updated_by_user_id` varchar(50) DEFAULT 'active',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `voyages`
--

LOCK TABLES `voyages` WRITE;
/*!40000 ALTER TABLE `voyages` DISABLE KEYS */;
/*!40000 ALTER TABLE `voyages` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-02-05 14:25:55
