/*
Navicat MySQL Data Transfer

Source Server         : 开发22
Source Server Version : 80017
Source Host           : 172.20.1.22:3306
Source Database       : dev_tools

Target Server Type    : MYSQL
Target Server Version : 80017
File Encoding         : 65001

Date: 2019-12-11 15:21:40
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for d_db_owner
-- ----------------------------
DROP TABLE IF EXISTS `d_db_owner`;
CREATE TABLE `d_db_owner` (
  `id` int(3) unsigned NOT NULL AUTO_INCREMENT COMMENT '数据库编号',
  `db_name` varchar(30) NOT NULL COMMENT '数据库名',
  `owner_name` varchar(30) NOT NULL COMMENT '负责人',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='数据库责任人表';

-- ----------------------------
-- Records of d_db_owner
-- ----------------------------
INSERT INTO `d_db_owner` VALUES ('1', 'erpdb', '刘能凯', '2019-12-07 22:16:27', null);

-- ----------------------------
-- Table structure for d_dev_db_change
-- ----------------------------
DROP TABLE IF EXISTS `d_dev_db_change`;
CREATE TABLE `d_dev_db_change` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '数据库变更编号(开发)',
  `user_id` int(5) unsigned NOT NULL COMMENT '变更者',
  `db_id` int(3) unsigned NOT NULL COMMENT '变更哪个数据库',
  `db_item` varchar(100) NOT NULL COMMENT '数据库变更内容',
  `db_reason` varchar(50) NOT NULL COMMENT '变更缘由',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '发布状态(1:缺省，未发布，2:发布到测试)',
  `create_time` datetime NOT NULL COMMENT '变更时间',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `db_id` (`db_id`),
  CONSTRAINT `d_dev_db_change_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `d_user` (`id`),
  CONSTRAINT `d_dev_db_change_ibfk_2` FOREIGN KEY (`db_id`) REFERENCES `d_db_owner` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='开发数据库变更';

-- ----------------------------
-- Records of d_dev_db_change
-- ----------------------------

-- ----------------------------
-- Table structure for d_online_publish_info
-- ----------------------------
DROP TABLE IF EXISTS `d_online_publish_info`;
CREATE TABLE `d_online_publish_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `test_id` int(10) unsigned NOT NULL COMMENT '发布到测试变更id',
  `online_id` int(10) unsigned NOT NULL COMMENT '发布到线上id',
  PRIMARY KEY (`id`),
  KEY `test_id` (`test_id`),
  KEY `online_id` (`online_id`),
  CONSTRAINT `d_online_publish_info_ibfk_1` FOREIGN KEY (`test_id`) REFERENCES `d_test_publish_info` (`id`),
  CONSTRAINT `d_online_publish_info_ibfk_2` FOREIGN KEY (`online_id`) REFERENCES `d_online_publish_version` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='发布到线上详情表';

-- ----------------------------
-- Records of d_online_publish_info
-- ----------------------------

-- ----------------------------
-- Table structure for d_online_publish_version
-- ----------------------------
DROP TABLE IF EXISTS `d_online_publish_version`;
CREATE TABLE `d_online_publish_version` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `version_id` varchar(50) NOT NULL COMMENT '线上版本号',
  `user_id` int(5) unsigned NOT NULL COMMENT '发布者',
  `publish_time` datetime NOT NULL COMMENT '发布时间',
  PRIMARY KEY (`id`),
  KEY `version_id` (`version_id`) USING BTREE,
  KEY `user_id` (`user_id`),
  CONSTRAINT `d_online_publish_version_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `d_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='线上发布版本表';

-- ----------------------------
-- Records of d_online_publish_version
-- ----------------------------

-- ----------------------------
-- Table structure for d_test_publish_info
-- ----------------------------
DROP TABLE IF EXISTS `d_test_publish_info`;
CREATE TABLE `d_test_publish_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '发布状态(1:缺省，未发布，2:发布到测试)',
  `dev_id` int(10) unsigned NOT NULL COMMENT '开发数据库变动内容id',
  `test_id` int(10) unsigned NOT NULL COMMENT '发布到测试详情',
  PRIMARY KEY (`id`),
  KEY `dev_id` (`dev_id`),
  KEY `test_id` (`test_id`),
  CONSTRAINT `d_test_publish_info_ibfk_1` FOREIGN KEY (`dev_id`) REFERENCES `d_dev_db_change` (`id`),
  CONSTRAINT `d_test_publish_info_ibfk_2` FOREIGN KEY (`test_id`) REFERENCES `d_test_publish_version` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='发布到测试详情表';

-- ----------------------------
-- Records of d_test_publish_info
-- ----------------------------

-- ----------------------------
-- Table structure for d_test_publish_version
-- ----------------------------
DROP TABLE IF EXISTS `d_test_publish_version`;
CREATE TABLE `d_test_publish_version` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `version_id` varchar(50) NOT NULL COMMENT '测试版本号',
  `user_id` int(5) unsigned NOT NULL COMMENT '发布者',
  `publish_time` datetime NOT NULL COMMENT '发布时间',
  PRIMARY KEY (`id`),
  KEY `version_id` (`version_id`) USING BTREE,
  KEY `user_id` (`user_id`),
  CONSTRAINT `d_test_publish_version_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `d_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='测试发布版本表';

-- ----------------------------
-- Records of d_test_publish_version
-- ----------------------------

-- ----------------------------
-- Table structure for d_user
-- ----------------------------
DROP TABLE IF EXISTS `d_user`;
CREATE TABLE `d_user` (
  `id` int(5) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `user_account` varchar(20) NOT NULL COMMENT '用户账号',
  `user_name` varchar(30) NOT NULL COMMENT '用户名',
  `user_password` varchar(128) NOT NULL COMMENT '密码',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `user_account` (`user_account`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';

-- ----------------------------
-- Records of d_user
-- ----------------------------
INSERT INTO `d_user` VALUES ('1', 'saya', '刘能凯', '96e79218965eb72c92a549dd5a330112', '2019-12-07 22:16:51', null);

-- ----------------------------
-- Table structure for d_user_mana_db
-- ----------------------------
DROP TABLE IF EXISTS `d_user_mana_db`;
CREATE TABLE `d_user_mana_db` (
  `id` int(5) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(5) unsigned NOT NULL COMMENT '用户id',
  `db_id` int(3) unsigned NOT NULL COMMENT '数据库id',
  PRIMARY KEY (`id`),
  KEY `db_id` (`db_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `d_user_mana_db_ibfk_1` FOREIGN KEY (`db_id`) REFERENCES `d_db_owner` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `d_user_mana_db_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `d_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户管理数据库表';

-- ----------------------------
-- Records of d_user_mana_db
-- ----------------------------
INSERT INTO `d_user_mana_db` VALUES ('1', '1', '1');
