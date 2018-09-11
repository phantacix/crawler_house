CREATE DATABASE szrem;

CREATE TABLE `licence` (
  `id` int(11) NOT NULL,
  `licence_id` varchar(45) DEFAULT '' COMMENT '许可证号',
  `land_agent` varchar(45) NOT NULL COMMENT '房地产商',
  `location` varchar(45) DEFAULT '' COMMENT '区域',
  `licencecol` varchar(45) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='许可证详细信息';

CREATE TABLE `project` (
  `id` int(11) NOT NULL,
  `licence_id` varchar(45) DEFAULT '' COMMENT '已经入库的预售或现售许可证号',
  `name` varchar(45) DEFAULT '' COMMENT '项目名称',
  `location` varchar(45) DEFAULT '' COMMENT '宗地位置',
  `building_id` varchar(45) DEFAULT '' COMMENT '楼名',
  `land_use_time` int(11) DEFAULT '0' COMMENT '使用年限',
  `pre_sale_count` int(11) DEFAULT '0' COMMENT '预售总套数',
  `on_sale_count` int(11) DEFAULT '0' COMMENT '现售总套数',
  `assignee_date` date DEFAULT NULL COMMENT '受让日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='项目表';


CREATE TABLE `house` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `building` varchar(255) DEFAULT '' COMMENT '项目楼栋情况',
  `block` varchar(255) DEFAULT NULL COMMENT '座号',
  `type` varchar(255) DEFAULT NULL COMMENT '户型',
  `contract` varchar(255) DEFAULT NULL COMMENT '合同号',
  `filing_price` int(11) DEFAULT NULL COMMENT '备案价格',
  `room` int(11) DEFAULT NULL COMMENT '房号     ',
  `house_type` varchar(255) DEFAULT NULL COMMENT '用途',
  `pre_sales_area` decimal(5,2) DEFAULT NULL COMMENT '预售查丈 建筑面积',
  `pre_unit_cst_area` decimal(5,2) DEFAULT NULL COMMENT '预售查丈 户内面积',
  `pre_shared_public_area` decimal(5,2) DEFAULT NULL COMMENT '预售查丈 分摊面积',
  `c_sales_area` decimal(5,2) DEFAULT NULL COMMENT '竣工查丈 建筑面积',
  `c_unit_cst_area` decimal(5,2) DEFAULT NULL COMMENT '竣工查丈 户内面积',
  `c_shared_public_area` decimal(5,2) DEFAULT NULL COMMENT '竣工查丈 分摊面积',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='套房详细信息\n'