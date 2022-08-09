-- Adminer 4.8.0 MySQL 5.5.5-10.7.3-MariaDB dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

CREATE DATABASE `db_task` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;
USE `db_task`;

DROP TABLE IF EXISTS `tbl_tasking`;
CREATE TABLE `tbl_tasking` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Task` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Assignee` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `Deadline` date NOT NULL,
  `Status` enum('open','cancel','close') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'open',
  `Note` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `tbl_tasking` (`Id`, `Task`, `Assignee`, `Deadline`, `Status`, `Note`) VALUES
(1,	'Membuat Database dan table product di mysql',	'uuh',	'2022-08-09',	'close',	''),
(2,	'Input data product dan Customer',	'Yuri',	'2022-08-09',	'close',	''),
(4,	'membuat web sakti ',	'yuyu',	'2022-09-07',	'close',	NULL),
(5,	'ayo kita buat baru us yayayaa',	'karedok',	'2022-09-01',	'cancel',	NULL),
(18,	'wqeqeq',	'eqweqweq',	'2022-09-29',	'close',	NULL),
(19,	'aduh',	'aa',	'2022-09-02',	'close',	NULL),
(20,	'qweqweq',	'weqwe',	'2022-08-31',	'cancel',	NULL);

-- 2022-08-09 09:12:37
