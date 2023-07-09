/*
 Navicat MySQL Data Transfer

 Source Server         : 119.45.140.172
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : 119.45.140.172:3306
 Source Schema         : yangfan

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 09/07/2023 18:46:33
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for api_case_relationships
-- ----------------------------
DROP TABLE IF EXISTS `api_case_relationships`;
CREATE TABLE `api_case_relationships` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `api_case_step_id` bigint unsigned DEFAULT NULL COMMENT '测试步骤',
  `api_case_id` bigint unsigned DEFAULT NULL COMMENT '测试用例',
  `sort` bigint unsigned DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_api_case_relationships_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_case_relationships
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_case_step_relationships
-- ----------------------------
DROP TABLE IF EXISTS `api_case_step_relationships`;
CREATE TABLE `api_case_step_relationships` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `api_step_id` bigint unsigned DEFAULT NULL COMMENT '测试步骤',
  `api_case_step_id` bigint unsigned DEFAULT NULL COMMENT '测试步骤',
  `sort` bigint unsigned DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_api_case_step_relationships_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_case_step_relationships
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_case_steps
-- ----------------------------
DROP TABLE IF EXISTS `api_case_steps`;
CREATE TABLE `api_case_steps` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `update_by_id` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `delete_by_id` bigint unsigned DEFAULT NULL COMMENT '删除者',
  `project_id` bigint unsigned DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '步骤名称',
  `front_case` tinyint(1) DEFAULT NULL COMMENT '允许设置为前置用例',
  `run_config_id` bigint unsigned DEFAULT NULL COMMENT '运行配置',
  `run_config_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '运行配置名称',
  `api_menu_id` bigint unsigned DEFAULT NULL COMMENT '所属菜单',
  `type` bigint DEFAULT NULL COMMENT '接口类型',
  `api_step_type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '性能测试step类型',
  `api_env_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '所属环境名称',
  `api_env_id` bigint unsigned DEFAULT NULL COMMENT '所属环境',
  PRIMARY KEY (`id`),
  KEY `idx_api_case_steps_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_case_steps
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_cases
-- ----------------------------
DROP TABLE IF EXISTS `api_cases`;
CREATE TABLE `api_cases` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `update_by_id` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `delete_by_id` bigint unsigned DEFAULT NULL COMMENT '删除者',
  `project_id` bigint unsigned DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用例名称',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `describe` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  `run_number` bigint DEFAULT NULL COMMENT '运行次数',
  `RunConfigID` bigint unsigned DEFAULT NULL COMMENT '运行配置',
  `front_case` tinyint(1) DEFAULT NULL COMMENT '是否为前置用例',
  `entry_id` bigint DEFAULT NULL,
  `api_menu_id` bigint unsigned DEFAULT NULL COMMENT '所属菜单',
  `api_env_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '所属环境名称',
  `api_env_id` bigint unsigned DEFAULT NULL COMMENT '所属环境',
  PRIMARY KEY (`id`),
  KEY `idx_api_cases_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_cases
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_configs
-- ----------------------------
DROP TABLE IF EXISTS `api_configs`;
CREATE TABLE `api_configs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `update_by_id` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `delete_by_id` bigint unsigned DEFAULT NULL COMMENT '删除者',
  `project_id` bigint unsigned DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '配置名称',
  `base_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '默认域名',
  `variables` json DEFAULT NULL COMMENT '变量',
  `headers` json DEFAULT NULL COMMENT '请求头',
  `parameters` json DEFAULT NULL COMMENT '参数',
  `variables_json` json DEFAULT NULL,
  `headers_json` json DEFAULT NULL,
  `weight` bigint DEFAULT NULL,
  `default` tinyint(1) DEFAULT NULL COMMENT '默认配置',
  `timeout` float DEFAULT NULL COMMENT '超时时间',
  `allow_redirects` tinyint(1) DEFAULT NULL,
  `verify` tinyint(1) DEFAULT NULL,
  `export` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '导出参数',
  `setup_case_id` bigint unsigned DEFAULT NULL COMMENT '前置用例',
  PRIMARY KEY (`id`),
  KEY `idx_api_configs_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_configs
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_debug_talks
-- ----------------------------
DROP TABLE IF EXISTS `api_debug_talks`;
CREATE TABLE `api_debug_talks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `update_by_id` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `delete_by_id` bigint unsigned DEFAULT NULL COMMENT '删除者',
  `file_type` bigint unsigned DEFAULT NULL COMMENT '文件类型',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '文件内容',
  `project_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_api_debug_talks_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of api_debug_talks
-- ----------------------------
BEGIN;
INSERT INTO `api_debug_talks` (`id`, `created_at`, `updated_at`, `deleted_at`, `created_by_id`, `update_by_id`, `delete_by_id`, `file_type`, `content`, `project_id`) VALUES (28, '2023-01-06 18:36:47.823', '2023-01-06 18:36:47.823', NULL, 1, NULL, NULL, 1, 'import logging\nimport time\nfrom typing import List\n\n\n# commented out function will be filtered\n# def get_headers():\n#     return {\"User-Agent\": \"hrp\"}\n\n\ndef get_user_agent():\n   return \"hrp/funppy\"\n\n\ndef sleep(n_secs):\n   time.sleep(n_secs)\n\n\ndef sum(*args):\n   result = 0\n   for arg in args:\n       result += arg\n   return result\n\n\ndef sum_ints(*args: List[int]) -> int:\n   result = 0\n   for arg in args:\n       result += arg\n   return result\n\n\ndef sum_two_int(a: int, b: int) -> int:\n   return a + b\n\n\ndef sum_two_string(a: str, b: str) -> str:\n   return a + b\n\n\ndef sum_strings(*args: List[str]) -> str:\n   result = \"\"\n   for arg in args:\n       result += arg\n   return result\n\n\ndef concatenate(*args: List[str]) -> str:\n   result = \"\"\n   for arg in args:\n       result += str(arg)\n   return result\n\n\ndef setup_hook_example(name):\n   logging.warning(\"setup_hook_example\")\n   return f\"setup_hook_example: {name}\"\n\n\ndef teardown_hook_example(name):\n   logging.warning(\"teardown_hook_example\")\n   return f\"teardown_hook_example: {name}\"\n\ndef return_string():\n   return \"demo\"\n\ndef setup_hook_encryption(request):\n    request[\"body\"][\"setup_hook_encryption_request\"] = \"setup_hook_encryption_request\"\n    return request\n\n\ndef setup_hook_decrypt(response):\n    response[\"body\"][\"setup_hook_decrypt\"] = \"setup_hook_encryption_response\"\n    return response', 1);
INSERT INTO `api_debug_talks` (`id`, `created_at`, `updated_at`, `deleted_at`, `created_by_id`, `update_by_id`, `delete_by_id`, `file_type`, `content`, `project_id`) VALUES (29, '2023-01-08 01:43:27.148', '2023-01-08 01:43:27.148', NULL, NULL, NULL, NULL, 1, 'import logging\nimport time\nfrom typing import List\n\n\n# commented out function will be filtered\n# def get_headers():\n#     return {\"User-Agent\": \"hrp\"}\n\n\ndef get_user_agent():\n   return \"hrp/funppy\"\n\n\ndef sleep(n_secs):\n   time.sleep(n_secs)\n\n\ndef sum(*args):\n   result = 0\n   for arg in args:\n       result += arg\n   return result\n\n\ndef sum_ints(*args: List[int]) -> int:\n   result = 0\n   for arg in args:\n       result += arg\n   return result\n\n\ndef sum_two_int(a: int, b: int) -> int:\n   return a + b\n\n\ndef sum_two_string(a: str, b: str) -> str:\n   return a + b\n\n\ndef sum_strings(*args: List[str]) -> str:\n   result = \"\"\n   for arg in args:\n       result += arg\n   return result\n\n\ndef concatenate(*args: List[str]) -> str:\n   result = \"\"\n   for arg in args:\n       result += str(arg)\n   return result\n\n\ndef setup_hook_example(name):\n   logging.warning(\"setup_hook_example\")\n   return f\"setup_hook_example: {name}\"\n\n\ndef teardown_hook_example(name):\n   logging.warning(\"teardown_hook_example\")\n   return f\"teardown_hook_example: {name}\"\n\ndef return_string():\n   return \"demo\"\n\ndef setup_hook_encryption(request):\n    request[\"body\"][\"setup_hook_encryption_request\"] = \"setup_hook_encryption_request\"\n    return request\n\n\ndef setup_hook_decrypt(response):\n    response[\"body\"][\"setup_hook_decrypt\"] = \"setup_hook_encryption_response\"\n    return response\n\ndef setup_hook_encryption(request):\n    request[\"body\"][\"setup_hook_encryption_request\"] = \"setup_hook_encryption_request\"\n    return request\n\n\ndef setup_hook_decrypt(response):\n    response[\"body\"][\"setup_hook_decrypt\"] = \"setup_hook_encryption_response\"\n    return response', 1);
INSERT INTO `api_debug_talks` (`id`, `created_at`, `updated_at`, `deleted_at`, `created_by_id`, `update_by_id`, `delete_by_id`, `file_type`, `content`, `project_id`) VALUES (30, '2023-01-08 01:44:30.642', '2023-01-08 01:44:30.642', NULL, NULL, NULL, NULL, 1, 'import logging\nimport time\nfrom typing import List\n\n\n# commented out function will be filtered\n# def get_headers():\n#     return {\"User-Agent\": \"hrp\"}\n\n\ndef get_user_agent():\n   return \"hrp/funppy\"\n\n\ndef sleep(n_secs):\n   time.sleep(n_secs)\n\n\ndef sum(*args):\n   result = 0\n   for arg in args:\n       result += arg\n   return result\n\n\ndef sum_ints(*args: List[int]) -> int:\n   result = 0\n   for arg in args:\n       result += arg\n   return result\n\n\ndef sum_two_int(a: int, b: int) -> int:\n   return a + b\n\n\ndef sum_two_string(a: str, b: str) -> str:\n   return a + b\n\n\ndef sum_strings(*args: List[str]) -> str:\n   result = \"\"\n   for arg in args:\n       result += arg\n   return result\n\n\ndef concatenate(*args: List[str]) -> str:\n   result = \"\"\n   for arg in args:\n       result += str(arg)\n   return result\n\n\ndef setup_hook_example(name):\n   logging.warning(\"setup_hook_example\")\n   return f\"setup_hook_example: {name}\"\n\n\ndef teardown_hook_example(name):\n   logging.warning(\"teardown_hook_example\")\n   return f\"teardown_hook_example: {name}\"\n\ndef return_string():\n   return \"demo\"\n', 14);
COMMIT;

-- ----------------------------
-- Table structure for api_env_details
-- ----------------------------
DROP TABLE IF EXISTS `api_env_details`;
CREATE TABLE `api_env_details` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `update_by_id` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `delete_by_id` bigint unsigned DEFAULT NULL COMMENT '删除者',
  `project_id` bigint unsigned DEFAULT NULL,
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `value` json DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_api_env_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_env_details
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_envs
-- ----------------------------
DROP TABLE IF EXISTS `api_envs`;
CREATE TABLE `api_envs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `update_by_id` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `delete_by_id` bigint unsigned DEFAULT NULL COMMENT '删除者',
  `project_id` bigint unsigned DEFAULT NULL,
  `default` tinyint(1) DEFAULT NULL COMMENT '默认环境',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `remarks` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_api_envs_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_envs
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_grpcs
-- ----------------------------
DROP TABLE IF EXISTS `api_grpcs`;
CREATE TABLE `api_grpcs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求地址',
  `headers` json DEFAULT NULL COMMENT '请求头',
  `headers_json` json DEFAULT NULL COMMENT '请求头json数据格式',
  `body` json DEFAULT NULL COMMENT '请求体',
  `timeout` float DEFAULT NULL COMMENT '超时时间',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求类型',
  `detail` json DEFAULT NULL COMMENT '请求服务器详情',
  `parent` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_api_grpcs_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_grpcs
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_menus
-- ----------------------------
DROP TABLE IF EXISTS `api_menus`;
CREATE TABLE `api_menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `update_by_id` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `delete_by_id` bigint unsigned DEFAULT NULL COMMENT '删除者',
  `project_id` bigint unsigned DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '菜单名称',
  `parent` bigint unsigned DEFAULT NULL COMMENT '父节点id',
  `menu_type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '菜单类型',
  PRIMARY KEY (`id`),
  KEY `idx_api_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_menus
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_py_pkg
-- ----------------------------
DROP TABLE IF EXISTS `api_py_pkg`;
CREATE TABLE `api_py_pkg` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `version` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `is_uninstall` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`),
  KEY `idx_api_py_pkg_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_py_pkg
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_report_details
-- ----------------------------
DROP TABLE IF EXISTS `api_report_details`;
CREATE TABLE `api_report_details` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `case_id` bigint unsigned DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `success` tinyint(1) DEFAULT NULL,
  `stat` json DEFAULT NULL,
  `time` json DEFAULT NULL,
  `in_out` json DEFAULT NULL,
  `api_records_id` bigint unsigned DEFAULT NULL,
  `root_dir` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `api_report_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_api_report_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_report_details
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_report_details_records
-- ----------------------------
DROP TABLE IF EXISTS `api_report_details_records`;
CREATE TABLE `api_report_details_records` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `parnt_id` bigint unsigned DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `step_type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `success` tinyint(1) DEFAULT NULL,
  `elapsed_ms` bigint DEFAULT NULL,
  `validate_number` bigint unsigned DEFAULT NULL,
  `export_vars` json DEFAULT NULL,
  `content_size` bigint DEFAULT NULL,
  `api_report_details_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_api_report_details_records_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_report_details_records
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_report_details_records_data
-- ----------------------------
DROP TABLE IF EXISTS `api_report_details_records_data`;
CREATE TABLE `api_report_details_records_data` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `parnt_id` bigint DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `step_type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `success` tinyint(1) DEFAULT NULL,
  `elapsed_ms` bigint DEFAULT NULL,
  `attachment` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `http_stat_id` bigint unsigned DEFAULT NULL,
  `data` json DEFAULT NULL,
  `export_vars` json DEFAULT NULL,
  `content_size` bigint DEFAULT NULL,
  `api_report_details_records_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_report_details_records_data
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_report_details_records_data_httpstats
-- ----------------------------
DROP TABLE IF EXISTS `api_report_details_records_data_httpstats`;
CREATE TABLE `api_report_details_records_data_httpstats` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `connect` bigint DEFAULT NULL,
  `content_transfer` bigint DEFAULT NULL,
  `dns_lookup` bigint DEFAULT NULL,
  `name_lookup` bigint DEFAULT NULL,
  `pretransfer` bigint DEFAULT NULL,
  `server_processing` bigint DEFAULT NULL,
  `start_transfer` bigint DEFAULT NULL,
  `tcp_connection` bigint DEFAULT NULL,
  `tls_handshake` bigint DEFAULT NULL,
  `total` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_api_report_details_records_data_httpstats_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_report_details_records_data_httpstats
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_report_stat_testcases
-- ----------------------------
DROP TABLE IF EXISTS `api_report_stat_testcases`;
CREATE TABLE `api_report_stat_testcases` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `total` bigint DEFAULT NULL,
  `success` bigint DEFAULT NULL,
  `fail` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_api_report_stat_testcases_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_report_stat_testcases
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_report_stat_teststeps
-- ----------------------------
DROP TABLE IF EXISTS `api_report_stat_teststeps`;
CREATE TABLE `api_report_stat_teststeps` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `total` bigint DEFAULT NULL,
  `successes` bigint DEFAULT NULL,
  `failures` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_api_report_stat_teststeps_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_report_stat_teststeps
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_report_stats
-- ----------------------------
DROP TABLE IF EXISTS `api_report_stats`;
CREATE TABLE `api_report_stats` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `test_cases_id` bigint unsigned DEFAULT NULL,
  `test_steps_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_api_report_stats_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_report_stats
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_report_times
-- ----------------------------
DROP TABLE IF EXISTS `api_report_times`;
CREATE TABLE `api_report_times` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `start_at` datetime(3) DEFAULT NULL,
  `duration` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_api_report_times_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_report_times
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_reports
-- ----------------------------
DROP TABLE IF EXISTS `api_reports`;
CREATE TABLE `api_reports` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `update_by_id` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `delete_by_id` bigint unsigned DEFAULT NULL COMMENT '删除者',
  `project_id` bigint unsigned DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `success` tinyint(1) DEFAULT NULL,
  `stat_id` bigint unsigned DEFAULT NULL,
  `time_id` bigint unsigned DEFAULT NULL,
  `platform` json DEFAULT NULL,
  `case_type` bigint DEFAULT NULL,
  `run_type` bigint DEFAULT NULL,
  `status` bigint DEFAULT NULL,
  `setup_case` tinyint(1) DEFAULT NULL,
  `describe` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `api_env_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '所属环境名称',
  `api_env_id` bigint unsigned DEFAULT NULL COMMENT '所属环境',
  `hostname` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '主机名',
  PRIMARY KEY (`id`),
  KEY `idx_api_reports_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_reports
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_requests
-- ----------------------------
DROP TABLE IF EXISTS `api_requests`;
CREATE TABLE `api_requests` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `agreement` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '协议',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求方法',
  `http2` tinyint(1) DEFAULT NULL COMMENT '是否为http2',
  `url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求url',
  `params` json DEFAULT NULL COMMENT 'url参数',
  `headers` json DEFAULT NULL COMMENT '请求头',
  `data` json DEFAULT NULL COMMENT 'request body data',
  `params_json` json DEFAULT NULL COMMENT 'url参数json数据格式',
  `headers_json` json DEFAULT NULL COMMENT '请求头json数据格式',
  `data_json` json DEFAULT NULL COMMENT 'request body data json数据格式',
  `json` json DEFAULT NULL,
  `timeout` float DEFAULT NULL COMMENT '超时时间',
  `allow_redirects` tinyint(1) DEFAULT NULL,
  `verify` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_api_requests_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_requests
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_step_rendezvous
-- ----------------------------
DROP TABLE IF EXISTS `api_step_rendezvous`;
CREATE TABLE `api_step_rendezvous` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `percent` float DEFAULT NULL,
  `number` bigint DEFAULT NULL,
  `timeout` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_api_step_rendezvous_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_step_rendezvous
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_step_think_times
-- ----------------------------
DROP TABLE IF EXISTS `api_step_think_times`;
CREATE TABLE `api_step_think_times` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `time` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_api_step_think_times_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_step_think_times
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_step_transactions
-- ----------------------------
DROP TABLE IF EXISTS `api_step_transactions`;
CREATE TABLE `api_step_transactions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_api_step_transactions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_step_transactions
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_steps
-- ----------------------------
DROP TABLE IF EXISTS `api_steps`;
CREATE TABLE `api_steps` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `update_by_id` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `delete_by_id` bigint unsigned DEFAULT NULL COMMENT '删除者',
  `project_id` bigint unsigned DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '接口名称',
  `api_type` bigint DEFAULT NULL COMMENT '接口类型',
  `think_time_id` bigint unsigned DEFAULT NULL COMMENT '思考时间',
  `transaction_id` bigint unsigned DEFAULT NULL COMMENT '事务',
  `rendezvous_id` bigint unsigned DEFAULT NULL COMMENT '集合点',
  `request_id` bigint unsigned DEFAULT NULL COMMENT 'http请求',
  `grpc_id` bigint unsigned DEFAULT NULL COMMENT 'grpc请求',
  `variables` json DEFAULT NULL COMMENT '变量',
  `extract` json DEFAULT NULL COMMENT '导出参数',
  `validate` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '断言',
  `validate_number` bigint unsigned DEFAULT NULL,
  `validate_json` json DEFAULT NULL,
  `extract_json` json DEFAULT NULL,
  `variables_json` json DEFAULT NULL,
  `hooks` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `setup_hooks` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `teardown_hooks` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `sort` bigint unsigned DEFAULT NULL,
  `export_header` json DEFAULT NULL COMMENT '导出请求头到全局config',
  `export_parameter` json DEFAULT NULL COMMENT '导出参数到全局config',
  `parent` bigint unsigned DEFAULT NULL,
  `api_menu_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_api_steps_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_steps
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_timer_task_relationships
-- ----------------------------
DROP TABLE IF EXISTS `api_timer_task_relationships`;
CREATE TABLE `api_timer_task_relationships` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `api_timer_task_id` bigint unsigned DEFAULT NULL COMMENT '定时任务',
  `api_case_id` bigint unsigned DEFAULT NULL COMMENT '测试用例',
  `sort` bigint unsigned DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_api_timer_task_relationships_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_timer_task_relationships
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_timer_task_tag_relationships
-- ----------------------------
DROP TABLE IF EXISTS `api_timer_task_tag_relationships`;
CREATE TABLE `api_timer_task_tag_relationships` (
  `api_timer_task_id` bigint unsigned DEFAULT NULL COMMENT '定时任务',
  `api_timer_task_tag_id` bigint unsigned DEFAULT NULL COMMENT '定时任务标签'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_timer_task_tag_relationships
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_timer_task_tags
-- ----------------------------
DROP TABLE IF EXISTS `api_timer_task_tags`;
CREATE TABLE `api_timer_task_tags` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `update_by_id` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `delete_by_id` bigint unsigned DEFAULT NULL COMMENT '删除者',
  `project_id` bigint unsigned DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `remarks` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_api_timer_task_tags_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_timer_task_tags
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_timer_tasks
-- ----------------------------
DROP TABLE IF EXISTS `api_timer_tasks`;
CREATE TABLE `api_timer_tasks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `update_by_id` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `delete_by_id` bigint unsigned DEFAULT NULL COMMENT '删除者',
  `project_id` bigint unsigned DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '任务名称',
  `run_time` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '运行时间cron',
  `next_run_time` datetime(3) DEFAULT NULL COMMENT '下次运行时间',
  `status` tinyint(1) DEFAULT NULL COMMENT '运行状态',
  `describe` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  `run_number` bigint DEFAULT NULL COMMENT '运行次数',
  `run_config_id` bigint unsigned DEFAULT NULL COMMENT '运行配置',
  `entry_id` bigint DEFAULT NULL,
  `api_env_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '所属环境名称',
  `api_env_id` bigint unsigned DEFAULT NULL COMMENT '所属环境',
  PRIMARY KEY (`id`),
  KEY `idx_api_timer_tasks_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_timer_tasks
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for api_user_configs
-- ----------------------------
DROP TABLE IF EXISTS `api_user_configs`;
CREATE TABLE `api_user_configs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `update_by_id` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `delete_by_id` bigint unsigned DEFAULT NULL COMMENT '删除者',
  `project_id` bigint unsigned DEFAULT NULL,
  `api_config_id` bigint unsigned DEFAULT NULL COMMENT '接口配置ID',
  `api_env_id` bigint unsigned DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户ID',
  PRIMARY KEY (`id`),
  KEY `idx_api_user_configs_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api_user_configs
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/user/admin_register', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/fileUploadAndDownload/editFileName', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '8881', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/user/admin_register', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/fileUploadAndDownload/editFileName', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/autoCode/createTemp', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9528', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/user/admin_register', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/user/setUserInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/sysDictionary/createSysDictionary', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/sysDictionary/deleteSysDictionary', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/sysDictionary/updateSysDictionary', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/sysDictionary/findSysDictionary', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/sysDictionary/getSysDictionaryList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/project/createProject', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/project/deleteProject', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/project/deleteProjectByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/project/updateProject', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/project/findProject', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/project/getProjectList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/createApiMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/deleteApiMenu', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/deleteApiMenuByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/updateApiMenu', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/findApiMenu', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/getApiMenuList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/createInterfaceTemplate', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/deleteInterfaceTemplate', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/deleteInterfaceTemplateByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/updateInterfaceTemplate', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/findInterfaceTemplate', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/getInterfaceTemplateList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/updateDebugTalk', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/getDebugTalk', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/ac/:project/createApiConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/ac/:project/deleteApiConfig', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/ac/:project/deleteApiConfigByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/ac/:project/updateApiConfig', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/ac/:project/findApiConfig', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/ac/:project/getApiConfigList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/step/createTestCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/step/deleteTestCase', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/step/deleteTestCaseByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/step/updateTestCase', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/step/findTestCase', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/step/getTestCaseList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/step/getTestCaseList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/step/addTestCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/step/delTestCase', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/:project/step/sortTestCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/run/:project/runTestCaseStep', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/run/:project/runApiCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/run/:project/runApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/run/:project/runTimerTask', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/run/:project/runBoomerDebug', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/run/:project/runBoomer', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/report/:project/getReportList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/case/report/:project/findReport', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/testcase/:project/createApiCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/testcase/:project/deleteApiCase', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/testcase/:project/deleteApiCaseByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/testcase/:project/updateApiCase', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/testcase/:project/getApiCaseList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/testcase/:project/findApiTestCase', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/testcase/:project/addApisCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/testcase/:project/setApisCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/testcase/:project/sortApisCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/testcase/:project/AddApiTestCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/testcase/:project/delApisCase', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/testcase/:project/findApiCase', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/task/:project/sortTaskCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/task/:project/addTaskCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/task/:project/delTaskCase', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/task/:project/findTaskTestCase', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/task/:project/addTaskTestCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/task/:project/setTaskCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/task/:project/createTimerTask', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/task/:project/deleteTimerTask', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/task/:project/deleteTimerTaskByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/task/:project/updateTimerTask', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/task/:project/findTimerTask', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/task/:project/getTimerTaskList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/performance/:project/createPerformance', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/performance/:project/getPerformanceList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/performance/:project/deletePerformance', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/performance/:project/findPerformance', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/performance/:project/updatePerformance', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/performance/:project/addPerformanceCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/performance/:project/sortPerformanceCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/performance/:project/delPerformanceCase', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/performance/:project/findPerformanceCase', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/performance/:project/addOperation', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/performance/:project/findPerformanceStep', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/performance/:project/getReportList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '9999', '/performance/:project/findReport', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/user/setSelfInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/api/createApiMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/api/updateApiMenu', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/api/findApiMenu', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/api/getApiMenuList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/sysDictionary/findSysDictionary', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/sysDictionary/getSysDictionaryList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/project/findProject', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/project/getProjectList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/:project/createApiMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/:project/updateApiMenu', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/:project/findApiMenu', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/:project/getApiMenuList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/:project/createInterfaceTemplate', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/:project/updateInterfaceTemplate', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/:project/findInterfaceTemplate', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/:project/getInterfaceTemplateList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/:project/getDebugTalk', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/ac/:project/createApiConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/ac/:project/deleteApiConfigByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/ac/:project/updateApiConfig', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/ac/:project/findApiConfig', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/ac/:project/getApiConfigList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/:project/step/createTestCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/:project/step/updateTestCase', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/:project/step/findTestCase', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/:project/step/getTestCaseList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/:project/step/getTestCaseList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/:project/step/addTestCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/:project/step/sortTestCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/run/:project/runTestCaseStep', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/run/:project/runApiCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/run/:project/runApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/run/:project/runTimerTask', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/run/:project/runBoomerDebug', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/run/:project/runBoomer', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/report/:project/getReportList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/case/report/:project/findReport', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/testcase/:project/createApiCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/testcase/:project/updateApiCase', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/testcase/:project/getApiCaseList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/testcase/:project/findApiTestCase', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/testcase/:project/addApisCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/testcase/:project/setApisCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/testcase/:project/sortApisCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/testcase/:project/AddApiTestCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/testcase/:project/findApiCase', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/task/:project/sortTaskCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/task/:project/addTaskCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/task/:project/findTaskTestCase', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/task/:project/addTaskTestCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/task/:project/setTaskCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/task/:project/createTimerTask', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/task/:project/updateTimerTask', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/task/:project/findTimerTask', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/task/:project/getTimerTaskList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/performance/:project/createPerformance', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/performance/:project/getPerformanceList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/performance/:project/deletePerformance', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/performance/:project/findPerformance', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/performance/:project/updatePerformance', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/performance/:project/addPerformanceCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/performance/:project/sortPerformanceCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/performance/:project/delPerformanceCase', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/performance/:project/findPerformanceCase', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/performance/:project/addOperation', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/performance/:project/findPerformanceStep', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/performance/:project/getReportList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '666', '/performance/:project/findReport', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/user/deleteUser', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/user/admin_register', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/user/setUserInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/user/setSelfInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/user/setUserAuthorities', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/user/resetPassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/user/setUserProjects', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/api/deleteApisByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/api/createApiMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/api/deleteApiMenu', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/api/deleteApiMenuByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/api/updateApiMenu', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/api/findApiMenu', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/api/getApiMenuList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/authority/copyAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/authority/updateAuthority', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/fileUploadAndDownload/findFile', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/fileUploadAndDownload/breakpointContinue', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/fileUploadAndDownload/breakpointContinueFinish', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/fileUploadAndDownload/removeChunk', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/fileUploadAndDownload/editFileName', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/system/getServerInfo', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/autoCode/getDB', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/autoCode/getTables', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/autoCode/createTemp', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/autoCode/preview', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/autoCode/getColumn', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/autoCode/createPlug', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/autoCode/installPlugin', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/autoCode/createPackage', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/autoCode/getPackage', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/autoCode/delPackage', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/autoCode/getMeta', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/autoCode/rollback', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/autoCode/getSysHistory', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/autoCode/delSysHistory', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/sysDictionary/createSysDictionary', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/sysDictionary/deleteSysDictionary', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/sysDictionary/updateSysDictionary', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/sysDictionary/findSysDictionary', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/sysDictionary/getSysDictionaryList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/sysOperationRecord/createSysOperationRecord', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/sysOperationRecord/findSysOperationRecord', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/sysOperationRecord/getSysOperationRecordList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/simpleUploader/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/simpleUploader/checkFileMd5', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/simpleUploader/mergeFileMd5', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/email/emailTest', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/email/emailSend', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/excel/importExcel', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/excel/loadExcel', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/excel/exportExcel', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/excel/downloadTemplate', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/authorityBtn/setAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/authorityBtn/getAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/authorityBtn/canRemoveAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/project/createProject', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/project/deleteProject', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/project/deleteProjectByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/project/updateProject', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/project/findProject', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/project/getProjectList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/createApiMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/deleteApiMenu', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/deleteApiMenuByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/updateApiMenu', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/findApiMenu', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/getApiMenuList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/createInterfaceTemplate', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/deleteInterfaceTemplate', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/deleteInterfaceTemplateByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/updateInterfaceTemplate', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/findInterfaceTemplate', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/getInterfaceTemplateList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/updateDebugTalk', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/getDebugTalk', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/createDebugTalk', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/deleteDebugTalk', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/getDebugTalkList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/getGrpc', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/createUserConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/getUserConfig', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/ac/:project/createApiConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/ac/:project/deleteApiConfig', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/ac/:project/deleteApiConfigByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/ac/:project/updateApiConfig', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/ac/:project/findApiConfig', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/ac/:project/getApiConfigList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/step/createTestCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/step/deleteTestCase', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/step/deleteTestCaseByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/step/updateTestCase', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/step/findTestCase', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/step/getTestCaseList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/step/getTestCaseList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/step/addTestCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/step/delTestCase', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/step/sortTestCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/run/:project/runTestCaseStep', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/run/:project/runApiCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/run/:project/runApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/run/:project/runTimerTask', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/run/:project/runBoomerDebug', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/run/:project/runBoomer', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/run/:project/rebalance', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/run/:project/stop', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/report/:project/getReportList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/report/:project/findReport', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/report/:project/delReport', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/testcase/:project/createApiCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/testcase/:project/deleteApiCase', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/testcase/:project/deleteApiCaseByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/testcase/:project/updateApiCase', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/testcase/:project/getApiCaseList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/testcase/:project/findApiTestCase', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/testcase/:project/addApisCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/testcase/:project/setApisCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/testcase/:project/sortApisCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/testcase/:project/AddApiTestCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/testcase/:project/delApisCase', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/testcase/:project/findApiCase', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/task/:project/sortTaskCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/task/:project/addTaskCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/task/:project/delTaskCase', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/task/:project/findTaskTestCase', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/task/:project/addTaskTestCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/task/:project/setTaskCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/task/:project/createTimerTask', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/task/:project/deleteTimerTask', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/task/:project/deleteTimerTaskByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/task/:project/updateTimerTask', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/task/:project/findTimerTask', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/task/:project/getTimerTaskList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/task/:project/getTimerTaskTagList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/task/:project/createTimerTaskTag', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/task/:project/deleteTimerTaskTag', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/performance/:project/createPerformance', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/performance/:project/getPerformanceList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/performance/:project/deletePerformance', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/performance/:project/findPerformance', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/performance/:project/updatePerformance', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/performance/:project/addPerformanceCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/performance/:project/sortPerformanceCase', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/performance/:project/delPerformanceCase', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/performance/:project/findPerformanceCase', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/performance/:project/addOperation', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/performance/:project/findPerformanceStep', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/performance/:project/getReportList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/performance/:project/findReport', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/pyPkg/installPyPkg', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/pyPkg/uninstallPyPkg', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/pyPkg/updatePyPkg', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/pyPkg/getPkgVersionList', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/case/:project/pyPkg/pyPkgList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/env/:project/createEnv', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/env/:project/updateEnv', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/env/:project/deleteEnv', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/env/:project/findEnv', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/env/:project/getEnvList', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/env/:project/createEnvVariable', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/env/:project/deleteEnvVariable', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/env/:project/findEnvVariable', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/env/:project/getEnvVariableList', 'GET', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for exa_customers
-- ----------------------------
DROP TABLE IF EXISTS `exa_customers`;
CREATE TABLE `exa_customers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customer_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint unsigned DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` bigint unsigned DEFAULT NULL COMMENT '管理角色ID',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_exa_customers_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of exa_customers
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for exa_file_chunks
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_chunks`;
CREATE TABLE `exa_file_chunks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `exa_file_id` bigint unsigned DEFAULT NULL,
  `file_chunk_number` bigint DEFAULT NULL,
  `file_chunk_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_exa_file_chunks_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of exa_file_chunks
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for exa_file_upload_and_downloads
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
CREATE TABLE `exa_file_upload_and_downloads` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of exa_file_upload_and_downloads
-- ----------------------------
BEGIN;
INSERT INTO `exa_file_upload_and_downloads` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `url`, `tag`, `key`) VALUES (1, '2022-07-10 15:12:16.481', '2022-07-10 15:12:16.481', NULL, '10.png', 'https://qmplusimg.henrongyi.top/gvalogo.png', 'png', '158787308910.png');
INSERT INTO `exa_file_upload_and_downloads` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `url`, `tag`, `key`) VALUES (2, '2022-07-10 15:12:16.481', '2022-07-10 15:12:16.481', NULL, 'logo.png', 'https://qmplusimg.henrongyi.top/1576554439myAvatar.png', 'png', '1587973709logo.png');
COMMIT;

-- ----------------------------
-- Table structure for exa_files
-- ----------------------------
DROP TABLE IF EXISTS `exa_files`;
CREATE TABLE `exa_files` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `file_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `file_md5` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `file_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `chunk_total` bigint DEFAULT NULL,
  `is_finish` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_exa_files_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of exa_files
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for fs_user_info
-- ----------------------------
DROP TABLE IF EXISTS `fs_user_info`;
CREATE TABLE `fs_user_info` (
  `gva_user_id` bigint unsigned DEFAULT NULL,
  `sub` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `picture` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `open_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `union_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `en_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `tenant_key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `avatar_url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `avatar_thumb` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `avatar_middle` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `avatar_big` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `user_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `mobile` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of fs_user_info
-- ----------------------------
BEGIN;
INSERT INTO `fs_user_info` (`gva_user_id`, `sub`, `name`, `picture`, `open_id`, `union_id`, `en_name`, `tenant_key`, `avatar_url`, `avatar_thumb`, `avatar_middle`, `avatar_big`, `email`, `user_id`, `mobile`) VALUES (7, 'ou_1871e4d2a088021b3369cea8b0cc2dce', 'yangfan', 'https://s1-imfile.feishucdn.com/static-resource/v1/v2_d517651d-6101-4e41-baba-b7d29c701e8g~?image_size=72x72&cut_type=&quality=&format=png&sticker_format=.webp', 'ou_1871e4d2a088021b3369cea8b0cc2dce', 'on_1446c319c1f837dd5674aab6a9740707', '', '1200a3d7a30e575e', 'https://s1-imfile.feishucdn.com/static-resource/v1/v2_d517651d-6101-4e41-baba-b7d29c701e8g~?image_size=72x72&cut_type=&quality=&format=png&sticker_format=.webp', 'https://s1-imfile.feishucdn.com/static-resource/v1/v2_d517651d-6101-4e41-baba-b7d29c701e8g~?image_size=72x72&cut_type=&quality=&format=png&sticker_format=.webp', 'https://s3-imfile.feishucdn.com/static-resource/v1/v2_d517651d-6101-4e41-baba-b7d29c701e8g~?image_size=240x240&cut_type=&quality=&format=png&sticker_format=.webp', 'https://s3-imfile.feishucdn.com/static-resource/v1/v2_d517651d-6101-4e41-baba-b7d29c701e8g~?image_size=640x640&cut_type=&quality=&format=png&sticker_format=.webp', '', '', '');
INSERT INTO `fs_user_info` (`gva_user_id`, `sub`, `name`, `picture`, `open_id`, `union_id`, `en_name`, `tenant_key`, `avatar_url`, `avatar_thumb`, `avatar_middle`, `avatar_big`, `email`, `user_id`, `mobile`) VALUES (9, '', '', '', '', '', '', '', '', '', '', '', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT 'jwt',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=98 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------
BEGIN;
INSERT INTO `jwt_blacklists` (`id`, `created_at`, `updated_at`, `deleted_at`, `jwt`) VALUES (88, '2022-11-25 21:33:59.623', '2022-11-25 21:33:59.623', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZmZiOWY3YzQtODA2Zi00MzRiLThjOGItNzRmM2FhYzI5ODlmIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6IkEyMjMzNyIsIkF1dGhvcml0eUlkIjo2NjYsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2Njk5ODc1NTAsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY2OTM4MTc1MH0.rML5bg-QZCxc571u3NsEqJ6bBIejM9H8Su1ZzaXbdPU');
INSERT INTO `jwt_blacklists` (`id`, `created_at`, `updated_at`, `deleted_at`, `jwt`) VALUES (89, '2023-01-08 01:18:02.534', '2023-01-08 01:18:02.534', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzI4NDYxODktN2RlNy00YjU3LThlOGItMTJmMjcxNmI5NzU5IiwiSUQiOjgsIlVzZXJuYW1lIjoiY2hlZXRhaCIsIk5pY2tOYW1lIjoiY2hlZXRhaCIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2NzM3MTY1MTMsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY3MzExMDcxM30.IVWmMTGNMXMYw6WUWH0k7VNd3B8WPciZd7PRtAKAJg4');
INSERT INTO `jwt_blacklists` (`id`, `created_at`, `updated_at`, `deleted_at`, `jwt`) VALUES (90, '2023-01-08 01:18:28.510', '2023-01-08 01:18:28.510', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZmZiOWY3YzQtODA2Zi00MzRiLThjOGItNzRmM2FhYzI5ODlmIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6IkEyMjMzNyIsIkF1dGhvcml0eUlkIjo2NjYsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2NzM3MTY2MzEsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY3MzExMDgzMX0.6tn__Ka_A1Gd1vTP0_d_MIR0Gy1jkJGbHH3dtVewm8c');
INSERT INTO `jwt_blacklists` (`id`, `created_at`, `updated_at`, `deleted_at`, `jwt`) VALUES (91, '2023-01-08 01:19:28.548', '2023-01-08 01:19:28.548', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzI4NDYxODktN2RlNy00YjU3LThlOGItMTJmMjcxNmI5NzU5IiwiSUQiOjgsIlVzZXJuYW1lIjoiY2hlZXRhaCIsIk5pY2tOYW1lIjoiY2hlZXRhaCIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2NzM3MTY3NTEsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY3MzExMDk1MX0.DAHvBrZWnM5TMXcKJ56b-d-fieIFjqrM-XV0nQW0mNw');
INSERT INTO `jwt_blacklists` (`id`, `created_at`, `updated_at`, `deleted_at`, `jwt`) VALUES (92, '2023-01-08 01:20:10.319', '2023-01-08 01:20:10.319', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZmZiOWY3YzQtODA2Zi00MzRiLThjOGItNzRmM2FhYzI5ODlmIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6IkEyMjMzNyIsIkF1dGhvcml0eUlkIjo2NjYsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2NzM3MTY2OTUsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY3MzExMDg5NX0.nfGCgtW8pNHWIpdSmuTkJbMVRqyyhrYHOwecz_gglas');
INSERT INTO `jwt_blacklists` (`id`, `created_at`, `updated_at`, `deleted_at`, `jwt`) VALUES (93, '2023-01-08 01:23:29.206', '2023-01-08 01:23:29.206', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZmZiOWY3YzQtODA2Zi00MzRiLThjOGItNzRmM2FhYzI5ODlmIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6IkEyMjMzNyIsIkF1dGhvcml0eUlkIjo2NjYsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2NzM3MTY5ODIsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY3MzExMTE4Mn0.P5yfKFmOb2uoTEF-S6vsuav3NAwGE0R_yxwQspIuinU');
INSERT INTO `jwt_blacklists` (`id`, `created_at`, `updated_at`, `deleted_at`, `jwt`) VALUES (94, '2023-05-09 17:57:47.227', '2023-05-09 17:57:47.227', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzI4NDYxODktN2RlNy00YjU3LThlOGItMTJmMjcxNmI5NzU5IiwiSUQiOjgsIlVzZXJuYW1lIjoiY2hlZXRhaCIsIk5pY2tOYW1lIjoiY2hlZXRhaCIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2ODM4MTk5MjMsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY4MzIxNDEyM30.L5qdqk2rNdIy3ZZQOBLWZ-hjWtWsWJhPBj2lG7WcGz8');
INSERT INTO `jwt_blacklists` (`id`, `created_at`, `updated_at`, `deleted_at`, `jwt`) VALUES (95, '2023-05-09 17:58:18.135', '2023-05-09 17:58:18.135', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZmZiOWY3YzQtODA2Zi00MzRiLThjOGItNzRmM2FhYzI5ODlmIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6IkEyMjMzNyIsIkF1dGhvcml0eUlkIjo2NjYsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2ODQyNTk4ODIsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY4MzY1NDA4Mn0.ifg2Lhvb9nrqjv8vAYTNf6go-tPY1MWf3rzlIwU9eHs');
INSERT INTO `jwt_blacklists` (`id`, `created_at`, `updated_at`, `deleted_at`, `jwt`) VALUES (96, '2023-05-09 17:59:01.593', '2023-05-09 17:59:01.593', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYzI4NDYxODktN2RlNy00YjU3LThlOGItMTJmMjcxNmI5NzU5IiwiSUQiOjgsIlVzZXJuYW1lIjoieWFuZ2ZhbiIsIk5pY2tOYW1lIjoieWFuZ2ZhbiIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2ODQyNTk5MTgsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY4MzY1NDExOH0.uRkqyz_pP7PpK3fHQv1lNMgqK68U_JChNQLS9ohCqoQ');
INSERT INTO `jwt_blacklists` (`id`, `created_at`, `updated_at`, `deleted_at`, `jwt`) VALUES (97, '2023-05-18 23:32:11.539', '2023-05-18 23:32:11.539', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZmZiOWY3YzQtODA2Zi00MzRiLThjOGItNzRmM2FhYzI5ODlmIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6IkEyMjMzNyIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2ODUwMjg2MTUsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY4NDQyMjgxNX0.ry9gaDUgs8JfD3OwhamJay1MkQN0wzmHi8Tj_yMUCZE');
COMMIT;

-- ----------------------------
-- Table structure for performance_relationships
-- ----------------------------
DROP TABLE IF EXISTS `performance_relationships`;
CREATE TABLE `performance_relationships` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `performance_id` bigint unsigned DEFAULT NULL COMMENT '性能任务',
  `api_case_step_id` bigint unsigned DEFAULT NULL COMMENT '测试步骤',
  `sort` bigint unsigned DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_performance_relationships_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of performance_relationships
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for performance_report_details
-- ----------------------------
DROP TABLE IF EXISTS `performance_report_details`;
CREATE TABLE `performance_report_details` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_count` bigint DEFAULT NULL,
  `state` bigint DEFAULT NULL,
  `stats_total` json DEFAULT NULL,
  `transactions_passed` bigint DEFAULT NULL,
  `transactions_failed` bigint DEFAULT NULL,
  `total_avg_response_time` double DEFAULT NULL,
  `total_min_response_time` double DEFAULT NULL,
  `total_max_response_time` double DEFAULT NULL,
  `total_rps` double DEFAULT NULL,
  `total_fail_ratio` double DEFAULT NULL,
  `total_fail_per_sec` double DEFAULT NULL,
  `duration` double DEFAULT NULL,
  `errors` json DEFAULT NULL,
  `performance_report_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_performance_report_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of performance_report_details
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for performance_report_masters
-- ----------------------------
DROP TABLE IF EXISTS `performance_report_masters`;
CREATE TABLE `performance_report_masters` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `state` int DEFAULT NULL,
  `workers` int DEFAULT NULL,
  `target_users` bigint DEFAULT NULL,
  `current_users` int DEFAULT NULL,
  `performance_report_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_performance_report_masters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of performance_report_masters
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for performance_report_total_stats
-- ----------------------------
DROP TABLE IF EXISTS `performance_report_total_stats`;
CREATE TABLE `performance_report_total_stats` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `start_time` bigint DEFAULT NULL,
  `num_failures` bigint DEFAULT NULL,
  `num_requests` bigint DEFAULT NULL,
  `max_response_time` bigint DEFAULT NULL,
  `min_response_time` bigint DEFAULT NULL,
  `num_none_requests` bigint DEFAULT NULL,
  `total_response_time` bigint DEFAULT NULL,
  `total_content_length` bigint DEFAULT NULL,
  `last_request_timestamp` bigint DEFAULT NULL,
  `current_rps` double DEFAULT NULL,
  `current_fail_per_sec` double DEFAULT NULL,
  `performance_report_detail_id` bigint unsigned DEFAULT NULL,
  `response_timer` json DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_performance_report_total_stats_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of performance_report_total_stats
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for performance_report_workers
-- ----------------------------
DROP TABLE IF EXISTS `performance_report_workers`;
CREATE TABLE `performance_report_workers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `performance_report_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_performance_report_workers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of performance_report_workers
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for performance_report_works
-- ----------------------------
DROP TABLE IF EXISTS `performance_report_works`;
CREATE TABLE `performance_report_works` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `work_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `os` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `arch` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `state` int DEFAULT NULL,
  `heartbeat` int DEFAULT NULL,
  `user_count` bigint DEFAULT NULL,
  `worker_cpu_usage` double DEFAULT NULL,
  `cpu_usage` double DEFAULT NULL,
  `cpu_warning_emitted` tinyint(1) DEFAULT NULL,
  `worker_memory_usage` double DEFAULT NULL,
  `memory_usage` double DEFAULT NULL,
  `performance_report_worker_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_performance_report_works_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of performance_report_works
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for performance_reports
-- ----------------------------
DROP TABLE IF EXISTS `performance_reports`;
CREATE TABLE `performance_reports` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `state` bigint DEFAULT NULL,
  `performance_id` bigint unsigned DEFAULT NULL,
  `project_id` bigint unsigned DEFAULT NULL,
  `stats_total` json DEFAULT NULL,
  `total_avg_response_time` double DEFAULT NULL,
  `total_min_response_time` double DEFAULT NULL,
  `total_max_response_time` double DEFAULT NULL,
  `total_rps` double DEFAULT NULL,
  `total_fail_ratio` double DEFAULT NULL,
  `total_fail_per_sec` double DEFAULT NULL,
  `user_count` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_performance_reports_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of performance_reports
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for performances
-- ----------------------------
DROP TABLE IF EXISTS `performances`;
CREATE TABLE `performances` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `created_by_id` bigint unsigned DEFAULT NULL COMMENT '创建者',
  `update_by_id` bigint unsigned DEFAULT NULL COMMENT '更新者',
  `delete_by_id` bigint unsigned DEFAULT NULL COMMENT '删除者',
  `project_id` bigint unsigned DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '性能任务名称',
  `status` bigint DEFAULT NULL COMMENT '状态',
  `describe` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  `RunConfigID` bigint unsigned DEFAULT NULL COMMENT '运行配置',
  `front_case` tinyint(1) DEFAULT NULL,
  `entry_id` bigint DEFAULT NULL,
  `api_menu_id` bigint unsigned DEFAULT NULL COMMENT '所属菜单',
  `performance_report_id` bigint unsigned DEFAULT NULL,
  `api_env_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '所属环境名称',
  `api_env_id` bigint unsigned DEFAULT NULL COMMENT '所属环境',
  PRIMARY KEY (`id`),
  KEY `idx_performances_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of performances
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for projects
-- ----------------------------
DROP TABLE IF EXISTS `projects`;
CREATE TABLE `projects` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '项目名称',
  `admin` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '项目管理员',
  `creator` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `describe` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sys_user_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_projects_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of projects
-- ----------------------------
BEGIN;
INSERT INTO `projects` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `admin`, `creator`, `describe`, `sys_user_id`) VALUES (1, '2022-05-30 17:59:35.586', '2022-05-30 17:59:35.586', NULL, '测试项目', '', '', '', 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_apis_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=216 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
BEGIN;
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (1, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/base/login', '用户登录(必选)', 'base', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (2, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单(退出，必选)', 'jwt', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (3, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/user/deleteUser', '删除用户', '系统用户', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (4, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/user/admin_register', '用户注册', '系统用户', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (5, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/user/getUserList', '获取用户列表', '系统用户', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (6, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/user/setUserInfo', '设置用户信息', '系统用户', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (7, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/user/setSelfInfo', '设置自身信息(必选)', '系统用户', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (8, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/user/getUserInfo', '获取自身信息(必选)', '系统用户', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (9, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/user/setUserAuthorities', '设置权限组', '系统用户', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (10, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/user/changePassword', '修改密码（建议选择)', '系统用户', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (11, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/user/setUserAuthority', '修改用户角色(必选)', '系统用户', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (12, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/user/resetPassword', '重置用户密码', '系统用户', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (13, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/api/createApi', '创建api', 'api', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (14, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/api/deleteApi', '删除Api', 'api', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (15, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/api/updateApi', '更新Api', 'api', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (16, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/api/getApiList', '获取api列表', 'api', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (17, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/api/getAllApis', '获取所有api', 'api', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (18, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/api/getApiById', '获取api详细信息', 'api', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (19, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/api/deleteApisByIds', '批量删除api', 'api', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (20, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/authority/copyAuthority', '拷贝角色', '角色', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (21, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/authority/createAuthority', '创建角色', '角色', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (22, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/authority/deleteAuthority', '删除角色', '角色', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (23, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/authority/updateAuthority', '更新角色信息', '角色', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (24, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/authority/getAuthorityList', '获取角色列表', '角色', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (25, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/authority/setDataAuthority', '设置角色资源权限', '角色', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (26, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/casbin/updateCasbin', '更改角色api权限', 'casbin', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (27, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (28, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/menu/addBaseMenu', '新增菜单', '菜单', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (29, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/menu/getMenu', '获取菜单树(必选)', '菜单', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (30, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/menu/deleteBaseMenu', '删除菜单', '菜单', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (31, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/menu/updateBaseMenu', '更新菜单', '菜单', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (32, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/menu/getBaseMenuById', '根据id获取菜单', '菜单', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (33, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/menu/getMenuList', '分页获取基础menu列表', '菜单', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (34, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/menu/getBaseMenuTree', '获取用户动态路由', '菜单', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (35, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/menu/getMenuAuthority', '获取指定角色menu', '菜单', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (36, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', '菜单', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (37, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/fileUploadAndDownload/findFile', '寻找目标文件（秒传）', '分片上传', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (38, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/fileUploadAndDownload/breakpointContinue', '断点续传', '分片上传', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (39, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/fileUploadAndDownload/breakpointContinueFinish', '断点续传完成', '分片上传', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (40, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/fileUploadAndDownload/removeChunk', '上传完成移除文件', '分片上传', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (41, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/fileUploadAndDownload/upload', '文件上传示例', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (42, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/fileUploadAndDownload/deleteFile', '删除文件', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (43, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/fileUploadAndDownload/editFileName', '文件名或者备注编辑', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (44, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/fileUploadAndDownload/getFileList', '获取上传文件列表', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (45, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/system/getServerInfo', '获取服务器信息', '系统服务', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (46, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/system/getSystemConfig', '获取配置文件内容', '系统服务', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (47, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/system/setSystemConfig', '设置配置文件内容', '系统服务', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (48, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/customer/customer', '更新客户', '客户', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (49, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/customer/customer', '创建客户', '客户', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (50, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/customer/customer', '删除客户', '客户', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (51, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/customer/customer', '获取单一客户', '客户', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (52, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/customer/customerList', '获取客户列表', '客户', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (53, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/autoCode/getDB', '获取所有数据库', '代码生成器', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (54, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/autoCode/getTables', '获取数据库表', '代码生成器', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (55, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/autoCode/createTemp', '自动化代码', '代码生成器', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (56, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/autoCode/preview', '预览自动化代码', '代码生成器', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (57, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/autoCode/getColumn', '获取所选table的所有字段', '代码生成器', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (58, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/autoCode/createPlug', '自动创建插件包', '代码生成器', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (59, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/autoCode/installPlugin', '安装插件', '代码生成器', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (60, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/autoCode/createPackage', '生成包(package)', '包（pkg）生成器', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (61, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/autoCode/getPackage', '获取所有包(package)', '包（pkg）生成器', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (62, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/autoCode/delPackage', '删除包(package)', '包（pkg）生成器', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (63, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/autoCode/getMeta', '获取meta信息', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (64, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/autoCode/rollback', '回滚自动生成代码', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (65, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/autoCode/getSysHistory', '查询回滚记录', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (66, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/autoCode/delSysHistory', '删除回滚记录', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (67, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/sysDictionaryDetail/updateSysDictionaryDetail', '更新字典内容', '系统字典详情', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (68, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/sysDictionaryDetail/createSysDictionaryDetail', '新增字典内容', '系统字典详情', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (69, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/sysDictionaryDetail/deleteSysDictionaryDetail', '删除字典内容', '系统字典详情', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (70, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/sysDictionaryDetail/findSysDictionaryDetail', '根据ID获取字典内容', '系统字典详情', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (71, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/sysDictionaryDetail/getSysDictionaryDetailList', '获取字典内容列表', '系统字典详情', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (72, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/sysDictionary/createSysDictionary', '新增字典', '系统字典', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (73, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/sysDictionary/deleteSysDictionary', '删除字典', '系统字典', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (74, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/sysDictionary/updateSysDictionary', '更新字典', '系统字典', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (75, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/sysDictionary/findSysDictionary', '根据ID获取字典', '系统字典', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (76, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/sysDictionary/getSysDictionaryList', '获取字典列表', '系统字典', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (77, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/sysOperationRecord/createSysOperationRecord', '新增操作记录', '操作记录', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (78, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/sysOperationRecord/findSysOperationRecord', '根据ID获取操作记录', '操作记录', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (79, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/sysOperationRecord/getSysOperationRecordList', '获取操作记录列表', '操作记录', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (80, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/sysOperationRecord/deleteSysOperationRecord', '删除操作记录', '操作记录', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (81, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/sysOperationRecord/deleteSysOperationRecordByIds', '批量删除操作历史', '操作记录', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (82, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/simpleUploader/upload', '插件版分片上传', '断点续传(插件版)', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (83, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/simpleUploader/checkFileMd5', '文件完整度验证', '断点续传(插件版)', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (84, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/simpleUploader/mergeFileMd5', '上传完成合并文件', '断点续传(插件版)', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (85, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/email/emailTest', '发送测试邮件', 'email', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (86, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/email/emailSend', '发送邮件示例', 'email', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (87, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/excel/importExcel', '导入excel', 'excel', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (88, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/excel/loadExcel', '下载excel', 'excel', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (89, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/excel/exportExcel', '导出excel', 'excel', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (90, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/excel/downloadTemplate', '下载excel模板', 'excel', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (91, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/authorityBtn/setAuthorityBtn', '设置按钮权限', '按钮权限', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (92, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/authorityBtn/getAuthorityBtn', '获取已有按钮权限', '按钮权限', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (93, '2022-07-10 15:12:15.279', '2022-07-10 15:12:15.279', NULL, '/authorityBtn/canRemoveAuthorityBtn', '删除按钮', '按钮权限', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (94, '2022-01-19 00:48:30.190', '2022-05-21 14:25:47.773', NULL, '/project/createProject', '新增项目管理', '项目管理', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (95, '2022-01-19 00:48:30.214', '2022-05-21 14:25:51.405', NULL, '/project/deleteProject', '删除项目管理', '项目管理', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (96, '2022-01-19 00:48:30.239', '2022-05-21 14:25:54.101', NULL, '/project/deleteProjectByIds', '批量删除项目管理', '项目管理', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (97, '2022-01-19 00:48:30.263', '2022-05-21 14:25:56.958', NULL, '/project/updateProject', '更新项目管理', '项目管理', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (98, '2022-01-19 00:48:30.286', '2022-05-21 14:25:59.693', NULL, '/project/findProject', '根据ID获取项目管理', '项目管理', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (99, '2022-01-19 00:48:30.307', '2022-05-21 14:26:02.470', NULL, '/project/getProjectList', '获取项目管理列表', '项目管理', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (100, '2022-01-21 16:35:37.713', '2022-01-21 16:35:37.713', NULL, '/user/setUserProjects', '分配项目', '系统用户', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (101, '2022-01-23 00:50:35.674', '2022-01-23 00:50:35.674', NULL, '/api/createApiMenu', '新增接口菜单', 'api', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (102, '2022-01-23 00:50:35.704', '2022-01-23 00:50:35.704', NULL, '/api/deleteApiMenu', '删除接口菜单', 'api', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (103, '2022-01-23 00:50:35.735', '2022-01-23 00:50:35.735', NULL, '/api/deleteApiMenuByIds', '批量删除接口菜单', 'api', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (104, '2022-01-23 00:50:35.776', '2022-01-23 00:50:35.776', NULL, '/api/updateApiMenu', '更新接口菜单', 'api', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (105, '2022-01-23 00:50:35.804', '2022-01-23 00:50:35.804', NULL, '/api/findApiMenu', '根据ID获取接口菜单', 'api', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (106, '2022-01-23 00:50:35.832', '2022-01-23 00:50:35.832', NULL, '/api/getApiMenuList', '获取接口菜单列表', 'api', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (107, '2022-01-23 01:55:02.958', '2022-05-21 14:27:30.182', NULL, '/case/:project/createApiMenu', '新增接口菜单', '接口分组树形菜单', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (108, '2022-01-23 01:55:03.000', '2022-05-21 14:27:35.109', NULL, '/case/:project/deleteApiMenu', '删除接口菜单', '接口分组树形菜单', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (109, '2022-01-23 01:55:03.043', '2022-05-21 14:27:40.213', NULL, '/case/:project/deleteApiMenuByIds', '批量删除接口菜单', '接口分组树形菜单', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (110, '2022-01-23 01:55:03.086', '2022-05-21 14:27:44.765', NULL, '/case/:project/updateApiMenu', '更新接口菜单', '接口分组树形菜单', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (111, '2022-01-23 01:55:03.127', '2022-05-21 14:27:49.289', NULL, '/case/:project/findApiMenu', '根据ID获取接口菜单', '接口分组树形菜单', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (112, '2022-01-23 01:55:03.168', '2022-05-21 14:27:54.733', NULL, '/case/:project/getApiMenuList', '获取接口菜单列表', '接口分组树形菜单', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (113, '2022-01-23 20:20:09.138', '2022-05-21 14:28:10.606', NULL, '/case/:project/createInterfaceTemplate', '新增api 模版', 'api模板', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (114, '2022-01-23 20:20:09.164', '2022-05-21 14:28:15.253', NULL, '/case/:project/deleteInterfaceTemplate', '删除api 模版', 'api模板', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (115, '2022-01-23 20:20:09.187', '2022-05-21 14:28:19.830', NULL, '/case/:project/deleteInterfaceTemplateByIds', '批量删除api 模版', 'api模板', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (116, '2022-01-23 20:20:09.210', '2022-05-21 14:28:25.148', NULL, '/case/:project/updateInterfaceTemplate', '更新api 模版', 'api模板', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (117, '2022-01-23 20:20:09.233', '2022-05-21 14:28:30.660', NULL, '/case/:project/findInterfaceTemplate', '根据ID获取api 模版', 'api模板', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (118, '2022-01-23 20:20:09.254', '2022-05-21 14:28:43.543', NULL, '/case/:project/getInterfaceTemplateList', '获取api 模版列表', 'api模板', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (119, '2022-03-09 17:12:46.493', '2022-03-09 17:12:46.493', NULL, '/excel/downloadTemplate', '下载excel模板', 'excel', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (120, '2022-05-11 20:22:06.298', '2022-05-21 14:23:51.146', NULL, '/ac/:project/createApiConfig', '新增配置管理', '测试配置', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (121, '2022-05-11 20:22:06.300', '2022-05-21 14:23:59.601', NULL, '/ac/:project/deleteApiConfig', '删除配置管理', '测试配置', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (122, '2022-05-11 20:22:06.301', '2022-05-21 14:24:03.807', NULL, '/ac/:project/deleteApiConfigByIds', '批量删除配置管理', '测试配置', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (123, '2022-05-11 20:22:06.302', '2022-05-21 14:24:08.535', NULL, '/ac/:project/updateApiConfig', '更新配置管理', '测试配置', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (124, '2022-05-11 20:22:06.302', '2022-05-21 14:24:12.518', NULL, '/ac/:project/findApiConfig', '根据ID获取配置管理', '测试配置', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (125, '2022-05-11 20:22:06.303', '2022-05-21 14:24:17.261', NULL, '/ac/:project/getApiConfigList', '获取配置管理列表', '测试配置', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (126, '2022-05-16 11:52:09.222', '2022-05-21 14:29:00.845', NULL, '/case/:project/step/createTestCase', '新增测试步骤', '测试步骤', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (127, '2022-05-16 11:52:09.224', '2022-05-21 14:29:05.597', NULL, '/case/:project/step/deleteTestCase', '删除测试步骤', '测试步骤', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (128, '2022-05-16 11:52:09.225', '2022-05-21 14:29:11.933', NULL, '/case/:project/step/deleteTestCaseByIds', '批量删除测试步骤', '测试步骤', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (129, '2022-05-16 11:52:09.226', '2022-05-21 14:29:18.309', NULL, '/case/:project/step/updateTestCase', '更新测试步骤', '测试步骤', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (130, '2022-05-16 11:52:09.227', '2022-05-21 14:29:25.028', NULL, '/case/:project/step/findTestCase', '根据ID获取测试步骤', '测试步骤', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (131, '2022-05-16 11:52:09.227', '2022-05-21 14:29:32.149', NULL, '/case/:project/step/getTestCaseList', '获取测试步骤列表', '测试步骤', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (132, '2022-05-19 11:06:19.655', '2022-05-21 14:48:49.099', NULL, '/case/:project/step/getTestCaseList', '测试步骤列表排序', '测试步骤', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (133, '2022-05-25 16:47:34.031', '2022-05-28 11:39:27.033', NULL, '/case/:project/step/addTestCase', '测试步骤添加api', '测试步骤', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (134, '2022-05-25 19:09:32.489', '2022-05-28 11:39:18.524', NULL, '/case/:project/step/delTestCase', '删除测试步骤关联的api', '测试步骤', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (135, '2022-05-26 19:22:17.771', '2022-08-24 08:35:44.133', NULL, '/case/run/:project/runTestCaseStep', '运行测试用例', '运行', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (136, '2022-06-07 13:48:52.585', '2022-06-07 13:48:52.585', NULL, '/case/report/:project/getReportList', '测试报告列表', '测试报告', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (137, '2022-06-08 18:45:04.313', '2022-06-08 18:45:04.313', NULL, '/case/report/:project/findReport', '测试报告详情', '测试报告', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (138, '2022-06-21 17:30:46.541', '2022-06-21 17:54:26.165', NULL, '/testcase/:project/createApiCase', '新增测试用例', '测试用例', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (139, '2022-06-21 17:30:46.572', '2022-06-21 17:54:20.640', NULL, '/testcase/:project/deleteApiCase', '删除测试用例', '测试用例', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (140, '2022-06-21 17:30:46.603', '2022-06-21 17:54:33.600', NULL, '/testcase/:project/deleteApiCaseByIds', '批量删除测试用例', '测试用例', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (141, '2022-06-21 17:30:46.636', '2022-06-21 17:54:38.151', NULL, '/testcase/:project/updateApiCase', '更新测试用例', '测试用例', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (142, '2022-06-21 17:30:46.665', '2022-08-28 21:02:33.070', '2022-08-28 21:02:34.978', '/testcase/:project/findApiTestCaseA', '根据ID获取测试用例2', '测试用例', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (143, '2022-06-21 17:30:46.696', '2022-06-21 17:54:47.086', NULL, '/testcase/:project/getApiCaseList', '获取测试用例列表', '测试用例', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (144, '2022-06-24 11:32:22.281', '2022-06-24 11:38:09.393', NULL, '/case/:project/step/sortTestCase', '测试步骤排序', '测试步骤', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (145, '2022-06-24 20:28:15.933', '2022-06-24 20:28:15.933', NULL, '/testcase/:project/findApiTestCase', '测试用例用例', '测试用例', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (146, '2022-06-24 21:39:07.629', '2022-06-24 21:39:07.629', NULL, '/testcase/:project/addApisCase', '添加测试用例', '测试用例', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (147, '2022-06-25 17:43:56.636', '2022-06-25 17:43:56.636', NULL, '/testcase/:project/setApisCase', '测试用例设置测试步骤', '测试用例', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (148, '2022-06-28 10:40:46.075', '2022-06-28 10:40:46.075', NULL, '/case/run/:project/runApiCase', '运行定时任务', '运行', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (149, '2022-07-12 20:08:21.281', '2022-07-12 20:10:23.118', NULL, '/case/:project/updateDebugTalk', '更新DebugTalk文件', 'api模板', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (150, '2022-07-12 20:08:58.940', '2022-07-12 20:13:40.945', NULL, '/case/:project/getDebugTalk', '获取DebugTalk文件', 'api模板', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (151, '2022-08-01 11:33:43.475', '2022-08-01 11:33:43.475', NULL, '/testcase/:project/sortApisCase', '测试步骤排序', '测试用例', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (152, '2022-08-01 15:42:51.648', '2022-08-01 15:42:51.648', NULL, '/testcase/:project/AddApiTestCase', '任务添加测试步骤', '测试用例', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (153, '2022-08-01 16:45:18.502', '2022-08-01 16:45:18.502', NULL, '/testcase/:project/delApisCase', '定时任务删除测试步骤', '测试用例', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (154, '2022-08-10 12:30:44.986', '2022-08-10 12:30:44.986', NULL, '/case/run/:project/runApi', '调试接口', '运行', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (155, '2022-08-10 12:31:08.302', '2022-08-10 12:31:08.302', '2022-08-10 16:36:28.503', '/case/run/:project/runApiSave', '调试并保存接口', '运行', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (156, '2022-08-01 11:33:43.475', '2022-08-01 11:33:43.475', NULL, '/task/:project/sortTaskCase', '测试用例排序', '定时任务', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (157, '2022-08-01 15:42:51.648', '2022-08-01 15:42:51.648', NULL, '/task/:project/addTaskCase', '任务添加测试用例', '定时任务', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (158, '2022-08-01 16:45:18.502', '2022-08-01 16:45:18.502', NULL, '/task/:project/delTaskCase', '定时任务删除测试用例', '定时任务', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (159, '2022-06-24 20:28:15.933', '2022-06-24 20:28:15.933', NULL, '/task/:project/findTaskTestCase', '定时任务用例', '定时任务', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (160, '2022-06-24 21:39:07.629', '2022-06-24 21:39:07.629', NULL, '/task/:project/addTaskTestCase', '添加测试用例', '定时任务', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (161, '2022-06-25 17:43:56.636', '2022-06-25 17:43:56.636', NULL, '/task/:project/setTaskCase', '定时任务设置测试用例', '定时任务', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (162, '2022-06-28 10:40:46.075', '2022-08-30 00:00:04.328', NULL, '/case/run/:project/runTimerTask', '运行定时任务', '运行', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (163, '2022-06-21 17:30:46.541', '2022-06-21 17:54:26.165', NULL, '/task/:project/createTimerTask', '新增定时任务', '定时任务', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (164, '2022-06-21 17:30:46.572', '2022-06-21 17:54:20.640', NULL, '/task/:project/deleteTimerTask', '删除定时任务', '定时任务', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (165, '2022-06-21 17:30:46.603', '2022-06-21 17:54:33.600', NULL, '/task/:project/deleteTimerTaskByIds', '批量删除定时任务', '定时任务', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (166, '2022-06-21 17:30:46.636', '2022-06-21 17:54:38.151', NULL, '/task/:project/updateTimerTask', '更新定时任务', '定时任务', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (167, '2022-06-21 17:30:46.665', '2022-06-21 17:54:42.496', NULL, '/task/:project/findTimerTask', '根据ID获取定时任务', '定时任务', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (168, '2022-06-21 17:30:46.696', '2022-06-21 17:54:47.086', NULL, '/task/:project/getTimerTaskList', '获取定时任务列表', '定时任务', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (169, '2022-08-28 21:00:08.358', '2022-08-28 21:00:08.358', '2022-08-28 21:00:13.879', '123', '123', '123', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (170, '2022-08-28 21:42:08.583', '2022-08-28 21:42:08.583', NULL, '/testcase/:project/findApiCase', '根据ID获取测试用例', '测试用例', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (174, '2022-10-10 16:04:15.897', '2022-10-10 16:25:06.064', NULL, '/performance/:project/createPerformance', '创建性能测试任务', '性能测试', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (175, '2022-10-10 18:13:10.936', '2022-10-10 18:13:10.936', NULL, '/performance/:project/getPerformanceList', '获取性能任务列表', '性能测试', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (176, '2022-10-10 19:39:08.317', '2022-10-10 19:39:08.317', NULL, '/performance/:project/deletePerformance', '删除性能任务', '性能测试', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (177, '2022-10-10 20:00:51.862', '2022-10-10 20:00:51.862', NULL, '/performance/:project/findPerformance', '通过id查找性能任务', '性能测试', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (178, '2022-10-11 11:56:49.534', '2022-10-11 11:58:15.854', NULL, '/performance/:project/updatePerformance', '更新性能测试任务', '性能测试', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (179, '2022-10-11 16:24:50.743', '2022-10-11 16:24:50.743', NULL, '/performance/:project/addPerformanceCase', '性能任务添加测试用例', '性能测试', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (180, '2022-10-11 16:25:14.114', '2022-10-11 16:25:14.114', NULL, '/performance/:project/sortPerformanceCase', '性能测试用例排序', '性能测试', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (181, '2022-10-11 16:25:30.016', '2022-10-11 16:25:30.016', NULL, '/performance/:project/delPerformanceCase', '性能测试删除测试用例', '性能测试', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (182, '2022-10-11 19:42:48.656', '2022-10-11 19:42:48.656', NULL, '/performance/:project/findPerformanceCase', '性能测试通过id查找用例', '性能测试', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (183, '2022-10-30 09:13:03.545', '2022-10-30 09:13:03.545', NULL, '/performance/:project/addOperation', '添加事务、集合点', '性能测试', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (184, '2022-10-31 16:53:40.593', '2022-10-31 16:53:40.593', NULL, '/performance/:project/findPerformanceStep', '查看性能测试步骤', '性能测试', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (185, '2022-11-01 16:46:08.149', '2022-11-01 16:46:08.149', NULL, '/case/run/:project/runBoomerDebug', '性能测试调试运行', '运行', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (186, '2022-11-01 22:42:48.532', '2022-11-01 22:42:48.532', NULL, '/case/run/:project/runBoomer', '运行性能测试', '运行', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (187, '2022-11-03 16:03:02.156', '2022-11-03 19:17:36.449', NULL, '/performance/:project/getReportList', '性能测试报告列表', '性能测试', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (188, '2022-11-03 19:17:46.378', '2022-11-03 19:17:46.378', NULL, '/performance/:project/findReport', '性能测试报告详情', '性能测试', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (189, '2022-12-20 01:25:19.442', '2022-12-20 01:25:19.442', NULL, '/case/report/:project/delReport', '删除测试报告', '测试报告', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (190, '2023-05-09 17:39:19.296', '2023-05-09 17:39:19.296', NULL, '/case/:project/createDebugTalk', '创建DebugTalk文件', 'api模板', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (191, '2023-05-09 17:39:19.305', '2023-05-09 17:39:19.305', NULL, '/case/:project/deleteDebugTalk', '删除DebugTalk文件', 'api模板', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (192, '2023-05-09 17:39:19.315', '2023-05-09 17:39:19.315', NULL, '/case/:project/getDebugTalkList', '获取DebugTalk列表', 'api模板', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (193, '2023-05-09 17:39:19.325', '2023-05-09 17:39:19.325', NULL, '/case/:project/getGrpc', '获取grpc信息', 'api模板', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (194, '2023-05-09 17:39:19.334', '2023-05-09 17:39:19.334', NULL, '/case/:project/createUserConfig', '创建/更新用户配置', 'api模板', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (195, '2023-05-09 17:39:19.345', '2023-05-09 17:39:19.345', NULL, '/case/:project/getUserConfig', '获取用户配置', 'api模板', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (196, '2023-05-09 17:39:19.354', '2023-05-09 17:39:19.354', NULL, '/case/:project/pyPkg/installPyPkg', '安装python第三方库', 'py库管理', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (197, '2023-05-09 17:39:19.364', '2023-05-09 17:39:19.364', NULL, '/case/:project/pyPkg/uninstallPyPkg', '卸载ython第三方库', 'py库管理', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (198, '2023-05-09 17:39:19.374', '2023-05-09 17:39:19.374', NULL, '/case/:project/pyPkg/updatePyPkg', '更新ython第三方库', 'py库管理', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (199, '2023-05-09 17:39:19.383', '2023-05-09 17:39:19.383', NULL, '/case/:project/pyPkg/getPkgVersionList', '获取python第三方库版本信息', 'py库管理', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (200, '2023-05-09 17:39:19.392', '2023-05-09 17:39:19.392', NULL, '/case/:project/pyPkg/pyPkgList', '获取python第三方库列表', 'py库管理', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (201, '2023-05-09 17:39:19.450', '2023-05-09 17:39:19.450', NULL, '/task/:project/getTimerTaskTagList', '获取TimerTaskTag列表', '定时任务', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (202, '2023-05-09 17:39:19.479', '2023-05-09 17:39:19.479', NULL, '/task/:project/createTimerTaskTag', '创建定时任务标签', '定时任务', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (203, '2023-05-09 17:39:19.489', '2023-05-09 17:39:19.489', NULL, '/task/:project/deleteTimerTaskTag', '删除定时任务标签', '定时任务', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (204, '2023-05-09 17:39:19.763', '2023-05-09 17:39:19.763', NULL, '/env/:project/createEnv', '新增环境', '环境变量', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (205, '2023-05-09 17:39:19.773', '2023-05-09 17:39:19.773', NULL, '/env/:project/updateEnv', '修改环境', '环境变量', 'PUT');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (206, '2023-05-09 17:39:19.782', '2023-05-09 17:39:19.782', NULL, '/env/:project/deleteEnv', '删除环境', '环境变量', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (207, '2023-05-09 17:39:19.791', '2023-05-09 17:39:19.791', NULL, '/env/:project/findEnv', '通过id查找环境', '环境变量', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (208, '2023-05-09 17:39:19.801', '2023-05-09 17:39:19.801', NULL, '/env/:project/getEnvList', '查询环境列表', '环境变量', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (209, '2023-05-09 17:39:19.810', '2023-05-09 17:39:19.810', NULL, '/env/:project/createEnvVariable', '新增变量', '环境变量', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (210, '2023-05-09 17:39:19.820', '2023-05-09 17:39:19.820', NULL, '/env/:project/deleteEnvVariable', '新增变量', '环境变量', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (211, '2023-05-09 17:39:19.831', '2023-05-09 17:39:19.831', NULL, '/env/:project/findEnvVariable', '通过id查找变量', '环境变量', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (212, '2023-05-09 17:39:19.841', '2023-05-09 17:39:19.841', NULL, '/env/:project/getEnvVariableList', '查询变量列表', '环境变量', 'GET');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (213, '2023-07-09 18:23:47.022', '2023-07-09 18:23:47.022', NULL, '/performance/:project/deleteReport', '删除性能测试报告', '性能测试', 'DELETE');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (214, '2023-07-09 18:23:47.062', '2023-07-09 18:23:47.062', NULL, '/case/run/:project/rebalance', '调整性能测试参数', '运行', 'POST');
INSERT INTO `sys_apis` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`) VALUES (215, '2023-07-09 18:23:47.067', '2023-07-09 18:23:47.067', NULL, '/case/run/:project/stop', '停止性能测试', '运行', 'GET');
COMMIT;

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `authority_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `authority_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`) USING BTREE,
  UNIQUE KEY `authority_id` (`authority_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
BEGIN;
INSERT INTO `sys_authorities` (`created_at`, `updated_at`, `deleted_at`, `authority_id`, `authority_name`, `parent_id`, `default_router`) VALUES ('2022-07-11 17:39:20.502', '2023-07-09 18:36:51.534', NULL, 666, '飞书登录默认角色', 0, '关于我们');
INSERT INTO `sys_authorities` (`created_at`, `updated_at`, `deleted_at`, `authority_id`, `authority_name`, `parent_id`, `default_router`) VALUES ('2022-07-10 15:12:15.356', '2023-07-09 18:36:42.288', NULL, 888, '普通用户', 0, '关于我们');
INSERT INTO `sys_authorities` (`created_at`, `updated_at`, `deleted_at`, `authority_id`, `authority_name`, `parent_id`, `default_router`) VALUES ('2022-07-10 15:12:15.356', '2022-07-10 15:12:16.430', NULL, 8881, '普通用户子角色', 888, 'dashboard');
INSERT INTO `sys_authorities` (`created_at`, `updated_at`, `deleted_at`, `authority_id`, `authority_name`, `parent_id`, `default_router`) VALUES ('2022-07-10 15:12:15.356', '2022-07-10 15:12:16.310', NULL, 9528, '测试角色', 0, 'dashboard');
INSERT INTO `sys_authorities` (`created_at`, `updated_at`, `deleted_at`, `authority_id`, `authority_name`, `parent_id`, `default_router`) VALUES ('2022-11-22 11:50:24.089', '2022-11-22 11:58:42.664', NULL, 9999, '自动化测试', 0, 'dashboard');
COMMIT;

-- ----------------------------
-- Table structure for sys_authority_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_btns`;
CREATE TABLE `sys_authority_btns` (
  `authority_id` bigint unsigned DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint unsigned DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_authority_btns
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus` (
  `sys_base_menu_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`,`sys_authority_authority_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
BEGIN;
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (2, 8881);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (2, 9528);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (3, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (4, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (4, 8881);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (5, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (5, 8881);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (6, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (6, 8881);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (7, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (7, 8881);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (8, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (8, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (8, 8881);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (8, 9528);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (8, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (9, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (9, 8881);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (10, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (10, 8881);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (11, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (11, 8881);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (12, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (12, 8881);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (13, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (14, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (14, 8881);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (15, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (15, 8881);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (16, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (16, 8881);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (17, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (17, 8881);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (18, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (19, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (20, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (22, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (22, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (23, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (24, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (25, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (26, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (27, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (28, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (29, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (30, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (31, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (32, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (33, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (33, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (33, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (34, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (34, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (34, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (35, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (35, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (35, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (36, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (36, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (36, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (37, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (37, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (37, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (38, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (38, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (38, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (42, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (42, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (42, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (43, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (43, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (43, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (44, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (44, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (44, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (46, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (46, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (46, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (47, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (47, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (47, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (49, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (49, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (49, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (51, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (51, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (51, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (53, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (53, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (53, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (55, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (55, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (55, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (56, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (56, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (56, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (57, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (57, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (57, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (58, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (58, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (58, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (59, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (59, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (59, 9999);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (60, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (61, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (62, 888);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (63, 666);
INSERT INTO `sys_authority_menus` (`sys_base_menu_id`, `sys_authority_authority_id`) VALUES (63, 888);
COMMIT;

-- ----------------------------
-- Table structure for sys_auto_code_histories
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_code_histories`;
CREATE TABLE `sys_auto_code_histories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `package` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `request_meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `auto_code_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `injection_meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `struct_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `struct_cn_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `api_ids` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `flag` bigint DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_auto_code_histories_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_auto_code_histories
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_auto_codes
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_codes`;
CREATE TABLE `sys_auto_codes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `package_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '包名',
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '展示名',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_auto_codes_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_auto_codes
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_base_menu_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_btns`;
CREATE TABLE `sys_base_menu_btns` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_base_menu_btns_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_base_menu_btns
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_base_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
CREATE TABLE `sys_base_menu_parameters` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned DEFAULT NULL,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_base_menu_parameters
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `menu_level` bigint unsigned DEFAULT NULL,
  `parent_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
BEGIN;
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (1, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', '2023-07-09 18:27:08.299', 0, '0', 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, 0, 0, '仪表盘', 'odometer', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (3, '2022-07-10 15:12:16.165', '2022-07-12 18:24:59.477', NULL, 0, '0', 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 2, 0, 0, '超级管理员', 'user', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (4, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '3', 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, 0, 0, '角色管理', 'avatar', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (5, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '3', 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 2, 1, 0, '菜单管理', 'tickets', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (6, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '3', 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, 1, 0, 'api管理', 'platform', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (7, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '3', 'user', 'user', 0, 'view/superAdmin/user/user.vue', 4, 0, 0, '用户管理', 'coordinate', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (8, '2022-07-10 15:12:16.165', '2022-07-12 18:25:05.048', NULL, 0, '0', 'person', 'person', 1, 'view/person/person.vue', 3, 0, 0, '个人信息', 'message', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (9, '2022-07-10 15:12:16.165', '2022-11-22 14:50:42.311', NULL, 0, '0', 'example', 'example', 0, 'view/example/index.vue', 17, 0, 0, '示例文件', 'management', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (10, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '9', 'excel', 'excel', 0, 'view/example/excel/excel.vue', 4, 0, 0, 'excel导入导出', 'takeaway-box', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (11, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '9', 'upload', 'upload', 0, 'view/example/upload/upload.vue', 5, 0, 0, '媒体库（上传下载）', 'upload', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (12, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '9', 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', 6, 0, 0, '断点续传', 'upload-filled', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (13, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '9', 'customer', 'customer', 0, 'view/example/customer/customer.vue', 7, 0, 0, '客户列表（资源示例）', 'avatar', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (14, '2022-07-10 15:12:16.165', '2022-11-22 14:50:24.888', NULL, 0, '0', 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 15, 0, 0, '系统工具', 'tools', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (15, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '14', 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, 1, 0, '代码生成器', 'cpu', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (16, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '14', 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 2, 1, 0, '表单生成器', 'magic-stick', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (17, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '14', 'system', 'system', 0, 'view/systemTools/system/system.vue', 3, 0, 0, '系统配置', 'operation', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (18, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '3', 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', 5, 0, 0, '字典管理', 'notebook', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (19, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '3', 'dictionaryDetail/:id', 'dictionaryDetail', 1, 'view/superAdmin/dictionary/sysDictionaryDetail.vue', 1, 0, 0, '字典详情-${id}', 'order', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (20, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '3', 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, 0, 0, '操作历史', 'pie-chart', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (21, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '9', 'simpleUploader', 'simpleUploader', 0, 'view/example/simpleUploader/simpleUploader', 6, 0, 0, '断点续传（插件版）', 'upload', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (22, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '0', 'http://www.yangfan.gd.cn/', 'http://www.yangfan.gd.cn/', 0, '/', 0, 0, 0, '官方网站', 'home-filled', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (23, '2022-07-10 15:12:16.165', '2022-11-22 14:50:49.106', NULL, 0, '0', 'state', 'state', 0, 'view/system/state.vue', 18, 0, 0, '服务器状态', 'cloudy', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (24, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '14', 'autoCodeAdmin', 'autoCodeAdmin', 0, 'view/systemTools/autoCodeAdmin/index.vue', 1, 0, 0, '自动化代码管理', 'magic-stick', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (25, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '14', 'autoCodeEdit/:id', 'autoCodeEdit', 1, 'view/systemTools/autoCode/index.vue', 0, 0, 0, '自动化代码-${id}', 'magic-stick', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (26, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '14', 'autoPkg', 'autoPkg', 0, 'view/systemTools/autoPkg/autoPkg.vue', 0, 0, 0, '自动化package', 'folder', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (27, '2022-07-10 15:12:16.165', '2022-11-22 14:50:34.751', NULL, 0, '0', 'plugin', 'plugin', 0, 'view/routerHolder.vue', 16, 0, 0, '插件系统', 'cherry', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (28, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '27', 'https://plugin.gin-vue-admin.com/', 'https://plugin.gin-vue-admin.com/', 0, 'https://plugin.gin-vue-admin.com/', 0, 0, 0, '插件市场', 'shop', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (29, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '27', 'installPlugin', 'installPlugin', 0, 'view/systemTools/installPlugin/index.vue', 1, 0, 0, '插件安装', 'box', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (30, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '27', 'autoPlug', 'autoPlug', 0, 'view/systemTools/autoPlug/autoPlug.vue', 2, 0, 0, '插件模板', 'folder', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (31, '2022-07-10 15:12:16.165', '2022-07-10 15:12:16.165', NULL, 0, '27', 'plugin-email', 'plugin-email', 0, 'plugin/email/view/index.vue', 3, 0, 0, '邮件插件', 'message', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (32, '2022-01-19 01:09:37.274', '2023-07-09 18:23:46.921', NULL, 0, '3', 'project', 'project', 0, 'view/project/project.vue', 1, 0, 0, '项目管理', 'aim', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (33, '2022-01-23 18:34:58.647', '2023-07-09 18:23:46.928', NULL, 0, '0', 'interfaces', 'interfaces', 0, 'view/interface/index.vue', 4, 0, 0, '接口自动化', 'box', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (34, '2022-01-23 18:37:56.434', '2023-07-09 18:23:46.934', NULL, 0, '33', 'interfacetemplate', 'interfacetemplate', 0, 'view/interface/interfaceTemplate/interfaceTemplate.vue', 200, 0, 0, '接口管理', 'coin', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (35, '2022-01-28 13:29:16.282', '2023-07-09 18:23:46.941', NULL, 0, '33', 'DebugReport', 'DebugReport', 1, 'view/interface/interfaceReport/DebugReport.vue', 99999, 0, 0, 'DebugReport', 'aim', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (36, '2022-05-11 19:34:31.473', '2023-07-09 18:23:46.945', NULL, 0, '33', 'apiConfig', 'apiConfig', 0, 'view/interface/interfaceTemplate/apiConfig.vue', 100, 0, 0, '配置管理', 'expand', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (37, '2022-05-16 12:15:29.299', '2023-07-09 18:23:46.950', NULL, 0, '33', 'testCaseStep', 'testCaseStep', 0, 'view/interface/testCaseStep/testCaseStep.vue', 300, 0, 0, '测试步骤', 'suitcase', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (38, '2022-05-17 12:54:39.571', '2023-07-09 18:23:46.955', NULL, 0, '33', 'testCaseStepDetail/:id', 'testCaseStepDetail', 1, 'view/interface/testCaseStep/testCaseStepDetail.vue', 99999, 0, 0, '步骤详情-${id}', 'finished', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (39, '2022-05-18 19:30:05.572', '2022-05-18 19:34:27.917', '2022-07-12 11:43:22.410', 0, '33', 'apiTest', 'apiTest', 0, 'view/interface/test.vue', 0, 0, 0, '测试', 'aim', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (40, '2022-05-21 15:15:37.413', '2022-05-21 15:16:32.325', '2022-07-12 11:43:26.991', 0, '33', 'testCaseAdd', 'testCaseAdd', 0, 'view/interface/testCase/testCaseAdd.vue', 0, 0, 0, '添加测试用例', 'aim', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (41, '2022-05-18 19:30:05.572', '2022-05-18 19:34:27.917', '2022-07-12 11:43:33.305', 0, '33', 'apiTest2', 'apiTest2', 0, 'view/interface/test2.vue', 0, 0, 0, '测试2', 'aim', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (42, '2022-06-07 13:37:19.135', '2023-07-09 18:23:46.961', NULL, 0, '33', 'report', 'report', 0, 'view/interface/Reports/report.vue', 600, 0, 0, '测试报告', 'compass', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (43, '2022-06-08 17:40:11.776', '2023-07-09 18:23:46.969', NULL, 0, '33', 'reportDetail/:report_id', 'reportDetail', 1, 'view/interface/Reports/reportDetail.vue', 99999, 0, 0, '测试报告详情-${report_id}', 'document', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (44, '2022-06-21 17:34:55.263', '2023-07-09 18:23:46.974', NULL, 0, '33', 'testCase', 'testCase', 0, 'view/interface/apiCase/apiCase.vue', 400, 0, 0, '测试用例', 'briefcase', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (45, '2022-07-12 11:38:23.076', '2022-07-12 18:22:14.945', '2022-07-14 08:25:52.291', 0, '33', 'debugtalk', 'debugtalk', 0, 'view/interface/debugtalk/index.vue', 99, 0, 0, 'debugtalk', 'document', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (46, '2022-07-12 11:39:29.347', '2023-07-09 18:23:46.979', NULL, 0, '33', 'debugtalk', 'debugtalk', 0, 'view/interface/debugtalk/debugtalk.vue', 700, 0, 0, '驱动函数', 'reading', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (47, '2022-07-12 11:40:03.210', '2023-07-09 18:23:46.984', NULL, 0, '33', 'debugtalkGen', 'debugtalkGen', 1, 'view/interface/debugtalk/debugtalkGen.vue', 99999, 0, 0, 'debugtalkGen', 'aim', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (48, '2022-07-29 19:33:59.953', '2022-07-29 19:33:59.953', '2022-08-01 17:00:43.663', 0, '33', 'taskAddCase', 'taskAddCase', 0, 'view/interface/timerTask/taskAddCase.vue', 0, 0, 0, '任务添加测试用例', 'aim', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (49, '2022-08-01 06:08:02.383', '2023-07-09 18:23:46.989', NULL, 0, '33', 'apisCaseDetail/:id', 'apisCaseDetail', 1, 'view/interface/apiCase/apisCaseDetail.vue', 99999, 0, 0, '用例详情-${id}', 'aim', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (50, '2022-08-10 14:34:55.331', '2022-08-10 14:34:55.331', '2022-08-10 14:48:12.813', 0, '33', 'envConfig', 'envConfig', 0, 'view/interface/interfaceComponents/envConfig.vue', 1, 0, 0, 'envConfig', 'aim', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (51, '2022-08-01 06:08:02.383', '2023-07-09 18:23:46.994', NULL, 0, '33', 'taskCaseDetail/:id', 'taskCaseDetail', 1, 'view/interface/timerTask/taskCaseDetail.vue', 99999, 0, 0, '任务详情-${id}', 'aim', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (52, '2022-07-29 19:33:59.953', '2022-07-29 19:33:59.953', '2022-08-01 17:00:43.663', 0, '33', 'taskAddCase', 'taskAddCase', 0, 'view/interface/timerTask/taskAddCase.vue', 0, 0, 0, '任务添加测试用例', 'aim', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (53, '2022-06-21 17:34:55.263', '2023-07-09 18:23:47.003', NULL, 0, '33', 'timerTask', 'timerTask', 0, 'view/interface/timerTask/timerTask.vue', 500, 0, 0, '定时任务', 'timer', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (54, '2022-09-04 09:57:40.935', '2022-09-04 09:57:40.935', '2023-05-18 23:31:26.297', 0, '33', 'testTree', 'testTree', 0, 'view/interface/interfaceComponents/tree.vue', 0, 0, 0, 'arcoTree', 'aim', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (55, '2022-10-08 19:30:47.868', '2023-07-09 18:23:47.017', NULL, 0, '0', 'performance', 'performance', 0, 'view/performance/index.vue', 5, 0, 0, '性能测试', 'stopwatch', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (56, '2022-10-08 19:51:05.009', '2023-07-09 18:23:47.022', NULL, 0, '55', 'performanceTask', 'performanceTask', 0, 'view/performance/task/index.vue', 1, 0, 0, '性能任务', 'cpu', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (57, '2022-10-11 14:27:10.146', '2023-07-09 18:23:47.028', NULL, 0, '55', 'performanceDetail/:id', 'performanceDetail', 1, 'view/performance/task/taskDetail.vue', 99999, 0, 0, '性能任务详情-${id}', 'aim', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (58, '2022-11-03 15:49:19.428', '2023-07-09 18:23:47.033', NULL, 0, '55', 'pReport', 'pReport', 0, 'view/performance/report.vue', 2, 0, 0, '性能测试报告', 'compass', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (59, '2022-11-03 18:59:54.997', '2023-07-09 18:23:47.037', NULL, 0, '55', 'pReportDetail/:id', 'pReportDetail', 1, 'view/performance/reportDetail.vue', 999, 0, 0, '性能测试报告详情-${id}', 'document', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (60, '2023-05-09 17:39:19.121', '2023-07-09 18:23:47.042', NULL, 0, '33', 'jsonCompare', 'jsonCompare', 1, 'view/interface/interfaceComponents/jsonCompare.vue', 99999, 0, 0, 'json', 'aim', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (61, '2023-05-09 17:39:19.132', '2023-07-09 18:23:47.049', NULL, 0, '33', 'env', 'env', 0, 'view/interface/environment/environment.vue', 0, 0, 0, '环境变量', 'grid', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (62, '2023-05-09 17:39:19.144', '2023-07-09 18:23:47.057', NULL, 0, '33', 'py_pkg', 'py_pkg', 0, 'view/py_pkg/py_pkg.vue', 680, 0, 0, 'py库管理', 'office-building', 0);
INSERT INTO `sys_base_menus` (`id`, `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`, `close_tab`) VALUES (63, '2023-07-09 18:35:48.941', '2023-07-09 18:35:48.941', NULL, 0, '0', '关于我们', '关于我们', 0, 'view/about/index.vue', 0, 1, 0, '关于我们', 'info-filled', 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
BEGIN;
INSERT INTO `sys_data_authority_id` (`sys_authority_authority_id`, `data_authority_id_authority_id`) VALUES (888, 888);
INSERT INTO `sys_data_authority_id` (`sys_authority_authority_id`, `data_authority_id_authority_id`) VALUES (888, 8881);
INSERT INTO `sys_data_authority_id` (`sys_authority_authority_id`, `data_authority_id_authority_id`) VALUES (888, 9528);
INSERT INTO `sys_data_authority_id` (`sys_authority_authority_id`, `data_authority_id_authority_id`) VALUES (9528, 8881);
INSERT INTO `sys_data_authority_id` (`sys_authority_authority_id`, `data_authority_id_authority_id`) VALUES (9528, 9528);
COMMIT;

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_dictionaries_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
BEGIN;
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (1, '2022-07-10 15:12:15.509', '2022-07-10 15:12:15.548', NULL, '性别', 'gender', 1, '性别字典');
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (2, '2022-07-10 15:12:15.509', '2022-07-10 15:12:15.600', NULL, '数据库int类型', 'int', 1, 'int类型对应的数据库类型');
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (3, '2022-07-10 15:12:15.509', '2022-07-10 15:12:15.650', NULL, '数据库时间日期类型', 'time.Time', 1, '数据库时间日期类型');
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (4, '2022-07-10 15:12:15.509', '2022-07-10 15:12:15.703', NULL, '数据库浮点型', 'float64', 1, '数据库浮点型');
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (5, '2022-07-10 15:12:15.509', '2022-07-10 15:12:15.755', NULL, '数据库字符串', 'string', 1, '数据库字符串');
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (6, '2022-07-10 15:12:15.509', '2022-07-10 15:12:15.807', NULL, '数据库bool类型', 'bool', 1, '数据库bool类型');
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (7, '2022-05-09 17:16:48.006', '2022-05-09 17:16:48.006', NULL, '断言', 'assert', 1, '断言');
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (8, '2022-05-09 17:19:43.351', '2022-05-09 17:19:43.351', NULL, '请求方式', 'method', 1, '请求方式');
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (9, '2022-07-04 10:55:49.346', '2022-07-04 10:55:49.346', NULL, '断言类型', 'assertType', 1, '断言类型');
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (10, '2022-07-08 18:18:39.059', '2022-07-08 18:26:34.933', NULL, '变量', 'variablesType', 1, '需要导出的变量');
INSERT INTO `sys_dictionaries` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `type`, `status`, `desc`) VALUES (11, '2022-07-28 20:15:46.758', '2022-07-28 20:15:46.758', NULL, '请求头', 'requestHeader', 1, '请求头');
COMMIT;

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '展示值',
  `value` bigint DEFAULT NULL COMMENT '字典值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint unsigned DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=135 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_dictionary_details
-- ----------------------------
BEGIN;
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (1, '2022-07-10 15:12:15.559', '2022-07-10 15:12:15.559', NULL, '男', 1, 1, 1, 1);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (2, '2022-07-10 15:12:15.559', '2022-07-10 15:12:15.559', NULL, '女', 2, 1, 2, 1);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (3, '2022-07-10 15:12:15.610', '2022-07-10 15:12:15.610', NULL, 'smallint', 1, 1, 1, 2);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (4, '2022-07-10 15:12:15.610', '2022-07-10 15:12:15.610', NULL, 'mediumint', 2, 1, 2, 2);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (5, '2022-07-10 15:12:15.610', '2022-07-10 15:12:15.610', NULL, 'int', 3, 1, 3, 2);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (6, '2022-07-10 15:12:15.610', '2022-07-10 15:12:15.610', NULL, 'bigint', 4, 1, 4, 2);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (7, '2022-07-10 15:12:15.661', '2022-07-10 15:12:15.661', NULL, 'date', 0, 1, 0, 3);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (8, '2022-07-10 15:12:15.661', '2022-07-10 15:12:15.661', NULL, 'time', 1, 1, 1, 3);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (9, '2022-07-10 15:12:15.661', '2022-07-10 15:12:15.661', NULL, 'year', 2, 1, 2, 3);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (10, '2022-07-10 15:12:15.661', '2022-07-10 15:12:15.661', NULL, 'datetime', 3, 1, 3, 3);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (11, '2022-07-10 15:12:15.661', '2022-07-10 15:12:15.661', NULL, 'timestamp', 5, 1, 5, 3);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (12, '2022-07-10 15:12:15.714', '2022-07-10 15:12:15.714', NULL, 'float', 0, 1, 0, 4);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (13, '2022-07-10 15:12:15.714', '2022-07-10 15:12:15.714', NULL, 'double', 1, 1, 1, 4);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (14, '2022-07-10 15:12:15.714', '2022-07-10 15:12:15.714', NULL, 'decimal', 2, 1, 2, 4);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (15, '2022-07-10 15:12:15.766', '2022-07-10 15:12:15.766', NULL, 'char', 0, 1, 0, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (16, '2022-07-10 15:12:15.766', '2022-07-10 15:12:15.766', NULL, 'varchar', 1, 1, 1, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (17, '2022-07-10 15:12:15.766', '2022-07-10 15:12:15.766', NULL, 'tinyblob', 2, 1, 2, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (18, '2022-07-10 15:12:15.766', '2022-07-10 15:12:15.766', NULL, 'tinytext', 3, 1, 3, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (19, '2022-07-10 15:12:15.766', '2022-07-10 15:12:15.766', NULL, 'text', 4, 1, 4, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (20, '2022-07-10 15:12:15.766', '2022-07-10 15:12:15.766', NULL, 'blob', 5, 1, 5, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (21, '2022-07-10 15:12:15.766', '2022-07-10 15:12:15.766', NULL, 'mediumblob', 6, 1, 6, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (22, '2022-07-10 15:12:15.766', '2022-07-10 15:12:15.766', NULL, 'mediumtext', 7, 1, 7, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (23, '2022-07-10 15:12:15.766', '2022-07-10 15:12:15.766', NULL, 'longblob', 8, 1, 8, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (24, '2022-07-10 15:12:15.766', '2022-07-10 15:12:15.766', NULL, 'longtext', 9, 1, 9, 5);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (25, '2022-07-10 15:12:15.816', '2022-07-10 15:12:15.816', NULL, 'tinyint', 0, 1, 0, 6);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (65, '2022-05-09 17:17:16.214', '2022-07-01 19:34:53.387', NULL, 'equals', 0, 1, 0, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (66, '2022-05-09 17:17:29.332', '2022-07-01 19:34:58.862', NULL, 'less_than', 1, 1, 1, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (67, '2022-05-09 17:17:43.377', '2022-07-01 19:35:05.146', NULL, 'less_or_equals', 2, 1, 2, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (68, '2022-05-09 17:17:55.141', '2022-07-01 19:35:12.841', NULL, 'greater_than', 3, 1, 3, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (69, '2022-05-09 17:19:17.666', '2022-07-01 19:35:37.513', NULL, 'greater_or_equals', 4, 1, 4, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (70, '2022-05-09 17:20:22.523', '2022-05-09 17:20:22.523', NULL, 'GET', 0, 1, 0, 8);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (71, '2022-05-09 17:20:32.833', '2022-05-09 17:20:32.833', NULL, 'POST', 1, 1, 1, 8);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (72, '2022-05-09 17:20:44.944', '2022-05-09 17:20:44.944', NULL, 'PUT', 2, 1, 2, 8);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (73, '2022-05-09 17:20:52.568', '2022-05-09 17:20:52.568', NULL, 'DELETE', 3, 1, 3, 8);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (74, '2022-05-09 17:21:00.473', '2022-05-09 17:21:00.473', NULL, 'HEAD', 4, 1, 4, 8);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (75, '2022-05-09 17:21:09.770', '2022-05-09 17:21:09.770', NULL, 'OPTIONS', 5, 1, 5, 8);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (76, '2022-05-09 17:21:18.312', '2022-05-09 17:21:18.312', NULL, 'PATCH', 6, 1, 6, 8);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (77, '2022-05-09 19:39:08.442', '2022-07-01 19:36:00.277', NULL, 'not_equal', 5, 1, 5, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (78, '2022-05-09 19:39:18.872', '2022-07-01 19:36:09.501', NULL, 'string_equals', 6, 1, 6, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (79, '2022-05-09 19:39:26.640', '2022-07-01 19:36:16.050', NULL, 'length_equals', 7, 1, 7, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (80, '2022-05-09 19:39:35.063', '2022-07-01 19:36:23.659', NULL, 'length_greater_than', 8, 1, 8, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (81, '2022-05-09 19:39:43.399', '2022-07-01 19:36:33.168', NULL, 'length_greater_or_equals', 9, 1, 9, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (82, '2022-05-09 19:39:55.143', '2022-05-09 19:39:55.143', NULL, 'length_less_than', 10, 1, 10, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (83, '2022-05-09 19:40:03.511', '2022-05-09 19:40:03.511', NULL, 'length_less_or_equals', 11, 1, 11, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (84, '2022-05-09 19:40:11.110', '2022-05-09 19:40:11.110', NULL, 'contains', 12, 1, 12, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (85, '2022-05-09 19:40:29.305', '2022-05-09 19:40:29.305', NULL, 'contained_by', 14, 1, 14, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (86, '2022-05-09 19:40:35.407', '2022-05-09 19:40:35.407', NULL, 'type_match', 15, 1, 15, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (87, '2022-05-09 19:40:42.703', '2022-05-09 19:40:42.703', NULL, 'regex_match', 16, 1, 16, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (88, '2022-05-09 19:40:51.383', '2022-05-09 19:40:51.383', NULL, 'startswith', 17, 1, 17, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (90, '2022-05-09 19:40:51.383', '2022-05-09 19:40:51.383', NULL, 'endswith', 18, 1, 18, 7);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (91, '2022-07-04 10:56:04.819', '2022-07-04 10:56:04.819', NULL, 'String', 1, 1, 1, 9);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (92, '2022-07-04 10:56:14.858', '2022-07-04 10:56:14.858', NULL, 'Integer', 2, 1, 2, 9);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (93, '2022-07-04 10:56:24.400', '2022-07-04 10:56:24.400', NULL, 'Float', 3, 1, 3, 9);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (94, '2022-07-04 10:56:33.618', '2022-07-04 10:56:33.618', NULL, 'Boolean', 4, 1, 4, 9);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (95, '2022-07-04 10:56:41.632', '2022-07-04 10:56:41.632', NULL, 'List', 5, 1, 5, 9);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (96, '2022-07-04 10:56:50.469', '2022-07-04 10:56:50.469', NULL, 'Dict', 6, 1, 6, 9);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (97, '2022-07-04 10:56:57.926', '2022-07-04 10:56:57.926', NULL, 'None', 7, 1, 7, 9);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (98, '2022-07-08 18:19:00.537', '2022-07-08 18:19:00.537', NULL, 'String', 1, 1, 1, 10);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (99, '2022-07-08 18:19:08.034', '2022-07-08 18:19:08.034', NULL, 'Integer', 2, 1, 2, 10);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (100, '2022-07-08 18:19:15.967', '2022-07-08 18:19:15.967', NULL, 'Float', 3, 1, 3, 10);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (101, '2022-07-08 18:19:24.003', '2022-07-08 18:19:24.003', NULL, 'Boolean', 4, 1, 4, 10);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (102, '2022-07-08 18:19:31.851', '2022-07-08 18:19:31.851', NULL, 'List', 5, 1, 5, 10);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (103, '2022-07-08 18:19:40.409', '2022-07-08 18:19:40.409', NULL, 'Dict', 6, 1, 6, 10);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (104, '2022-07-28 20:16:09.494', '2022-07-28 20:16:09.494', NULL, 'Accept', 1, 1, 1, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (105, '2022-07-28 20:16:25.095', '2022-07-28 20:16:35.183', NULL, 'Accept-Charset', 2, 1, 2, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (106, '2022-07-28 20:16:44.685', '2022-07-28 20:16:44.685', NULL, 'Accept-Language', 3, 1, 3, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (107, '2022-07-28 20:16:52.205', '2022-07-28 20:16:52.205', NULL, 'Accept-Datetime', 4, 1, 4, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (108, '2022-07-28 20:17:01.597', '2022-07-28 20:17:01.597', NULL, 'Authorization', 5, 1, 5, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (109, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Cache-Control', 6, 1, 6, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (110, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Connection', 7, 1, 7, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (111, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Cookie', 8, 1, 8, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (112, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Content-Length', 9, 1, 9, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (113, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Content-MD5', 10, 1, 10, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (114, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Content-Type', 9, 1, 9, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (115, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Expect', 10, 1, 10, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (116, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Date', 11, 1, 11, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (117, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'From', 12, 1, 12, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (118, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Host', 13, 1, 13, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (119, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'If-Match', 14, 1, 14, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (120, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'If-Modified-Since', 15, 1, 15, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (121, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'If-None-Match', 16, 1, 16, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (122, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'If-Range', 17, 1, 17, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (123, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'If-Unmodified-Since', 18, 1, 18, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (124, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Max-Forwards', 19, 1, 19, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (125, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Origin', 20, 1, 20, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (126, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Pragma', 21, 1, 21, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (127, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Proxy-Authorization', 22, 1, 22, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (128, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Range', 23, 1, 23, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (129, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Referer', 24, 1, 24, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (130, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'TE', 25, 1, 25, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (131, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'User-Agent', 26, 1, 26, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (132, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Upgrade', 27, 1, 27, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (133, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Via', 28, 1, 28, 11);
INSERT INTO `sys_dictionary_details` (`id`, `created_at`, `updated_at`, `deleted_at`, `label`, `value`, `status`, `sort`, `sys_dictionary_id`) VALUES (134, '2022-07-28 20:17:13.342', '2022-07-28 20:17:13.342', NULL, 'Warning', 29, 1, 29, 11);
COMMIT;

-- ----------------------------
-- Table structure for sys_operation_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_records`;
CREATE TABLE `sys_operation_records` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `ip` varchar(191) DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) DEFAULT NULL COMMENT '请求路径',
  `status` bigint DEFAULT NULL COMMENT '请求状态',
  `latency` bigint DEFAULT NULL COMMENT '延迟',
  `agent` varchar(191) DEFAULT NULL COMMENT '代理',
  `error_message` varchar(191) DEFAULT NULL COMMENT '错误信息',
  `body` text COMMENT '请求Body',
  `resp` text COMMENT '响应Body',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `idx_sys_operation_records_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sys_operation_records
-- ----------------------------
BEGIN;
INSERT INTO `sys_operation_records` (`id`, `created_at`, `updated_at`, `deleted_at`, `ip`, `method`, `path`, `status`, `latency`, `agent`, `error_message`, `body`, `resp`, `user_id`) VALUES (1, '2023-07-09 18:27:08.309', '2023-07-09 18:27:08.309', NULL, '219.136.74.95', 'POST', '/menu/deleteBaseMenu', 200, 11401791, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36', '', '{\"ID\":1}', '{\"code\":0,\"data\":{},\"msg\":\"删除成功\"}', 1);
INSERT INTO `sys_operation_records` (`id`, `created_at`, `updated_at`, `deleted_at`, `ip`, `method`, `path`, `status`, `latency`, `agent`, `error_message`, `body`, `resp`, `user_id`) VALUES (2, '2023-07-09 18:35:48.945', '2023-07-09 18:35:48.945', NULL, '219.136.74.95', 'POST', '/menu/addBaseMenu', 200, 5182427, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36', '', '{\"ID\":0,\"path\":\"关于我们\",\"name\":\"关于我们\",\"hidden\":false,\"parentId\":\"0\",\"component\":\"view/about/index.vue\",\"meta\":{\"title\":\"关于我们\",\"icon\":\"info-filled\",\"defaultMenu\":false,\"closeTab\":false,\"keepAlive\":true},\"parameters\":[],\"menuBtn\":[],\"sort\":0}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` (`id`, `created_at`, `updated_at`, `deleted_at`, `ip`, `method`, `path`, `status`, `latency`, `agent`, `error_message`, `body`, `resp`, `user_id`) VALUES (3, '2023-07-09 18:36:02.952', '2023-07-09 18:36:02.952', NULL, '219.136.74.95', 'POST', '/menu/addMenuAuthority', 200, 9002084, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36', '', '{\"menus\":[{\"ID\":63,\"CreatedAt\":\"2023-07-09T18:35:48.941+08:00\",\"UpdatedAt\":\"2023-07-09T18:35:48.941+08:00\",\"parentId\":\"0\",\"path\":\"关于我们\",\"name\":\"关于我们\",\"hidden\":false,\"component\":\"view/about/index.vue\",\"sort\":0,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"关于我们\",\"icon\":\"info-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":22,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"0\",\"path\":\"http://www.yangfan.gd.cn/\",\"name\":\"http://www.yangfan.gd.cn/\",\"hidden\":false,\"component\":\"/\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"官方网站\",\"icon\":\"home-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":3,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-12T18:24:59.477+08:00\",\"parentId\":\"0\",\"path\":\"admin\",\"name\":\"superAdmin\",\"hidden\":false,\"component\":\"view/superAdmin/index.vue\",\"sort\":2,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"超级管理员\",\"icon\":\"user\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":32,\"CreatedAt\":\"2022-01-19T01:09:37.274+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.921+08:00\",\"parentId\":\"3\",\"path\":\"project\",\"name\":\"project\",\"hidden\":false,\"component\":\"view/project/project.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"项目管理\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":19,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"dictionaryDetail/:id\",\"name\":\"dictionaryDetail\",\"hidden\":true,\"component\":\"view/superAdmin/dictionary/sysDictionaryDetail.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典详情-${id}\",\"icon\":\"order\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":4,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"authority\",\"name\":\"authority\",\"hidden\":false,\"component\":\"view/superAdmin/authority/authority.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色管理\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":5,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"menu\",\"name\":\"menu\",\"hidden\":false,\"component\":\"view/superAdmin/menu/menu.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"菜单管理\",\"icon\":\"tickets\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":6,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"api\",\"name\":\"api\",\"hidden\":false,\"component\":\"view/superAdmin/api/api.vue\",\"sort\":3,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"api管理\",\"icon\":\"platform\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":7,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"user\",\"name\":\"user\",\"hidden\":false,\"component\":\"view/superAdmin/user/user.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户管理\",\"icon\":\"coordinate\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":18,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"dictionary\",\"name\":\"dictionary\",\"hidden\":false,\"component\":\"view/superAdmin/dictionary/sysDictionary.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典管理\",\"icon\":\"notebook\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":20,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"operation\",\"name\":\"operation\",\"hidden\":false,\"component\":\"view/superAdmin/operation/sysOperationRecord.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"操作历史\",\"icon\":\"pie-chart\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":32,\"CreatedAt\":\"2022-01-19T01:09:37.274+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.921+08:00\",\"parentId\":\"3\",\"path\":\"project\",\"name\":\"project\",\"hidden\":false,\"component\":\"view/project/project.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"项目管理\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":19,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"dictionaryDetail/:id\",\"name\":\"dictionaryDetail\",\"hidden\":true,\"component\":\"view/superAdmin/dictionary/sysDictionaryDetail.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典详情-${id}\",\"icon\":\"order\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":4,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"authority\",\"name\":\"authority\",\"hidden\":false,\"component\":\"view/superAdmin/authority/authority.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色管理\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":5,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"menu\",\"name\":\"menu\",\"hidden\":false,\"component\":\"view/superAdmin/menu/menu.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"菜单管理\",\"icon\":\"tickets\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":6,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"api\",\"name\":\"api\",\"hidden\":false,\"component\":\"view/superAdmin/api/api.vue\",\"sort\":3,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"api管理\",\"icon\":\"platform\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":7,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"user\",\"name\":\"user\",\"hidden\":false,\"component\":\"view/superAdmin/user/user.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户管理\",\"icon\":\"coordinate\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":18,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"dictionary\",\"name\":\"dictionary\",\"hidden\":false,\"component\":\"view/superAdmin/dictionary/sysDictionary.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典管理\",\"icon\":\"notebook\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":20,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"operation\",\"name\":\"operation\",\"hidden\":false,\"component\":\"view/superAdmin/operation/sysOperationRecord.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"操作历史\",\"icon\":\"pie-chart\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":8,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-12T18:25:05.048+08:00\",\"parentId\":\"0\",\"path\":\"person\",\"name\":\"person\",\"hidden\":true,\"component\":\"view/person/person.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"个人信息\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":33,\"CreatedAt\":\"2022-01-23T18:34:58.647+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.928+08:00\",\"parentId\":\"0\",\"path\":\"interfaces\",\"name\":\"interfaces\",\"hidden\":false,\"component\":\"view/interface/index.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"接口自动化\",\"icon\":\"box\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":61,\"CreatedAt\":\"2023-05-09T17:39:19.132+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.049+08:00\",\"parentId\":\"33\",\"path\":\"env\",\"name\":\"env\",\"hidden\":false,\"component\":\"view/interface/environment/environment.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"环境变量\",\"icon\":\"grid\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":36,\"CreatedAt\":\"2022-05-11T19:34:31.473+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.945+08:00\",\"parentId\":\"33\",\"path\":\"apiConfig\",\"name\":\"apiConfig\",\"hidden\":false,\"component\":\"view/interface/interfaceTemplate/apiConfig.vue\",\"sort\":100,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"配置管理\",\"icon\":\"expand\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":34,\"CreatedAt\":\"2022-01-23T18:37:56.434+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.934+08:00\",\"parentId\":\"33\",\"path\":\"interfacetemplate\",\"name\":\"interfacetemplate\",\"hidden\":false,\"component\":\"view/interface/interfaceTemplate/interfaceTemplate.vue\",\"sort\":200,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"接口管理\",\"icon\":\"coin\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":37,\"CreatedAt\":\"2022-05-16T12:15:29.299+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.95+08:00\",\"parentId\":\"33\",\"path\":\"testCaseStep\",\"name\":\"testCaseStep\",\"hidden\":false,\"component\":\"view/interface/testCaseStep/testCaseStep.vue\",\"sort\":300,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试步骤\",\"icon\":\"suitcase\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":44,\"CreatedAt\":\"2022-06-21T17:34:55.263+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.974+08:00\",\"parentId\":\"33\",\"path\":\"testCase\",\"name\":\"testCase\",\"hidden\":false,\"component\":\"view/interface/apiCase/apiCase.vue\",\"sort\":400,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试用例\",\"icon\":\"briefcase\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":53,\"CreatedAt\":\"2022-06-21T17:34:55.263+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.003+08:00\",\"parentId\":\"33\",\"path\":\"timerTask\",\"name\":\"timerTask\",\"hidden\":false,\"component\":\"view/interface/timerTask/timerTask.vue\",\"sort\":500,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"定时任务\",\"icon\":\"timer\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":42,\"CreatedAt\":\"2022-06-07T13:37:19.135+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.961+08:00\",\"parentId\":\"33\",\"path\":\"report\",\"name\":\"report\",\"hidden\":false,\"component\":\"view/interface/Reports/report.vue\",\"sort\":600,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试报告\",\"icon\":\"compass\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":62,\"CreatedAt\":\"2023-05-09T17:39:19.144+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.057+08:00\",\"parentId\":\"33\",\"path\":\"py_pkg\",\"name\":\"py_pkg\",\"hidden\":false,\"component\":\"view/py_pkg/py_pkg.vue\",\"sort\":680,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"py库管理\",\"icon\":\"office-building\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":46,\"CreatedAt\":\"2022-07-12T11:39:29.347+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.979+08:00\",\"parentId\":\"33\",\"path\":\"debugtalk\",\"name\":\"debugtalk\",\"hidden\":false,\"component\":\"view/interface/debugtalk/debugtalk.vue\",\"sort\":700,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"驱动函数\",\"icon\":\"reading\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":47,\"CreatedAt\":\"2022-07-12T11:40:03.21+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.984+08:00\",\"parentId\":\"33\",\"path\":\"debugtalkGen\",\"name\":\"debugtalkGen\",\"hidden\":true,\"component\":\"view/interface/debugtalk/debugtalkGen.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"debugtalkGen\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":49,\"CreatedAt\":\"2022-08-01T06:08:02.383+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.989+08:00\",\"parentId\":\"33\",\"path\":\"apisCaseDetail/:id\",\"name\":\"apisCaseDetail\",\"hidden\":true,\"component\":\"view/interface/apiCase/apisCaseDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用例详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":51,\"CreatedAt\":\"2022-08-01T06:08:02.383+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.994+08:00\",\"parentId\":\"33\",\"path\":\"taskCaseDetail/:id\",\"name\":\"taskCaseDetail\",\"hidden\":true,\"component\":\"view/interface/timerTask/taskCaseDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"任务详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":35,\"CreatedAt\":\"2022-01-28T13:29:16.282+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.941+08:00\",\"parentId\":\"33\",\"path\":\"DebugReport\",\"name\":\"DebugReport\",\"hidden\":true,\"component\":\"view/interface/interfaceReport/DebugReport.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"DebugReport\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":43,\"CreatedAt\":\"2022-06-08T17:40:11.776+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.969+08:00\",\"parentId\":\"33\",\"path\":\"reportDetail/:report_id\",\"name\":\"reportDetail\",\"hidden\":true,\"component\":\"view/interface/Reports/reportDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试报告详情-${report_id}\",\"icon\":\"document\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":60,\"CreatedAt\":\"2023-05-09T17:39:19.121+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.042+08:00\",\"parentId\":\"33\",\"path\":\"jsonCompare\",\"name\":\"jsonCompare\",\"hidden\":true,\"component\":\"view/interface/interfaceComponents/jsonCompare.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"json\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":38,\"CreatedAt\":\"2022-05-17T12:54:39.571+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.955+08:00\",\"parentId\":\"33\",\"path\":\"testCaseStepDetail/:id\",\"name\":\"testCaseStepDetail\",\"hidden\":true,\"component\":\"view/interface/testCaseStep/testCaseStepDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"步骤详情-${id}\",\"icon\":\"finished\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":61,\"CreatedAt\":\"2023-05-09T17:39:19.132+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.049+08:00\",\"parentId\":\"33\",\"path\":\"env\",\"name\":\"env\",\"hidden\":false,\"component\":\"view/interface/environment/environment.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"环境变量\",\"icon\":\"grid\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":36,\"CreatedAt\":\"2022-05-11T19:34:31.473+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.945+08:00\",\"parentId\":\"33\",\"path\":\"apiConfig\",\"name\":\"apiConfig\",\"hidden\":false,\"component\":\"view/interface/interfaceTemplate/apiConfig.vue\",\"sort\":100,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"配置管理\",\"icon\":\"expand\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":34,\"CreatedAt\":\"2022-01-23T18:37:56.434+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.934+08:00\",\"parentId\":\"33\",\"path\":\"interfacetemplate\",\"name\":\"interfacetemplate\",\"hidden\":false,\"component\":\"view/interface/interfaceTemplate/interfaceTemplate.vue\",\"sort\":200,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"接口管理\",\"icon\":\"coin\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":37,\"CreatedAt\":\"2022-05-16T12:15:29.299+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.95+08:00\",\"parentId\":\"33\",\"path\":\"testCaseStep\",\"name\":\"testCaseStep\",\"hidden\":false,\"component\":\"view/interface/testCaseStep/testCaseStep.vue\",\"sort\":300,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试步骤\",\"icon\":\"suitcase\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":44,\"CreatedAt\":\"2022-06-21T17:34:55.263+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.974+08:00\",\"parentId\":\"33\",\"path\":\"testCase\",\"name\":\"testCase\",\"hidden\":false,\"component\":\"view/interface/apiCase/apiCase.vue\",\"sort\":400,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试用例\",\"icon\":\"briefcase\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":53,\"CreatedAt\":\"2022-06-21T17:34:55.263+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.003+08:00\",\"parentId\":\"33\",\"path\":\"timerTask\",\"name\":\"timerTask\",\"hidden\":false,\"component\":\"view/interface/timerTask/timerTask.vue\",\"sort\":500,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"定时任务\",\"icon\":\"timer\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":42,\"CreatedAt\":\"2022-06-07T13:37:19.135+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.961+08:00\",\"parentId\":\"33\",\"path\":\"report\",\"name\":\"report\",\"hidden\":false,\"component\":\"view/interface/Reports/report.vue\",\"sort\":600,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试报告\",\"icon\":\"compass\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":62,\"CreatedAt\":\"2023-05-09T17:39:19.144+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.057+08:00\",\"parentId\":\"33\",\"path\":\"py_pkg\",\"name\":\"py_pkg\",\"hidden\":false,\"component\":\"view/py_pkg/py_pkg.vue\",\"sort\":680,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"py库管理\",\"icon\":\"office-building\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":46,\"CreatedAt\":\"2022-07-12T11:39:29.347+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.979+08:00\",\"parentId\":\"33\",\"path\":\"debugtalk\",\"name\":\"debugtalk\",\"hidden\":false,\"component\":\"view/interface/debugtalk/debugtalk.vue\",\"sort\":700,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"驱动函数\",\"icon\":\"reading\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":47,\"CreatedAt\":\"2022-07-12T11:40:03.21+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.984+08:00\",\"parentId\":\"33\",\"path\":\"debugtalkGen\",\"name\":\"debugtalkGen\",\"hidden\":true,\"component\":\"view/interface/debugtalk/debugtalkGen.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"debugtalkGen\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":49,\"CreatedAt\":\"2022-08-01T06:08:02.383+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.989+08:00\",\"parentId\":\"33\",\"path\":\"apisCaseDetail/:id\",\"name\":\"apisCaseDetail\",\"hidden\":true,\"component\":\"view/interface/apiCase/apisCaseDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用例详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":51,\"CreatedAt\":\"2022-08-01T06:08:02.383+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.994+08:00\",\"parentId\":\"33\",\"path\":\"taskCaseDetail/:id\",\"name\":\"taskCaseDetail\",\"hidden\":true,\"component\":\"view/interface/timerTask/taskCaseDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"任务详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":35,\"CreatedAt\":\"2022-01-28T13:29:16.282+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.941+08:00\",\"parentId\":\"33\",\"path\":\"DebugReport\",\"name\":\"DebugReport\",\"hidden\":true,\"component\":\"view/interface/interfaceReport/DebugReport.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"DebugReport\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":43,\"CreatedAt\":\"2022-06-08T17:40:11.776+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.969+08:00\",\"parentId\":\"33\",\"path\":\"reportDetail/:report_id\",\"name\":\"reportDetail\",\"hidden\":true,\"component\":\"view/interface/Reports/reportDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试报告详情-${report_id}\",\"icon\":\"document\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":60,\"CreatedAt\":\"2023-05-09T17:39:19.121+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.042+08:00\",\"parentId\":\"33\",\"path\":\"jsonCompare\",\"name\":\"jsonCompare\",\"hidden\":true,\"component\":\"view/interface/interfaceComponents/jsonCompare.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"json\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":38,\"CreatedAt\":\"2022-05-17T12:54:39.571+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.955+08:00\",\"parentId\":\"33\",\"path\":\"testCaseStepDetail/:id\",\"name\":\"testCaseStepDetail\",\"hidden\":true,\"component\":\"view/interface/testCaseStep/testCaseStepDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"步骤详情-${id}\",\"icon\":\"finished\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":55,\"CreatedAt\":\"2022-10-08T19:30:47.868+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.017+08:00\",\"parentId\":\"0\",\"path\":\"performance\",\"name\":\"performance\",\"hidden\":false,\"component\":\"view/performance/index.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试\",\"icon\":\"stopwatch\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":56,\"CreatedAt\":\"2022-10-08T19:51:05.009+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.022+08:00\",\"parentId\":\"55\",\"path\":\"performanceTask\",\"name\":\"performanceTask\",\"hidden\":false,\"component\":\"view/performance/task/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能任务\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":58,\"CreatedAt\":\"2022-11-03T15:49:19.428+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.033+08:00\",\"parentId\":\"55\",\"path\":\"pReport\",\"name\":\"pReport\",\"hidden\":false,\"component\":\"view/performance/report.vue\",\"sort\":2,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试报告\",\"icon\":\"compass\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":59,\"CreatedAt\":\"2022-11-03T18:59:54.997+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.037+08:00\",\"parentId\":\"55\",\"path\":\"pReportDetail/:id\",\"name\":\"pReportDetail\",\"hidden\":true,\"component\":\"view/performance/reportDetail.vue\",\"sort\":999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试报告详情-${id}\",\"icon\":\"document\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":57,\"CreatedAt\":\"2022-10-11T14:27:10.146+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.028+08:00\",\"parentId\":\"55\",\"path\":\"performanceDetail/:id\",\"name\":\"performanceDetail\",\"hidden\":true,\"component\":\"view/performance/task/taskDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能任务详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":56,\"CreatedAt\":\"2022-10-08T19:51:05.009+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.022+08:00\",\"parentId\":\"55\",\"path\":\"performanceTask\",\"name\":\"performanceTask\",\"hidden\":false,\"component\":\"view/performance/task/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能任务\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":58,\"CreatedAt\":\"2022-11-03T15:49:19.428+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.033+08:00\",\"parentId\":\"55\",\"path\":\"pReport\",\"name\":\"pReport\",\"hidden\":false,\"component\":\"view/performance/report.vue\",\"sort\":2,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试报告\",\"icon\":\"compass\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":59,\"CreatedAt\":\"2022-11-03T18:59:54.997+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.037+08:00\",\"parentId\":\"55\",\"path\":\"pReportDetail/:id\",\"name\":\"pReportDetail\",\"hidden\":true,\"component\":\"view/performance/reportDetail.vue\",\"sort\":999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试报告详情-${id}\",\"icon\":\"document\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":57,\"CreatedAt\":\"2022-10-11T14:27:10.146+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.028+08:00\",\"parentId\":\"55\",\"path\":\"performanceDetail/:id\",\"name\":\"performanceDetail\",\"hidden\":true,\"component\":\"view/performance/task/taskDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能任务详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":14,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-11-22T14:50:24.888+08:00\",\"parentId\":\"0\",\"path\":\"systemTools\",\"name\":\"systemTools\",\"hidden\":false,\"component\":\"view/systemTools/index.vue\",\"sort\":15,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统工具\",\"icon\":\"tools\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":26,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"autoPkg\",\"name\":\"autoPkg\",\"hidden\":false,\"component\":\"view/systemTools/autoPkg/autoPkg.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化package\",\"icon\":\"folder\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":25,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeEdit/:id\",\"name\":\"autoCodeEdit\",\"hidden\":true,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码-${id}\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":24,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeAdmin\",\"name\":\"autoCodeAdmin\",\"hidden\":false,\"component\":\"view/systemTools/autoCodeAdmin/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码管理\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":15,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"autoCode\",\"name\":\"autoCode\",\"hidden\":false,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"代码生成器\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":16,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"formCreate\",\"name\":\"formCreate\",\"hidden\":false,\"component\":\"view/systemTools/formCreate/index.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"表单生成器\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":17,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"system\",\"name\":\"system\",\"hidden\":false,\"component\":\"view/systemTools/system/system.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统配置\",\"icon\":\"operation\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":26,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"autoPkg\",\"name\":\"autoPkg\",\"hidden\":false,\"component\":\"view/systemTools/autoPkg/autoPkg.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化package\",\"icon\":\"folder\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":25,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeEdit/:id\",\"name\":\"autoCodeEdit\",\"hidden\":true,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码-${id}\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":24,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeAdmin\",\"name\":\"autoCodeAdmin\",\"hidden\":false,\"component\":\"view/systemTools/autoCodeAdmin/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码管理\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":15,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"autoCode\",\"name\":\"autoCode\",\"hidden\":false,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"代码生成器\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":16,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"formCreate\",\"name\":\"formCreate\",\"hidden\":false,\"component\":\"view/systemTools/formCreate/index.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"表单生成器\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":17,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"system\",\"name\":\"system\",\"hidden\":false,\"component\":\"view/systemTools/system/system.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统配置\",\"icon\":\"operation\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":27,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-11-22T14:50:34.751+08:00\",\"parentId\":\"0\",\"path\":\"plugin\",\"name\":\"plugin\",\"hidden\":false,\"component\":\"view/routerHolder.vue\",\"sort\":16,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件系统\",\"icon\":\"cherry\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":28,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"27\",\"path\":\"https://plugin.gin-vue-admin.com/\",\"name\":\"https://plugin.gin-vue-admin.com/\",\"hidden\":false,\"component\":\"https://plugin.gin-vue-admin.com/\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件市场\",\"icon\":\"shop\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":29,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"27\",\"path\":\"installPlugin\",\"name\":\"installPlugin\",\"hidden\":false,\"component\":\"view/systemTools/installPlugin/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件安装\",\"icon\":\"box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":30,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"27\",\"path\":\"autoPlug\",\"name\":\"autoPlug\",\"hidden\":false,\"component\":\"view/systemTools/autoPlug/autoPlug.vue\",\"sort\":2,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件模板\",\"icon\":\"folder\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":31,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"27\",\"path\":\"plugin-email\",\"name\":\"plugin-email\",\"hidden\":false,\"component\":\"plugin/email/view/index.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"邮件插件\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":28,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"27\",\"path\":\"https://plugin.gin-vue-admin.com/\",\"name\":\"https://plugin.gin-vue-admin.com/\",\"hidden\":false,\"component\":\"https://plugin.gin-vue-admin.com/\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件市场\",\"icon\":\"shop\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":29,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"27\",\"path\":\"installPlugin\",\"name\":\"installPlugin\",\"hidden\":false,\"component\":\"view/systemTools/installPlugin/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件安装\",\"icon\":\"box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":30,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"27\",\"path\":\"autoPlug\",\"name\":\"autoPlug\",\"hidden\":false,\"component\":\"view/systemTools/autoPlug/autoPlug.vue\",\"sort\":2,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件模板\",\"icon\":\"folder\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":31,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"27\",\"path\":\"plugin-email\",\"name\":\"plugin-email\",\"hidden\":false,\"component\":\"plugin/email/view/index.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"邮件插件\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":9,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-11-22T14:50:42.311+08:00\",\"parentId\":\"0\",\"path\":\"example\",\"name\":\"example\",\"hidden\":false,\"component\":\"view/example/index.vue\",\"sort\":17,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"示例文件\",\"icon\":\"management\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":10,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"excel\",\"name\":\"excel\",\"hidden\":false,\"component\":\"view/example/excel/excel.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"excel导入导出\",\"icon\":\"takeaway-box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":11,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"upload\",\"name\":\"upload\",\"hidden\":false,\"component\":\"view/example/upload/upload.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"媒体库（上传下载）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":12,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"breakpoint\",\"name\":\"breakpoint\",\"hidden\":false,\"component\":\"view/example/breakpoint/breakpoint.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传\",\"icon\":\"upload-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":21,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"simpleUploader\",\"name\":\"simpleUploader\",\"hidden\":false,\"component\":\"view/example/simpleUploader/simpleUploader\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传（插件版）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":13,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"customer\",\"name\":\"customer\",\"hidden\":false,\"component\":\"view/example/customer/customer.vue\",\"sort\":7,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"客户列表（资源示例）\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":10,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"excel\",\"name\":\"excel\",\"hidden\":false,\"component\":\"view/example/excel/excel.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"excel导入导出\",\"icon\":\"takeaway-box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":11,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"upload\",\"name\":\"upload\",\"hidden\":false,\"component\":\"view/example/upload/upload.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"媒体库（上传下载）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":12,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"breakpoint\",\"name\":\"breakpoint\",\"hidden\":false,\"component\":\"view/example/breakpoint/breakpoint.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传\",\"icon\":\"upload-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":13,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"customer\",\"name\":\"customer\",\"hidden\":false,\"component\":\"view/example/customer/customer.vue\",\"sort\":7,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"客户列表（资源示例）\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":23,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-11-22T14:50:49.106+08:00\",\"parentId\":\"0\",\"path\":\"state\",\"name\":\"state\",\"hidden\":false,\"component\":\"view/system/state.vue\",\"sort\":18,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"服务器状态\",\"icon\":\"cloudy\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"authorityId\":888}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` (`id`, `created_at`, `updated_at`, `deleted_at`, `ip`, `method`, `path`, `status`, `latency`, `agent`, `error_message`, `body`, `resp`, `user_id`) VALUES (4, '2023-07-09 18:36:14.737', '2023-07-09 18:36:14.737', NULL, '219.136.74.95', 'POST', '/casbin/updateCasbin', 200, 56126802, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36', '', '{\"authorityId\":888,\"casbinInfos\":[{\"path\":\"/base/login\",\"method\":\"POST\"},{\"path\":\"/jwt/jsonInBlacklist\",\"method\":\"POST\"},{\"path\":\"/user/deleteUser\",\"method\":\"DELETE\"},{\"path\":\"/user/admin_register\",\"method\":\"POST\"},{\"path\":\"/user/getUserList\",\"method\":\"POST\"},{\"path\":\"/user/setUserInfo\",\"method\":\"PUT\"},{\"path\":\"/user/setSelfInfo\",\"method\":\"PUT\"},{\"path\":\"/user/getUserInfo\",\"method\":\"GET\"},{\"path\":\"/user/setUserAuthorities\",\"method\":\"POST\"},{\"path\":\"/user/changePassword\",\"method\":\"POST\"},{\"path\":\"/user/setUserAuthority\",\"method\":\"POST\"},{\"path\":\"/user/resetPassword\",\"method\":\"POST\"},{\"path\":\"/user/setUserProjects\",\"method\":\"POST\"},{\"path\":\"/api/createApi\",\"method\":\"POST\"},{\"path\":\"/api/deleteApi\",\"method\":\"POST\"},{\"path\":\"/api/updateApi\",\"method\":\"POST\"},{\"path\":\"/api/getApiList\",\"method\":\"POST\"},{\"path\":\"/api/getAllApis\",\"method\":\"POST\"},{\"path\":\"/api/getApiById\",\"method\":\"POST\"},{\"path\":\"/api/deleteApisByIds\",\"method\":\"DELETE\"},{\"path\":\"/api/createApiMenu\",\"method\":\"POST\"},{\"path\":\"/api/deleteApiMenu\",\"method\":\"DELETE\"},{\"path\":\"/api/deleteApiMenuByIds\",\"method\":\"DELETE\"},{\"path\":\"/api/updateApiMenu\",\"method\":\"PUT\"},{\"path\":\"/api/findApiMenu\",\"method\":\"GET\"},{\"path\":\"/api/getApiMenuList\",\"method\":\"GET\"},{\"path\":\"/authority/copyAuthority\",\"method\":\"POST\"},{\"path\":\"/authority/createAuthority\",\"method\":\"POST\"},{\"path\":\"/authority/deleteAuthority\",\"method\":\"POST\"},{\"path\":\"/authority/updateAuthority\",\"method\":\"PUT\"},{\"path\":\"/authority/getAuthorityList\",\"method\":\"POST\"},{\"path\":\"/authority/setDataAuthority\",\"method\":\"POST\"},{\"path\":\"/casbin/updateCasbin\",\"method\":\"POST\"},{\"path\":\"/casbin/getPolicyPathByAuthorityId\",\"method\":\"POST\"},{\"path\":\"/menu/addBaseMenu\",\"method\":\"POST\"},{\"path\":\"/menu/getMenu\",\"method\":\"POST\"},{\"path\":\"/menu/deleteBaseMenu\",\"method\":\"POST\"},{\"path\":\"/menu/updateBaseMenu\",\"method\":\"POST\"},{\"path\":\"/menu/getBaseMenuById\",\"method\":\"POST\"},{\"path\":\"/menu/getMenuList\",\"method\":\"POST\"},{\"path\":\"/menu/getBaseMenuTree\",\"method\":\"POST\"},{\"path\":\"/menu/getMenuAuthority\",\"method\":\"POST\"},{\"path\":\"/menu/addMenuAuthority\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/findFile\",\"method\":\"GET\"},{\"path\":\"/fileUploadAndDownload/breakpointContinue\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/breakpointContinueFinish\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/removeChunk\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/upload\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/deleteFile\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/editFileName\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/getFileList\",\"method\":\"POST\"},{\"path\":\"/system/getServerInfo\",\"method\":\"POST\"},{\"path\":\"/system/getSystemConfig\",\"method\":\"POST\"},{\"path\":\"/system/setSystemConfig\",\"method\":\"POST\"},{\"path\":\"/customer/customer\",\"method\":\"PUT\"},{\"path\":\"/customer/customer\",\"method\":\"POST\"},{\"path\":\"/customer/customer\",\"method\":\"DELETE\"},{\"path\":\"/customer/customer\",\"method\":\"GET\"},{\"path\":\"/customer/customerList\",\"method\":\"GET\"},{\"path\":\"/autoCode/getDB\",\"method\":\"GET\"},{\"path\":\"/autoCode/getTables\",\"method\":\"GET\"},{\"path\":\"/autoCode/createTemp\",\"method\":\"POST\"},{\"path\":\"/autoCode/preview\",\"method\":\"POST\"},{\"path\":\"/autoCode/getColumn\",\"method\":\"GET\"},{\"path\":\"/autoCode/createPlug\",\"method\":\"POST\"},{\"path\":\"/autoCode/installPlugin\",\"method\":\"POST\"},{\"path\":\"/autoCode/createPackage\",\"method\":\"POST\"},{\"path\":\"/autoCode/getPackage\",\"method\":\"POST\"},{\"path\":\"/autoCode/delPackage\",\"method\":\"POST\"},{\"path\":\"/autoCode/getMeta\",\"method\":\"POST\"},{\"path\":\"/autoCode/rollback\",\"method\":\"POST\"},{\"path\":\"/autoCode/getSysHistory\",\"method\":\"POST\"},{\"path\":\"/autoCode/delSysHistory\",\"method\":\"POST\"},{\"path\":\"/sysDictionaryDetail/updateSysDictionaryDetail\",\"method\":\"PUT\"},{\"path\":\"/sysDictionaryDetail/createSysDictionaryDetail\",\"method\":\"POST\"},{\"path\":\"/sysDictionaryDetail/deleteSysDictionaryDetail\",\"method\":\"DELETE\"},{\"path\":\"/sysDictionaryDetail/findSysDictionaryDetail\",\"method\":\"GET\"},{\"path\":\"/sysDictionaryDetail/getSysDictionaryDetailList\",\"method\":\"GET\"},{\"path\":\"/sysDictionary/createSysDictionary\",\"method\":\"POST\"},{\"path\":\"/sysDictionary/deleteSysDictionary\",\"method\":\"DELETE\"},{\"path\":\"/sysDictionary/updateSysDictionary\",\"method\":\"PUT\"},{\"path\":\"/sysDictionary/findSysDictionary\",\"method\":\"GET\"},{\"path\":\"/sysDictionary/getSysDictionaryList\",\"method\":\"GET\"},{\"path\":\"/sysOperationRecord/createSysOperationRecord\",\"method\":\"POST\"},{\"path\":\"/sysOperationRecord/findSysOperationRecord\",\"method\":\"GET\"},{\"path\":\"/sysOperationRecord/getSysOperationRecordList\",\"method\":\"GET\"},{\"path\":\"/sysOperationRecord/deleteSysOperationRecord\",\"method\":\"DELETE\"},{\"path\":\"/sysOperationRecord/deleteSysOperationRecordByIds\",\"method\":\"DELETE\"},{\"path\":\"/simpleUploader/upload\",\"method\":\"POST\"},{\"path\":\"/simpleUploader/checkFileMd5\",\"method\":\"GET\"},{\"path\":\"/simpleUploader/mergeFileMd5\",\"method\":\"GET\"},{\"path\":\"/email/emailTest\",\"method\":\"POST\"},{\"path\":\"/email/emailSend\",\"method\":\"POST\"},{\"path\":\"/excel/importExcel\",\"method\":\"POST\"},{\"path\":\"/excel/loadExcel\",\"method\":\"GET\"},{\"path\":\"/excel/exportExcel\",\"method\":\"POST\"},{\"path\":\"/excel/downloadTemplate\",\"method\":\"GET\"},{\"path\":\"/authorityBtn/setAuthorityBtn\",\"method\":\"POST\"},{\"path\":\"/authorityBtn/getAuthorityBtn\",\"method\":\"POST\"},{\"path\":\"/authorityBtn/canRemoveAuthorityBtn\",\"method\":\"POST\"},{\"path\":\"/project/createProject\",\"method\":\"POST\"},{\"path\":\"/project/deleteProject\",\"method\":\"DELETE\"},{\"path\":\"/project/deleteProjectByIds\",\"method\":\"DELETE\"},{\"path\":\"/project/updateProject\",\"method\":\"PUT\"},{\"path\":\"/project/findProject\",\"method\":\"GET\"},{\"path\":\"/project/getProjectList\",\"method\":\"GET\"},{\"path\":\"/case/:project/createApiMenu\",\"method\":\"POST\"},{\"path\":\"/case/:project/deleteApiMenu\",\"method\":\"DELETE\"},{\"path\":\"/case/:project/deleteApiMenuByIds\",\"method\":\"DELETE\"},{\"path\":\"/case/:project/updateApiMenu\",\"method\":\"PUT\"},{\"path\":\"/case/:project/findApiMenu\",\"method\":\"GET\"},{\"path\":\"/case/:project/getApiMenuList\",\"method\":\"GET\"},{\"path\":\"/case/:project/createInterfaceTemplate\",\"method\":\"POST\"},{\"path\":\"/case/:project/deleteInterfaceTemplate\",\"method\":\"DELETE\"},{\"path\":\"/case/:project/deleteInterfaceTemplateByIds\",\"method\":\"DELETE\"},{\"path\":\"/case/:project/updateInterfaceTemplate\",\"method\":\"PUT\"},{\"path\":\"/case/:project/findInterfaceTemplate\",\"method\":\"GET\"},{\"path\":\"/case/:project/getInterfaceTemplateList\",\"method\":\"GET\"},{\"path\":\"/case/:project/updateDebugTalk\",\"method\":\"PUT\"},{\"path\":\"/case/:project/getDebugTalk\",\"method\":\"POST\"},{\"path\":\"/case/:project/createDebugTalk\",\"method\":\"POST\"},{\"path\":\"/case/:project/deleteDebugTalk\",\"method\":\"POST\"},{\"path\":\"/case/:project/getDebugTalkList\",\"method\":\"POST\"},{\"path\":\"/case/:project/getGrpc\",\"method\":\"POST\"},{\"path\":\"/case/:project/createUserConfig\",\"method\":\"POST\"},{\"path\":\"/case/:project/getUserConfig\",\"method\":\"GET\"},{\"path\":\"/ac/:project/createApiConfig\",\"method\":\"POST\"},{\"path\":\"/ac/:project/deleteApiConfig\",\"method\":\"DELETE\"},{\"path\":\"/ac/:project/deleteApiConfigByIds\",\"method\":\"DELETE\"},{\"path\":\"/ac/:project/updateApiConfig\",\"method\":\"PUT\"},{\"path\":\"/ac/:project/findApiConfig\",\"method\":\"GET\"},{\"path\":\"/ac/:project/getApiConfigList\",\"method\":\"GET\"},{\"path\":\"/case/:project/step/createTestCase\",\"method\":\"POST\"},{\"path\":\"/case/:project/step/deleteTestCase\",\"method\":\"DELETE\"},{\"path\":\"/case/:project/step/deleteTestCaseByIds\",\"method\":\"DELETE\"},{\"path\":\"/case/:project/step/updateTestCase\",\"method\":\"PUT\"},{\"path\":\"/case/:project/step/findTestCase\",\"method\":\"GET\"},{\"path\":\"/case/:project/step/getTestCaseList\",\"method\":\"GET\"},{\"path\":\"/case/:project/step/getTestCaseList\",\"method\":\"POST\"},{\"path\":\"/case/:project/step/addTestCase\",\"method\":\"POST\"},{\"path\":\"/case/:project/step/delTestCase\",\"method\":\"DELETE\"},{\"path\":\"/case/:project/step/sortTestCase\",\"method\":\"POST\"},{\"path\":\"/case/run/:project/runTestCaseStep\",\"method\":\"POST\"},{\"path\":\"/case/run/:project/runApiCase\",\"method\":\"POST\"},{\"path\":\"/case/run/:project/runApi\",\"method\":\"POST\"},{\"path\":\"/case/run/:project/runTimerTask\",\"method\":\"POST\"},{\"path\":\"/case/run/:project/runBoomerDebug\",\"method\":\"POST\"},{\"path\":\"/case/run/:project/runBoomer\",\"method\":\"POST\"},{\"path\":\"/case/run/:project/rebalance\",\"method\":\"POST\"},{\"path\":\"/case/run/:project/stop\",\"method\":\"GET\"},{\"path\":\"/case/report/:project/getReportList\",\"method\":\"GET\"},{\"path\":\"/case/report/:project/findReport\",\"method\":\"GET\"},{\"path\":\"/case/report/:project/delReport\",\"method\":\"DELETE\"},{\"path\":\"/testcase/:project/createApiCase\",\"method\":\"POST\"},{\"path\":\"/testcase/:project/deleteApiCase\",\"method\":\"DELETE\"},{\"path\":\"/testcase/:project/deleteApiCaseByIds\",\"method\":\"DELETE\"},{\"path\":\"/testcase/:project/updateApiCase\",\"method\":\"PUT\"},{\"path\":\"/testcase/:project/getApiCaseList\",\"method\":\"GET\"},{\"path\":\"/testcase/:project/findApiTestCase\",\"method\":\"GET\"},{\"path\":\"/testcase/:project/addApisCase\",\"method\":\"POST\"},{\"path\":\"/testcase/:project/setApisCase\",\"method\":\"POST\"},{\"path\":\"/testcase/:project/sortApisCase\",\"method\":\"POST\"},{\"path\":\"/testcase/:project/AddApiTestCase\",\"method\":\"POST\"},{\"path\":\"/testcase/:project/delApisCase\",\"method\":\"DELETE\"},{\"path\":\"/testcase/:project/findApiCase\",\"method\":\"GET\"},{\"path\":\"/task/:project/sortTaskCase\",\"method\":\"POST\"},{\"path\":\"/task/:project/addTaskCase\",\"method\":\"POST\"},{\"path\":\"/task/:project/delTaskCase\",\"method\":\"DELETE\"},{\"path\":\"/task/:project/findTaskTestCase\",\"method\":\"GET\"},{\"path\":\"/task/:project/addTaskTestCase\",\"method\":\"POST\"},{\"path\":\"/task/:project/setTaskCase\",\"method\":\"POST\"},{\"path\":\"/task/:project/createTimerTask\",\"method\":\"POST\"},{\"path\":\"/task/:project/deleteTimerTask\",\"method\":\"DELETE\"},{\"path\":\"/task/:project/deleteTimerTaskByIds\",\"method\":\"DELETE\"},{\"path\":\"/task/:project/updateTimerTask\",\"method\":\"PUT\"},{\"path\":\"/task/:project/findTimerTask\",\"method\":\"GET\"},{\"path\":\"/task/:project/getTimerTaskList\",\"method\":\"GET\"},{\"path\":\"/task/:project/getTimerTaskTagList\",\"method\":\"GET\"},{\"path\":\"/task/:project/createTimerTaskTag\",\"method\":\"POST\"},{\"path\":\"/task/:project/deleteTimerTaskTag\",\"method\":\"DELETE\"},{\"path\":\"/performance/:project/createPerformance\",\"method\":\"POST\"},{\"path\":\"/performance/:project/getPerformanceList\",\"method\":\"GET\"},{\"path\":\"/performance/:project/deletePerformance\",\"method\":\"DELETE\"},{\"path\":\"/performance/:project/findPerformance\",\"method\":\"GET\"},{\"path\":\"/performance/:project/updatePerformance\",\"method\":\"PUT\"},{\"path\":\"/performance/:project/addPerformanceCase\",\"method\":\"POST\"},{\"path\":\"/performance/:project/sortPerformanceCase\",\"method\":\"POST\"},{\"path\":\"/performance/:project/delPerformanceCase\",\"method\":\"DELETE\"},{\"path\":\"/performance/:project/findPerformanceCase\",\"method\":\"GET\"},{\"path\":\"/performance/:project/addOperation\",\"method\":\"POST\"},{\"path\":\"/performance/:project/findPerformanceStep\",\"method\":\"GET\"},{\"path\":\"/performance/:project/getReportList\",\"method\":\"GET\"},{\"path\":\"/performance/:project/findReport\",\"method\":\"GET\"},{\"path\":\"/case/:project/pyPkg/installPyPkg\",\"method\":\"POST\"},{\"path\":\"/case/:project/pyPkg/uninstallPyPkg\",\"method\":\"POST\"},{\"path\":\"/case/:project/pyPkg/updatePyPkg\",\"method\":\"POST\"},{\"path\":\"/case/:project/pyPkg/getPkgVersionList\",\"method\":\"POST\"},{\"path\":\"/case/:project/pyPkg/pyPkgList\",\"method\":\"GET\"},{\"path\":\"/env/:project/createEnv\",\"method\":\"POST\"},{\"path\":\"/env/:project/updateEnv\",\"method\":\"PUT\"},{\"path\":\"/env/:project/deleteEnv\",\"method\":\"DELETE\"},{\"path\":\"/env/:project/findEnv\",\"method\":\"POST\"},{\"path\":\"/env/:project/getEnvList\",\"method\":\"GET\"},{\"path\":\"/env/:project/createEnvVariable\",\"method\":\"POST\"},{\"path\":\"/env/:project/deleteEnvVariable\",\"method\":\"DELETE\"},{\"path\":\"/env/:project/findEnvVariable\",\"method\":\"GET\"},{\"path\":\"/env/:project/getEnvVariableList\",\"method\":\"GET\"}]}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` (`id`, `created_at`, `updated_at`, `deleted_at`, `ip`, `method`, `path`, `status`, `latency`, `agent`, `error_message`, `body`, `resp`, `user_id`) VALUES (5, '2023-07-09 18:36:40.420', '2023-07-09 18:36:40.420', NULL, '219.136.74.95', 'PUT', '/authority/updateAuthority', 200, 6528500, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36', '', '{\"authorityId\":888,\"AuthorityName\":\"普通用户\",\"parentId\":0,\"defaultRouter\":\"关于我们\"}', '{\"code\":0,\"data\":{\"authority\":{\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"authorityId\":888,\"authorityName\":\"普通用户\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"关于我们\"}},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` (`id`, `created_at`, `updated_at`, `deleted_at`, `ip`, `method`, `path`, `status`, `latency`, `agent`, `error_message`, `body`, `resp`, `user_id`) VALUES (6, '2023-07-09 18:36:42.297', '2023-07-09 18:36:42.297', NULL, '219.136.74.95', 'POST', '/menu/addMenuAuthority', 200, 11010720, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36', '', '{\"menus\":[{\"ID\":63,\"CreatedAt\":\"2023-07-09T18:35:48.941+08:00\",\"UpdatedAt\":\"2023-07-09T18:35:48.941+08:00\",\"parentId\":\"0\",\"path\":\"关于我们\",\"name\":\"关于我们\",\"hidden\":false,\"component\":\"view/about/index.vue\",\"sort\":0,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"关于我们\",\"icon\":\"info-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":22,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"0\",\"path\":\"http://www.yangfan.gd.cn/\",\"name\":\"http://www.yangfan.gd.cn/\",\"hidden\":false,\"component\":\"/\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"官方网站\",\"icon\":\"home-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":3,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-12T18:24:59.477+08:00\",\"parentId\":\"0\",\"path\":\"admin\",\"name\":\"superAdmin\",\"hidden\":false,\"component\":\"view/superAdmin/index.vue\",\"sort\":2,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"超级管理员\",\"icon\":\"user\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":32,\"CreatedAt\":\"2022-01-19T01:09:37.274+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.921+08:00\",\"parentId\":\"3\",\"path\":\"project\",\"name\":\"project\",\"hidden\":false,\"component\":\"view/project/project.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"项目管理\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":19,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"dictionaryDetail/:id\",\"name\":\"dictionaryDetail\",\"hidden\":true,\"component\":\"view/superAdmin/dictionary/sysDictionaryDetail.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典详情-${id}\",\"icon\":\"order\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":4,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"authority\",\"name\":\"authority\",\"hidden\":false,\"component\":\"view/superAdmin/authority/authority.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色管理\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":5,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"menu\",\"name\":\"menu\",\"hidden\":false,\"component\":\"view/superAdmin/menu/menu.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"菜单管理\",\"icon\":\"tickets\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":6,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"api\",\"name\":\"api\",\"hidden\":false,\"component\":\"view/superAdmin/api/api.vue\",\"sort\":3,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"api管理\",\"icon\":\"platform\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":7,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"user\",\"name\":\"user\",\"hidden\":false,\"component\":\"view/superAdmin/user/user.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户管理\",\"icon\":\"coordinate\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":18,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"dictionary\",\"name\":\"dictionary\",\"hidden\":false,\"component\":\"view/superAdmin/dictionary/sysDictionary.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典管理\",\"icon\":\"notebook\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":20,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"operation\",\"name\":\"operation\",\"hidden\":false,\"component\":\"view/superAdmin/operation/sysOperationRecord.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"操作历史\",\"icon\":\"pie-chart\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":32,\"CreatedAt\":\"2022-01-19T01:09:37.274+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.921+08:00\",\"parentId\":\"3\",\"path\":\"project\",\"name\":\"project\",\"hidden\":false,\"component\":\"view/project/project.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"项目管理\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":19,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"dictionaryDetail/:id\",\"name\":\"dictionaryDetail\",\"hidden\":true,\"component\":\"view/superAdmin/dictionary/sysDictionaryDetail.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典详情-${id}\",\"icon\":\"order\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":4,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"authority\",\"name\":\"authority\",\"hidden\":false,\"component\":\"view/superAdmin/authority/authority.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色管理\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":5,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"menu\",\"name\":\"menu\",\"hidden\":false,\"component\":\"view/superAdmin/menu/menu.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"菜单管理\",\"icon\":\"tickets\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":6,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"api\",\"name\":\"api\",\"hidden\":false,\"component\":\"view/superAdmin/api/api.vue\",\"sort\":3,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"api管理\",\"icon\":\"platform\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":7,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"user\",\"name\":\"user\",\"hidden\":false,\"component\":\"view/superAdmin/user/user.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户管理\",\"icon\":\"coordinate\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":18,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"dictionary\",\"name\":\"dictionary\",\"hidden\":false,\"component\":\"view/superAdmin/dictionary/sysDictionary.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典管理\",\"icon\":\"notebook\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":20,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"3\",\"path\":\"operation\",\"name\":\"operation\",\"hidden\":false,\"component\":\"view/superAdmin/operation/sysOperationRecord.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"操作历史\",\"icon\":\"pie-chart\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":8,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-12T18:25:05.048+08:00\",\"parentId\":\"0\",\"path\":\"person\",\"name\":\"person\",\"hidden\":true,\"component\":\"view/person/person.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"个人信息\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":33,\"CreatedAt\":\"2022-01-23T18:34:58.647+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.928+08:00\",\"parentId\":\"0\",\"path\":\"interfaces\",\"name\":\"interfaces\",\"hidden\":false,\"component\":\"view/interface/index.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"接口自动化\",\"icon\":\"box\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":61,\"CreatedAt\":\"2023-05-09T17:39:19.132+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.049+08:00\",\"parentId\":\"33\",\"path\":\"env\",\"name\":\"env\",\"hidden\":false,\"component\":\"view/interface/environment/environment.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"环境变量\",\"icon\":\"grid\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":36,\"CreatedAt\":\"2022-05-11T19:34:31.473+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.945+08:00\",\"parentId\":\"33\",\"path\":\"apiConfig\",\"name\":\"apiConfig\",\"hidden\":false,\"component\":\"view/interface/interfaceTemplate/apiConfig.vue\",\"sort\":100,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"配置管理\",\"icon\":\"expand\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":34,\"CreatedAt\":\"2022-01-23T18:37:56.434+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.934+08:00\",\"parentId\":\"33\",\"path\":\"interfacetemplate\",\"name\":\"interfacetemplate\",\"hidden\":false,\"component\":\"view/interface/interfaceTemplate/interfaceTemplate.vue\",\"sort\":200,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"接口管理\",\"icon\":\"coin\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":37,\"CreatedAt\":\"2022-05-16T12:15:29.299+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.95+08:00\",\"parentId\":\"33\",\"path\":\"testCaseStep\",\"name\":\"testCaseStep\",\"hidden\":false,\"component\":\"view/interface/testCaseStep/testCaseStep.vue\",\"sort\":300,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试步骤\",\"icon\":\"suitcase\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":44,\"CreatedAt\":\"2022-06-21T17:34:55.263+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.974+08:00\",\"parentId\":\"33\",\"path\":\"testCase\",\"name\":\"testCase\",\"hidden\":false,\"component\":\"view/interface/apiCase/apiCase.vue\",\"sort\":400,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试用例\",\"icon\":\"briefcase\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":53,\"CreatedAt\":\"2022-06-21T17:34:55.263+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.003+08:00\",\"parentId\":\"33\",\"path\":\"timerTask\",\"name\":\"timerTask\",\"hidden\":false,\"component\":\"view/interface/timerTask/timerTask.vue\",\"sort\":500,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"定时任务\",\"icon\":\"timer\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":42,\"CreatedAt\":\"2022-06-07T13:37:19.135+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.961+08:00\",\"parentId\":\"33\",\"path\":\"report\",\"name\":\"report\",\"hidden\":false,\"component\":\"view/interface/Reports/report.vue\",\"sort\":600,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试报告\",\"icon\":\"compass\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":62,\"CreatedAt\":\"2023-05-09T17:39:19.144+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.057+08:00\",\"parentId\":\"33\",\"path\":\"py_pkg\",\"name\":\"py_pkg\",\"hidden\":false,\"component\":\"view/py_pkg/py_pkg.vue\",\"sort\":680,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"py库管理\",\"icon\":\"office-building\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":46,\"CreatedAt\":\"2022-07-12T11:39:29.347+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.979+08:00\",\"parentId\":\"33\",\"path\":\"debugtalk\",\"name\":\"debugtalk\",\"hidden\":false,\"component\":\"view/interface/debugtalk/debugtalk.vue\",\"sort\":700,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"驱动函数\",\"icon\":\"reading\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":47,\"CreatedAt\":\"2022-07-12T11:40:03.21+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.984+08:00\",\"parentId\":\"33\",\"path\":\"debugtalkGen\",\"name\":\"debugtalkGen\",\"hidden\":true,\"component\":\"view/interface/debugtalk/debugtalkGen.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"debugtalkGen\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":49,\"CreatedAt\":\"2022-08-01T06:08:02.383+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.989+08:00\",\"parentId\":\"33\",\"path\":\"apisCaseDetail/:id\",\"name\":\"apisCaseDetail\",\"hidden\":true,\"component\":\"view/interface/apiCase/apisCaseDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用例详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":51,\"CreatedAt\":\"2022-08-01T06:08:02.383+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.994+08:00\",\"parentId\":\"33\",\"path\":\"taskCaseDetail/:id\",\"name\":\"taskCaseDetail\",\"hidden\":true,\"component\":\"view/interface/timerTask/taskCaseDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"任务详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":35,\"CreatedAt\":\"2022-01-28T13:29:16.282+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.941+08:00\",\"parentId\":\"33\",\"path\":\"DebugReport\",\"name\":\"DebugReport\",\"hidden\":true,\"component\":\"view/interface/interfaceReport/DebugReport.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"DebugReport\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":43,\"CreatedAt\":\"2022-06-08T17:40:11.776+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.969+08:00\",\"parentId\":\"33\",\"path\":\"reportDetail/:report_id\",\"name\":\"reportDetail\",\"hidden\":true,\"component\":\"view/interface/Reports/reportDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试报告详情-${report_id}\",\"icon\":\"document\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":60,\"CreatedAt\":\"2023-05-09T17:39:19.121+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.042+08:00\",\"parentId\":\"33\",\"path\":\"jsonCompare\",\"name\":\"jsonCompare\",\"hidden\":true,\"component\":\"view/interface/interfaceComponents/jsonCompare.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"json\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":38,\"CreatedAt\":\"2022-05-17T12:54:39.571+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.955+08:00\",\"parentId\":\"33\",\"path\":\"testCaseStepDetail/:id\",\"name\":\"testCaseStepDetail\",\"hidden\":true,\"component\":\"view/interface/testCaseStep/testCaseStepDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"步骤详情-${id}\",\"icon\":\"finished\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":61,\"CreatedAt\":\"2023-05-09T17:39:19.132+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.049+08:00\",\"parentId\":\"33\",\"path\":\"env\",\"name\":\"env\",\"hidden\":false,\"component\":\"view/interface/environment/environment.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"环境变量\",\"icon\":\"grid\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":36,\"CreatedAt\":\"2022-05-11T19:34:31.473+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.945+08:00\",\"parentId\":\"33\",\"path\":\"apiConfig\",\"name\":\"apiConfig\",\"hidden\":false,\"component\":\"view/interface/interfaceTemplate/apiConfig.vue\",\"sort\":100,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"配置管理\",\"icon\":\"expand\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":34,\"CreatedAt\":\"2022-01-23T18:37:56.434+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.934+08:00\",\"parentId\":\"33\",\"path\":\"interfacetemplate\",\"name\":\"interfacetemplate\",\"hidden\":false,\"component\":\"view/interface/interfaceTemplate/interfaceTemplate.vue\",\"sort\":200,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"接口管理\",\"icon\":\"coin\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":37,\"CreatedAt\":\"2022-05-16T12:15:29.299+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.95+08:00\",\"parentId\":\"33\",\"path\":\"testCaseStep\",\"name\":\"testCaseStep\",\"hidden\":false,\"component\":\"view/interface/testCaseStep/testCaseStep.vue\",\"sort\":300,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试步骤\",\"icon\":\"suitcase\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":44,\"CreatedAt\":\"2022-06-21T17:34:55.263+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.974+08:00\",\"parentId\":\"33\",\"path\":\"testCase\",\"name\":\"testCase\",\"hidden\":false,\"component\":\"view/interface/apiCase/apiCase.vue\",\"sort\":400,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试用例\",\"icon\":\"briefcase\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":53,\"CreatedAt\":\"2022-06-21T17:34:55.263+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.003+08:00\",\"parentId\":\"33\",\"path\":\"timerTask\",\"name\":\"timerTask\",\"hidden\":false,\"component\":\"view/interface/timerTask/timerTask.vue\",\"sort\":500,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"定时任务\",\"icon\":\"timer\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":42,\"CreatedAt\":\"2022-06-07T13:37:19.135+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.961+08:00\",\"parentId\":\"33\",\"path\":\"report\",\"name\":\"report\",\"hidden\":false,\"component\":\"view/interface/Reports/report.vue\",\"sort\":600,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试报告\",\"icon\":\"compass\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":62,\"CreatedAt\":\"2023-05-09T17:39:19.144+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.057+08:00\",\"parentId\":\"33\",\"path\":\"py_pkg\",\"name\":\"py_pkg\",\"hidden\":false,\"component\":\"view/py_pkg/py_pkg.vue\",\"sort\":680,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"py库管理\",\"icon\":\"office-building\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":46,\"CreatedAt\":\"2022-07-12T11:39:29.347+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.979+08:00\",\"parentId\":\"33\",\"path\":\"debugtalk\",\"name\":\"debugtalk\",\"hidden\":false,\"component\":\"view/interface/debugtalk/debugtalk.vue\",\"sort\":700,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"驱动函数\",\"icon\":\"reading\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":47,\"CreatedAt\":\"2022-07-12T11:40:03.21+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.984+08:00\",\"parentId\":\"33\",\"path\":\"debugtalkGen\",\"name\":\"debugtalkGen\",\"hidden\":true,\"component\":\"view/interface/debugtalk/debugtalkGen.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"debugtalkGen\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":49,\"CreatedAt\":\"2022-08-01T06:08:02.383+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.989+08:00\",\"parentId\":\"33\",\"path\":\"apisCaseDetail/:id\",\"name\":\"apisCaseDetail\",\"hidden\":true,\"component\":\"view/interface/apiCase/apisCaseDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用例详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":51,\"CreatedAt\":\"2022-08-01T06:08:02.383+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.994+08:00\",\"parentId\":\"33\",\"path\":\"taskCaseDetail/:id\",\"name\":\"taskCaseDetail\",\"hidden\":true,\"component\":\"view/interface/timerTask/taskCaseDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"任务详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":35,\"CreatedAt\":\"2022-01-28T13:29:16.282+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.941+08:00\",\"parentId\":\"33\",\"path\":\"DebugReport\",\"name\":\"DebugReport\",\"hidden\":true,\"component\":\"view/interface/interfaceReport/DebugReport.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"DebugReport\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":43,\"CreatedAt\":\"2022-06-08T17:40:11.776+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.969+08:00\",\"parentId\":\"33\",\"path\":\"reportDetail/:report_id\",\"name\":\"reportDetail\",\"hidden\":true,\"component\":\"view/interface/Reports/reportDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试报告详情-${report_id}\",\"icon\":\"document\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":60,\"CreatedAt\":\"2023-05-09T17:39:19.121+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.042+08:00\",\"parentId\":\"33\",\"path\":\"jsonCompare\",\"name\":\"jsonCompare\",\"hidden\":true,\"component\":\"view/interface/interfaceComponents/jsonCompare.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"json\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":38,\"CreatedAt\":\"2022-05-17T12:54:39.571+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.955+08:00\",\"parentId\":\"33\",\"path\":\"testCaseStepDetail/:id\",\"name\":\"testCaseStepDetail\",\"hidden\":true,\"component\":\"view/interface/testCaseStep/testCaseStepDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"步骤详情-${id}\",\"icon\":\"finished\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":55,\"CreatedAt\":\"2022-10-08T19:30:47.868+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.017+08:00\",\"parentId\":\"0\",\"path\":\"performance\",\"name\":\"performance\",\"hidden\":false,\"component\":\"view/performance/index.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试\",\"icon\":\"stopwatch\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":56,\"CreatedAt\":\"2022-10-08T19:51:05.009+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.022+08:00\",\"parentId\":\"55\",\"path\":\"performanceTask\",\"name\":\"performanceTask\",\"hidden\":false,\"component\":\"view/performance/task/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能任务\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":58,\"CreatedAt\":\"2022-11-03T15:49:19.428+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.033+08:00\",\"parentId\":\"55\",\"path\":\"pReport\",\"name\":\"pReport\",\"hidden\":false,\"component\":\"view/performance/report.vue\",\"sort\":2,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试报告\",\"icon\":\"compass\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":59,\"CreatedAt\":\"2022-11-03T18:59:54.997+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.037+08:00\",\"parentId\":\"55\",\"path\":\"pReportDetail/:id\",\"name\":\"pReportDetail\",\"hidden\":true,\"component\":\"view/performance/reportDetail.vue\",\"sort\":999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试报告详情-${id}\",\"icon\":\"document\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":57,\"CreatedAt\":\"2022-10-11T14:27:10.146+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.028+08:00\",\"parentId\":\"55\",\"path\":\"performanceDetail/:id\",\"name\":\"performanceDetail\",\"hidden\":true,\"component\":\"view/performance/task/taskDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能任务详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":56,\"CreatedAt\":\"2022-10-08T19:51:05.009+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.022+08:00\",\"parentId\":\"55\",\"path\":\"performanceTask\",\"name\":\"performanceTask\",\"hidden\":false,\"component\":\"view/performance/task/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能任务\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":58,\"CreatedAt\":\"2022-11-03T15:49:19.428+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.033+08:00\",\"parentId\":\"55\",\"path\":\"pReport\",\"name\":\"pReport\",\"hidden\":false,\"component\":\"view/performance/report.vue\",\"sort\":2,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试报告\",\"icon\":\"compass\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":59,\"CreatedAt\":\"2022-11-03T18:59:54.997+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.037+08:00\",\"parentId\":\"55\",\"path\":\"pReportDetail/:id\",\"name\":\"pReportDetail\",\"hidden\":true,\"component\":\"view/performance/reportDetail.vue\",\"sort\":999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试报告详情-${id}\",\"icon\":\"document\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":57,\"CreatedAt\":\"2022-10-11T14:27:10.146+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.028+08:00\",\"parentId\":\"55\",\"path\":\"performanceDetail/:id\",\"name\":\"performanceDetail\",\"hidden\":true,\"component\":\"view/performance/task/taskDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能任务详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":14,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-11-22T14:50:24.888+08:00\",\"parentId\":\"0\",\"path\":\"systemTools\",\"name\":\"systemTools\",\"hidden\":false,\"component\":\"view/systemTools/index.vue\",\"sort\":15,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统工具\",\"icon\":\"tools\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":26,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"autoPkg\",\"name\":\"autoPkg\",\"hidden\":false,\"component\":\"view/systemTools/autoPkg/autoPkg.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化package\",\"icon\":\"folder\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":25,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeEdit/:id\",\"name\":\"autoCodeEdit\",\"hidden\":true,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码-${id}\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":24,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeAdmin\",\"name\":\"autoCodeAdmin\",\"hidden\":false,\"component\":\"view/systemTools/autoCodeAdmin/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码管理\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":15,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"autoCode\",\"name\":\"autoCode\",\"hidden\":false,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"代码生成器\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":16,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"formCreate\",\"name\":\"formCreate\",\"hidden\":false,\"component\":\"view/systemTools/formCreate/index.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"表单生成器\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":17,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"system\",\"name\":\"system\",\"hidden\":false,\"component\":\"view/systemTools/system/system.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统配置\",\"icon\":\"operation\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":26,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"autoPkg\",\"name\":\"autoPkg\",\"hidden\":false,\"component\":\"view/systemTools/autoPkg/autoPkg.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化package\",\"icon\":\"folder\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":25,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeEdit/:id\",\"name\":\"autoCodeEdit\",\"hidden\":true,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码-${id}\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":24,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"autoCodeAdmin\",\"name\":\"autoCodeAdmin\",\"hidden\":false,\"component\":\"view/systemTools/autoCodeAdmin/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码管理\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":15,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"autoCode\",\"name\":\"autoCode\",\"hidden\":false,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"代码生成器\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":16,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"formCreate\",\"name\":\"formCreate\",\"hidden\":false,\"component\":\"view/systemTools/formCreate/index.vue\",\"sort\":2,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"表单生成器\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":17,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"14\",\"path\":\"system\",\"name\":\"system\",\"hidden\":false,\"component\":\"view/systemTools/system/system.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统配置\",\"icon\":\"operation\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":27,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-11-22T14:50:34.751+08:00\",\"parentId\":\"0\",\"path\":\"plugin\",\"name\":\"plugin\",\"hidden\":false,\"component\":\"view/routerHolder.vue\",\"sort\":16,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件系统\",\"icon\":\"cherry\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":28,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"27\",\"path\":\"https://plugin.gin-vue-admin.com/\",\"name\":\"https://plugin.gin-vue-admin.com/\",\"hidden\":false,\"component\":\"https://plugin.gin-vue-admin.com/\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件市场\",\"icon\":\"shop\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":29,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"27\",\"path\":\"installPlugin\",\"name\":\"installPlugin\",\"hidden\":false,\"component\":\"view/systemTools/installPlugin/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件安装\",\"icon\":\"box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":30,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"27\",\"path\":\"autoPlug\",\"name\":\"autoPlug\",\"hidden\":false,\"component\":\"view/systemTools/autoPlug/autoPlug.vue\",\"sort\":2,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件模板\",\"icon\":\"folder\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":31,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"27\",\"path\":\"plugin-email\",\"name\":\"plugin-email\",\"hidden\":false,\"component\":\"plugin/email/view/index.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"邮件插件\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":28,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"27\",\"path\":\"https://plugin.gin-vue-admin.com/\",\"name\":\"https://plugin.gin-vue-admin.com/\",\"hidden\":false,\"component\":\"https://plugin.gin-vue-admin.com/\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件市场\",\"icon\":\"shop\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":29,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"27\",\"path\":\"installPlugin\",\"name\":\"installPlugin\",\"hidden\":false,\"component\":\"view/systemTools/installPlugin/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件安装\",\"icon\":\"box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":30,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"27\",\"path\":\"autoPlug\",\"name\":\"autoPlug\",\"hidden\":false,\"component\":\"view/systemTools/autoPlug/autoPlug.vue\",\"sort\":2,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件模板\",\"icon\":\"folder\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":31,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"27\",\"path\":\"plugin-email\",\"name\":\"plugin-email\",\"hidden\":false,\"component\":\"plugin/email/view/index.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"邮件插件\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":9,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-11-22T14:50:42.311+08:00\",\"parentId\":\"0\",\"path\":\"example\",\"name\":\"example\",\"hidden\":false,\"component\":\"view/example/index.vue\",\"sort\":17,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"示例文件\",\"icon\":\"management\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":10,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"excel\",\"name\":\"excel\",\"hidden\":false,\"component\":\"view/example/excel/excel.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"excel导入导出\",\"icon\":\"takeaway-box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":11,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"upload\",\"name\":\"upload\",\"hidden\":false,\"component\":\"view/example/upload/upload.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"媒体库（上传下载）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":12,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"breakpoint\",\"name\":\"breakpoint\",\"hidden\":false,\"component\":\"view/example/breakpoint/breakpoint.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传\",\"icon\":\"upload-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":21,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"simpleUploader\",\"name\":\"simpleUploader\",\"hidden\":false,\"component\":\"view/example/simpleUploader/simpleUploader\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传（插件版）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":13,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"customer\",\"name\":\"customer\",\"hidden\":false,\"component\":\"view/example/customer/customer.vue\",\"sort\":7,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"客户列表（资源示例）\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":10,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"excel\",\"name\":\"excel\",\"hidden\":false,\"component\":\"view/example/excel/excel.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"excel导入导出\",\"icon\":\"takeaway-box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":11,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"upload\",\"name\":\"upload\",\"hidden\":false,\"component\":\"view/example/upload/upload.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"媒体库（上传下载）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":12,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"breakpoint\",\"name\":\"breakpoint\",\"hidden\":false,\"component\":\"view/example/breakpoint/breakpoint.vue\",\"sort\":6,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传\",\"icon\":\"upload-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":13,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"9\",\"path\":\"customer\",\"name\":\"customer\",\"hidden\":false,\"component\":\"view/example/customer/customer.vue\",\"sort\":7,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"客户列表（资源示例）\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":23,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-11-22T14:50:49.106+08:00\",\"parentId\":\"0\",\"path\":\"state\",\"name\":\"state\",\"hidden\":false,\"component\":\"view/system/state.vue\",\"sort\":18,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"服务器状态\",\"icon\":\"cloudy\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"authorityId\":888}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` (`id`, `created_at`, `updated_at`, `deleted_at`, `ip`, `method`, `path`, `status`, `latency`, `agent`, `error_message`, `body`, `resp`, `user_id`) VALUES (7, '2023-07-09 18:36:48.603', '2023-07-09 18:36:48.603', NULL, '219.136.74.95', 'PUT', '/authority/updateAuthority', 200, 3782912, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36', '', '{\"authorityId\":666,\"AuthorityName\":\"飞书登录默认角色\",\"parentId\":0,\"defaultRouter\":\"关于我们\"}', '{\"code\":0,\"data\":{\"authority\":{\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"authorityId\":666,\"authorityName\":\"飞书登录默认角色\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"关于我们\"}},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` (`id`, `created_at`, `updated_at`, `deleted_at`, `ip`, `method`, `path`, `status`, `latency`, `agent`, `error_message`, `body`, `resp`, `user_id`) VALUES (8, '2023-07-09 18:36:49.581', '2023-07-09 18:36:49.581', NULL, '219.136.74.95', 'POST', '/menu/addMenuAuthority', 200, 7861725, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36', '', '{\"menus\":[{\"ID\":63,\"CreatedAt\":\"2023-07-09T18:35:48.941+08:00\",\"UpdatedAt\":\"2023-07-09T18:35:48.941+08:00\",\"parentId\":\"0\",\"path\":\"关于我们\",\"name\":\"关于我们\",\"hidden\":false,\"component\":\"view/about/index.vue\",\"sort\":0,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"关于我们\",\"icon\":\"info-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":22,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"0\",\"path\":\"http://www.yangfan.gd.cn/\",\"name\":\"http://www.yangfan.gd.cn/\",\"hidden\":false,\"component\":\"/\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"官方网站\",\"icon\":\"home-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":8,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-12T18:25:05.048+08:00\",\"parentId\":\"0\",\"path\":\"person\",\"name\":\"person\",\"hidden\":true,\"component\":\"view/person/person.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"个人信息\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":33,\"CreatedAt\":\"2022-01-23T18:34:58.647+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.928+08:00\",\"parentId\":\"0\",\"path\":\"interfaces\",\"name\":\"interfaces\",\"hidden\":false,\"component\":\"view/interface/index.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"接口自动化\",\"icon\":\"box\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":61,\"CreatedAt\":\"2023-05-09T17:39:19.132+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.049+08:00\",\"parentId\":\"33\",\"path\":\"env\",\"name\":\"env\",\"hidden\":false,\"component\":\"view/interface/environment/environment.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"环境变量\",\"icon\":\"grid\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":36,\"CreatedAt\":\"2022-05-11T19:34:31.473+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.945+08:00\",\"parentId\":\"33\",\"path\":\"apiConfig\",\"name\":\"apiConfig\",\"hidden\":false,\"component\":\"view/interface/interfaceTemplate/apiConfig.vue\",\"sort\":100,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"配置管理\",\"icon\":\"expand\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":34,\"CreatedAt\":\"2022-01-23T18:37:56.434+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.934+08:00\",\"parentId\":\"33\",\"path\":\"interfacetemplate\",\"name\":\"interfacetemplate\",\"hidden\":false,\"component\":\"view/interface/interfaceTemplate/interfaceTemplate.vue\",\"sort\":200,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"接口管理\",\"icon\":\"coin\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":37,\"CreatedAt\":\"2022-05-16T12:15:29.299+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.95+08:00\",\"parentId\":\"33\",\"path\":\"testCaseStep\",\"name\":\"testCaseStep\",\"hidden\":false,\"component\":\"view/interface/testCaseStep/testCaseStep.vue\",\"sort\":300,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试步骤\",\"icon\":\"suitcase\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":44,\"CreatedAt\":\"2022-06-21T17:34:55.263+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.974+08:00\",\"parentId\":\"33\",\"path\":\"testCase\",\"name\":\"testCase\",\"hidden\":false,\"component\":\"view/interface/apiCase/apiCase.vue\",\"sort\":400,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试用例\",\"icon\":\"briefcase\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":53,\"CreatedAt\":\"2022-06-21T17:34:55.263+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.003+08:00\",\"parentId\":\"33\",\"path\":\"timerTask\",\"name\":\"timerTask\",\"hidden\":false,\"component\":\"view/interface/timerTask/timerTask.vue\",\"sort\":500,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"定时任务\",\"icon\":\"timer\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":42,\"CreatedAt\":\"2022-06-07T13:37:19.135+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.961+08:00\",\"parentId\":\"33\",\"path\":\"report\",\"name\":\"report\",\"hidden\":false,\"component\":\"view/interface/Reports/report.vue\",\"sort\":600,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试报告\",\"icon\":\"compass\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":62,\"CreatedAt\":\"2023-05-09T17:39:19.144+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.057+08:00\",\"parentId\":\"33\",\"path\":\"py_pkg\",\"name\":\"py_pkg\",\"hidden\":false,\"component\":\"view/py_pkg/py_pkg.vue\",\"sort\":680,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"py库管理\",\"icon\":\"office-building\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":46,\"CreatedAt\":\"2022-07-12T11:39:29.347+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.979+08:00\",\"parentId\":\"33\",\"path\":\"debugtalk\",\"name\":\"debugtalk\",\"hidden\":false,\"component\":\"view/interface/debugtalk/debugtalk.vue\",\"sort\":700,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"驱动函数\",\"icon\":\"reading\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":47,\"CreatedAt\":\"2022-07-12T11:40:03.21+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.984+08:00\",\"parentId\":\"33\",\"path\":\"debugtalkGen\",\"name\":\"debugtalkGen\",\"hidden\":true,\"component\":\"view/interface/debugtalk/debugtalkGen.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"debugtalkGen\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":49,\"CreatedAt\":\"2022-08-01T06:08:02.383+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.989+08:00\",\"parentId\":\"33\",\"path\":\"apisCaseDetail/:id\",\"name\":\"apisCaseDetail\",\"hidden\":true,\"component\":\"view/interface/apiCase/apisCaseDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用例详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":51,\"CreatedAt\":\"2022-08-01T06:08:02.383+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.994+08:00\",\"parentId\":\"33\",\"path\":\"taskCaseDetail/:id\",\"name\":\"taskCaseDetail\",\"hidden\":true,\"component\":\"view/interface/timerTask/taskCaseDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"任务详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":35,\"CreatedAt\":\"2022-01-28T13:29:16.282+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.941+08:00\",\"parentId\":\"33\",\"path\":\"DebugReport\",\"name\":\"DebugReport\",\"hidden\":true,\"component\":\"view/interface/interfaceReport/DebugReport.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"DebugReport\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":43,\"CreatedAt\":\"2022-06-08T17:40:11.776+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.969+08:00\",\"parentId\":\"33\",\"path\":\"reportDetail/:report_id\",\"name\":\"reportDetail\",\"hidden\":true,\"component\":\"view/interface/Reports/reportDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试报告详情-${report_id}\",\"icon\":\"document\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":60,\"CreatedAt\":\"2023-05-09T17:39:19.121+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.042+08:00\",\"parentId\":\"33\",\"path\":\"jsonCompare\",\"name\":\"jsonCompare\",\"hidden\":true,\"component\":\"view/interface/interfaceComponents/jsonCompare.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"json\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":38,\"CreatedAt\":\"2022-05-17T12:54:39.571+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.955+08:00\",\"parentId\":\"33\",\"path\":\"testCaseStepDetail/:id\",\"name\":\"testCaseStepDetail\",\"hidden\":true,\"component\":\"view/interface/testCaseStep/testCaseStepDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"步骤详情-${id}\",\"icon\":\"finished\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":36,\"CreatedAt\":\"2022-05-11T19:34:31.473+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.945+08:00\",\"parentId\":\"33\",\"path\":\"apiConfig\",\"name\":\"apiConfig\",\"hidden\":false,\"component\":\"view/interface/interfaceTemplate/apiConfig.vue\",\"sort\":100,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"配置管理\",\"icon\":\"expand\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":34,\"CreatedAt\":\"2022-01-23T18:37:56.434+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.934+08:00\",\"parentId\":\"33\",\"path\":\"interfacetemplate\",\"name\":\"interfacetemplate\",\"hidden\":false,\"component\":\"view/interface/interfaceTemplate/interfaceTemplate.vue\",\"sort\":200,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"接口管理\",\"icon\":\"coin\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":37,\"CreatedAt\":\"2022-05-16T12:15:29.299+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.95+08:00\",\"parentId\":\"33\",\"path\":\"testCaseStep\",\"name\":\"testCaseStep\",\"hidden\":false,\"component\":\"view/interface/testCaseStep/testCaseStep.vue\",\"sort\":300,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试步骤\",\"icon\":\"suitcase\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":44,\"CreatedAt\":\"2022-06-21T17:34:55.263+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.974+08:00\",\"parentId\":\"33\",\"path\":\"testCase\",\"name\":\"testCase\",\"hidden\":false,\"component\":\"view/interface/apiCase/apiCase.vue\",\"sort\":400,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试用例\",\"icon\":\"briefcase\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":53,\"CreatedAt\":\"2022-06-21T17:34:55.263+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.003+08:00\",\"parentId\":\"33\",\"path\":\"timerTask\",\"name\":\"timerTask\",\"hidden\":false,\"component\":\"view/interface/timerTask/timerTask.vue\",\"sort\":500,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"定时任务\",\"icon\":\"timer\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":42,\"CreatedAt\":\"2022-06-07T13:37:19.135+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.961+08:00\",\"parentId\":\"33\",\"path\":\"report\",\"name\":\"report\",\"hidden\":false,\"component\":\"view/interface/Reports/report.vue\",\"sort\":600,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试报告\",\"icon\":\"compass\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":46,\"CreatedAt\":\"2022-07-12T11:39:29.347+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.979+08:00\",\"parentId\":\"33\",\"path\":\"debugtalk\",\"name\":\"debugtalk\",\"hidden\":false,\"component\":\"view/interface/debugtalk/debugtalk.vue\",\"sort\":700,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"驱动函数\",\"icon\":\"reading\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":47,\"CreatedAt\":\"2022-07-12T11:40:03.21+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.984+08:00\",\"parentId\":\"33\",\"path\":\"debugtalkGen\",\"name\":\"debugtalkGen\",\"hidden\":true,\"component\":\"view/interface/debugtalk/debugtalkGen.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"debugtalkGen\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":49,\"CreatedAt\":\"2022-08-01T06:08:02.383+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.989+08:00\",\"parentId\":\"33\",\"path\":\"apisCaseDetail/:id\",\"name\":\"apisCaseDetail\",\"hidden\":true,\"component\":\"view/interface/apiCase/apisCaseDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用例详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":51,\"CreatedAt\":\"2022-08-01T06:08:02.383+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.994+08:00\",\"parentId\":\"33\",\"path\":\"taskCaseDetail/:id\",\"name\":\"taskCaseDetail\",\"hidden\":true,\"component\":\"view/interface/timerTask/taskCaseDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"任务详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":35,\"CreatedAt\":\"2022-01-28T13:29:16.282+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.941+08:00\",\"parentId\":\"33\",\"path\":\"DebugReport\",\"name\":\"DebugReport\",\"hidden\":true,\"component\":\"view/interface/interfaceReport/DebugReport.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"DebugReport\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":43,\"CreatedAt\":\"2022-06-08T17:40:11.776+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.969+08:00\",\"parentId\":\"33\",\"path\":\"reportDetail/:report_id\",\"name\":\"reportDetail\",\"hidden\":true,\"component\":\"view/interface/Reports/reportDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试报告详情-${report_id}\",\"icon\":\"document\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":38,\"CreatedAt\":\"2022-05-17T12:54:39.571+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.955+08:00\",\"parentId\":\"33\",\"path\":\"testCaseStepDetail/:id\",\"name\":\"testCaseStepDetail\",\"hidden\":true,\"component\":\"view/interface/testCaseStep/testCaseStepDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"步骤详情-${id}\",\"icon\":\"finished\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":55,\"CreatedAt\":\"2022-10-08T19:30:47.868+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.017+08:00\",\"parentId\":\"0\",\"path\":\"performance\",\"name\":\"performance\",\"hidden\":false,\"component\":\"view/performance/index.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试\",\"icon\":\"stopwatch\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":56,\"CreatedAt\":\"2022-10-08T19:51:05.009+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.022+08:00\",\"parentId\":\"55\",\"path\":\"performanceTask\",\"name\":\"performanceTask\",\"hidden\":false,\"component\":\"view/performance/task/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能任务\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":58,\"CreatedAt\":\"2022-11-03T15:49:19.428+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.033+08:00\",\"parentId\":\"55\",\"path\":\"pReport\",\"name\":\"pReport\",\"hidden\":false,\"component\":\"view/performance/report.vue\",\"sort\":2,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试报告\",\"icon\":\"compass\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":59,\"CreatedAt\":\"2022-11-03T18:59:54.997+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.037+08:00\",\"parentId\":\"55\",\"path\":\"pReportDetail/:id\",\"name\":\"pReportDetail\",\"hidden\":true,\"component\":\"view/performance/reportDetail.vue\",\"sort\":999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试报告详情-${id}\",\"icon\":\"document\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":57,\"CreatedAt\":\"2022-10-11T14:27:10.146+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.028+08:00\",\"parentId\":\"55\",\"path\":\"performanceDetail/:id\",\"name\":\"performanceDetail\",\"hidden\":true,\"component\":\"view/performance/task/taskDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能任务详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":56,\"CreatedAt\":\"2022-10-08T19:51:05.009+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.022+08:00\",\"parentId\":\"55\",\"path\":\"performanceTask\",\"name\":\"performanceTask\",\"hidden\":false,\"component\":\"view/performance/task/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能任务\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":58,\"CreatedAt\":\"2022-11-03T15:49:19.428+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.033+08:00\",\"parentId\":\"55\",\"path\":\"pReport\",\"name\":\"pReport\",\"hidden\":false,\"component\":\"view/performance/report.vue\",\"sort\":2,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试报告\",\"icon\":\"compass\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":59,\"CreatedAt\":\"2022-11-03T18:59:54.997+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.037+08:00\",\"parentId\":\"55\",\"path\":\"pReportDetail/:id\",\"name\":\"pReportDetail\",\"hidden\":true,\"component\":\"view/performance/reportDetail.vue\",\"sort\":999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试报告详情-${id}\",\"icon\":\"document\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":57,\"CreatedAt\":\"2022-10-11T14:27:10.146+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.028+08:00\",\"parentId\":\"55\",\"path\":\"performanceDetail/:id\",\"name\":\"performanceDetail\",\"hidden\":true,\"component\":\"view/performance/task/taskDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能任务详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"authorityId\":666}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` (`id`, `created_at`, `updated_at`, `deleted_at`, `ip`, `method`, `path`, `status`, `latency`, `agent`, `error_message`, `body`, `resp`, `user_id`) VALUES (9, '2023-07-09 18:36:51.540', '2023-07-09 18:36:51.540', NULL, '219.136.74.95', 'POST', '/menu/addMenuAuthority', 200, 6993766, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36', '', '{\"menus\":[{\"ID\":63,\"CreatedAt\":\"2023-07-09T18:35:48.941+08:00\",\"UpdatedAt\":\"2023-07-09T18:35:48.941+08:00\",\"parentId\":\"0\",\"path\":\"关于我们\",\"name\":\"关于我们\",\"hidden\":false,\"component\":\"view/about/index.vue\",\"sort\":0,\"meta\":{\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"关于我们\",\"icon\":\"info-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":22,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"parentId\":\"0\",\"path\":\"http://www.yangfan.gd.cn/\",\"name\":\"http://www.yangfan.gd.cn/\",\"hidden\":false,\"component\":\"/\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"官方网站\",\"icon\":\"home-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":8,\"CreatedAt\":\"2022-07-10T15:12:16.165+08:00\",\"UpdatedAt\":\"2022-07-12T18:25:05.048+08:00\",\"parentId\":\"0\",\"path\":\"person\",\"name\":\"person\",\"hidden\":true,\"component\":\"view/person/person.vue\",\"sort\":3,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"个人信息\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":33,\"CreatedAt\":\"2022-01-23T18:34:58.647+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.928+08:00\",\"parentId\":\"0\",\"path\":\"interfaces\",\"name\":\"interfaces\",\"hidden\":false,\"component\":\"view/interface/index.vue\",\"sort\":4,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"接口自动化\",\"icon\":\"box\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":61,\"CreatedAt\":\"2023-05-09T17:39:19.132+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.049+08:00\",\"parentId\":\"33\",\"path\":\"env\",\"name\":\"env\",\"hidden\":false,\"component\":\"view/interface/environment/environment.vue\",\"sort\":0,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"环境变量\",\"icon\":\"grid\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":36,\"CreatedAt\":\"2022-05-11T19:34:31.473+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.945+08:00\",\"parentId\":\"33\",\"path\":\"apiConfig\",\"name\":\"apiConfig\",\"hidden\":false,\"component\":\"view/interface/interfaceTemplate/apiConfig.vue\",\"sort\":100,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"配置管理\",\"icon\":\"expand\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":34,\"CreatedAt\":\"2022-01-23T18:37:56.434+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.934+08:00\",\"parentId\":\"33\",\"path\":\"interfacetemplate\",\"name\":\"interfacetemplate\",\"hidden\":false,\"component\":\"view/interface/interfaceTemplate/interfaceTemplate.vue\",\"sort\":200,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"接口管理\",\"icon\":\"coin\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":37,\"CreatedAt\":\"2022-05-16T12:15:29.299+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.95+08:00\",\"parentId\":\"33\",\"path\":\"testCaseStep\",\"name\":\"testCaseStep\",\"hidden\":false,\"component\":\"view/interface/testCaseStep/testCaseStep.vue\",\"sort\":300,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试步骤\",\"icon\":\"suitcase\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":44,\"CreatedAt\":\"2022-06-21T17:34:55.263+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.974+08:00\",\"parentId\":\"33\",\"path\":\"testCase\",\"name\":\"testCase\",\"hidden\":false,\"component\":\"view/interface/apiCase/apiCase.vue\",\"sort\":400,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试用例\",\"icon\":\"briefcase\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":53,\"CreatedAt\":\"2022-06-21T17:34:55.263+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.003+08:00\",\"parentId\":\"33\",\"path\":\"timerTask\",\"name\":\"timerTask\",\"hidden\":false,\"component\":\"view/interface/timerTask/timerTask.vue\",\"sort\":500,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"定时任务\",\"icon\":\"timer\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":42,\"CreatedAt\":\"2022-06-07T13:37:19.135+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.961+08:00\",\"parentId\":\"33\",\"path\":\"report\",\"name\":\"report\",\"hidden\":false,\"component\":\"view/interface/Reports/report.vue\",\"sort\":600,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试报告\",\"icon\":\"compass\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":62,\"CreatedAt\":\"2023-05-09T17:39:19.144+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.057+08:00\",\"parentId\":\"33\",\"path\":\"py_pkg\",\"name\":\"py_pkg\",\"hidden\":false,\"component\":\"view/py_pkg/py_pkg.vue\",\"sort\":680,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"py库管理\",\"icon\":\"office-building\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":46,\"CreatedAt\":\"2022-07-12T11:39:29.347+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.979+08:00\",\"parentId\":\"33\",\"path\":\"debugtalk\",\"name\":\"debugtalk\",\"hidden\":false,\"component\":\"view/interface/debugtalk/debugtalk.vue\",\"sort\":700,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"驱动函数\",\"icon\":\"reading\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":47,\"CreatedAt\":\"2022-07-12T11:40:03.21+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.984+08:00\",\"parentId\":\"33\",\"path\":\"debugtalkGen\",\"name\":\"debugtalkGen\",\"hidden\":true,\"component\":\"view/interface/debugtalk/debugtalkGen.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"debugtalkGen\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":49,\"CreatedAt\":\"2022-08-01T06:08:02.383+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.989+08:00\",\"parentId\":\"33\",\"path\":\"apisCaseDetail/:id\",\"name\":\"apisCaseDetail\",\"hidden\":true,\"component\":\"view/interface/apiCase/apisCaseDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用例详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":51,\"CreatedAt\":\"2022-08-01T06:08:02.383+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.994+08:00\",\"parentId\":\"33\",\"path\":\"taskCaseDetail/:id\",\"name\":\"taskCaseDetail\",\"hidden\":true,\"component\":\"view/interface/timerTask/taskCaseDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"任务详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":35,\"CreatedAt\":\"2022-01-28T13:29:16.282+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.941+08:00\",\"parentId\":\"33\",\"path\":\"DebugReport\",\"name\":\"DebugReport\",\"hidden\":true,\"component\":\"view/interface/interfaceReport/DebugReport.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"DebugReport\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":43,\"CreatedAt\":\"2022-06-08T17:40:11.776+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.969+08:00\",\"parentId\":\"33\",\"path\":\"reportDetail/:report_id\",\"name\":\"reportDetail\",\"hidden\":true,\"component\":\"view/interface/Reports/reportDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试报告详情-${report_id}\",\"icon\":\"document\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":60,\"CreatedAt\":\"2023-05-09T17:39:19.121+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.042+08:00\",\"parentId\":\"33\",\"path\":\"jsonCompare\",\"name\":\"jsonCompare\",\"hidden\":true,\"component\":\"view/interface/interfaceComponents/jsonCompare.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"json\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":38,\"CreatedAt\":\"2022-05-17T12:54:39.571+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.955+08:00\",\"parentId\":\"33\",\"path\":\"testCaseStepDetail/:id\",\"name\":\"testCaseStepDetail\",\"hidden\":true,\"component\":\"view/interface/testCaseStep/testCaseStepDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"步骤详情-${id}\",\"icon\":\"finished\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":36,\"CreatedAt\":\"2022-05-11T19:34:31.473+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.945+08:00\",\"parentId\":\"33\",\"path\":\"apiConfig\",\"name\":\"apiConfig\",\"hidden\":false,\"component\":\"view/interface/interfaceTemplate/apiConfig.vue\",\"sort\":100,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"配置管理\",\"icon\":\"expand\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":34,\"CreatedAt\":\"2022-01-23T18:37:56.434+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.934+08:00\",\"parentId\":\"33\",\"path\":\"interfacetemplate\",\"name\":\"interfacetemplate\",\"hidden\":false,\"component\":\"view/interface/interfaceTemplate/interfaceTemplate.vue\",\"sort\":200,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"接口管理\",\"icon\":\"coin\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":37,\"CreatedAt\":\"2022-05-16T12:15:29.299+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.95+08:00\",\"parentId\":\"33\",\"path\":\"testCaseStep\",\"name\":\"testCaseStep\",\"hidden\":false,\"component\":\"view/interface/testCaseStep/testCaseStep.vue\",\"sort\":300,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试步骤\",\"icon\":\"suitcase\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":44,\"CreatedAt\":\"2022-06-21T17:34:55.263+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.974+08:00\",\"parentId\":\"33\",\"path\":\"testCase\",\"name\":\"testCase\",\"hidden\":false,\"component\":\"view/interface/apiCase/apiCase.vue\",\"sort\":400,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试用例\",\"icon\":\"briefcase\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":53,\"CreatedAt\":\"2022-06-21T17:34:55.263+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.003+08:00\",\"parentId\":\"33\",\"path\":\"timerTask\",\"name\":\"timerTask\",\"hidden\":false,\"component\":\"view/interface/timerTask/timerTask.vue\",\"sort\":500,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"定时任务\",\"icon\":\"timer\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":42,\"CreatedAt\":\"2022-06-07T13:37:19.135+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.961+08:00\",\"parentId\":\"33\",\"path\":\"report\",\"name\":\"report\",\"hidden\":false,\"component\":\"view/interface/Reports/report.vue\",\"sort\":600,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试报告\",\"icon\":\"compass\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":46,\"CreatedAt\":\"2022-07-12T11:39:29.347+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.979+08:00\",\"parentId\":\"33\",\"path\":\"debugtalk\",\"name\":\"debugtalk\",\"hidden\":false,\"component\":\"view/interface/debugtalk/debugtalk.vue\",\"sort\":700,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"驱动函数\",\"icon\":\"reading\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":47,\"CreatedAt\":\"2022-07-12T11:40:03.21+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.984+08:00\",\"parentId\":\"33\",\"path\":\"debugtalkGen\",\"name\":\"debugtalkGen\",\"hidden\":true,\"component\":\"view/interface/debugtalk/debugtalkGen.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"debugtalkGen\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":49,\"CreatedAt\":\"2022-08-01T06:08:02.383+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.989+08:00\",\"parentId\":\"33\",\"path\":\"apisCaseDetail/:id\",\"name\":\"apisCaseDetail\",\"hidden\":true,\"component\":\"view/interface/apiCase/apisCaseDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用例详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":51,\"CreatedAt\":\"2022-08-01T06:08:02.383+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.994+08:00\",\"parentId\":\"33\",\"path\":\"taskCaseDetail/:id\",\"name\":\"taskCaseDetail\",\"hidden\":true,\"component\":\"view/interface/timerTask/taskCaseDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"任务详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":35,\"CreatedAt\":\"2022-01-28T13:29:16.282+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.941+08:00\",\"parentId\":\"33\",\"path\":\"DebugReport\",\"name\":\"DebugReport\",\"hidden\":true,\"component\":\"view/interface/interfaceReport/DebugReport.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"DebugReport\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":43,\"CreatedAt\":\"2022-06-08T17:40:11.776+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.969+08:00\",\"parentId\":\"33\",\"path\":\"reportDetail/:report_id\",\"name\":\"reportDetail\",\"hidden\":true,\"component\":\"view/interface/Reports/reportDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"测试报告详情-${report_id}\",\"icon\":\"document\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":38,\"CreatedAt\":\"2022-05-17T12:54:39.571+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:46.955+08:00\",\"parentId\":\"33\",\"path\":\"testCaseStepDetail/:id\",\"name\":\"testCaseStepDetail\",\"hidden\":true,\"component\":\"view/interface/testCaseStep/testCaseStepDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"步骤详情-${id}\",\"icon\":\"finished\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":55,\"CreatedAt\":\"2022-10-08T19:30:47.868+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.017+08:00\",\"parentId\":\"0\",\"path\":\"performance\",\"name\":\"performance\",\"hidden\":false,\"component\":\"view/performance/index.vue\",\"sort\":5,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试\",\"icon\":\"stopwatch\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":56,\"CreatedAt\":\"2022-10-08T19:51:05.009+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.022+08:00\",\"parentId\":\"55\",\"path\":\"performanceTask\",\"name\":\"performanceTask\",\"hidden\":false,\"component\":\"view/performance/task/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能任务\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":58,\"CreatedAt\":\"2022-11-03T15:49:19.428+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.033+08:00\",\"parentId\":\"55\",\"path\":\"pReport\",\"name\":\"pReport\",\"hidden\":false,\"component\":\"view/performance/report.vue\",\"sort\":2,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试报告\",\"icon\":\"compass\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":59,\"CreatedAt\":\"2022-11-03T18:59:54.997+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.037+08:00\",\"parentId\":\"55\",\"path\":\"pReportDetail/:id\",\"name\":\"pReportDetail\",\"hidden\":true,\"component\":\"view/performance/reportDetail.vue\",\"sort\":999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试报告详情-${id}\",\"icon\":\"document\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":57,\"CreatedAt\":\"2022-10-11T14:27:10.146+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.028+08:00\",\"parentId\":\"55\",\"path\":\"performanceDetail/:id\",\"name\":\"performanceDetail\",\"hidden\":true,\"component\":\"view/performance/task/taskDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能任务详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":56,\"CreatedAt\":\"2022-10-08T19:51:05.009+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.022+08:00\",\"parentId\":\"55\",\"path\":\"performanceTask\",\"name\":\"performanceTask\",\"hidden\":false,\"component\":\"view/performance/task/index.vue\",\"sort\":1,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能任务\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":58,\"CreatedAt\":\"2022-11-03T15:49:19.428+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.033+08:00\",\"parentId\":\"55\",\"path\":\"pReport\",\"name\":\"pReport\",\"hidden\":false,\"component\":\"view/performance/report.vue\",\"sort\":2,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试报告\",\"icon\":\"compass\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":59,\"CreatedAt\":\"2022-11-03T18:59:54.997+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.037+08:00\",\"parentId\":\"55\",\"path\":\"pReportDetail/:id\",\"name\":\"pReportDetail\",\"hidden\":true,\"component\":\"view/performance/reportDetail.vue\",\"sort\":999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能测试报告详情-${id}\",\"icon\":\"document\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":57,\"CreatedAt\":\"2022-10-11T14:27:10.146+08:00\",\"UpdatedAt\":\"2023-07-09T18:23:47.028+08:00\",\"parentId\":\"55\",\"path\":\"performanceDetail/:id\",\"name\":\"performanceDetail\",\"hidden\":true,\"component\":\"view/performance/task/taskDetail.vue\",\"sort\":99999,\"meta\":{\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"性能任务详情-${id}\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"authorityId\":666}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_authority`;
CREATE TABLE `sys_user_authority` (
  `sys_user_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`,`sys_authority_authority_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_user_authority
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_authority` (`sys_user_id`, `sys_authority_authority_id`) VALUES (1, 888);
INSERT INTO `sys_user_authority` (`sys_user_id`, `sys_authority_authority_id`) VALUES (3, 888);
INSERT INTO `sys_user_authority` (`sys_user_id`, `sys_authority_authority_id`) VALUES (3, 8881);
INSERT INTO `sys_user_authority` (`sys_user_id`, `sys_authority_authority_id`) VALUES (3, 9528);
INSERT INTO `sys_user_authority` (`sys_user_id`, `sys_authority_authority_id`) VALUES (6, 666);
INSERT INTO `sys_user_authority` (`sys_user_id`, `sys_authority_authority_id`) VALUES (6, 888);
INSERT INTO `sys_user_authority` (`sys_user_id`, `sys_authority_authority_id`) VALUES (6, 9528);
INSERT INTO `sys_user_authority` (`sys_user_id`, `sys_authority_authority_id`) VALUES (8, 888);
INSERT INTO `sys_user_authority` (`sys_user_id`, `sys_authority_authority_id`) VALUES (8, 8881);
INSERT INTO `sys_user_authority` (`sys_user_id`, `sys_authority_authority_id`) VALUES (8, 9528);
INSERT INTO `sys_user_authority` (`sys_user_id`, `sys_authority_authority_id`) VALUES (8, 9999);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_project
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_project`;
CREATE TABLE `sys_user_project` (
  `project_id` bigint unsigned NOT NULL,
  `sys_user_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`project_id`,`sys_user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_user_project
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (1, 0);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (1, 1);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (1, 2);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (1, 3);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (1, 5);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (1, 6);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (1, 7);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (1, 8);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (1, 9);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (1, 10);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (1, 11);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (1, 12);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (1, 13);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (1, 14);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (1, 15);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (3, 0);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (3, 2);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (4, 1);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (4, 2);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (5, 1);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (6, 1);
INSERT INTO `sys_user_project` (`project_id`, `sys_user_id`) VALUES (14, 8);
COMMIT;

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '系统用户' COMMENT '用户昵称',
  `side_mode` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'dark' COMMENT '用户侧边主题',
  `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `base_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '#fff' COMMENT '基础颜色',
  `active_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '#1890ff' COMMENT '活跃颜色',
  `authority_id` bigint unsigned DEFAULT '888' COMMENT '用户角色ID',
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户邮箱',
  `enable` bigint DEFAULT '1' COMMENT '用户是否被冻结 1正常 2冻结',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sys_users_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
BEGIN;
INSERT INTO `sys_users` (`id`, `created_at`, `updated_at`, `deleted_at`, `uuid`, `username`, `password`, `nick_name`, `side_mode`, `header_img`, `base_color`, `active_color`, `authority_id`, `phone`, `email`, `enable`) VALUES (1, '2022-07-10 15:12:16.014', '2023-05-18 23:32:43.297', NULL, 'ffb9f7c4-806f-434b-8c8b-74f3aac2989f', 'admin', '$2a$10$R4871PWFfWuJ5kessAMg0e2QLLu.L/C/wogK5AcRzb2KMP1GD9wUe', 'admin', 'dark', 'https://qmplusimg.henrongyi.top/gva_header.jpg', '#fff', '#1890ff', 888, '1111111111111111111111111111', 'qwdqwdq', 1);
INSERT INTO `sys_users` (`id`, `created_at`, `updated_at`, `deleted_at`, `uuid`, `username`, `password`, `nick_name`, `side_mode`, `header_img`, `base_color`, `active_color`, `authority_id`, `phone`, `email`, `enable`) VALUES (8, '2022-09-03 18:34:50.587', '2022-11-22 14:47:55.186', NULL, 'c2846189-7de7-4b57-8e8b-12f2716b9759', 'yangfan', '$2a$10$R4871PWFfWuJ5kessAMg0e2QLLu.L/C/wogK5AcRzb2KMP1GD9wUe', 'yangfan', 'dark', 'https://qmplusimg.henrongyi.top/gva_header.jpg', '#fff', '#1890ff', 888, '', '', 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
