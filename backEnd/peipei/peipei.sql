/*
 Navicat Premium Data Transfer

 Source Server         : peipei_test
 Source Server Type    : MySQL
 Source Server Version : 80015
 Source Host           : 127.0.0.1:3306
 Source Schema         : peipei

 Target Server Type    : MySQL
 Target Server Version : 80015
 File Encoding         : 65001

 Date: 15/11/2019 18:22:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for auth_group
-- ----------------------------
DROP TABLE IF EXISTS `auth_group`;
CREATE TABLE `auth_group`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for auth_group_permissions
-- ----------------------------
DROP TABLE IF EXISTS `auth_group_permissions`;
CREATE TABLE `auth_group_permissions`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `group_id` int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `auth_group_permissions_group_id_permission_id_0cd325b0_uniq`(`group_id`, `permission_id`) USING BTREE,
  INDEX `auth_group_permissio_permission_id_84c5c92e_fk_auth_perm`(`permission_id`) USING BTREE,
  CONSTRAINT `auth_group_permissio_permission_id_84c5c92e_fk_auth_perm` FOREIGN KEY (`permission_id`) REFERENCES `auth_permission` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `auth_group_permissions_group_id_b120cbf9_fk_auth_group_id` FOREIGN KEY (`group_id`) REFERENCES `auth_group` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for auth_permission
-- ----------------------------
DROP TABLE IF EXISTS `auth_permission`;
CREATE TABLE `auth_permission`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `content_type_id` int(11) NOT NULL,
  `codename` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `auth_permission_content_type_id_codename_01ab375a_uniq`(`content_type_id`, `codename`) USING BTREE,
  CONSTRAINT `auth_permission_content_type_id_2f476e4b_fk_django_co` FOREIGN KEY (`content_type_id`) REFERENCES `django_content_type` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 57 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of auth_permission
-- ----------------------------
INSERT INTO `auth_permission` VALUES (1, 'Can add log entry', 1, 'add_logentry');
INSERT INTO `auth_permission` VALUES (2, 'Can change log entry', 1, 'change_logentry');
INSERT INTO `auth_permission` VALUES (3, 'Can delete log entry', 1, 'delete_logentry');
INSERT INTO `auth_permission` VALUES (4, 'Can view log entry', 1, 'view_logentry');
INSERT INTO `auth_permission` VALUES (5, 'Can add group', 4, 'add_group');
INSERT INTO `auth_permission` VALUES (6, 'Can change group', 4, 'change_group');
INSERT INTO `auth_permission` VALUES (7, 'Can delete group', 4, 'delete_group');
INSERT INTO `auth_permission` VALUES (8, 'Can view group', 4, 'view_group');
INSERT INTO `auth_permission` VALUES (9, 'Can add permission', 2, 'add_permission');
INSERT INTO `auth_permission` VALUES (10, 'Can change permission', 2, 'change_permission');
INSERT INTO `auth_permission` VALUES (11, 'Can delete permission', 2, 'delete_permission');
INSERT INTO `auth_permission` VALUES (12, 'Can view permission', 2, 'view_permission');
INSERT INTO `auth_permission` VALUES (13, 'Can add user', 3, 'add_user');
INSERT INTO `auth_permission` VALUES (14, 'Can change user', 3, 'change_user');
INSERT INTO `auth_permission` VALUES (15, 'Can delete user', 3, 'delete_user');
INSERT INTO `auth_permission` VALUES (16, 'Can view user', 3, 'view_user');
INSERT INTO `auth_permission` VALUES (17, 'Can add content type', 5, 'add_contenttype');
INSERT INTO `auth_permission` VALUES (18, 'Can change content type', 5, 'change_contenttype');
INSERT INTO `auth_permission` VALUES (19, 'Can delete content type', 5, 'delete_contenttype');
INSERT INTO `auth_permission` VALUES (20, 'Can view content type', 5, 'view_contenttype');
INSERT INTO `auth_permission` VALUES (21, 'Can add session', 6, 'add_session');
INSERT INTO `auth_permission` VALUES (22, 'Can change session', 6, 'change_session');
INSERT INTO `auth_permission` VALUES (23, 'Can delete session', 6, 'delete_session');
INSERT INTO `auth_permission` VALUES (24, 'Can view session', 6, 'view_session');
INSERT INTO `auth_permission` VALUES (25, 'Can add  学生阅读小说信息表', 13, 'add_novels');
INSERT INTO `auth_permission` VALUES (26, 'Can change  学生阅读小说信息表', 13, 'change_novels');
INSERT INTO `auth_permission` VALUES (27, 'Can delete  学生阅读小说信息表', 13, 'delete_novels');
INSERT INTO `auth_permission` VALUES (28, 'Can view  学生阅读小说信息表', 13, 'view_novels');
INSERT INTO `auth_permission` VALUES (29, 'Can add  学生看剧类型信息表', 12, 'add_videos');
INSERT INTO `auth_permission` VALUES (30, 'Can change  学生看剧类型信息表', 12, 'change_videos');
INSERT INTO `auth_permission` VALUES (31, 'Can delete  学生看剧类型信息表', 12, 'delete_videos');
INSERT INTO `auth_permission` VALUES (32, 'Can view  学生看剧类型信息表', 12, 'view_videos');
INSERT INTO `auth_permission` VALUES (33, 'Can add  学生游戏类型信息表', 14, 'add_games');
INSERT INTO `auth_permission` VALUES (34, 'Can change  学生游戏类型信息表', 14, 'change_games');
INSERT INTO `auth_permission` VALUES (35, 'Can delete  学生游戏类型信息表', 14, 'delete_games');
INSERT INTO `auth_permission` VALUES (36, 'Can view  学生游戏类型信息表', 14, 'view_games');
INSERT INTO `auth_permission` VALUES (37, 'Can add  学生运动信息表', 11, 'add_sports');
INSERT INTO `auth_permission` VALUES (38, 'Can change  学生运动信息表', 11, 'change_sports');
INSERT INTO `auth_permission` VALUES (39, 'Can delete  学生运动信息表', 11, 'delete_sports');
INSERT INTO `auth_permission` VALUES (40, 'Can view  学生运动信息表', 11, 'view_sports');
INSERT INTO `auth_permission` VALUES (41, 'Can add  分配宿舍成员信息表', 8, 'add_dormitory');
INSERT INTO `auth_permission` VALUES (42, 'Can change  分配宿舍成员信息表', 8, 'change_dormitory');
INSERT INTO `auth_permission` VALUES (43, 'Can delete  分配宿舍成员信息表', 8, 'delete_dormitory');
INSERT INTO `auth_permission` VALUES (44, 'Can view  分配宿舍成员信息表', 8, 'view_dormitory');
INSERT INTO `auth_permission` VALUES (45, 'Can add 学校信息表', 7, 'add_school');
INSERT INTO `auth_permission` VALUES (46, 'Can change 学校信息表', 7, 'change_school');
INSERT INTO `auth_permission` VALUES (47, 'Can delete 学校信息表', 7, 'delete_school');
INSERT INTO `auth_permission` VALUES (48, 'Can view 学校信息表', 7, 'view_school');
INSERT INTO `auth_permission` VALUES (49, 'Can add 学校专业信息表', 9, 'add_major');
INSERT INTO `auth_permission` VALUES (50, 'Can change 学校专业信息表', 9, 'change_major');
INSERT INTO `auth_permission` VALUES (51, 'Can delete 学校专业信息表', 9, 'delete_major');
INSERT INTO `auth_permission` VALUES (52, 'Can view 学校专业信息表', 9, 'view_major');
INSERT INTO `auth_permission` VALUES (53, 'Can add  学生基本信息表', 10, 'add_student');
INSERT INTO `auth_permission` VALUES (54, 'Can change  学生基本信息表', 10, 'change_student');
INSERT INTO `auth_permission` VALUES (55, 'Can delete  学生基本信息表', 10, 'delete_student');
INSERT INTO `auth_permission` VALUES (56, 'Can view  学生基本信息表', 10, 'view_student');

-- ----------------------------
-- Table structure for auth_user
-- ----------------------------
DROP TABLE IF EXISTS `auth_user`;
CREATE TABLE `auth_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `last_login` datetime(6) NULL DEFAULT NULL,
  `is_superuser` tinyint(1) NOT NULL,
  `username` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `first_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `last_name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` varchar(254) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `is_staff` tinyint(1) NOT NULL,
  `is_active` tinyint(1) NOT NULL,
  `date_joined` datetime(6) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for auth_user_groups
-- ----------------------------
DROP TABLE IF EXISTS `auth_user_groups`;
CREATE TABLE `auth_user_groups`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `group_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `auth_user_groups_user_id_group_id_94350c0c_uniq`(`user_id`, `group_id`) USING BTREE,
  INDEX `auth_user_groups_group_id_97559544_fk_auth_group_id`(`group_id`) USING BTREE,
  CONSTRAINT `auth_user_groups_group_id_97559544_fk_auth_group_id` FOREIGN KEY (`group_id`) REFERENCES `auth_group` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `auth_user_groups_user_id_6a12ed8b_fk_auth_user_id` FOREIGN KEY (`user_id`) REFERENCES `auth_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for auth_user_user_permissions
-- ----------------------------
DROP TABLE IF EXISTS `auth_user_user_permissions`;
CREATE TABLE `auth_user_user_permissions`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `auth_user_user_permissions_user_id_permission_id_14a6b632_uniq`(`user_id`, `permission_id`) USING BTREE,
  INDEX `auth_user_user_permi_permission_id_1fbb5f2c_fk_auth_perm`(`permission_id`) USING BTREE,
  CONSTRAINT `auth_user_user_permi_permission_id_1fbb5f2c_fk_auth_perm` FOREIGN KEY (`permission_id`) REFERENCES `auth_permission` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `auth_user_user_permissions_user_id_a95ead1b_fk_auth_user_id` FOREIGN KEY (`user_id`) REFERENCES `auth_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for django_admin_log
-- ----------------------------
DROP TABLE IF EXISTS `django_admin_log`;
CREATE TABLE `django_admin_log`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `action_time` datetime(6) NOT NULL,
  `object_id` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `object_repr` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `action_flag` smallint(5) UNSIGNED NOT NULL,
  `change_message` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `content_type_id` int(11) NULL DEFAULT NULL,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `django_admin_log_content_type_id_c4bce8eb_fk_django_co`(`content_type_id`) USING BTREE,
  INDEX `django_admin_log_user_id_c564eba6_fk_auth_user_id`(`user_id`) USING BTREE,
  CONSTRAINT `django_admin_log_content_type_id_c4bce8eb_fk_django_co` FOREIGN KEY (`content_type_id`) REFERENCES `django_content_type` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `django_admin_log_user_id_c564eba6_fk_auth_user_id` FOREIGN KEY (`user_id`) REFERENCES `auth_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for django_content_type
-- ----------------------------
DROP TABLE IF EXISTS `django_content_type`;
CREATE TABLE `django_content_type`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app_label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `model` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `django_content_type_app_label_model_76bd3d3b_uniq`(`app_label`, `model`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of django_content_type
-- ----------------------------
INSERT INTO `django_content_type` VALUES (1, 'admin', 'logentry');
INSERT INTO `django_content_type` VALUES (4, 'auth', 'group');
INSERT INTO `django_content_type` VALUES (2, 'auth', 'permission');
INSERT INTO `django_content_type` VALUES (3, 'auth', 'user');
INSERT INTO `django_content_type` VALUES (5, 'contenttypes', 'contenttype');
INSERT INTO `django_content_type` VALUES (6, 'sessions', 'session');
INSERT INTO `django_content_type` VALUES (8, 'students', 'dormitory');
INSERT INTO `django_content_type` VALUES (14, 'students', 'games');
INSERT INTO `django_content_type` VALUES (9, 'students', 'major');
INSERT INTO `django_content_type` VALUES (13, 'students', 'novels');
INSERT INTO `django_content_type` VALUES (7, 'students', 'school');
INSERT INTO `django_content_type` VALUES (11, 'students', 'sports');
INSERT INTO `django_content_type` VALUES (10, 'students', 'student');
INSERT INTO `django_content_type` VALUES (12, 'students', 'videos');

-- ----------------------------
-- Table structure for django_migrations
-- ----------------------------
DROP TABLE IF EXISTS `django_migrations`;
CREATE TABLE `django_migrations`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `app` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `applied` datetime(6) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of django_migrations
-- ----------------------------
INSERT INTO `django_migrations` VALUES (1, 'contenttypes', '0001_initial', '2019-11-14 14:57:16.594021');
INSERT INTO `django_migrations` VALUES (2, 'auth', '0001_initial', '2019-11-14 14:57:16.840411');
INSERT INTO `django_migrations` VALUES (3, 'admin', '0001_initial', '2019-11-14 14:57:17.700275');
INSERT INTO `django_migrations` VALUES (4, 'admin', '0002_logentry_remove_auto_add', '2019-11-14 14:57:17.922926');
INSERT INTO `django_migrations` VALUES (5, 'admin', '0003_logentry_add_action_flag_choices', '2019-11-14 14:57:17.938885');
INSERT INTO `django_migrations` VALUES (6, 'contenttypes', '0002_remove_content_type_name', '2019-11-14 14:57:18.112157');
INSERT INTO `django_migrations` VALUES (7, 'auth', '0002_alter_permission_name_max_length', '2019-11-14 14:57:18.210893');
INSERT INTO `django_migrations` VALUES (8, 'auth', '0003_alter_user_email_max_length', '2019-11-14 14:57:18.251782');
INSERT INTO `django_migrations` VALUES (9, 'auth', '0004_alter_user_username_opts', '2019-11-14 14:57:18.271729');
INSERT INTO `django_migrations` VALUES (10, 'auth', '0005_alter_user_last_login_null', '2019-11-14 14:57:18.372969');
INSERT INTO `django_migrations` VALUES (11, 'auth', '0006_require_contenttypes_0002', '2019-11-14 14:57:18.382943');
INSERT INTO `django_migrations` VALUES (12, 'auth', '0007_alter_validators_add_error_messages', '2019-11-14 14:57:18.401891');
INSERT INTO `django_migrations` VALUES (13, 'auth', '0008_alter_user_username_max_length', '2019-11-14 14:57:18.503130');
INSERT INTO `django_migrations` VALUES (14, 'auth', '0009_alter_user_last_name_max_length', '2019-11-14 14:57:18.602864');
INSERT INTO `django_migrations` VALUES (15, 'auth', '0010_alter_group_name_max_length', '2019-11-14 14:57:18.640762');
INSERT INTO `django_migrations` VALUES (16, 'auth', '0011_update_proxy_permissions', '2019-11-14 14:57:18.656720');
INSERT INTO `django_migrations` VALUES (17, 'sessions', '0001_initial', '2019-11-14 14:57:18.696613');
INSERT INTO `django_migrations` VALUES (18, 'students', '0001_initial', '2019-11-14 14:57:19.265091');
INSERT INTO `django_migrations` VALUES (19, 'students', '0002_auto_20191115_1531', '2019-11-15 07:31:54.673984');
INSERT INTO `django_migrations` VALUES (20, 'students', '0003_auto_20191115_1615', '2019-11-15 08:15:35.233055');
INSERT INTO `django_migrations` VALUES (21, 'students', '0004_auto_20191115_1628', '2019-11-15 08:28:57.753460');
INSERT INTO `django_migrations` VALUES (22, 'students', '0005_auto_20191115_1632', '2019-11-15 08:32:20.687872');

-- ----------------------------
-- Table structure for django_session
-- ----------------------------
DROP TABLE IF EXISTS `django_session`;
CREATE TABLE `django_session`  (
  `session_key` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `session_data` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `expire_date` datetime(6) NOT NULL,
  PRIMARY KEY (`session_key`) USING BTREE,
  INDEX `django_session_expire_date_a5c62663`(`expire_date`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for students_dormitory
-- ----------------------------
DROP TABLE IF EXISTS `students_dormitory`;
CREATE TABLE `students_dormitory`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gender` tinyint(1) NOT NULL,
  `number` int(11) NOT NULL,
  `remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `addTime` datetime(6) NULL DEFAULT NULL,
  `school_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `students_dormitory_school_id_ff218711`(`school_id`) USING BTREE,
  CONSTRAINT `students_dormitory_school_id_ff218711_fk_students_school_id` FOREIGN KEY (`school_id`) REFERENCES `students_school` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for students_games
-- ----------------------------
DROP TABLE IF EXISTS `students_games`;
CREATE TABLE `students_games`  (
  `game_type` int(11) NOT NULL,
  `add_time` datetime(6) NOT NULL,
  PRIMARY KEY (`game_type`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of students_games
-- ----------------------------
INSERT INTO `students_games` VALUES (1, '2019-11-14 21:41:36.000000');
INSERT INTO `students_games` VALUES (2, '2019-11-15 15:05:09.000000');
INSERT INTO `students_games` VALUES (3, '2019-11-15 15:05:15.000000');
INSERT INTO `students_games` VALUES (4, '2019-11-15 15:05:20.000000');
INSERT INTO `students_games` VALUES (5, '2019-11-15 15:05:25.000000');
INSERT INTO `students_games` VALUES (6, '2019-11-15 15:05:31.000000');
INSERT INTO `students_games` VALUES (7, '2019-11-15 15:05:36.000000');
INSERT INTO `students_games` VALUES (8, '2019-11-15 15:05:46.000000');
INSERT INTO `students_games` VALUES (9, '2019-11-15 15:05:51.000000');
INSERT INTO `students_games` VALUES (10, '2019-11-15 15:05:56.000000');
INSERT INTO `students_games` VALUES (11, '2019-11-15 15:06:02.000000');
INSERT INTO `students_games` VALUES (12, '2019-11-15 15:06:14.000000');
INSERT INTO `students_games` VALUES (13, '2019-11-15 15:06:19.000000');
INSERT INTO `students_games` VALUES (14, '2019-11-15 15:06:25.000000');
INSERT INTO `students_games` VALUES (15, '2019-11-15 15:06:31.000000');
INSERT INTO `students_games` VALUES (16, '2019-11-15 15:32:38.000000');

-- ----------------------------
-- Table structure for students_major
-- ----------------------------
DROP TABLE IF EXISTS `students_major`;
CREATE TABLE `students_major`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `major` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `add_time` datetime(6) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of students_major
-- ----------------------------
INSERT INTO `students_major` VALUES (1, '计算机科学与技术', '2019-11-15 16:10:01.000000');

-- ----------------------------
-- Table structure for students_major_school
-- ----------------------------
DROP TABLE IF EXISTS `students_major_school`;
CREATE TABLE `students_major_school`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `major_id` int(11) NOT NULL,
  `school_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `students_major_school_major_id_school_id_807fb4cc_uniq`(`major_id`, `school_id`) USING BTREE,
  INDEX `students_major_school_school_id_3b84a530_fk_students_school_id`(`school_id`) USING BTREE,
  CONSTRAINT `students_major_school_major_id_a66ce177_fk_students_major_id` FOREIGN KEY (`major_id`) REFERENCES `students_major` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `students_major_school_school_id_3b84a530_fk_students_school_id` FOREIGN KEY (`school_id`) REFERENCES `students_school` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for students_novels
-- ----------------------------
DROP TABLE IF EXISTS `students_novels`;
CREATE TABLE `students_novels`  (
  `novel_type` int(11) NOT NULL,
  `add_time` datetime(6) NOT NULL,
  PRIMARY KEY (`novel_type`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of students_novels
-- ----------------------------
INSERT INTO `students_novels` VALUES (1, '2019-11-15 07:23:29.040567');
INSERT INTO `students_novels` VALUES (2, '2019-11-15 07:23:29.046287');
INSERT INTO `students_novels` VALUES (3, '2019-11-15 07:23:29.052270');
INSERT INTO `students_novels` VALUES (4, '2019-11-15 07:23:29.059251');
INSERT INTO `students_novels` VALUES (5, '2019-11-15 07:23:29.068038');
INSERT INTO `students_novels` VALUES (6, '2019-11-15 07:23:29.074905');
INSERT INTO `students_novels` VALUES (7, '2019-11-15 07:23:29.082317');
INSERT INTO `students_novels` VALUES (8, '2019-11-15 07:23:29.090398');
INSERT INTO `students_novels` VALUES (9, '2019-11-15 07:23:29.098398');
INSERT INTO `students_novels` VALUES (10, '2019-11-15 07:23:29.105236');
INSERT INTO `students_novels` VALUES (11, '2019-11-15 07:23:29.113780');

-- ----------------------------
-- Table structure for students_school
-- ----------------------------
DROP TABLE IF EXISTS `students_school`;
CREATE TABLE `students_school`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `school` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `province` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `city` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `remark` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `add_time` datetime(6) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `school`(`school`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of students_school
-- ----------------------------
INSERT INTO `students_school` VALUES (1, '北京交通大学威海校区', '山东', '威海', NULL, '2019-11-14 22:58:30.000000');

-- ----------------------------
-- Table structure for students_sports
-- ----------------------------
DROP TABLE IF EXISTS `students_sports`;
CREATE TABLE `students_sports`  (
  `sport_type` int(11) NOT NULL,
  `add_time` datetime(6) NOT NULL,
  PRIMARY KEY (`sport_type`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of students_sports
-- ----------------------------
INSERT INTO `students_sports` VALUES (1, '2019-11-14 21:36:58.000000');
INSERT INTO `students_sports` VALUES (2, '2019-11-15 15:21:15.000000');
INSERT INTO `students_sports` VALUES (3, '2019-11-15 07:22:43.127867');
INSERT INTO `students_sports` VALUES (4, '2019-11-15 07:22:43.134852');
INSERT INTO `students_sports` VALUES (5, '2019-11-15 07:22:43.142030');
INSERT INTO `students_sports` VALUES (6, '2019-11-15 07:22:43.148178');
INSERT INTO `students_sports` VALUES (7, '2019-11-15 07:22:43.155564');
INSERT INTO `students_sports` VALUES (8, '2019-11-15 07:22:43.161453');
INSERT INTO `students_sports` VALUES (9, '2019-11-15 07:22:43.167453');

-- ----------------------------
-- Table structure for students_student
-- ----------------------------
DROP TABLE IF EXISTS `students_student`;
CREATE TABLE `students_student`  (
  `id` int(11) NOT NULL,
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `is_super` tinyint(1) NULL DEFAULT NULL,
  `gender` tinyint(1) NULL DEFAULT NULL,
  `province` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `south_north` tinyint(1) NULL DEFAULT NULL,
  `south_north_weight` int(11) NULL DEFAULT NULL,
  `sports_type_weight` int(11) NULL DEFAULT NULL,
  `novels_type_weight` int(11) NULL DEFAULT NULL,
  `videos_type_weight` int(11) NULL DEFAULT NULL,
  `games_type_weight` int(11) NULL DEFAULT NULL,
  `winter_temperature` int(11) NULL DEFAULT NULL,
  `summer_temperature` int(11) NULL DEFAULT NULL,
  `temperature_weight` int(11) NULL DEFAULT NULL,
  `sleep_time` int(11) NULL DEFAULT NULL,
  `wake_time` int(11) NULL DEFAULT NULL,
  `sleep_weight` int(11) NULL DEFAULT NULL,
  `smoke` tinyint(1) NULL DEFAULT NULL,
  `smoke_weight` int(11) NULL DEFAULT NULL,
  `add_time` datetime(6) NULL DEFAULT NULL,
  `dormitory_id` int(11) NULL DEFAULT NULL,
  `major_id` int(11) NULL DEFAULT NULL,
  `school_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `students_student_major_id_1e02fe7d_fk_students_major_id`(`major_id`) USING BTREE,
  INDEX `students_student_dormitory_id_5ec7d6d5_fk_students_dormitory_id`(`dormitory_id`) USING BTREE,
  INDEX `students_student_school_id_be4a7ab9_fk_students_school_id`(`school_id`) USING BTREE,
  CONSTRAINT `students_student_dormitory_id_5ec7d6d5_fk_students_dormitory_id` FOREIGN KEY (`dormitory_id`) REFERENCES `students_dormitory` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `students_student_major_id_1e02fe7d_fk_students_major_id` FOREIGN KEY (`major_id`) REFERENCES `students_major` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `students_student_school_id_be4a7ab9_fk_students_school_id` FOREIGN KEY (`school_id`) REFERENCES `students_school` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of students_student
-- ----------------------------
INSERT INTO `students_student` VALUES (17722017, '彭奏章', 1, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '2019-11-14 22:59:11.000000', NULL, NULL, 1);
INSERT INTO `students_student` VALUES (17722027, '王逸润', 1, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '2019-11-14 22:58:37.000000', NULL, NULL, 1);
INSERT INTO `students_student` VALUES (17722028, '魏高林', 0, 1, '山东', 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 5, 0, 5, '2019-11-15 09:00:34.495685', NULL, 1, 1);

-- ----------------------------
-- Table structure for students_student_games_type
-- ----------------------------
DROP TABLE IF EXISTS `students_student_games_type`;
CREATE TABLE `students_student_games_type`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `student_id` int(11) NOT NULL,
  `games_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `students_student_games_type_student_id_games_id_00aebe32_uniq`(`student_id`, `games_id`) USING BTREE,
  INDEX `students_student_gam_games_id_51be53f5_fk_students_`(`games_id`) USING BTREE,
  CONSTRAINT `students_student_gam_games_id_51be53f5_fk_students_` FOREIGN KEY (`games_id`) REFERENCES `students_games` (`game_type`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `students_student_gam_student_id_45c4e500_fk_students_` FOREIGN KEY (`student_id`) REFERENCES `students_student` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of students_student_games_type
-- ----------------------------
INSERT INTO `students_student_games_type` VALUES (1, 17722028, 1);
INSERT INTO `students_student_games_type` VALUES (2, 17722028, 2);
INSERT INTO `students_student_games_type` VALUES (3, 17722028, 3);
INSERT INTO `students_student_games_type` VALUES (4, 17722028, 4);
INSERT INTO `students_student_games_type` VALUES (5, 17722028, 7);

-- ----------------------------
-- Table structure for students_student_novels_type
-- ----------------------------
DROP TABLE IF EXISTS `students_student_novels_type`;
CREATE TABLE `students_student_novels_type`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `student_id` int(11) NOT NULL,
  `novels_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `students_student_novels_type_student_id_novels_id_9da37111_uniq`(`student_id`, `novels_id`) USING BTREE,
  INDEX `students_student_nov_novels_id_1ab94d65_fk_students_`(`novels_id`) USING BTREE,
  CONSTRAINT `students_student_nov_novels_id_1ab94d65_fk_students_` FOREIGN KEY (`novels_id`) REFERENCES `students_novels` (`novel_type`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `students_student_nov_student_id_d11ab45e_fk_students_` FOREIGN KEY (`student_id`) REFERENCES `students_student` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of students_student_novels_type
-- ----------------------------
INSERT INTO `students_student_novels_type` VALUES (1, 17722028, 4);
INSERT INTO `students_student_novels_type` VALUES (2, 17722028, 5);
INSERT INTO `students_student_novels_type` VALUES (3, 17722028, 6);

-- ----------------------------
-- Table structure for students_student_sports_type
-- ----------------------------
DROP TABLE IF EXISTS `students_student_sports_type`;
CREATE TABLE `students_student_sports_type`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `student_id` int(11) NOT NULL,
  `sports_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `students_student_sports_type_student_id_sports_id_642d061d_uniq`(`student_id`, `sports_id`) USING BTREE,
  INDEX `students_student_spo_sports_id_13027bf4_fk_students_`(`sports_id`) USING BTREE,
  CONSTRAINT `students_student_spo_sports_id_13027bf4_fk_students_` FOREIGN KEY (`sports_id`) REFERENCES `students_sports` (`sport_type`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `students_student_spo_student_id_5796c500_fk_students_` FOREIGN KEY (`student_id`) REFERENCES `students_student` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of students_student_sports_type
-- ----------------------------
INSERT INTO `students_student_sports_type` VALUES (1, 17722028, 1);
INSERT INTO `students_student_sports_type` VALUES (2, 17722028, 2);
INSERT INTO `students_student_sports_type` VALUES (3, 17722028, 3);

-- ----------------------------
-- Table structure for students_student_videos_type
-- ----------------------------
DROP TABLE IF EXISTS `students_student_videos_type`;
CREATE TABLE `students_student_videos_type`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `student_id` int(11) NOT NULL,
  `videos_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `students_student_videos_type_student_id_videos_id_ffd44e40_uniq`(`student_id`, `videos_id`) USING BTREE,
  INDEX `students_student_vid_videos_id_e323029b_fk_students_`(`videos_id`) USING BTREE,
  CONSTRAINT `students_student_vid_student_id_26915963_fk_students_` FOREIGN KEY (`student_id`) REFERENCES `students_student` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `students_student_vid_videos_id_e323029b_fk_students_` FOREIGN KEY (`videos_id`) REFERENCES `students_videos` (`video_type`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of students_student_videos_type
-- ----------------------------
INSERT INTO `students_student_videos_type` VALUES (2, 17722028, 7);
INSERT INTO `students_student_videos_type` VALUES (1, 17722028, 8);

-- ----------------------------
-- Table structure for students_videos
-- ----------------------------
DROP TABLE IF EXISTS `students_videos`;
CREATE TABLE `students_videos`  (
  `video_type` int(11) NOT NULL,
  `addTime` datetime(6) NOT NULL,
  PRIMARY KEY (`video_type`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of students_videos
-- ----------------------------
INSERT INTO `students_videos` VALUES (1, '2019-11-15 07:24:43.311000');
INSERT INTO `students_videos` VALUES (2, '2019-11-15 07:24:43.318523');
INSERT INTO `students_videos` VALUES (3, '2019-11-15 07:24:43.324930');
INSERT INTO `students_videos` VALUES (4, '2019-11-15 07:24:43.331148');
INSERT INTO `students_videos` VALUES (5, '2019-11-15 07:24:43.339701');
INSERT INTO `students_videos` VALUES (6, '2019-11-15 07:24:43.345235');
INSERT INTO `students_videos` VALUES (7, '2019-11-15 07:24:43.353980');
INSERT INTO `students_videos` VALUES (8, '2019-11-15 07:24:43.360022');
INSERT INTO `students_videos` VALUES (9, '2019-11-15 07:24:43.371062');
INSERT INTO `students_videos` VALUES (10, '2019-11-15 07:24:43.378090');
INSERT INTO `students_videos` VALUES (11, '2019-11-15 07:24:43.384913');
INSERT INTO `students_videos` VALUES (12, '2019-11-15 07:24:43.392314');
INSERT INTO `students_videos` VALUES (13, '2019-11-15 07:24:43.398674');

SET FOREIGN_KEY_CHECKS = 1;
