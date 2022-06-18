/*
 Navicat Premium Data Transfer

 Source Server         : mysql8
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : localhost:3306
 Source Schema         : djwk_test

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 13/06/2022 18:16:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_open_copy
-- ----------------------------
DROP TABLE IF EXISTS `t_open_copy`;
CREATE TABLE `t_open_copy`  (
  `setl_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '36284113',
  `mdtrt_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '56741021',
  `psn_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '陈雪英',
  `psn_cert_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '01',
  `psn_cert_type_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '居民身份证（户口簿）',
  `certno` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '430121193310307923',
  `psn_no` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '43000020100030375675',
  `gend` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '2',
  `gend_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '女',
  `naty` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '01',
  `naty_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '汉族',
  `insutype` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '390',
  `insutype_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '城乡居民基本医疗保险',
  `psn_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '16',
  `psn_type_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '居民(老)',
  `begndate` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '2022-01-01',
  `enddate` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '2022-01-01',
  `setl_time` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '2022-01-01 08:09:49',
  `medfee_sumamt` double NULL DEFAULT NULL COMMENT '576.07',
  `hi_agre_sumfee` double NULL DEFAULT NULL COMMENT '576.07',
  `fund_pay_sumamt` bigint NULL DEFAULT NULL COMMENT '210',
  `psn_pay` double NULL DEFAULT NULL COMMENT '366.07',
  `acct_pay` bigint NULL DEFAULT NULL COMMENT '0',
  `cash_payamt` double NULL DEFAULT NULL COMMENT '366.07',
  `fulamt_ownpay_amt` bigint NULL DEFAULT NULL COMMENT '0',
  `preselfpay_amt` bigint NULL DEFAULT NULL COMMENT '0',
  `inscp_amt` double NULL DEFAULT NULL COMMENT '576.07',
  `dedc_std` bigint NULL DEFAULT NULL COMMENT '0',
  `crt_dedc` bigint NULL DEFAULT NULL COMMENT '0',
  `act_pay_dedc` bigint NULL DEFAULT NULL COMMENT '0',
  `hifp_pay` bigint NULL DEFAULT NULL COMMENT '210',
  `pool_prop_selfpay` double NULL DEFAULT NULL COMMENT '0.7',
  `cvlserv_pay` bigint NULL DEFAULT NULL COMMENT '0',
  `hifes_pay` bigint NULL DEFAULT NULL COMMENT '0',
  `hifmi_pay` bigint NULL DEFAULT NULL COMMENT '0',
  `hifob_pay` bigint NULL DEFAULT NULL COMMENT '0',
  `hifdm_pay` bigint NULL DEFAULT NULL COMMENT '0',
  `maf_pay` bigint NULL DEFAULT NULL COMMENT '0',
  `othfund_pay` bigint NULL DEFAULT NULL COMMENT '0',
  `cvlserv_flag` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '0',
  `cvlserv_flag_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '否',
  `cvlserv_lv` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'null',
  `cvlserv_lv_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'null',
  `sp_psn_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'null',
  `sp_psn_type_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'null',
  `sp_psn_type_lv` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'null',
  `sp_psn_type_lv_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'null',
  `clct_grde` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'null',
  `clct_grde_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'null',
  `flxempe_flag` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '0',
  `flxempe_flag_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '否',
  `nwb_flag` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'null',
  `nwb_flag_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'null',
  `insu_admdvs` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '430105',
  `insu_admdvs_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '开福区',
  `emp_no` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '430105040300',
  `emp_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '迎宾路社区',
  `emp_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '01',
  `emp_type_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '01',
  `emp_mgt_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'null',
  `emp_mgt_type_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'null',
  `pay_loc` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '2',
  `pay_loc_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '医疗机构',
  `fixmedins_code` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'P43010201038',
  `fixmedins_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '湖南达嘉维康医药产业股份有限公司五一路分店',
  `hosp_lv` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '0',
  `hosp_lv_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '0',
  `mdtrt_cert_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '02',
  `mdtrt_cert_type_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '居民身份证',
  `med_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '14',
  `med_type_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '门诊慢特病',
  `setl_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '2',
  `setl_type_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联网结算',
  `clr_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '11',
  `clr_type_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '门诊',
  `clr_way` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '1',
  `clr_way_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '按项目',
  `clr_optins` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '430199',
  `clr_optins_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '长沙市市本级',
  `refd_setl_flag` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '0',
  `refd_setl_flag_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '否',
  `mdtrt_cert_no` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '430121193310307923',
  `dise_no` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'M01603',
  `dise_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '糖尿病并发症',
  PRIMARY KEY (`setl_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
