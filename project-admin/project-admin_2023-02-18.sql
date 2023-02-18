# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 172.16.10.183 (MySQL 5.7.9-log)
# Database: project-admin
# Generation Time: 2023-02-18 09:57:15 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table application
# ------------------------------------------------------------

DROP TABLE IF EXISTS `application`;

CREATE TABLE `application` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `zh_name` char(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '中文名称',
  `en_name` char(20) COLLATE utf8mb4_bin NOT NULL COMMENT '英文名称，相当于程序名称',
  `ico` char(200) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '图标',
  `info` char(200) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '简介',
  `create_user` bigint(20) NOT NULL COMMENT '创建者id',
  `demand_ids` char(200) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '需求组ids',
  `doc_ids` char(200) COLLATE utf8mb4_bin NOT NULL COMMENT '文档组ids',
  `join_users` char(200) COLLATE utf8mb4_bin NOT NULL COMMENT '参与者ids',
  `join_groups` char(200) COLLATE utf8mb4_bin NOT NULL COMMENT '参与组ids',
  `project_id` char(200) COLLATE utf8mb4_bin NOT NULL COMMENT '所属项目id',
  `remark` varchar(250) COLLATE utf8mb4_bin NOT NULL COMMENT '备注',
  `rank` tinyint(100) NOT NULL DEFAULT '1' COMMENT '排序',
  `state` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，-2删除，-1禁用，待审核0，启用1',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_unique` (`create_user`,`en_name`,`project_id`),
  KEY `index_filter` (`create_time`,`state`,`zh_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='应用表';

LOCK TABLES `application` WRITE;
/*!40000 ALTER TABLE `application` DISABLE KEYS */;

INSERT INTO `application` (`id`, `zh_name`, `en_name`, `ico`, `info`, `create_user`, `demand_ids`, `doc_ids`, `join_users`, `join_groups`, `project_id`, `remark`, `rank`, `state`, `create_time`, `update_time`)
VALUES
	(149099121239,X'E9A1B9E79BAE3234',X'70726F6A656374',X'68747470733A2F2F706963782E7A68696D672E636F6D2F76322D65346463626434323630646137356461653533313765653435373332663361655F6C2E6A70673F736F757263653D3137326165313862',X'E794A8E4BA8EE8829DE88386E883B0E8BE85E58AA9E8AF8AE696ADE69CBAE599A8E4BABAE9A1B9E79BAE',1112212,X'323334343434',X'',X'3133343132342C',X'3133343231342C',X'3334353334353233',X'677364676473',1,1,'2022-12-05 18:02:40','2022-12-14 15:28:59'),
	(334923565655,X'E9A1B9E79BAE3235',X'70726F6A6563743235',X'68747470733A2F2F706963782E7A68696D672E636F6D2F76322D65346463626434323630646137356461653533313765653435373332663361655F6C2E6A70673F736F757263653D3137326165313862',X'E794A8E4BA8EE8829DE88386E883B0E8BE85E58AA9E8AF8AE696ADE69CBAE599A8E4BABAE9A1B9E79BAE',1112212,X'323334343434',X'',X'3133343132342C',X'3133343231342C',X'3334353334353233',X'677364676473',1,1,'2022-12-05 18:32:40','2022-12-05 18:32:40'),
	(248590979041879,X'E6B58BE8AF9531',X'7465737431',X'67686A67686A',X'6A68676A676A',1489111,X'67686A67686A',X'67686A67686A',X'67686A67686A',X'67686A67686A',X'31',X'67686A67686A',1,1,'2023-02-04 09:49:33','2023-02-04 09:49:33'),
	(248606615407191,X'E6B58BE8AF9532',X'7465737432',X'67686A67686A',X'6A68676A676A',1489111,X'67686A67686A',X'67686A67686A',X'67686A67686A',X'67686A67686A',X'31',X'67686A67686A',1,1,'2023-02-04 09:49:43','2023-02-04 16:23:37'),
	(248621547129431,X'E6B58BE8AF9533',X'7465737433',X'67686A67686A',X'6A68676A676A',1489111,X'67686A67686A',X'67686A67686A',X'67686A67686A',X'67686A67686A',X'31',X'67686A67686A',1,1,'2023-02-04 09:49:52','2023-02-04 16:23:36'),
	(248647417596503,X'E6B58BE8AF9534',X'7465737434',X'67686A67686A',X'6A68676A676A',1489111,X'67686A67686A',X'67686A67686A',X'67686A67686A',X'67686A67686A',X'31',X'67686A67686A',1,1,'2023-02-04 09:50:07','2023-02-04 16:23:36'),
	(248704258804311,X'E6B58BE8AF9535',X'7465737435',X'67686A67686A',X'6A68676A676A',1489111,X'67686A67686A',X'67686A67686A',X'67686A67686A',X'67686A67686A',X'31',X'67686A67686A',1,1,'2023-02-04 09:50:41','2023-02-04 16:23:35'),
	(249171621710423,X'E6B58BE8AF9536',X'7465737436',X'67686A67686A',X'6A68676A676A',1489111,X'67686A67686A',X'67686A67686A',X'67686A67686A',X'67686A67686A',X'31',X'67686A67686A',1,1,'2023-02-04 09:55:19','2023-02-04 16:23:35'),
	(249195898341975,X'E6B58BE8AF9537',X'7465737437',X'67686A67686A',X'6A68676A676A',1489111,X'67686A67686A',X'67686A67686A',X'67686A67686A',X'67686A67686A',X'31',X'67686A67686A',1,1,'2023-02-04 09:55:34','2023-02-04 16:23:34'),
	(249291041933911,X'E6B58BE8AF9538',X'7465737438',X'67686A67686A',X'6A68676A676A',1489111,X'67686A67686A',X'67686A67686A',X'67686A67686A',X'67686A67686A',X'31',X'67686A67686A',1,1,'2023-02-04 09:56:31','2023-02-04 16:23:33'),
	(249311040375383,X'E6B58BE8AF9539',X'7465737439',X'67686A67686A',X'6A68676A676A',1489111,X'67686A67686A',X'67686A67686A',X'67686A67686A',X'67686A67686A',X'31',X'67686A67686A',1,1,'2023-02-04 09:56:43','2023-02-04 16:23:33'),
	(288315450329687,X'E6B58BE8AF953130',X'746573743130',X'67686A67686A',X'6A68676A676A',1489111,X'67686A67686A',X'67686A67686A',X'67686A67686A',X'67686A67686A',X'31',X'67686A67686A',1,1,'2023-02-04 16:24:11','2023-02-04 16:24:11'),
	(288327395707479,X'E6B58BE8AF953131',X'746573743131',X'31343839313131',X'31343839313131',1489111,X'67686A67686A',X'67686A67686A',X'67686A67686A',X'67686A67686A',X'31',X'67686A67686A',1,1,'2023-02-04 16:24:18','2023-02-17 16:42:52');

/*!40000 ALTER TABLE `application` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table application_info
# ------------------------------------------------------------

DROP TABLE IF EXISTS `application_info`;

CREATE TABLE `application_info` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `create_user` bigint(20) NOT NULL COMMENT '所属用户',
  `application_id` bigint(20) NOT NULL COMMENT '应用id',
  `version` char(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '版本号',
  `config_ids` varchar(500) COLLATE utf8mb4_bin NOT NULL COMMENT '配置ids ","号分隔',
  `state` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，-2删除，-1禁用，待审核0，启用1',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_unique` (`create_user`,`application_id`,`version`),
  KEY `index_filter` (`create_time`,`state`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='应用附属表';

LOCK TABLES `application_info` WRITE;
/*!40000 ALTER TABLE `application_info` DISABLE KEYS */;

INSERT INTO `application_info` (`id`, `create_user`, `application_id`, `version`, `config_ids`, `state`, `create_time`, `update_time`)
VALUES
	(1111,111,23,X'302E312E31',X'312C322C33',1,'2022-12-13 15:40:39','2022-12-13 15:41:42'),
	(2326797744081495,1,1,X'76312E322E33',X'323232',1,'2023-02-18 17:54:41','2023-02-18 17:54:41');

/*!40000 ALTER TABLE `application_info` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table config
# ------------------------------------------------------------

DROP TABLE IF EXISTS `config`;

CREATE TABLE `config` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `key` char(100) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'key',
  `value` text COLLATE utf8mb4_bin NOT NULL COMMENT 'value',
  `state` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，-2删除，-1禁用，待审核0，启用1',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_unique` (`user_id`,`key`),
  KEY `index_filter` (`create_time`,`state`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='配置表';

LOCK TABLES `config` WRITE;
/*!40000 ALTER TABLE `config` DISABLE KEYS */;

INSERT INTO `config` (`id`, `user_id`, `key`, `value`, `state`, `create_time`, `update_time`)
VALUES
	(0,11,X'31',X'6666',1,'2022-12-13 15:41:00','2022-12-13 15:41:00'),
	(1112,4,X'34',X'797979',1,'2022-12-13 15:41:57','2022-12-13 15:41:57'),
	(33111,22,X'32',X'6868',1,'2022-12-13 15:41:21','2022-12-13 15:41:21'),
	(345345,3,X'33',X'6D6A6B2C',1,'2022-12-13 15:41:35','2022-12-13 15:41:35');

/*!40000 ALTER TABLE `config` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table doc
# ------------------------------------------------------------

DROP TABLE IF EXISTS `doc`;

CREATE TABLE `doc` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `name` char(50) COLLATE utf8mb4_bin NOT NULL COMMENT '文档标题',
  `create_user` int(20) NOT NULL COMMENT '所属用户',
  `pre_content` text COLLATE utf8mb4_bin NOT NULL COMMENT '编辑内容',
  `content` text COLLATE utf8mb4_bin NOT NULL COMMENT '文档内容',
  `parent_doc` int(20) NOT NULL COMMENT '上级文档',
  `group_id` int(20) NOT NULL COMMENT '所属文档组',
  `sort` int(20) NOT NULL COMMENT '排序',
  `editor_mode` tinyint(4) NOT NULL COMMENT '编辑器模式,1表示Editormd编辑器，2表示Vditor编辑器，3表示iceEditor编辑器',
  `open_children` tinyint(4) NOT NULL COMMENT '展开下级目录',
  `show_children` tinyint(4) NOT NULL COMMENT '显示下级文档',
  `state` tinyint(4) NOT NULL DEFAULT '1' COMMENT '文档状态，-2删除，-1禁用，待审核-草稿0，启用1',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_unique` (`name`,`create_user`),
  KEY `index_filter` (`create_time`,`state`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='文档表';



# Dump of table doc_history
# ------------------------------------------------------------

DROP TABLE IF EXISTS `doc_history`;

CREATE TABLE `doc_history` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `pre_content` text COLLATE utf8mb4_bin NOT NULL COMMENT '编辑内容',
  `create_user` int(20) NOT NULL COMMENT '所属用户',
  `doc_id` int(20) NOT NULL COMMENT '文档id',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `index_filter` (`create_time`,`create_user`,`doc_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='文档历史表';

LOCK TABLES `doc_history` WRITE;
/*!40000 ALTER TABLE `doc_history` DISABLE KEYS */;

INSERT INTO `doc_history` (`id`, `pre_content`, `create_user`, `doc_id`, `create_time`)
VALUES
	(2326817675414103,X'6A6A686A6B6A',1,1,'2023-02-18 17:54:53');

/*!40000 ALTER TABLE `doc_history` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table group
# ------------------------------------------------------------

DROP TABLE IF EXISTS `group`;

CREATE TABLE `group` (
  `id` bigint(20) unsigned NOT NULL COMMENT '自增主键',
  `name` char(20) COLLATE utf8mb4_bin NOT NULL COMMENT '组名称',
  `create_user` bigint(20) NOT NULL COMMENT '创建者id',
  `ico` char(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '组图标',
  `remark` varchar(250) COLLATE utf8mb4_bin NOT NULL COMMENT '备注',
  `parent_id` bigint(20) NOT NULL COMMENT '父级id',
  `group_type` tinyint(100) NOT NULL DEFAULT '1' COMMENT '类型: 1部门、2用户组、3群组、4圈子、5话题',
  `rank` tinyint(100) NOT NULL DEFAULT '1' COMMENT '排序',
  `state` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，-2删除，-1禁用，待审核0，启用1',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_unique` (`create_user`,`name`),
  KEY `index_filter` (`create_time`,`state`,`group_type`,`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='组织管理中心';

LOCK TABLES `group` WRITE;
/*!40000 ALTER TABLE `group` DISABLE KEYS */;

INSERT INTO `group` (`id`, `name`, `create_user`, `ico`, `remark`, `parent_id`, `group_type`, `rank`, `state`, `create_time`, `update_time`)
VALUES
	(2326895001602647,X'35343636353435',2,X'35363536',X'353436343536',2,1,2,1,'2023-02-18 17:55:39','2023-02-18 17:55:39');

/*!40000 ALTER TABLE `group` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table group_group_relation
# ------------------------------------------------------------

DROP TABLE IF EXISTS `group_group_relation`;

CREATE TABLE `group_group_relation` (
  `id` bigint(20) unsigned NOT NULL COMMENT '自增主键',
  `create_user` bigint(20) NOT NULL COMMENT '创建者id',
  `master_group_id` bigint(20) NOT NULL COMMENT '主组id',
  `from_group_id` bigint(20) NOT NULL COMMENT '从组id',
  `relation` tinyint(4) NOT NULL DEFAULT '1' COMMENT '关系，-1不允许，1允许读写(非自己)，2只允许读(非自己)',
  `state` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，-2删除，-1禁用，待审核0，启用1',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_unique` (`create_user`,`master_group_id`,`from_group_id`,`relation`),
  KEY `index_filter` (`create_time`,`state`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='组与组关系描述表';

LOCK TABLES `group_group_relation` WRITE;
/*!40000 ALTER TABLE `group_group_relation` DISABLE KEYS */;

INSERT INTO `group_group_relation` (`id`, `create_user`, `master_group_id`, `from_group_id`, `relation`, `state`, `create_time`, `update_time`)
VALUES
	(2326753921993303,1,2,3,2,1,'2023-02-18 17:54:15','2023-02-18 17:54:15');

/*!40000 ALTER TABLE `group_group_relation` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table project
# ------------------------------------------------------------

DROP TABLE IF EXISTS `project`;

CREATE TABLE `project` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `name` char(20) COLLATE utf8mb4_bin NOT NULL COMMENT '名称',
  `ico` char(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '图标',
  `info` char(20) COLLATE utf8mb4_bin NOT NULL COMMENT '简介',
  `project_type` tinyint(100) NOT NULL DEFAULT '1' COMMENT '类型: 1编程、2其他',
  `create_user` bigint(20) NOT NULL COMMENT '创建者id',
  `join_users` char(200) COLLATE utf8mb4_bin NOT NULL COMMENT '参与者ids',
  `join_groups` char(200) COLLATE utf8mb4_bin NOT NULL COMMENT '参与组ids',
  `remark` varchar(250) COLLATE utf8mb4_bin NOT NULL COMMENT '备注',
  `rank` tinyint(100) NOT NULL DEFAULT '1' COMMENT '排序',
  `state` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，-2删除，-1禁用，待审核0，启用1',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_unique` (`create_user`,`name`,`project_type`),
  KEY `index_filter` (`create_time`,`state`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='项目表';

LOCK TABLES `project` WRITE;
/*!40000 ALTER TABLE `project` DISABLE KEYS */;

INSERT INTO `project` (`id`, `name`, `ico`, `info`, `project_type`, `create_user`, `join_users`, `join_groups`, `remark`, `rank`, `state`, `create_time`, `update_time`)
VALUES
	(2173772538841687,X'E58CBBE79697E69CBAE599A8E4BABA',X'445644E9A696E983BDE58F91E7949FE79A84E697B6E4BBA3',X'445644E9A696E983BDE58F91E7949FE79A84E697B6E4BBA3',1,4,X'31313131',X'313131',X'445644E9A696E983BDE58F91E7949FE79A84E697B6E4BBA3',1,1,'2023-02-17 16:34:31','2023-02-17 16:34:31'),
	(2178297119115863,X'E4B99FE4B880E6A0B7',X'E789B9E591B5E591B5',X'E8AEA9E4BB96E5928CE5A5B9E5A682E4BD95',1,1,X'323434',X'32343234',X'34',1,-1,'2023-02-17 17:19:28','2023-02-17 17:19:32');

/*!40000 ALTER TABLE `project` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user_group
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user_group`;

CREATE TABLE `user_group` (
  `id` bigint(20) unsigned NOT NULL COMMENT '自增主键',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `group_id` bigint(20) NOT NULL COMMENT '组id',
  `state` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态，-2删除，-1禁用，待审核0，启用1',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_unique` (`user_id`,`group_id`),
  KEY `index_filter` (`create_time`,`state`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户、组关系表';

LOCK TABLES `user_group` WRITE;
/*!40000 ALTER TABLE `user_group` DISABLE KEYS */;

INSERT INTO `user_group` (`id`, `user_id`, `group_id`, `state`, `create_time`, `update_time`)
VALUES
	(2326920653965911,1,1,1,'2023-02-18 17:55:54','2023-02-18 17:55:54');

/*!40000 ALTER TABLE `user_group` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
