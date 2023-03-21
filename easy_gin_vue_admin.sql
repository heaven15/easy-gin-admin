/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : localhost:3306
 Source Schema         : easy_gin_vue_admin

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 21/03/2023 12:16:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '配置ID',
  `real_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '配置名称',
  `app_key` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '配置标识',
  `app_val` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '配置值',
  `permission_code` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '权限标识',
  `remark` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建用户',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `real_name`(`real_name`) USING BTREE,
  UNIQUE INDEX `app_key`(`app_key`) USING BTREE,
  INDEX `idx_real_name`(`real_name`) USING BTREE,
  INDEX `idx_app_key`(`app_key`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_config
-- ----------------------------

-- ----------------------------
-- Table structure for sys_department
-- ----------------------------
DROP TABLE IF EXISTS `sys_department`;
CREATE TABLE `sys_department`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '部门ID',
  `parent_id` bigint(20) NULL DEFAULT NULL COMMENT '父类部门ID',
  `real_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '部门名称',
  `full_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '部门全称',
  `is_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '部门类型：1公司，2子公司，3.部门，4.小组',
  `remark` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `sort` int(11) NOT NULL COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用，2禁用',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建用户',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `real_name`(`real_name`) USING BTREE,
  INDEX `idx_real_name`(`real_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_department
-- ----------------------------

-- ----------------------------
-- Table structure for sys_department_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_department_user`;
CREATE TABLE `sys_department_user`  (
  `sys_department_id` bigint(20) NOT NULL COMMENT '部门ID',
  `sys_user_id` bigint(20) NOT NULL COMMENT '用户ID',
  PRIMARY KEY (`sys_department_id`, `sys_user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_department_user
-- ----------------------------

-- ----------------------------
-- Table structure for sys_dictionary
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary`;
CREATE TABLE `sys_dictionary`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典ID',
  `real_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典名称',
  `code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典标识',
  `is_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典类型',
  `remark` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用，2禁用',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建用户',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `code`(`code`) USING BTREE,
  INDEX `idx_code`(`code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dictionary
-- ----------------------------

-- ----------------------------
-- Table structure for sys_element
-- ----------------------------
DROP TABLE IF EXISTS `sys_element`;
CREATE TABLE `sys_element`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '页面元素ID',
  `real_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '页面名称',
  `code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '页面元素标识',
  `page_id` bigint(20) NULL DEFAULT NULL COMMENT '页面ID',
  `remark` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用，2禁用',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建用户',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `code`(`code`) USING BTREE,
  INDEX `idx_code`(`code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_element
-- ----------------------------

-- ----------------------------
-- Table structure for sys_element_permissions
-- ----------------------------
DROP TABLE IF EXISTS `sys_element_permissions`;
CREATE TABLE `sys_element_permissions`  (
  `sys_element_id` bigint(20) NOT NULL COMMENT '页面元素ID',
  `sys_permission_id` bigint(20) NOT NULL COMMENT '权限ID',
  PRIMARY KEY (`sys_element_id`, `sys_permission_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_element_permissions
-- ----------------------------

-- ----------------------------
-- Table structure for sys_file
-- ----------------------------
DROP TABLE IF EXISTS `sys_file`;
CREATE TABLE `sys_file`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文件ID',
  `real_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件名称',
  `path` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件路径',
  `remark` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用，2禁用',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建用户',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_file
-- ----------------------------

-- ----------------------------
-- Table structure for sys_file_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_file_permission`;
CREATE TABLE `sys_file_permission`  (
  `sys_file_id` bigint(20) NOT NULL COMMENT '文件ID',
  `sys_permission_id` bigint(20) NOT NULL COMMENT '权限ID',
  PRIMARY KEY (`sys_file_id`, `sys_permission_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_file_permission
-- ----------------------------

-- ----------------------------
-- Table structure for sys_group
-- ----------------------------
DROP TABLE IF EXISTS `sys_group`;
CREATE TABLE `sys_group`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户组ID',
  `parent_id` bigint(20) NULL DEFAULT NULL COMMENT '父类用户组ID',
  `real_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户组名称',
  `remark` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `sort` int(11) NOT NULL COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用，2禁用',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建用户',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_group
-- ----------------------------

-- ----------------------------
-- Table structure for sys_group_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_group_role`;
CREATE TABLE `sys_group_role`  (
  `sys_group_id` bigint(20) NOT NULL COMMENT '用户组ID',
  `sys_role_id` bigint(20) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_group_id`, `sys_role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_group_role
-- ----------------------------

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `parent_id` bigint(20) NULL DEFAULT NULL COMMENT '父类菜单ID',
  `real_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `redirect` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '重定向',
  `page_id` bigint(20) NULL DEFAULT NULL COMMENT '页面ID',
  `meta` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '元数据',
  `sort` int(11) NOT NULL COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用，2禁用',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建用户',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `real_name`(`real_name`) USING BTREE,
  INDEX `idx_real_name`(`real_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------

-- ----------------------------
-- Table structure for sys_operation
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation`;
CREATE TABLE `sys_operation`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '功能操作ID',
  `parent_id` bigint(20) NULL DEFAULT NULL COMMENT '父类功能操作ID',
  `real_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '功能操作名称',
  `code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '功能操作标识',
  `url` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '拦截URL前缀',
  `remark` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用，2禁用',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建用户',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `real_name`(`real_name`) USING BTREE,
  UNIQUE INDEX `code`(`code`) USING BTREE,
  INDEX `idx_real_name`(`real_name`) USING BTREE,
  INDEX `idx_code`(`code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_operation
-- ----------------------------

-- ----------------------------
-- Table structure for sys_operation_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_permission`;
CREATE TABLE `sys_operation_permission`  (
  `sys_operation_id` bigint(20) NOT NULL COMMENT '功能操作ID',
  `sys_permission_id` bigint(20) NOT NULL COMMENT '权限ID',
  PRIMARY KEY (`sys_operation_id`, `sys_permission_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_operation_permission
-- ----------------------------

-- ----------------------------
-- Table structure for sys_page
-- ----------------------------
DROP TABLE IF EXISTS `sys_page`;
CREATE TABLE `sys_page`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '页面ID',
  `real_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '页面名称',
  `url` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件路径',
  `sort` int(11) NOT NULL COMMENT '排序',
  `is_type` tinyint(4) NOT NULL COMMENT '页面类型：1.菜单 2.iframe 3.外链 4.按钮',
  `remark` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `permission_id` bigint(20) NULL DEFAULT NULL COMMENT '权限ID',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用，2禁用',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建用户',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `real_name`(`real_name`) USING BTREE,
  INDEX `idx_real_name`(`real_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_page
-- ----------------------------

-- ----------------------------
-- Table structure for sys_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_permission`;
CREATE TABLE `sys_permission`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '权限ID',
  `real_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '权限名称',
  `code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '权限标识',
  `remark` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用，2禁用',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建用户',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `real_name`(`real_name`) USING BTREE,
  UNIQUE INDEX `code`(`code`) USING BTREE,
  INDEX `idx_real_name`(`real_name`) USING BTREE,
  INDEX `idx_code`(`code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_permission
-- ----------------------------

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `real_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '岗位名称',
  `sort` int(11) NOT NULL COMMENT '排序',
  `remark` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用，2禁用',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建用户',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `real_name`(`real_name`) USING BTREE,
  INDEX `idx_real_name`(`real_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_post
-- ----------------------------

-- ----------------------------
-- Table structure for sys_rank
-- ----------------------------
DROP TABLE IF EXISTS `sys_rank`;
CREATE TABLE `sys_rank`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '职级ID',
  `real_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '职级名称',
  `sort` int(11) NOT NULL COMMENT '排序',
  `remark` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用，2禁用',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建用户',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `real_name`(`real_name`) USING BTREE,
  INDEX `idx_real_name`(`real_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_rank
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `real_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色标识',
  `remark` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `sort` int(11) NOT NULL COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用，2禁用',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建用户',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `real_name`(`real_name`) USING BTREE,
  UNIQUE INDEX `code`(`code`) USING BTREE,
  INDEX `idx_real_name`(`real_name`) USING BTREE,
  INDEX `idx_code`(`code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_permission`;
CREATE TABLE `sys_role_permission`  (
  `sys_role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `sys_permission_id` bigint(20) NOT NULL COMMENT '权限ID',
  PRIMARY KEY (`sys_role_id`, `sys_permission_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_permission
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `account_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '账号名称',
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户姓名',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户昵称',
  `password` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `gender` tinyint(1) NOT NULL DEFAULT 1 COMMENT '性别：1男，2女，3保密',
  `birthday` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '出生日期',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱',
  `mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户手机号',
  `salt` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '干扰码',
  `post_id` bigint(20) NOT NULL COMMENT '岗位ID',
  `rank_id` bigint(20) NOT NULL COMMENT '职级ID',
  `group_id` bigint(20) NOT NULL COMMENT '用户分组ID',
  `avatar` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '头像',
  `city` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '城市json数据',
  `address` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '地址详情',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态：1启用，2禁用',
  `ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'ip地址',
  `remark` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
  `create_user` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建用户',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `account_name`(`account_name`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE,
  INDEX `idx_account_name`(`account_name`) USING BTREE,
  INDEX `idx_username`(`username`) USING BTREE,
  INDEX `idx_mobile`(`mobile`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', '18079160220', '超级管理员', 'c567ae329f9929b518759d3bea13f492', 1, '0', '1060236395@qq.com', '18079160220', 'f9aZTAa8yz', 1, 1, 1, 'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-07-19/ccwpeuqz1i2s769hua.jpeg', NULL, ' ', 1, '127.0.0.1', ' ', 0, NULL, NULL);

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `sys_user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `sys_role_id` bigint(20) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`, `sys_role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
