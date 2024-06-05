/*
 Navicat MySQL Data Transfer

 Source Server         : localhost-root
 Source Server Type    : MySQL
 Source Server Version : 100121
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 100121
 File Encoding         : 65001

 Date: 20/09/2022 10:19:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `tag_id` int UNSIGNED NULL DEFAULT 0 COMMENT '标签ID',
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '简述',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '内容',
  `cover_image_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '封面图片地址',
  `created_on` int UNSIGNED NULL DEFAULT 0 COMMENT '新建时间',
  `created_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int UNSIGNED NULL DEFAULT 0,
  `state` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文章管理' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of blog_article
-- ----------------------------
INSERT INTO `blog_article` VALUES (1, 1, '测试添加文章', '文章描述', '文章内容', '', 0, '张聪', 0, '', 0, 1);
INSERT INTO `blog_article` VALUES (2, 1, '11111', '文章描述', '文章内容', 'http://127.0.0.1:8000/upload/images/4e90a70ebc3e2555fe607efde66bc563.jpg', 1648036456, '张聪', 1648036456, '', 0, 1);
INSERT INTO `blog_article` VALUES (3, 1, '11111', '文章描述', '文章内容', 'http://127.0.0.1:8000/upload/images/4e90a70ebc3e2555fe607efde66bc563.jpg', 1648036547, '张聪', 1648036547, '', 0, 1);

-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '账号',
  `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of blog_auth
-- ----------------------------
INSERT INTO `blog_auth` VALUES (1, 'test', 'test123');

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '标签名称',
  `created_on` int UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
  `created_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建人',
  `modified_on` int UNSIGNED NULL DEFAULT 0 COMMENT '修改时间',
  `modified_by` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `state` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文章标签管理' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of blog_tag
-- ----------------------------
INSERT INTO `blog_tag` VALUES (1, '标签a', 1614335202, '张聪', 1646962575, 'zc1', 0, 1);
INSERT INTO `blog_tag` VALUES (2, '标签2', 1614332202, '张聪1', 0, '张聪1', 0, 1);
INSERT INTO `blog_tag` VALUES (3, 'a', 0, 'zc', 0, '', 0, 1);
INSERT INTO `blog_tag` VALUES (4, 'b', 0, 'zc', 0, '', 0, 1);
INSERT INTO `blog_tag` VALUES (5, 'd', 1648208703, 'zc', 1648208703, '', 0, 1);
INSERT INTO `blog_tag` VALUES (6, '666666', 1648215198, 'zc', 1648215198, '', 0, 1);
INSERT INTO `blog_tag` VALUES (7, '77777', 1648215198, '', 1648215198, '', 0, 1);
INSERT INTO `blog_tag` VALUES (8, '88888', 1648215198, '', 1648215198, '', 0, 1);

SET FOREIGN_KEY_CHECKS = 1;
