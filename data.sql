-- MySQL dump 10.13  Distrib 8.4.8-8, for Linux (x86_64)
--
-- Host: localhost    Database: dashboard-ape
-- ------------------------------------------------------
-- Server version	8.4.5

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
/*!50717 SELECT COUNT(*) INTO @rocksdb_has_p_s_session_variables FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'performance_schema' AND TABLE_NAME = 'session_variables' */;
/*!50717 SET @rocksdb_get_is_supported = IF (@rocksdb_has_p_s_session_variables, 'SELECT COUNT(*) INTO @rocksdb_is_supported FROM performance_schema.session_variables WHERE VARIABLE_NAME=\'rocksdb_bulk_load\'', 'SELECT 0') */;
/*!50717 PREPARE s FROM @rocksdb_get_is_supported */;
/*!50717 EXECUTE s */;
/*!50717 DEALLOCATE PREPARE s */;
/*!50717 SET @rocksdb_enable_bulk_load = IF (@rocksdb_is_supported, 'SET SESSION rocksdb_bulk_load = 1', 'SET @rocksdb_dummy_bulk_load = 0') */;
/*!50717 PREPARE s FROM @rocksdb_enable_bulk_load */;
/*!50717 EXECUTE s */;
/*!50717 DEALLOCATE PREPARE s */;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `name` varchar(100) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_user_category` (`user_id`,`name`),
  KEY `idx_categories_user` (`user_id`),
  CONSTRAINT `fk_categories_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES ('0a031532-c1d3-4353-aa0f-6d721924f24a','e6e3d341-67bd-4964-882d-ada5c9511ef8','Cair','2026-06-01 00:45:55'),('2a562e9a-eff7-46b8-84f5-3d83f4b20eea','06ada16c-5e76-4faf-befe-d9346599eae2','Jajan','2026-06-04 03:05:04'),('589a9880-28d8-4308-93c6-8497c6ff49e4','e6e3d341-67bd-4964-882d-ada5c9511ef8','Jajan With Alizza','2026-05-31 13:16:06'),('7b02d701-0f67-4414-b1e1-0bf14b23c98b','e6e3d341-67bd-4964-882d-ada5c9511ef8','Keperluan Motor','2026-05-31 13:16:46'),('9ada2679-4416-4aec-b5f6-8757c327ccce','06ada16c-5e76-4faf-befe-d9346599eae2','berkah','2026-06-04 13:33:40'),('ae411d28-e35d-4413-9e5d-cb70b84218b5','06ada16c-5e76-4faf-befe-d9346599eae2','Gift','2026-06-04 03:05:52'),('cf1bd9d9-41bb-4890-a439-49233635a401','e6e3d341-67bd-4964-882d-ada5c9511ef8','Tanpa Keterangan ','2026-07-04 15:36:59'),('d2a6d382-9234-405d-839d-ff6206682074','e6e3d341-67bd-4964-882d-ada5c9511ef8','Jajan Harian','2026-05-31 13:16:38'),('f8aea20b-a93b-4580-9294-d17ed9f855e2','06ada16c-5e76-4faf-befe-d9346599eae2','Skincare & Make Up','2026-06-04 03:05:42');
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `habit_logs`
--

DROP TABLE IF EXISTS `habit_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `habit_logs` (
  `id` char(36) NOT NULL,
  `habit_id` char(36) NOT NULL,
  `log_date` date NOT NULL,
  `completed` tinyint(1) DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_log` (`habit_id`,`log_date`),
  KEY `idx_habit_logs_habit` (`habit_id`),
  KEY `idx_habit_logs_date` (`log_date`),
  KEY `idx_habit_logs_habit_date` (`habit_id`,`log_date`),
  CONSTRAINT `habit_logs_ibfk_1` FOREIGN KEY (`habit_id`) REFERENCES `habits` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `habit_logs`
--

LOCK TABLES `habit_logs` WRITE;
/*!40000 ALTER TABLE `habit_logs` DISABLE KEYS */;
INSERT INTO `habit_logs` VALUES ('0a4d8292-d6fb-42b3-a2e0-4cc319446c65','414a293d-ec1f-4053-a4f1-b2c690d0323e','2026-07-04',1,'2026-07-04 09:02:42'),('0a8b57f8-3a38-4d68-8fc9-0b6ddf2f5423','5a2f34e4-424b-46c2-b0e5-991fc823da19','2026-07-04',1,'2026-07-04 09:02:41'),('125324a2-9b02-49e1-9202-309b5218cb5a','659c1fdb-9969-42d8-9fc0-061cf322919b','2026-07-05',1,'2026-07-05 05:31:56'),('159bd134-885b-4b70-b31b-ee342f9deed9','eb871067-af3f-475c-a50b-7210a37c6aee','2026-07-04',0,'2026-07-04 08:57:54'),('17f4232b-cd4e-4726-85ec-1ebf1b44f8a7','69b576b3-defb-4fb0-8de4-efd8910d16ea','2026-07-04',1,'2026-07-04 13:58:29'),('184dd50d-ddf9-4a2e-ad5c-db66dff79e4d','69b576b3-defb-4fb0-8de4-efd8910d16ea','2026-07-05',1,'2026-07-05 12:14:29'),('24f82c40-58e2-40ee-983b-eb66062dbf38','a3a2e414-e639-4ddb-8587-460ceb129a3e','2026-07-05',1,'2026-07-05 12:14:26'),('372f5988-72ab-4360-911f-c6f087d6edc3','fd8938bd-4703-4927-a6bc-c6a33f6fcb92','2026-07-04',1,'2026-07-04 09:02:43'),('3fa4cd73-8c7e-4510-a343-e9c822e7ffe6','fd8938bd-4703-4927-a6bc-c6a33f6fcb92','2026-07-05',1,'2026-07-05 02:28:56'),('41831473-79d1-4937-ba05-9a3aadfd221c','e233bd32-e5d6-41c2-a25b-b55356b9248c','2026-07-05',1,'2026-07-05 09:04:15'),('6b4fee20-2dc0-4924-929a-9e0f37e96505','a3a2e414-e639-4ddb-8587-460ceb129a3e','2026-07-04',1,'2026-07-04 11:26:32'),('6bd06d58-18f7-4830-ae25-12847e4ef6e2','cd9a5fcc-a274-48fc-8851-1e7be9d1c8ee','2026-07-05',1,'2026-07-05 12:14:27'),('8164818e-c975-4395-b80f-5b36019df013','659c1fdb-9969-42d8-9fc0-061cf322919b','2026-07-04',1,'2026-07-04 09:02:45'),('88238e0b-5889-4d56-8a92-90eb6868fb6e','e233bd32-e5d6-41c2-a25b-b55356b9248c','2026-07-04',1,'2026-07-04 09:02:47'),('8d169824-6d06-44d1-9ed1-de188b7cd202','0199f15f-3691-4313-9cf0-4be2d8983815','2026-07-05',1,'2026-07-05 10:26:53'),('9359c784-8bab-4cb6-acb3-58ccd2472043','82144d57-54c7-4473-b069-4c110a6d3367','2026-07-05',1,'2026-07-05 09:04:16'),('94154058-35a7-4ce4-8191-e0019a59caec','414a293d-ec1f-4053-a4f1-b2c690d0323e','2026-07-05',1,'2026-07-05 02:28:55'),('a2badddc-761e-4a14-9536-715edc8b214f','82144d57-54c7-4473-b069-4c110a6d3367','2026-07-04',1,'2026-07-04 10:33:59'),('d811b996-22bc-4a80-b64b-b05faeb6a7aa','cd9a5fcc-a274-48fc-8851-1e7be9d1c8ee','2026-07-04',1,'2026-07-04 12:39:54'),('e59f5a2f-c702-418f-b7cc-b4050c6bebb4','f9f83626-1786-45fd-a793-bbfb84c0a890','2026-07-04',1,'2026-07-04 09:02:58'),('e8ad285b-3eec-457d-904e-218390658272','f9f83626-1786-45fd-a793-bbfb84c0a890','2026-07-05',1,'2026-07-05 12:26:37'),('ea1d7613-0350-41d0-bce9-ad0bbdbefaca','5a2f34e4-424b-46c2-b0e5-991fc823da19','2026-07-05',1,'2026-07-05 02:28:53');
/*!40000 ALTER TABLE `habit_logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `habits`
--

DROP TABLE IF EXISTS `habits`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `habits` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `frequency` enum('daily','weekly') DEFAULT 'daily',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `days` json DEFAULT NULL,
  `reminder_time` time DEFAULT NULL,
  `reminder_enabled` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `idx_reminder_time` (`reminder_time`),
  CONSTRAINT `habits_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `habits`
--

LOCK TABLES `habits` WRITE;
/*!40000 ALTER TABLE `habits` DISABLE KEYS */;
INSERT INTO `habits` VALUES ('0199f15f-3691-4313-9cf0-4be2d8983815','e6e3d341-67bd-4964-882d-ada5c9511ef8','Mandi Sore','daily','2026-07-05 10:26:50',NULL,'17:00:00',1),('414a293d-ec1f-4053-a4f1-b2c690d0323e','e6e3d341-67bd-4964-882d-ada5c9511ef8','Solat Shubuh','daily','2026-07-04 09:02:25',NULL,'05:10:00',1),('5a2f34e4-424b-46c2-b0e5-991fc823da19','e6e3d341-67bd-4964-882d-ada5c9511ef8','Mandi Pagi','daily','2026-07-04 09:02:34',NULL,'05:00:00',1),('659c1fdb-9969-42d8-9fc0-061cf322919b','e6e3d341-67bd-4964-882d-ada5c9511ef8','Solat Dzuhur','daily','2026-07-04 09:01:39',NULL,'12:00:00',1),('69b576b3-defb-4fb0-8de4-efd8910d16ea','e6e3d341-67bd-4964-882d-ada5c9511ef8','Makan Malam','daily','2026-07-04 08:59:36',NULL,'19:45:00',1),('82144d57-54c7-4473-b069-4c110a6d3367','e6e3d341-67bd-4964-882d-ada5c9511ef8','Solat Ashar','daily','2026-07-04 09:00:47',NULL,'15:25:00',1),('a3a2e414-e639-4ddb-8587-460ceb129a3e','e6e3d341-67bd-4964-882d-ada5c9511ef8','Solat Magrib','daily','2026-07-04 09:00:15',NULL,'18:05:00',1),('cd9a5fcc-a274-48fc-8851-1e7be9d1c8ee','e6e3d341-67bd-4964-882d-ada5c9511ef8','Solat Isya','daily','2026-07-04 08:59:47',NULL,'19:20:00',1),('e233bd32-e5d6-41c2-a25b-b55356b9248c','e6e3d341-67bd-4964-882d-ada5c9511ef8','Makan Siang','daily','2026-07-04 09:01:29',NULL,'12:30:00',1),('eb871067-af3f-475c-a50b-7210a37c6aee','e6e3d341-67bd-4964-882d-ada5c9511ef8','Cuci Muka','daily','2026-07-04 08:57:37',NULL,'22:30:00',1),('f9f83626-1786-45fd-a793-bbfb84c0a890','e6e3d341-67bd-4964-882d-ada5c9511ef8','Siapin Buku Sekolah','daily','2026-07-04 08:58:39',NULL,'22:00:00',1),('fd8938bd-4703-4927-a6bc-c6a33f6fcb92','e6e3d341-67bd-4964-882d-ada5c9511ef8','Makan Pagi','daily','2026-07-04 09:02:10',NULL,'05:20:00',1);
/*!40000 ALTER TABLE `habits` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `journals`
--

DROP TABLE IF EXISTS `journals`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `journals` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `mood` int NOT NULL,
  `content` text NOT NULL,
  `entry_date` date NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_user_entry_date` (`user_id`,`entry_date`),
  CONSTRAINT `journals_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `journals`
--

LOCK TABLES `journals` WRITE;
/*!40000 ALTER TABLE `journals` DISABLE KEYS */;
INSERT INTO `journals` VALUES ('14093200-986d-406e-be91-05498a7404ce','e6e3d341-67bd-4964-882d-ada5c9511ef8',5,'sangat senang sekali','2026-06-28','2026-06-28 12:33:28','2026-06-28 12:33:28'),('21c4d4c4-7c4d-44fd-b75a-cff44f4ee1a5','e6e3d341-67bd-4964-882d-ada5c9511ef8',5,'SERU MAIN VOLLY','2026-06-23','2026-06-23 05:28:53','2026-06-23 05:28:53'),('22fb311b-8f48-4e3d-b4da-099217edf6aa','e6e3d341-67bd-4964-882d-ada5c9511ef8',5,'menegangkan','2026-06-09','2026-06-09 03:45:00','2026-06-09 03:45:00'),('31d6d820-7d27-4a74-8f09-ba8f4b9e97b9','e6e3d341-67bd-4964-882d-ada5c9511ef8',5,'udah selesaii woyyyyyyy','2026-06-12','2026-06-12 06:07:32','2026-06-12 06:07:32'),('404fad69-9964-4f24-9106-00d2b1d3af1c','e6e3d341-67bd-4964-882d-ada5c9511ef8',5,'amaazingfgg','2026-06-19','2026-06-19 08:20:18','2026-06-19 08:20:18'),('5cdf36a9-93bb-47f6-ac93-d06ce515f5d3','e6e3d341-67bd-4964-882d-ada5c9511ef8',5,'Supra natural','2026-07-05','2026-07-04 22:27:38','2026-07-05 12:14:57'),('5e210885-573c-4003-8995-b766eeac65ad','e6e3d341-67bd-4964-882d-ada5c9511ef8',5,'seruu bangett nie','2026-06-22','2026-06-22 00:49:24','2026-06-22 00:49:24'),('668a5c1c-6a1c-4841-948b-0b6715c1c87e','e6e3d341-67bd-4964-882d-ada5c9511ef8',5,'SERU BNGT','2026-06-26','2026-06-26 15:44:59','2026-06-26 15:44:59'),('6d6c606d-660b-4645-98f1-c57863a5092d','e6e3d341-67bd-4964-882d-ada5c9511ef8',5,'SERU BNGT WOI ','2026-06-25','2026-06-25 06:18:04','2026-06-25 06:18:04'),('7918f783-8394-4aeb-a506-3126f4326cab','e6e3d341-67bd-4964-882d-ada5c9511ef8',5,'gue sangat amat senang','2026-06-27','2026-06-27 15:08:59','2026-06-27 15:08:59'),('822a628c-f5ba-4f6a-b284-4a6109ef2f43','e6e3d341-67bd-4964-882d-ada5c9511ef8',5,'seru banget coiiii','2026-06-14','2026-06-14 12:18:14','2026-06-14 12:18:14'),('87afa088-484d-46f5-ab5b-6feda40692a8','e6e3d341-67bd-4964-882d-ada5c9511ef8',4,'seperti biasanya','2026-06-15','2026-06-15 13:02:54','2026-06-15 13:02:54'),('a7c92d08-721c-496e-a4dc-917f29eca19e','e6e3d341-67bd-4964-882d-ada5c9511ef8',4,'good seru banget hari ini suka deh','2026-06-03','2026-06-03 06:27:45','2026-06-03 06:27:45'),('a945dc97-ec9b-41bf-932d-e8a982887764','e6e3d341-67bd-4964-882d-ada5c9511ef8',5,'mau selesai ujiannn','2026-06-11','2026-06-11 14:05:43','2026-06-11 14:05:43'),('b3608063-ecd7-4d68-b067-c9b99a7c7aff','e6e3d341-67bd-4964-882d-ada5c9511ef8',5,'amaazingggggg','2026-06-17','2026-06-17 05:48:20','2026-06-17 05:48:20'),('c6da0fc8-a8c2-41f5-bf76-6d307aee9455','e6e3d341-67bd-4964-882d-ada5c9511ef8',4,'ulangan hari ini menyenangkan','2026-06-08','2026-06-08 15:37:41','2026-06-08 15:37:41'),('d3910557-a344-4e09-a9db-4db478dbf997','e6e3d341-67bd-4964-882d-ada5c9511ef8',4,'senang sekalii pada hari ini aku sangat senang','2026-06-04','2026-06-04 09:34:08','2026-06-04 09:34:08'),('e9121138-14af-493b-8f71-c60f6bd7db2d','e6e3d341-67bd-4964-882d-ada5c9511ef8',5,'belajar coiii','2026-06-07','2026-06-07 15:07:25','2026-06-07 15:07:25'),('fa17165b-e84d-4884-bfc6-a5eff7f72cbd','e6e3d341-67bd-4964-882d-ada5c9511ef8',5,'ini sayaa lagi seneng banget','2026-06-13','2026-06-13 05:37:39','2026-06-13 05:37:39');
/*!40000 ALTER TABLE `journals` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schema_migrations`
--

DROP TABLE IF EXISTS `schema_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schema_migrations` (
  `version` bigint NOT NULL,
  `dirty` tinyint(1) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schema_migrations`
--

LOCK TABLES `schema_migrations` WRITE;
/*!40000 ALTER TABLE `schema_migrations` DISABLE KEYS */;
INSERT INTO `schema_migrations` VALUES (5,0),(8,0);
/*!40000 ALTER TABLE `schema_migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tasks`
--

DROP TABLE IF EXISTS `tasks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tasks` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` text,
  `quadrant` int NOT NULL DEFAULT '1',
  `is_completed` tinyint(1) NOT NULL DEFAULT '0',
  `due_date` datetime DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `tasks_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tasks`
--

LOCK TABLES `tasks` WRITE;
/*!40000 ALTER TABLE `tasks` DISABLE KEYS */;
INSERT INTO `tasks` VALUES ('361fad9c-7bb0-47ae-9483-8f5afa30fab2','e6e3d341-67bd-4964-882d-ada5c9511ef8','Project Rapat Penus','urgent kepala sekolah',1,1,NULL,'2026-06-05 04:30:16','2026-06-27 15:08:46'),('ad830f8a-1582-4445-a3c6-8f8d71c6fb67','e6e3d341-67bd-4964-882d-ada5c9511ef8','Project Rohis','website catatan alquran',2,0,NULL,'2026-06-13 09:07:16','2026-06-13 09:07:16'),('c4e5343e-21e8-479e-aa9a-d6331bba7c1e','e6e3d341-67bd-4964-882d-ada5c9511ef8','ABSENSI SISWA','project sekolah',2,0,NULL,'2026-06-05 04:30:43','2026-06-13 10:39:57'),('fcf6d66e-281e-45d0-87d8-5c4cde118fc9','e6e3d341-67bd-4964-882d-ada5c9511ef8','Project Pt','mencoba membuat sela ',1,0,NULL,'2026-06-14 12:18:50','2026-06-14 12:18:50');
/*!40000 ALTER TABLE `tasks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `timeblocks`
--

DROP TABLE IF EXISTS `timeblocks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `timeblocks` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `activity_name` varchar(255) NOT NULL,
  `start_time` varchar(5) NOT NULL,
  `end_time` varchar(5) NOT NULL,
  `color_code` varchar(7) NOT NULL DEFAULT '#4F46E5',
  `day_of_week` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `date` date DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `timeblocks_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `timeblocks`
--

LOCK TABLES `timeblocks` WRITE;
/*!40000 ALTER TABLE `timeblocks` DISABLE KEYS */;
INSERT INTO `timeblocks` VALUES ('7d09ff18-80d2-4304-abf6-f4ccdd35e535','e6e3d341-67bd-4964-882d-ada5c9511ef8','Curug with friend ','07:00','07:00','#4F46E5',4,'2026-06-24 11:28:05','2026-06-24 11:28:05','2026-06-25'),('cd4078a5-e777-4bba-9bd2-8f5413b6913d','e6e3d341-67bd-4964-882d-ada5c9511ef8','Masuk sekolah first kelas XII','06:00','09:00','#4F46E5',1,'2026-06-27 15:06:04','2026-06-27 15:06:04','2026-07-13');
/*!40000 ALTER TABLE `timeblocks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `category_id` char(36) DEFAULT NULL,
  `amount` decimal(15,2) NOT NULL,
  `type` enum('income','expense') NOT NULL,
  `description` text,
  `transaction_date` date NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_transactions_user` (`user_id`),
  KEY `idx_transactions_category` (`category_id`),
  KEY `idx_transactions_date` (`transaction_date`),
  CONSTRAINT `fk_transactions_category` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE SET NULL,
  CONSTRAINT `fk_transactions_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES ('02b5eecd-c692-430c-ad0e-839ebcab14ef','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',10000.00,'expense','jajan mie dll','2026-06-04','2026-06-04 05:22:44'),('051a2597-d6dd-4706-a113-bc02e221bc15','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',20000.00,'income','jajan harian','2026-06-29','2026-06-29 05:41:44'),('0917135d-5507-4b3c-9a57-43a1e1cafc81','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',20000.00,'income','jajan harian','2026-07-05','2026-07-05 06:30:30'),('0c2e8e51-3c6c-47b3-ac38-c9a07d6418ed','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',25000.00,'income','jajan haruan','2026-06-24','2026-06-24 11:27:42'),('0dc921c6-0782-4bbb-bae6-87e4e9d0c5a5','e6e3d341-67bd-4964-882d-ada5c9511ef8',NULL,5000.00,'expense','bebeb','2026-06-04','2026-06-04 13:04:51'),('12e41e85-d354-4de8-bffa-fa9123309b3a','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',30000.00,'income','jajan harian','2026-06-09','2026-06-09 02:15:13'),('13f5c140-4cf9-4838-b912-502cd247f927','e6e3d341-67bd-4964-882d-ada5c9511ef8','589a9880-28d8-4308-93c6-8497c6ff49e4',3000.00,'expense','hansaplash','2026-06-25','2026-06-24 23:54:27'),('146541e5-d69a-40ff-9189-47a538d970c2','e6e3d341-67bd-4964-882d-ada5c9511ef8','0a031532-c1d3-4353-aa0f-6d721924f24a',70000.00,'income','cair bosquu','2026-06-16','2026-06-16 02:17:42'),('15abb66d-8087-47e5-8eb8-6302aac991c2','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',30000.00,'expense','gacoan','2026-07-04','2026-07-04 16:37:25'),('1a0c40d1-bad1-42d1-be47-38fa8a7fe44f','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',30000.00,'income','Jajan harian','2026-06-05','2026-06-05 04:27:22'),('1ba87361-d013-491c-b0a2-3876ad993941','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',20000.00,'income','pemasukan cui','2026-06-13','2026-06-13 04:14:48'),('1e1dbf0e-8857-4e8f-99f6-df415df2ef26','e6e3d341-67bd-4964-882d-ada5c9511ef8','589a9880-28d8-4308-93c6-8497c6ff49e4',124000.00,'expense','jajan curugnyaa','2026-06-25','2026-06-25 06:33:29'),('215c2e5b-ca31-4829-9f4e-45fef7199358','e6e3d341-67bd-4964-882d-ada5c9511ef8','7b02d701-0f67-4414-b1e1-0bf14b23c98b',20000.00,'income','uang cuci motor','2026-06-21','2026-06-21 09:33:26'),('233d632c-0d00-40e2-9497-aca409160ed0','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',13000.00,'expense','jajan','2026-06-05','2026-06-05 12:18:43'),('2615f85a-e022-43bf-bd23-18374a63de8f','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',9000.00,'expense','jajannn','2026-06-13','2026-06-13 10:17:18'),('26455569-c3bd-4200-9eac-10e439efcf04','e6e3d341-67bd-4964-882d-ada5c9511ef8','7b02d701-0f67-4414-b1e1-0bf14b23c98b',20000.00,'expense','cuci','2026-06-21','2026-06-21 10:00:03'),('2754dbf9-846d-4b88-a72d-10db92f8277f','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',3000.00,'expense','parkir','2026-06-22','2026-06-22 10:25:51'),('2e772cfc-ea66-496c-877d-33a7a6e5185f','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',16000.00,'expense','kopi','2026-06-30','2026-06-30 06:43:13'),('390c956c-3a71-4df9-a208-e99d89e18799','e6e3d341-67bd-4964-882d-ada5c9511ef8',NULL,5000.00,'expense','ke anak jalanan','2026-06-04','2026-06-04 13:04:35'),('3c2ead99-621a-4a72-9898-2bedcb88a1d5','e6e3d341-67bd-4964-882d-ada5c9511ef8','0a031532-c1d3-4353-aa0f-6d721924f24a',30000.00,'income','hasil cair','2026-06-16','2026-06-16 03:33:38'),('3f6cd2f8-9cbd-4f33-a439-a0f16a0160e0','e6e3d341-67bd-4964-882d-ada5c9511ef8','0a031532-c1d3-4353-aa0f-6d721924f24a',10000.00,'income','alizza bayar','2026-06-08','2026-06-08 09:10:55'),('41e63857-cfb8-44ce-b3e8-7c25f5f3ea75','e6e3d341-67bd-4964-882d-ada5c9511ef8','0a031532-c1d3-4353-aa0f-6d721924f24a',7000.00,'income','cari ','2026-06-18','2026-06-18 02:49:36'),('47eaf6fe-df7d-4b1d-bf5b-4a5433b4a281','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',25000.00,'income','jajan','2026-06-26','2026-06-26 04:41:01'),('4ba9856e-5486-43a5-8f75-cbb89fc87a4a','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',10000.00,'expense','jajans sate','2026-06-14','2026-06-14 09:04:38'),('50cc8bd1-e75b-40da-8779-7816a3699de3','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',6000.00,'expense','es teh','2026-06-15','2026-06-14 22:23:38'),('5207b4d8-ea8e-4b83-8a56-3942f816d497','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',30000.00,'income','jajan harian','2026-06-08','2026-06-08 00:36:35'),('573e23e9-4591-449a-8204-92d8036bb608','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',8000.00,'expense','jajan cireng + es mahal','2026-06-13','2026-06-13 04:16:35'),('5a442e40-1451-4475-8db5-9eb9124d6629','e6e3d341-67bd-4964-882d-ada5c9511ef8',NULL,128.00,'expense','sedekah','2026-06-05','2026-06-05 12:19:30'),('5b580540-0bdc-40ec-83c8-0f5684565dc8','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',30000.00,'income','ini jajan harian','2026-06-12','2026-06-11 22:28:20'),('5b5c6d91-69d5-49f0-bbdd-05d9957f2d28','e6e3d341-67bd-4964-882d-ada5c9511ef8','0a031532-c1d3-4353-aa0f-6d721924f24a',38128.00,'income','cairrrr','2026-06-04','2026-06-04 02:29:24'),('5f23136b-88d4-4fd5-be02-b80d3bc86464','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',36000.00,'expense','jajan harian','2026-06-23','2026-06-23 08:20:08'),('699f1243-05c6-4e07-9e29-a94c59328663','06ada16c-5e76-4faf-befe-d9346599eae2',NULL,13000.00,'expense','buat mamah (hrusnya diganti) ','2026-06-04','2026-06-04 13:34:12'),('6b056b47-0a4b-4a7d-be9d-67873bb362dd','06ada16c-5e76-4faf-befe-d9346599eae2','2a562e9a-eff7-46b8-84f5-3d83f4b20eea',16000.00,'expense','jajan sekolah','2026-06-04','2026-06-04 11:18:18'),('6bcae45f-78ce-4085-97b1-0b790850c80a','e6e3d341-67bd-4964-882d-ada5c9511ef8','cf1bd9d9-41bb-4890-a439-49233635a401',26000.00,'expense','hilang entah kemana ','2026-07-04','2026-07-04 11:33:50'),('6d7f9274-31bf-422a-87c2-a9d85a396d55','e6e3d341-67bd-4964-882d-ada5c9511ef8',NULL,30000.00,'expense','jajan abis dari keperluan motor , jajan with alizza , smaa jajan harian ','2026-06-08','2026-06-08 09:10:38'),('6fa78b91-9bb5-41ed-b298-d3c839e8a077','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',20000.00,'income','jajan harian','2026-06-30','2026-06-30 04:12:44'),('77729602-121e-4928-b25b-8201a3aff56b','e6e3d341-67bd-4964-882d-ada5c9511ef8','589a9880-28d8-4308-93c6-8497c6ff49e4',68000.00,'expense','jajannnn','2026-06-30','2026-06-30 06:43:03'),('78bb3485-30fe-4743-8653-9ad4dcbebfe4','e6e3d341-67bd-4964-882d-ada5c9511ef8','589a9880-28d8-4308-93c6-8497c6ff49e4',3000.00,'expense','','2026-06-04','2026-06-04 05:30:28'),('7ffcb515-22bc-4a1d-a123-65aa8293e76c','e6e3d341-67bd-4964-882d-ada5c9511ef8','0a031532-c1d3-4353-aa0f-6d721924f24a',250000.00,'income','cairrrr','2026-06-25','2026-06-24 23:54:15'),('823aa627-f55e-46a4-a0e6-b68618c0111b','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',45000.00,'income','jajan','2026-06-14','2026-06-14 09:04:11'),('82f8d4a7-0a8c-4baf-8db8-997029d679b7','e6e3d341-67bd-4964-882d-ada5c9511ef8','7b02d701-0f67-4414-b1e1-0bf14b23c98b',5000.00,'expense','parkir','2026-06-25','2026-06-25 06:04:51'),('863005db-2c41-4fe2-a2ec-64a721cc0844','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',30000.00,'income','jajan harian','2026-06-04','2026-06-04 02:28:57'),('8776a0de-c3db-48d4-bc39-3fb1934f0458','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',20000.00,'income','jajan harian ','2026-06-07','2026-06-07 03:23:07'),('89e03ef7-d597-4402-bfa0-a490096d115d','e6e3d341-67bd-4964-882d-ada5c9511ef8','589a9880-28d8-4308-93c6-8497c6ff49e4',105000.00,'expense','jajan with bebeb','2026-06-30','2026-06-30 05:40:41'),('91db44ca-61dc-43d9-8597-8fda32d73aba','06ada16c-5e76-4faf-befe-d9346599eae2','9ada2679-4416-4aec-b5f6-8757c327ccce',4000.00,'expense','bismillah berjah','2026-06-04','2026-06-04 13:34:39'),('9bb35c6e-cff1-470a-993f-cd41371b80ec','e6e3d341-67bd-4964-882d-ada5c9511ef8',NULL,20000.00,'expense','ilangg GG ','2026-06-10','2026-06-09 22:41:46'),('9e823837-96aa-4db3-934e-a09ce4611983','e6e3d341-67bd-4964-882d-ada5c9511ef8','589a9880-28d8-4308-93c6-8497c6ff49e4',18000.00,'expense','beli chiken sama es kopi','2026-06-15','2026-06-15 13:01:21'),('9fda6959-51ee-4f97-808d-470be8b4850e','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',60000.00,'income','jajan harian ','2026-06-23','2026-06-23 00:57:43'),('a26f8573-7858-49da-96f0-e1d73164493d','e6e3d341-67bd-4964-882d-ada5c9511ef8','589a9880-28d8-4308-93c6-8497c6ff49e4',15000.00,'expense','ke alizza','2026-06-05','2026-06-05 12:19:13'),('a7075349-a6fa-4a94-a719-faff105a1e70','e6e3d341-67bd-4964-882d-ada5c9511ef8','0a031532-c1d3-4353-aa0f-6d721924f24a',9000.00,'income','namabahh','2026-06-06','2026-06-06 12:26:22'),('aae3eff2-5826-4bfd-b315-07aca1228434','e6e3d341-67bd-4964-882d-ada5c9511ef8','589a9880-28d8-4308-93c6-8497c6ff49e4',34000.00,'expense','Exo and coklat','2026-06-23','2026-06-23 08:50:54'),('abc58fda-6d61-4d01-9758-28380d0a39a7','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',5000.00,'expense','kentang','2026-06-07','2026-06-07 03:23:17'),('ac8e654b-f929-4158-95f6-6331782f71d7','e6e3d341-67bd-4964-882d-ada5c9511ef8',NULL,25000.00,'expense','ini seputar hari ini ','2026-06-11','2026-06-11 14:05:22'),('aeb69d1c-1b07-4005-a9c3-1aa51529bcec','06ada16c-5e76-4faf-befe-d9346599eae2','2a562e9a-eff7-46b8-84f5-3d83f4b20eea',7000.00,'expense','lapar','2026-06-04','2026-06-04 13:35:01'),('b069dcf0-c35a-435a-b672-1885022d4271','e6e3d341-67bd-4964-882d-ada5c9511ef8','589a9880-28d8-4308-93c6-8497c6ff49e4',32000.00,'expense','misuuuuuu','2026-06-04','2026-06-04 11:18:24'),('b4fda8df-ed88-42db-86c9-35504816401e','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',4000.00,'expense','bihun','2026-06-13','2026-06-13 04:15:02'),('b76a2858-84d4-4bc0-be30-305d4b1e0342','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',16000.00,'expense','beli micuu','2026-06-21','2026-06-21 09:41:26'),('b8746c6b-0b55-4e23-8bfb-181572d776a5','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',26000.00,'expense','nasi goreng','2026-06-14','2026-06-14 11:34:35'),('bb677948-82b2-4502-9af9-33a2411eb96a','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',12000.00,'expense','minyak kayu putih ','2026-06-24','2026-06-24 13:36:25'),('bc2812d5-820a-451d-bb5c-c163cce675a0','e6e3d341-67bd-4964-882d-ada5c9511ef8','0a031532-c1d3-4353-aa0f-6d721924f24a',2000.00,'income','nemu','2026-06-28','2026-06-28 15:17:24'),('bce8256c-1cfe-408d-a853-7312f4a4a5d9','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',13000.00,'expense','jusssss','2026-06-18','2026-06-18 03:38:58'),('be967090-d6bd-4766-a8b2-427ffc68591c','e6e3d341-67bd-4964-882d-ada5c9511ef8','0a031532-c1d3-4353-aa0f-6d721924f24a',20000.00,'income','jajan harian ','2026-06-18','2026-06-18 02:49:24'),('c05e797d-094d-4981-8c58-c50a46ca04d5','e6e3d341-67bd-4964-882d-ada5c9511ef8','0a031532-c1d3-4353-aa0f-6d721924f24a',10000.00,'income','caur bossqu','2026-06-27','2026-06-26 17:16:05'),('c06a4ef3-82dc-48f5-a488-dc11e7f7b334','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',10000.00,'expense','beli cilor','2026-06-18','2026-06-18 02:57:40'),('c3643d40-755d-40f6-b051-42ddb7a76226','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',31000.00,'expense','dimsum','2026-06-22','2026-06-22 10:26:00'),('c98b27c8-5005-424d-8826-890828da6f4f','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',20000.00,'income','jajan','2026-06-06','2026-06-06 01:32:25'),('ca345adb-622b-4d31-b07f-01ff74f69235','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',16000.00,'expense','mixue','2026-06-12','2026-06-12 06:51:55'),('cc970c92-2403-4c9f-8022-4085de44938c','e6e3d341-67bd-4964-882d-ada5c9511ef8','0a031532-c1d3-4353-aa0f-6d721924f24a',50000.00,'income','jajan coyyy','2026-06-10','2026-06-10 09:32:53'),('cd44ff16-1126-453d-babe-3e49697b8be9','e6e3d341-67bd-4964-882d-ada5c9511ef8','7b02d701-0f67-4414-b1e1-0bf14b23c98b',3000.00,'expense','parkir','2026-06-19','2026-06-19 01:27:17'),('cda8d1d7-9821-45f8-a275-d6704bac52c9','e6e3d341-67bd-4964-882d-ada5c9511ef8','0a031532-c1d3-4353-aa0f-6d721924f24a',15000.00,'income','cairrrr','2026-06-06','2026-06-06 14:06:52'),('d05a78d4-d444-4f5f-ba80-014b84cd5940','e6e3d341-67bd-4964-882d-ada5c9511ef8','7b02d701-0f67-4414-b1e1-0bf14b23c98b',4000.00,'expense','parkirnya ','2026-06-25','2026-06-25 06:17:09'),('d59776eb-39f0-4ba2-95ba-78cd47c6eacf','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',25000.00,'income','cairrrr','2026-06-21','2026-06-21 05:11:28'),('d5b8aa00-ab5a-4ea9-b58e-259bb1b675ff','e6e3d341-67bd-4964-882d-ada5c9511ef8','589a9880-28d8-4308-93c6-8497c6ff49e4',38000.00,'expense','momoyo + parkir','2026-06-10','2026-06-10 09:33:22'),('d7e4875b-09b2-41c0-8820-689cc17ab522','e6e3d341-67bd-4964-882d-ada5c9511ef8','589a9880-28d8-4308-93c6-8497c6ff49e4',31000.00,'expense','jajan alizza and mee','2026-06-25','2026-06-25 08:07:29'),('dc078254-d56f-4edd-bfc1-a6488094e6ff','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',20000.00,'expense','gacoan','2026-06-06','2026-06-06 14:04:43'),('dc8858f7-58ef-4c8b-95a3-09852817fbde','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',13000.00,'expense','beli jus sama kopi','2026-06-09','2026-06-09 03:44:14'),('e02d9f1a-9404-4976-aaa7-da285e45a566','e6e3d341-67bd-4964-882d-ada5c9511ef8','0a031532-c1d3-4353-aa0f-6d721924f24a',10000.00,'income','caur','2026-06-24','2026-06-24 15:54:26'),('e0489203-e750-4581-be62-d80ca80f6aa5','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',10000.00,'expense','kentang','2026-06-04','2026-06-04 13:04:11'),('e40818c2-4763-4a7c-8045-95aa3e406b58','e6e3d341-67bd-4964-882d-ada5c9511ef8','7b02d701-0f67-4414-b1e1-0bf14b23c98b',30000.00,'expense','bensin','2026-06-23','2026-06-23 08:58:11'),('e52d03b5-da01-4a29-958c-66345751a489','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',30000.00,'income','jajan harian','2026-06-22','2026-06-22 00:29:59'),('e5959a6c-95c2-4bcf-90ae-f82416d59c7b','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',20000.00,'income','jajan harian','2026-06-20','2026-06-20 01:02:48'),('ead2e8f6-54b0-4b29-8503-d74d00fe4099','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',20000.00,'income','jajan harian ','2026-06-15','2026-06-15 08:51:34'),('f5a571de-bd35-4e13-98b0-b7080c5ad310','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',10000.00,'expense','ultra milk','2026-06-22','2026-06-22 10:26:15'),('f615c315-7362-4a7b-8e29-c5b4de4dabf4','e6e3d341-67bd-4964-882d-ada5c9511ef8','7b02d701-0f67-4414-b1e1-0bf14b23c98b',5000.00,'expense','parkir','2026-06-07','2026-06-07 03:23:26'),('f9bdd87b-e08a-4366-8829-3bab3e2c861b','e6e3d341-67bd-4964-882d-ada5c9511ef8','7b02d701-0f67-4414-b1e1-0bf14b23c98b',4000.00,'expense','for parkiran','2026-06-12','2026-06-12 05:35:23'),('fe5b0bac-a4b4-42ad-8a91-e7d3c826c4bd','e6e3d341-67bd-4964-882d-ada5c9511ef8','589a9880-28d8-4308-93c6-8497c6ff49e4',50000.00,'expense','jajan bubub','2026-06-28','2026-06-28 15:17:07'),('fe7f6589-3428-4c4e-a3a8-7610d4baa639','06ada16c-5e76-4faf-befe-d9346599eae2',NULL,40000.00,'income','uang masuk','2026-06-04','2026-06-04 03:07:22'),('ff1117b1-ab8c-4db0-83b6-58a215951688','e6e3d341-67bd-4964-882d-ada5c9511ef8','0a031532-c1d3-4353-aa0f-6d721924f24a',114000.00,'income','cair coyyy ','2026-06-28','2026-06-28 12:28:13'),('ffbb3d96-8b5c-412d-8915-b73a9f1bf43e','e6e3d341-67bd-4964-882d-ada5c9511ef8','d2a6d382-9234-405d-839d-ff6206682074',10000.00,'expense','jajan ','2026-06-29','2026-06-29 05:50:44');
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_devices`
--

DROP TABLE IF EXISTS `user_devices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_devices` (
  `id` char(36) NOT NULL,
  `user_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `fmc_token` varchar(255) NOT NULL,
  `device_type` enum('ios','android') NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  CONSTRAINT `user_devices_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_devices`
--

LOCK TABLES `user_devices` WRITE;
/*!40000 ALTER TABLE `user_devices` DISABLE KEYS */;
INSERT INTO `user_devices` VALUES ('7822a4e5-8238-4d42-89b3-a102af7d9daf','e6e3d341-67bd-4964-882d-ada5c9511ef8','dFHzpGwzQJS5I6DK5mHxc_:APA91bE1MdCIolJnqIAy1bx9XZeTxmpRWK1n9OYn6VQQqQ6HPk5Q4CY9ViRSDiYPdBZFVp8fzzhM5vmoyuIX-tOMjC85w7Slg4EMe7mN3axH2G5iG7Wj3w8','android','2026-06-27 12:07:14','2026-07-05 13:07:02');
/*!40000 ALTER TABLE `user_devices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` char(36) NOT NULL,
  `name` varchar(100) NOT NULL,
  `email` varchar(150) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `nomor_hp` varchar(20) DEFAULT NULL,
  `bio` text,
  `email_verified` tinyint(1) DEFAULT '0',
  `verification_code` varchar(255) DEFAULT NULL,
  `verification_expire_at` datetime DEFAULT NULL,
  `reset_password_code` varchar(255) DEFAULT NULL,
  `reset_password_expire_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_email` (`email`),
  KEY `idx_verification_code` (`verification_code`),
  KEY `idx_reset_password_code` (`reset_password_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('0472ec5d-280f-458f-b2a8-6a1a1f196669','piyaah','luthpiyyahalipah11@gmail.com','$2a$10$GWYg0d6lrFZaxd4V90doTOPjtFPc2uTHhpKWtlI5zOxpJU0AQgHee','2026-06-21 06:04:45','2026-06-21 06:05:07','0754649','jelek bngt',1,'',NULL,'',NULL),('06ada16c-5e76-4faf-befe-d9346599eae2','Alizza Pasha Fahira','afhiraa@gmail.com','$2a$10$uF5kZle.kRXcF/6UVYrCnOBSsFtP0Fet08ZHmMYV5NyuuX1pzI/CG','2026-06-04 03:03:27','2026-06-04 03:03:27',NULL,NULL,0,NULL,NULL,NULL,NULL),('e6e3d341-67bd-4964-882d-ada5c9511ef8','Reihan','rhanssap@gmail.com','$2a$10$.ju6e2xoqIx18q6hVZ.5.u/RHBbOMKCvOQf64F8thT2DH58jZHmZe','2026-05-31 12:50:04','2026-07-04 09:44:33','0895345570902','my bio is keren',1,'',NULL,'',NULL),('f81f8176-55be-4a4c-82c6-f6bbe0615c66','ALVIN Rizky','alvinr120409@gmail.com','$2a$10$g2LgMf/nER/ZpaGmQj7XvO7t1cVVLlQdoofmiSuSrhyCo4kvgimUe','2026-06-22 07:27:32','2026-06-22 07:33:26','081398036877','totot',1,'',NULL,'',NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!50112 SET @disable_bulk_load = IF (@is_rocksdb_supported, 'SET SESSION rocksdb_bulk_load = @old_rocksdb_bulk_load', 'SET @dummy_rocksdb_bulk_load = 0') */;
/*!50112 PREPARE s FROM @disable_bulk_load */;
/*!50112 EXECUTE s */;
/*!50112 DEALLOCATE PREPARE s */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-07-05 15:36:13
