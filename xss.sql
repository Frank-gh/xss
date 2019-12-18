-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        10.3.15-MariaDB-log - Source distribution
-- 服务器OS:                        Linux
-- HeidiSQL 版本:                  10.2.0.5599
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for xss
CREATE DATABASE IF NOT EXISTS `xss` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `xss`;

-- Dumping structure for table xss.xss_ad
DROP TABLE IF EXISTS `xss_ad`;
CREATE TABLE IF NOT EXISTS `xss_ad` (
  `ad_position_id` int(11) NOT NULL DEFAULT 0,
  `content` varchar(255) NOT NULL DEFAULT '',
  `enabled` int(11) NOT NULL DEFAULT 0,
  `end_time` int(11) NOT NULL DEFAULT 0,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `image_url` varchar(255) NOT NULL DEFAULT '',
  `link` varchar(255) NOT NULL DEFAULT '',
  `media_type` int(11) NOT NULL DEFAULT 0,
  `name` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table xss.xss_ad: ~3 rows (大约)
/*!40000 ALTER TABLE `xss_ad` DISABLE KEYS */;
INSERT INTO `xss_ad` (`ad_position_id`, `content`, `enabled`, `end_time`, `id`, `image_url`, `link`, `media_type`, `name`) VALUES
	(1, '合作 谁是你的菜', 1, 0, 1, 'http://www.dtworm.com/images/banner01.jpg', '', 1, '合作 谁是你的菜'),
	(1, '活动 美食节', 1, 0, 2, 'http://www.dtworm.com/images/banner02.jpg', '', 1, '活动 美食节'),
	(1, '活动 母亲节', 1, 0, 3, 'http://www.dtworm.com/images/banner03.jpg', '', 1, '活动 母亲节');
/*!40000 ALTER TABLE `xss_ad` ENABLE KEYS */;

-- Dumping structure for table xss.xss_channel
DROP TABLE IF EXISTS `xss_channel`;
CREATE TABLE IF NOT EXISTS `xss_channel` (
  `icon_url` varchar(255) NOT NULL DEFAULT '',
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `sort_order` int(11) NOT NULL DEFAULT 0,
  `url` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table xss.xss_channel: ~5 rows (大约)
/*!40000 ALTER TABLE `xss_channel` DISABLE KEYS */;
INSERT INTO `xss_channel` (`icon_url`, `id`, `name`, `sort_order`, `url`) VALUES
	('http://www.dtworm.com/images/service01.png', 1, '上衣', 1, '/pages/category/category?id=1005000'),
	('http://www.dtworm.com/images/service02.png', 2, '裤子', 2, '/pages/category/category?id=1005001'),
	('http://www.dtworm.com/images/service03.png', 3, '鞋子', 3, '/pages/category/category?id=1008000'),
	('http://www.dtworm.com/images/service04.png', 4, '包', 4, '/pages/category/category?id=1005002'),
	('http://www.dtworm.com/images/service05.png', 5, '奢侈品', 5, '/pages/category/category?id=1019000');
/*!40000 ALTER TABLE `xss_channel` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
